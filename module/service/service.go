package service

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"relay-chain/module/server"
	"relay-chain/module/types/pb"
	"time"
)

const (
	maxTry  = 10
	timeOut = time.Second * time.Duration(20)
)

type Service struct {
	pb.UnimplementedServiceServer
}

// Invoke business contract on this chain
func (s *Service) Invoke(ctx context.Context, request *pb.Request) (*pb.RPCReply, error) {
	log.Printf("[receive invoke] [%s]", request.CrossId)
	chainPlugin := server.DefaultServer.Plugin
	// verify original prof
	err := server.DefaultServer.Verifier.VerifyRequest(request)
	if err != nil {
		log.Printf("[verify original prof fail] [%s] [%s]", request.CrossId, err)
		return nil, err
	}
	// Execute
	err = chainPlugin.Execute(request)
	if err != nil {
		log.Printf("[invoke failed] [%s] [%s]", request.CrossId, err)
		return nil, err
	}

	reply := pb.RPCReply{
		ChainId: chainPlugin.GetChainId(),
		Success: true,
		Prof:    request.Profs[chainPlugin.GetChainId()],
	}
	log.Printf("[%s] [invoke success]=>[%v]", request.CrossId, &reply)
	return &reply, nil

}

// Commit response on this chain
func (s *Service) Commit(ctx context.Context, response *pb.Response) (*pb.RPCReply, error) {
	log.Printf("[receive commit] [%s]", response.CrossId)
	err := server.DefaultServer.Plugin.Commit(response)
	if err != nil {
		log.Printf("commit for [%s] failed [%s]", response.CrossId, err)
		return nil, err
	}
	log.Printf("commit for [%s] success", response.CrossId)
	return &pb.RPCReply{
		ChainId: server.DefaultServer.Plugin.GetChainId(),
		Success: true,
		Prof:    nil,
	}, nil
}

// Rollback business contract on this chain
func (s *Service) Rollback(ctx context.Context, request *pb.Request) (*pb.RPCReply, error) {
	log.Printf("[receive rollback] [%s]", request.CrossId)
	err := server.DefaultServer.Plugin.Rollback(request)
	if err != nil {
		log.Printf("rollback for [%s] failed [%s]", request.CrossId, err)
		return nil, err
	}
	log.Printf("rollback for [%s] success", request.CrossId)
	return &pb.RPCReply{
		ChainId: server.DefaultServer.Plugin.GetChainId(),
		Success: true,
		Prof:    nil,
	}, nil
}

// Prove the request is valid or not
func (s *Service) Prove(ctx context.Context, prof *pb.Prof) (*pb.Prof, error) {
	log.Printf("[receive verify] [%s]", prof.Key)
	resp := server.DefaultServer.Plugin.GetTxByTxId(prof.Key)
	if resp == nil {
		return nil, errors.New("get tx by txId failed")
	}
	return resp, nil
}

// GetResponse return the response of crossId
func (s *Service) GetResponse(ctx context.Context, prof *pb.Prof) (*pb.Prof, error) {
	response, err := server.DefaultServer.DB.GetResponse(prof.Key)
	if err != nil {
		return nil, err
	}
	return &pb.Prof{
		Key:    prof.Key,
		Result: response,
	}, nil
}

