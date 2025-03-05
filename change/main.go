package main

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"log"
)

var (
	clients        []*sdk.ChainClient
	chain1TokenCli *sdk.ChainClient
	chain2TokenCli *sdk.ChainClient

	chain1SdkConfig = "./configs/chainmaker/chain1/sdk_config_org1_client1.yml"
	chain2SdkConfig = "./configs/chainmaker/chain2/sdk_config_org2_client1.yml"
	relaySdkConfig  = "./configs/chainmaker/chain3/sdk_config_org3_client1.yml"
)

const (
	ConcurrenceNum = 1
)

func generateClients() {
	var err error
	if chain1TokenCli, err = sdk.NewChainClient(sdk.WithConfPath(chain1SdkConfig)); err != nil {
		log.Fatalln(err)
	}
	if chain2TokenCli, err = sdk.NewChainClient(sdk.WithConfPath(chain2SdkConfig)); err != nil {
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

var param1 = map[string][]byte{
	"number": []byte("9999900000"),
}

var param2 = map[string][]byte{
	"number": []byte("10000000000"),
}

func main() {
	generateClients()
	_, err := chain1TokenCli.InvokeContract("CrossExample", "minus", "", toKvs(param1), -1, false)
	if err != nil {
		log.Fatal(err)
	}
	_, err = chain2TokenCli.InvokeContract("CrossExample", "minus", "", toKvs(param2), -1, false)
	if err != nil {
		log.Fatal(err)
	}
}
