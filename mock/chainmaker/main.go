package main

import (
	"chainmaker.org/chainmaker/common/v2/random/uuid"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	"encoding/json"
	"log"
	"relay-chain/module/plugin"
	"relay-chain/module/plugin/chainmaker"
	"relay-chain/module/types/pb"
)

var (
	CrossExampleContractName = "CrossExample"
	Chain1IP                 = "127.0.0.1:1234"
	Chain1IPConsensus        = "127.0.0.1:1238"
	IpFabric                 = "127.0.0.1:1233"
	Chain2IP                 = "127.0.0.1:1235"
	Chain3IP                 = "127.0.0.1:1236"
	Chain4IP                 = "127.0.0.1:1237"
)

func main() {
	cmClient := chainmaker.ChainMaker{
		ChainId: "chain1",
	}
	err := cmClient.InitClient("./configs/chainmaker/chain1/sdk_config_org1_client1.yml")
	if err != nil {
		log.Fatalln(err)
	}

	crossId := uuid.GetUUID()
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
	/*chain3TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "100"},
	}
	chain3TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "100"},
	}

	chain5TxExe := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "minus",
		Params:       map[string]string{"number": "400"},
	}
	chain5TxRoll := &pb.Transaction{
		ContractName: CrossExampleContractName,
		Method:       "plus",
		Params:       map[string]string{"number": "400"},
	}*/
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
	/*chain3Txs := &pb.ChainTransaction{
		ChainId:             "chain3",
		ExecuteTransaction:  chain3TxExe,
		RollbackTransaction: chain3TxRoll,
	}

	chain5Txs := &pb.ChainTransaction{
		ChainId:             "mychannel",
		ExecuteTransaction:  chain5TxExe,
		RollbackTransaction: chain5TxRoll,
	}*/
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
	/*chain3Add := &pb.ChainRPCAdd{
		ListenIP: Chain3IP,
	}

	chain5Add := &pb.ChainRPCAdd{
		ListenIP: IpFabric,
	}*/
	request := &pb.Request{
		CrossId: crossId,
		// to be sure txs and RPCAdds have identical chain name
		Transactions: map[string]*pb.ChainTransaction{"chain1": chain1Txs, "chain2": chain2Txs},
		RPCAdds:      map[string]*pb.ChainRPCAdd{"chain1": chain1Add, "chain2": chain2Add},
	}
	payload, err := json.Marshal(request)
	log.Printf("payload [%s]", string(payload))
	if err != nil {
		log.Fatalf("fail marshal request,err:[%s]", err)
	}
	kvs := []*common.KeyValuePair{
		{
			Key:   "CrossId",
			Value: []byte(crossId),
		},
		{
			Key:   "Payload",
			Value: payload,
		},
	}

	resp, err := cmClient.Client.InvokeContract(plugin.CrossTransactionContractName, plugin.StartCross, "", kvs, -1, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("crossId:[%s] txid:[%s], response:[%s]", crossId, resp.TxId, resp.ContractResult.Result)
}
