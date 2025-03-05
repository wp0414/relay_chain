package main

import (
	"chainmaker.org/chainmaker/common/v2/random/uuid"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"relay-chain/module/types/pb"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// test const var
const (
	SendTime            = time.Minute * 20
	SendSpanTime        = time.Second * 1
	SendSpanTimeoutTime = time.Second * 10
	logSpanTime         = time.Second * 10
	ConcurrenceNum      = 1
)

var (
	isTpsFile = true
	tpsFile   *os.File

	startTime int64

	errSend     = new(int64)
	successSend = new(int64)
	totalSend   = new(int64)
	wg          sync.WaitGroup
	ctx         context.Context

	clients        []*sdk.ChainClient
	chain1TokenCli *sdk.ChainClient
	chain2TokenCli *sdk.ChainClient
	relayTokenCli  *sdk.ChainClient
	kvs            []*common.KeyValuePair

	chain1SdkConfig = "./configs/chainmaker/chain1/sdk_config_org1_client1.yml"
	chain2SdkConfig = "./configs/chainmaker/chain2/sdk_config_org2_client1.yml"
	relaySdkConfig  = "./configs/chainmaker/chain3/sdk_config_org3_client1.yml"
)

var (
	CrossExampleContractName = "CrossExample"
	Chain1IP                 = "127.0.0.1:1234"
	Chain1IPConsensus        = "127.0.0.1:1238"
	Chain2IP                 = "127.0.0.1:1235"
	Chain4IP                 = "127.0.0.1:1237"
	crossId                  string
)

var resChan = make(chan string, 20)

func generateClients() {
	var err error
	if chain1TokenCli, err = sdk.NewChainClient(sdk.WithConfPath(chain1SdkConfig)); err != nil {
		log.Fatalln(err)
	}
	if chain2TokenCli, err = sdk.NewChainClient(sdk.WithConfPath(chain2SdkConfig)); err != nil {
		log.Fatalln(err)
	}

	if relayTokenCli, err = sdk.NewChainClient(sdk.WithConfPath(relaySdkConfig)); err != nil {
		log.Fatalln(err)
	}

	// only send to chain1
	for i := 0; i < ConcurrenceNum; i++ {
		if temp, err := sdk.NewChainClient(sdk.WithConfPath(chain1SdkConfig)); err != nil {
			log.Fatalln(err)
		} else {
			clients = append(clients, temp)
		}
	}
	log.Printf("Success init %d cm clients", ConcurrenceNum)
}
func toKvs(params map[string][]byte) []*common.KeyValuePair {
	var kvs []*common.KeyValuePair
	for key, value := range params {
		kvs = append(kvs, &common.KeyValuePair{Key: key, Value: value})
	}
	return kvs
}

var queryParams = map[string][]byte{
	"method": []byte("show"),
}

func getToken(chain1 bool) (int, error) {
	var (
		resp *common.TxResponse
		err  error
	)
	if chain1 {
		resp, err = chain1TokenCli.QueryContract(CrossExampleContractName, "invoke_contract", toKvs(queryParams), -1)
	} else {
		resp, err = chain2TokenCli.QueryContract(CrossExampleContractName, "invoke_contract", toKvs(queryParams), -1)
	}

	if err != nil {
		return 0, err
	}
	token, err := strconv.ParseFloat(string(resp.ContractResult.Result), 64)
	if err != nil {
		return 0, err
	}

	return int(token), nil
}

func checkCross() {
	ticker := time.NewTicker(logSpanTime)
	for {
		select {
		case <-ticker.C:
			//infoResTimeOut()
			info()
		}
	}
}

func info() {
	timeResult := float64(time.Now().UnixNano()-startTime) / 1e9
	token1, err1 := getToken(true)  //链1的token
	token2, err2 := getToken(false) //链2的token
	if err1 != nil || err2 != nil {
		log.Printf("failed get token, chain1 err:%s, chain2 err:%s", err1, err2)
		return
	}
	token1 = token1 - 10000000000 + 10000
	token2 = token2 - 10000000000
	count := float64(token2)
	tokenDifference := token1 - token2
	str := fmt.Sprintln("[totalSend]:", atomic.LoadInt64(totalSend), "[SuccessSend]:", atomic.LoadInt64(successSend), "[errSend]:", atomic.LoadInt64(errSend), "[Duration]:", strconv.FormatFloat(timeResult, 'g', 9, 64)+" s", "[TPS]:", count/timeResult, "[SuccessCross]:", count, "[Token1]:", token1, "[Token2]:", token2, "[TokenDiff]:", tokenDifference)
	if isTpsFile {
		if _, err := tpsFile.WriteString(str); err != nil {
			log.Println(err)
		}
	}
	log.Println(str)
}

func infoResTimeOut() {
	for resReq := range resChan {
		go func(resReq string) {
			timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
			defer cancelFunc()
			params := map[string][]byte{
				"CrossId": []byte(resReq),
				"Type":    []byte("Response"),
			}
			begin := time.Now().UnixNano()
		loop:
			select {
			case <-timeout.Done():
				log.Printf("failed get checkRes res for %s,err:%s", resReq, timeout.Err())
				return
			case <-time.After(time.Millisecond * 100):
				resp, err := relayTokenCli.QueryContract("Cross", "ShowCross", toKvs(params), -1)
				if err != nil || resp.Code != common.TxStatusCode_SUCCESS || len(resp.ContractResult.Result) < 1 {
					goto loop
				}
				log.Printf("query result:%s\n", resp)
			}

			timeRes := float64(time.Now().UnixNano()-begin) / 1e9
			str := fmt.Sprintln("[totalSend]:", atomic.LoadInt64(totalSend), "[SuccessSend]:", atomic.LoadInt64(successSend), "[errSend]:", atomic.LoadInt64(errSend), "[Timeout]:", strconv.FormatFloat(timeRes, 'g', 9, 64)+" s", "[TPS]:", 0, "[SuccessCross]:", 0, "[Token1]:", 0, "[Token2]:", 0, "[TokenDiff]:", 0)

			if _, err := tpsFile.WriteString(str); err != nil {
				log.Println(err)
			}

			log.Println(str)
		}(resReq)
	}
}

func test() {
	startTime = time.Now().UnixNano()
	for i := 0; i < ConcurrenceNum; i++ {
		go func(index int) {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					//log.Printf("goroutine[%d] exit", index)
					return
				default:
					//log.Printf("goroutine[%d] exec", index)
					//transferTimeout(index)
					transfer(index)
				}
			}
		}(i)
	}
}

