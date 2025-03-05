package chainmaker

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"relay-chain/module/crypto"
	"relay-chain/module/listener"
	"relay-chain/module/plugin"
	"relay-chain/module/server"
	"relay-chain/module/types/pb"
)

var _ plugin.ChainPlugin = &ChainMaker{}

type ChainMaker struct {
	ChainId string
	Client  *sdk.ChainClient
}

func (cm *ChainMaker) GetChainId() string {
	return cm.ChainId
}

func (cm *ChainMaker) InitClient(configYmlPath string) error {
	client, err := sdk.NewChainClient(sdk.WithConfPath(configYmlPath))
	if err != nil {
		return err
	}
	cm.Client = client
	return nil
}

func (cm *ChainMaker) InvokeContract(contractName, method, txId string,
	params map[string]string, withSyncResult bool, limit *common.Limit) (*common.TxResponse, error) {

	resp, err := cm.Client.InvokeContractWithLimit(contractName, method, txId, cm.ToKvs(params), -1, withSyncResult, limit)
	if err != nil {
		return nil, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return nil, fmt.Errorf("invoke contract failed, [code:%d] [msg:%s]\n", resp.Code, resp.Message)
	}

	if !withSyncResult {
		log.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
	} else {
		log.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	}

	return resp, nil
}

func (cm *ChainMaker) Execute(request *pb.Request) error {
	crossTx, ok := request.Transactions[cm.ChainId]
	if !ok {
		return errors.New("failed find crossTx")
	}
	exeTx := crossTx.ExecuteTransaction
	resp, err := cm.InvokeContract(exeTx.ContractName, exeTx.Method, "", exeTx.Params, true, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", exeTx.ContractName, exeTx.Method, err))
	}
	if request.Profs == nil {
		request.Profs = make(map[string]*pb.Prof)
	}
	request.Profs[cm.ChainId] = &pb.Prof{Key: resp.TxId, Result: string(resp.ContractResult.Result)}
	return nil
}

func (cm *ChainMaker) ToKvs(params map[string]string) []*common.KeyValuePair {
	var kvs []*common.KeyValuePair
	for key, value := range params {
		kvs = append(kvs, &common.KeyValuePair{Key: key, Value: []byte(value)})
	}
	return kvs
}

type FinalRPCReply struct {
	ChainId string
	Success bool
	Prof    []byte
}

type FinalResponse struct {
	CrossId string
	Success bool
	Done    bool
	Result  map[string]FinalRPCReply
}

func (cm *ChainMaker) Commit(response *pb.Response) error {
	finalResult := make(map[string]FinalRPCReply)
	key := []byte("abcdef12345678901234567890abcdef")
	for chainId, reply := range response.Result {
		plaintext, _ := json.Marshal(reply.Prof)

		// 加密
		ciphertext, err := crypto.AESEncryptCTR(plaintext, key)
		if err != nil {
			fmt.Println("Error encrypting:", err)
			return err
		}
		finalResult[chainId] = FinalRPCReply{
			ChainId: reply.ChainId,
			Success: reply.Success,
			Prof:    ciphertext,
		}
	}

	finalResponse := &FinalResponse{
		CrossId: response.CrossId,
		Success: response.Success,
		Done:    response.Done,
		Result:  finalResult,
	}

	responseBytes, err := json.Marshal(finalResponse)
	if err != nil {
		return err
	}

	err = server.DefaultServer.DB.PutResponse(response.CrossId, string(responseBytes))
	if err != nil {
		log.Println("put response failed")
	} else {
		log.Println("put response success")
	}

	params := map[string]string{
		plugin.CrossId:       response.CrossId,
		plugin.Status:        plugin.SuccessStatus,
		plugin.ResponseParam: string(responseBytes),
	}
	_, err = cm.InvokeContract(plugin.CrossTransactionContractName, plugin.SetCrossStatusResponse, "", params, true, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", plugin.CrossTransactionContractName, plugin.SetCrossStatusResponse, err))
	}
	return nil
}

func (cm *ChainMaker) Rollback(request *pb.Request) error {
	err := server.DefaultServer.DB.PutResponse(request.CrossId, plugin.RollbackResponse)
	if err != nil {
		log.Println("put Rollback response failed")
	} else {
		log.Println("put Rollback response success")
	}

	crossTx, ok := request.Transactions[cm.ChainId]
	if !ok {
		return errors.New("fail find crossTx")
	}
	rollTx := crossTx.RollbackTransaction
	_, err = cm.InvokeContract(rollTx.ContractName, rollTx.Method, "", rollTx.Params, true, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", rollTx.ContractName, rollTx.Method, err))
	}
	return nil
}

func (cm *ChainMaker) GetTxByTxId(txId string) *pb.Prof {
	transactionInfo, err := cm.Client.GetTxByTxId(txId)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &pb.Prof{
		Key:    txId,
		Result: string(transactionInfo.Transaction.Result.ContractResult.Result),
	}
}

func (cm *ChainMaker) SubscribeCrossEvent() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//订阅实时最新区块的Cross合约的CrossStart主题
	c, err := cm.Client.SubscribeContractEvent(ctx, -1, -1, plugin.CrossTransactionContractName, plugin.CrossStartTopic)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		select {
		case event, ok := <-c:
			if !ok {
				log.Fatalln("listen chan is close!")
			}
			if event == nil {
				log.Println("[listen] [waring] nil event")
				continue
			}
			//断言，event为接口类型
			contractEventInfo, ok := event.(*common.ContractEventInfo)
			if !ok {
				log.Println("[listen] [waring] type conversion failed")
				continue
			}
			log.Printf("[listen] [receive] [%d] => %+v\n", contractEventInfo.BlockHeight, contractEventInfo)
			//EventData接收CrossStart主题提交的字符串数组（[]string{string(payloadModified)}）
			go listener.BeginCrossChain([]byte(contractEventInfo.EventData[0]))
			//if err := client.Stop(); err != nil {
			//	return
			//}
			//return
		case <-ctx.Done():
			log.Print(ctx.Err())
			return
		}
	}
}
