package listener

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"relay-chain/module/server"
	"relay-chain/module/types/pb"
	"time"
)

const (
	timeOut  = time.Second * time.Duration(20)
	relayAdd = "localhost:1240"
)

func BeginCrossChain(eventData []byte) {
	request := new(pb.Request)
	err := json.Unmarshal(eventData, request)
	if err != nil {
		log.Printf("fail unmarshal request [%s]\n", err)
		return
	}

	log.Printf("============>start process cross [%s]==========\n", request.CrossId)

	db := server.DefaultServer.DB
	err = db.PutRequest(request.CrossId, string(eventData))
	if err != nil {
		log.Printf("put request failed, err:[%s]", err)
	}
	sendRequest(request)
}

func sendRequest(request *pb.Request) {
	//获取中继链rpc服务的客户端
	client, err := rpcClient(relayAdd)
	if err != nil {
		log.Printf("[%s] fail new rpc client [%s]", request.CrossId, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	_, err = client.DealRequest(ctx, request)
	if err != nil {
		log.Printf("The crossChain request [%s] fail, [%s]", request.CrossId, err)
		return
	}
}

func rpcClient(add string) (pb.ServiceClient, error) {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return nil, err
	}
	c := pb.NewServiceClient(conn)
	return c, nil
}