func (s *Service) DealRequest(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	var originChainId string
	log.Printf("[receive request] [%s]", request.CrossId)

	verifier := server.DefaultServer.Verifier
	chainPlugin := server.DefaultServer.Plugin

	for chainId, _ := range request.Profs {
		originChainId = chainId
	}

	//todo 中继网关进行交易验证

	response := pb.Response{
		CrossId: request.CrossId,
		Success: true,
		Done:    false,
		Result:  make(map[string]*pb.RPCReply),
	}

	// Profs is nil, which mean CrossTransaction do not invoke business contract, so need invoke on this chain
	/*if request.Profs == nil {
		log.Printf("==>CrossTransaction did not invoke business contract, need to start invoke on this chain...")
		request.Origin = thisChainId
		// start invoke business contract on this chain
		err = chainAdapter.Execute(request)
		if err != nil {
			response.Done = true
			response.Result[thisChainId] = &pb.RPCReply{
				ChainId: thisChainId,
				Success: false,
				Prof:    &pb.Prof{},
			}
			log.Printf("[%s] fail execute in this chain [%s], [terminate cross transaction execution] \n", request.CrossId, err)
			err = chainAdapter.Commit(&response)
			if err != nil {
				log.Printf("[%s] [%s] commit fail response failed [%s]\n", request.CrossId, thisChainId, err)
			}
			return
		}
		log.Printf("invoke business contract on this chain success<==")
	}*/

	// invoke business contract on this chain done by Cross contract or done by the `above code`
	response.Result[originChainId] = &pb.RPCReply{
		ChainId: originChainId,
		Success: true,
		Prof:    request.Profs[originChainId],
	}

	// start invoke business contract for other chain
	callNum := len(request.RPCAdds) - 1
	log.Println("[callNum]=", callNum)

	rpcReplies := make(chan *pb.RPCReply, callNum)

	for chainId, RPCAdd := range request.RPCAdds {

		if chainId == originChainId {
			continue
		}

		if len(RPCAdd.ListenIP) == 0 {
			log.Printf("[%s] [%s] no listen ip", request.CrossId, chainId)
			callNum--
			continue
		}
		go dispatchRequest(rpcReplies, request, RPCAdd.ListenIP, chainId)
	}

	hasReceived := 0

	// hasSuccVerified mean the number of success verifyRPCReply reply, only when hasSuccVerified==callNum can be regard as crossTx success,
	// and if it will be modified by more than one goroutine, so use sync.Mutex or sync.RWMutex or atomic package,
	// commit or rollback should be after verifyRPCReply done, so do not exist concurrent
	hasSuccVerified := 0

	// start collect responses of the other chains
	for reply := range rpcReplies {
		hasReceived++
		response.Result[reply.ChainId] = reply

		if !reply.Success {
			response.Success = false
			log.Printf("[rpc call fail] on [%s]", reply.ChainId)
		} else {
			log.Printf("[rpc reply]:[%s] [%v] [%s] [%v] => start verifyRPCReply...\n", reply.ChainId, reply.Success, reply.Prof.Key, reply.Prof.Result)
			// this if statement is necessary, in case of users`s default, such as chain_name in txs and rpcAddr is different
			// at least one node, even if CensusIP is nil
			if len(request.RPCAdds[reply.ChainId].ListenIP) != 0 {
				success := verifier.VerifyRPCReply(request.RPCAdds[reply.ChainId], reply)
				if success {
					hasSuccVerified++
				}
			}
		}

		if hasReceived == callNum {
			response.Done = true
			close(rpcReplies)
		}

		// when commit and rollback, ignore rpc call fail, that may cause inconsistency, use Max_Try to alleviate

		// start rollback
		if (response.Done && !response.Success) || (response.Done && hasSuccVerified != callNum) {
			log.Printf("[start rollback] [%s] [%v]\n", request.CrossId, &response)
			// rollback
			for chainId, RPCAdd := range request.RPCAdds {
				// only invoke success need to rollback
				if !response.Result[chainId].Success {
					continue
				}

				log.Printf("start rollback on [%s]", chainId)

				if len(RPCAdd.ListenIP) == 0 {
					log.Printf("[%s] [%s] no listen ip", request.CrossId, chainId)
					continue
				}
				go dispatchRollback(request, RPCAdd.ListenIP, chainId)
			}
		}

		// TODO commit consensus
		// start commit
		if response.Done && response.Success && hasSuccVerified == callNum {
			log.Printf("start [commit] for [%s] [%#v]\n", request.CrossId, &response)
			// commit
			for i := 0; i < maxTry; i++ {
				err := chainPlugin.Commit(&response)
				if err != nil {
					log.Printf("[%s] fail commit [%s]", response.CrossId, err)
					continue
				} else {
					break
				}

			}
		}
	}
	log.Printf("==========[%s] collect [%d] rpc replies done========", request.CrossId, callNum)
	return &response, nil
}
func dispatchRequest(respChan chan *pb.RPCReply, request *pb.Request, ip string, chainId string) {
	// fail call rpc, give a failure reply
	failReply := &pb.RPCReply{
		ChainId: chainId,
		Success: false,
		Prof: &pb.Prof{
			Key:    "",
			Result: "",
		},
	}

	//获取rpc服务的客户端
	client, err := rpcClient(ip)
	if err != nil {
		log.Printf("[%s] fail new rpc client [%s]", request.CrossId, err)
		respChan <- failReply
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	rpcReply, err := client.Invoke(ctx, request)
	if err != nil {
		log.Printf("[%s] [%s] fail disaptch request [%s]", request.CrossId, chainId, err)
		respChan <- failReply
		return
	}

	respChan <- rpcReply
}

func dispatchRollback(request *pb.Request, ip string, chainId string) {
	for i := 0; i < maxTry; i++ {
		client, err := rpcClient(ip)
		if err != nil {
			log.Printf("[%s] fail new rpc client [%s]", request.CrossId, err)
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeOut)
		defer cancel()
		rpcReply, err := client.Rollback(ctx, request)
		if err != nil {
			log.Printf("[%s] [%s] fail dispatch rollback [%s]", request.CrossId, chainId, err)
			continue
		}
		log.Printf("[%s] dispatch rollback to [%s] [%s] success=> [%v]", request.CrossId, chainId, ip, &rpcReply)
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
