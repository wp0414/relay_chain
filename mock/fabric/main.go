package main

import (
	"chainmaker.org/chainmaker/common/v2/random/uuid"
	"cross-listener/module/adapter/fabric"
	"cross-listener/module/types/pb"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"log"
)

var (
	CrossExampleAddressFlato = "0x3bf478f2472798311fbb3d45737b4909ca92b1ef"
	CrossExampleContractName = "CrossExample"
	IpFabric                 = "127.0.0.1:1233"
	Chain1IP                 = "127.0.0.1:1234"
	Chain2IP                 = "127.0.0.1:1235"
	Chain3IP                 = "127.0.0.1:1236"
	Chain4IP                 = "127.0.0.1:1237"
	FlatoIp                  = "127.0.0.1:1238"
)

func fabricAndFiveChainmaker() {

	fabClient := fabric.Fabric{
		ChainId: "mychannel",
	}
	err := fabClient.InitClient("./configs/fabric/fabric_sdk.yml")
	if err != nil {
		log.Fatal(err)
	}

	crossId := uuid.GetUUID()
	chain1TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain1TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}
	chain2TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain2TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}
	chain3TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain3TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}
	chain4TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain4TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}
	chain5TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "1"},
	}
	chain5TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "400"},
	}
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
	chain3Txs := &pb.ChainTransaction{
		ChainId:             "chain3",
		ExecuteTransaction:  chain3TxExe,
		RollbackTransaction: chain3TxRoll,
	}
	chain4Txs := &pb.ChainTransaction{
		ChainId:             "chain4",
		ExecuteTransaction:  chain4TxExe,
		RollbackTransaction: chain4TxRoll,
	}
	chain5Txs := &pb.ChainTransaction{
		ChainId:             "mychannel",
		ExecuteTransaction:  chain5TxExe,
		RollbackTransaction: chain5TxRoll,
	}
	chain1Add := &pb.ChainRPCAdd{
		ListenIP: Chain1IP,
	}
	chain2Add := &pb.ChainRPCAdd{
		ListenIP: Chain2IP,
	}
	chain3Add := &pb.ChainRPCAdd{
		ListenIP: Chain3IP,
	}
	chain4Add := &pb.ChainRPCAdd{
		ListenIP: Chain4IP,
	}
	chain5Add := &pb.ChainRPCAdd{
		ListenIP: IpFabric,
	}
	request := &pb.Request{
		CrossId: crossId,
		// to be sure txs and RPCAdds have identical chain name
		Transactions: map[string]*pb.ChainTransaction{"chain1": chain1Txs, "chain2": chain2Txs, "chain3": chain3Txs, "chain4": chain4Txs, "mychannel": chain5Txs},
		RPCAdds:      map[string]*pb.ChainRPCAdd{"chain1": chain1Add, "chain2": chain2Add, "chain3": chain3Add, "chain4": chain4Add, "mychannel": chain5Add},
	}
	payload, err := json.Marshal(request)
	log.Printf("payload [%s]", string(payload))
	if err != nil {
		log.Fatalf("fail marshal request,err:[%s]", err)
	}
	params := [][]byte{
		[]byte(crossId),
		payload,
	}
	req := channel.Request{
		ChaincodeID: "Cross",
		Fcn:         "StartCross",
		Args:        params,
	}
	response, err := fabClient.ChannelClient.Execute(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("crossId:[%s] txid:[%s], response:[%s]", crossId, response.TransactionID, response.Payload)
}

func fabricAndChainmaker() {

	fabClient := fabric.Fabric{
		ChainId: "mychannel",
	}
	err := fabClient.InitClient("./configs/fabric/fabric_sdk.yml")
	if err != nil {
		log.Fatal(err)
	}

	crossId := uuid.GetUUID()
	chain1TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain1TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}
	chain2TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "1"},
	}
	chain2TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "400"},
	}
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
	chain1Add := &pb.ChainRPCAdd{
		ListenIP: Chain1IP,
	}
	chain2Add := &pb.ChainRPCAdd{
		ListenIP: IpFabric,
	}
	request := &pb.Request{
		CrossId: crossId,
		// to be sure txs and RPCAdds have identical chain name
		Transactions: map[string]*pb.ChainTransaction{"chain1": chain1Txs, "mychannel": chain2Txs},
		RPCAdds:      map[string]*pb.ChainRPCAdd{"chain1": chain1Add, "mychannel": chain2Add},
	}
	payload, err := json.Marshal(request)
	log.Printf("payload [%s]", string(payload))
	if err != nil {
		log.Fatalf("fail marshal request,err:[%s]", err)
	}
	params := [][]byte{
		[]byte(crossId),
		payload,
	}
	req := channel.Request{
		ChaincodeID: "Cross",
		Fcn:         "StartCross",
		Args:        params,
	}
	response, err := fabClient.ChannelClient.Execute(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("crossId:[%s] txid:[%s], response:[%s]", crossId, response.TransactionID, response.Payload)
}

func fabricAndFlato() {

	fabClient := fabric.Fabric{
		ChainId: "mychannel",
	}
	err := fabClient.InitClient("./configs/fabric/fabric_sdk.yml")
	if err != nil {
		log.Fatal(err)
	}

	crossId := uuid.GetUUID()
	chain1TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain1TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}

	// flato Params`s key is the data type of value
	chain2TxExe := &pb.Transaction{
		ContractName: CrossExampleAddressFlato,
		Method:       "minus",
		Params:       map[string]string{"int": "100"},
	}
	chain2TxRoll := &pb.Transaction{
		ContractName: CrossExampleAddressFlato,
		Method:       "plus",
		Params:       map[string]string{"int": "100"},
	}
	chain1Txs := &pb.ChainTransaction{
		ChainId:             "mychannel",
		ExecuteTransaction:  chain1TxExe,
		RollbackTransaction: chain1TxRoll,
	}
	chain2Txs := &pb.ChainTransaction{
		ChainId:             "flato",
		ExecuteTransaction:  chain2TxExe,
		RollbackTransaction: chain2TxRoll,
	}

	chain1Add := &pb.ChainRPCAdd{
		ListenIP: IpFabric,
	}
	chain2Add := &pb.ChainRPCAdd{
		ListenIP: FlatoIp,
	}

	request := &pb.Request{
		CrossId: crossId,
		// to be sure txs and RPCAdds have identical chain name
		Transactions: map[string]*pb.ChainTransaction{"mychannel": chain1Txs, "flato": chain2Txs},
		RPCAdds:      map[string]*pb.ChainRPCAdd{"mychannel": chain1Add, "flato": chain2Add},
	}
	payload, err := json.Marshal(request)
	log.Printf("payload [%s]", string(payload))
	if err != nil {
		log.Fatalf("fail marshal request,err:[%s]", err)
	}
	params := [][]byte{
		[]byte(crossId),
		payload,
	}
	req := channel.Request{
		ChaincodeID: "Cross",
		Fcn:         "StartCross",
		Args:        params,
	}
	response, err := fabClient.ChannelClient.Execute(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("crossId:[%s] txid:[%s], response:[%s]", crossId, response.TransactionID, response.Payload)
}

func main() {
	fabricAndChainmaker()
}