func construct() {
	crossId = uuid.GetUUID()
	chain1TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "1"},
	}
	chain1TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "1"},
	}
	chain2TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "1"},
	}
	chain2TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "1"},
	}
	//chain4TxExe := &pb.Transaction{
	//	ContractName: CrossExampleContractName,
	//	Method:       "plus",
	//	Params:       map[string]string{"number": "1"},
	//}
	//chain4TxRoll := &pb.Transaction{
	//	ContractName: CrossExampleContractName,
	//	Method:       "minus",
	//	Params:       map[string]string{"number": "1"},
	//}

	chain1Txs := &pb.ChainTransaction{
		ChainId:             "chain1",
		ExecuteTransaction:  chain1TxExe,
		RollbackTransaction: chain1TxRoll,
	}
	chain2Txs := &pb.ChainTransaction{
		ChainId:             "chain2",
		ExecuteTransaction:  chain2TxExe,
		RollbackTransaction: chain2TxRoll,
	}
	//chain4Txs := &pb.ChainTransaction{
	//	ChainId:             "chain4",
	//	ExecuteTransaction:  chain4TxExe,
	//	RollbackTransaction: chain4TxRoll,
	//}

	chain1Add := &pb.ChainRPCAdd{
		ListenIP:     Chain1IP,
		ConsensusIPs: []string{Chain1IPConsensus},
	}
	chain2Add := &pb.ChainRPCAdd{
		ListenIP: Chain2IP,
	}
	//chain4Add := &pb.ChainRPCAdd{
	//	ListenIP: Chain4IP,
	//}

	request := &pb.Request{
		CrossId: crossId,
		// to be sure txs and RPCAdds have identical chain name
		Transactions: map[string]*pb.ChainTransaction{"chain1": chain1Txs, "chain2": chain2Txs},
		RPCAdds:      map[string]*pb.ChainRPCAdd{"chain1": chain1Add, "chain2": chain2Add},
	}
	payload, err := json.Marshal(request)
	//log.Printf("payload [%s]", string(payload))
	if err != nil {
		log.Fatalf("fail marshal request,err:[%s]", err)
	}
	kvs = []*common.KeyValuePair{
		{
			Key:   "CrossId",
			Value: []byte(crossId),
		},
		{
			Key:   "Payload",
			Value: payload,
		},
	}
}
func transfer(index int) {
	construct()
	if _, err := clients[index].InvokeContractWithLimit("Cross", "StartCross", "", kvs, -1, false, nil); err != nil {
		atomic.AddInt64(errSend, 1)
		log.Println("err send: ", err)
	} else {
		atomic.AddInt64(successSend, 1)
	}
	atomic.AddInt64(totalSend, 1)
	time.Sleep(SendSpanTime)

}

func transferTimeout(index int) {
	construct()
	time.Sleep(SendSpanTimeoutTime)
	if _, err := clients[index].InvokeContractWithLimit("Cross", "StartCross", "", kvs, -1, false, nil); err != nil {
		atomic.AddInt64(errSend, 1)
		log.Println("err send: ", err)
	} else {
		resChan <- crossId
		atomic.AddInt64(successSend, 1)
	}
	atomic.AddInt64(totalSend, 1)

}

func initToken() {
	token1, err := getToken(true)
	token2, err := getToken(false)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("init, token1:", token1, "token2:", token2)
}

func main() {
	if isTpsFile {
		logFilename := "cmTPS1.log"
		var err error
		tpsFile, err = os.OpenFile("./"+logFilename, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			if tpsFile != nil {
				tpsFile.Close()
			}
		}()
	}

	generateClients()
	//initToken()
	//construct()

	log.Println("Testing...")
	var cancelFunc context.CancelFunc
	ctx, cancelFunc = context.WithDeadline(context.Background(), time.Now().Add(SendTime))
	defer cancelFunc()
	wg.Add(ConcurrenceNum)
	go test()
	go checkCross()
	wg.Wait()
	log.Printf("Done! Test Program has been sending for %d minutes,  %d goroutines exit", SendTime/time.Minute, ConcurrenceNum)
	time.Sleep(logSpanTime)
}
