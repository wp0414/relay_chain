package flato

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/meshplus/gosdk/abi"
	"github.com/meshplus/gosdk/account"
	"github.com/meshplus/gosdk/common"
	"github.com/meshplus/gosdk/rpc"
	"io/ioutil"
	"log"
	"relay-chain/module/plugin"
	"relay-chain/module/server"
	"relay-chain/module/types/pb"
	"strconv"
	"strings"
)

type Flato struct {
	ChainId string
	Client  struct {
		RPC *rpc.RPC
		Key *account.SM2Key
	}
	CrossTransactionABI abi.ABI
}

var (
	_ plugin.ChainPlugin = &Flato{}

	accountFileName                 = "/account.json"
	CrossTransactionContractAddress = ""
	confPath                        = ""
)

func getAbiPath(abi string) string {
	return confPath + "/abi/" + abi + ".abi"
}

func (f *Flato) GetChainId() string {
	return f.ChainId
}

func (f *Flato) InitClient(configPath string) error {
	confPath = configPath

	abiStr, err := common.ReadFileAsString(getAbiPath(CrossTransactionContractAddress))
	if err != nil {
		return err
	}

	ABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return err
	}

	f.CrossTransactionABI = ABI

	accJsonPath := configPath + accountFileName
	var accountJson string
	var err1, err2 error
	accountJson, err1 = common.ReadFileAsString(accJsonPath)
	isNew := false
	if err1 != nil {
		log.Printf("read [%s] failed, will use a new account!", accJsonPath)
		isNew = true
	}

	if isNew {
		accountJson, err2 = account.NewAccountSm2("")
		if err2 != nil {
			return err2
		}

		err := ioutil.WriteFile(accJsonPath, []byte(accountJson), 0644)
		if err != nil {
			log.Printf("write account file failed [%s]", err)
		}
	}
	f.Client.Key, err2 = account.NewAccountSm2FromAccountJSON(accountJson, "")
	if err2 != nil {
		return err2
	}
	f.Client.RPC = rpc.NewRPCWithPath(configPath)
	log.Printf("your account address is [%s]", f.Client.Key.GetAddress().String())
	return nil
}

func (f *Flato) InvokeContract(contractAddress string, method string, kvs map[string]string) (*rpc.TxReceipt, error) {
	abiPath := getAbiPath(contractAddress)

	abiStr, err := common.ReadFileAsString(abiPath)
	if err != nil {
		return nil, err
	}

	ABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}

	packed, err := ABI.Pack(method, f.Parse(kvs)...)
	if err != nil {
		return nil, err
	}

	tranInvoke := rpc.NewTransaction(f.Client.Key.GetAddress().String()).Invoke(contractAddress, packed)
	tranInvoke.Sign(f.Client.Key)
	txReceipt, stdErr := f.Client.RPC.InvokeContract(tranInvoke)
	if stdErr != nil || !txReceipt.Valid {
		log.Printf("call [%s]-[%s] failed, resp: [ %#v ]", contractAddress, method, txReceipt)
		return nil, errors.New(fmt.Sprintf("invoke contract failed: [%s], txReceipt.Valid maybe false", stdErr))
	}

	return txReceipt, nil
}

func (f *Flato) Execute(request *pb.Request) error {
	crossTx, ok := request.Transactions[f.ChainId]
	if !ok {
		return errors.New("fail find crossTx")
	}
	exeTx := crossTx.ExecuteTransaction
	resp, err := f.InvokeContract(exeTx.ContractName, exeTx.Method, exeTx.Params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", exeTx.ContractName, exeTx.Method, err))
	}
	if request.Profs == nil {
		request.Profs = make(map[string]*pb.Prof)
	}
	request.Profs[f.ChainId] = &pb.Prof{Key: resp.TxHash, Result: resp.Ret}
	return nil
}

// Parse , for the flato, assume the key of parmas is the value data type.
func (f *Flato) Parse(params map[string]string) []interface{} {
	var kvs []interface{}
	for key, value := range params {
		if key == "int" {
			temp, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalln(err)
			}
			kvs = append(kvs, uint64(temp))
			continue
		}
		kvs = append(kvs, value)
	}
	return kvs
}

func (f *Flato) Commit(response *pb.Response) error {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = server.DefaultServer.DB.PutResponse(response.CrossId, string(responseBytes))
	if err != nil {
		log.Printf("put response failed, err:[%s]", err)
	}

	params := map[string]string{
		plugin.CrossId:       response.CrossId,
		plugin.Status:        plugin.SuccessStatus,
		plugin.ResponseParam: string(responseBytes),
	}
	_, err = f.InvokeContract(CrossTransactionContractAddress, plugin.SetCrossStatusResponse, params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", CrossTransactionContractAddress, plugin.SetCrossStatusResponse, err))
	}
	return nil
}

func (f *Flato) Rollback(request *pb.Request) error {
	err := server.DefaultServer.DB.PutResponse(request.CrossId, plugin.RollbackResponse)
	if err != nil {
		log.Printf("put Rollback response failed, err:[%s]", err)
	}

	crossTx, ok := request.Transactions[f.ChainId]
	if !ok {
		return errors.New("fail find crossTx")
	}
	rollTx := crossTx.RollbackTransaction
	_, err = f.InvokeContract(rollTx.ContractName, rollTx.Method, rollTx.Params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", rollTx.ContractName, rollTx.Method, err))
	}
	return nil
}

func (f *Flato) GetTxByTxId(txId string) *pb.Prof {
	txReceipt, err := f.Client.RPC.GetTxReceipt(txId, false)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &pb.Prof{
		Key:    txId,
		Result: txReceipt.Ret,
	}
}

func (f *Flato) SubscribeCrossEvent() {
	wsCli := f.Client.RPC.GetWebSocketClient()

	filter := rpc.NewLogsFilter().
		AddAddress(CrossTransactionContractAddress).
		SetTopic(0, f.CrossTransactionABI.Events["CrossStart"].Id())

	_, err := wsCli.Subscribe(1, filter, listenerHandler{})
	if err != nil {
		log.Println(err)
	}
	log.Println("[[[flato dose not support websocket, can not be a chain listener!]]]")
}

type listenerHandler struct {
}

func (lH listenerHandler) OnSubscribe() {
	log.Println("[OnSubscribe]")
}

func (lH listenerHandler) OnMessage(data []byte) {
	log.Println(data)
}

func (lH listenerHandler) OnUnSubscribe() {
	log.Println("[OnUnSubscribe]")
}

func (lH listenerHandler) OnClose() {
	log.Println("[OnClose]")
}
