package verifier

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"relay-chain/module/types/pb"
	"time"
)

const timeOut = time.Second * 10

type Verifier struct {
}

// VerifyRequest request from other chain is valid or not
func (v *Verifier) VerifyRequest(request *pb.Request) error {
	var originProf *pb.Prof
	var originChainId string

	i := 0
	// find the only one prof, which comes from the first caller
	for chainId, prof := range request.Profs {
		i++
		originProf = prof
		originChainId = chainId
	}

	if i > 1 {
		return errors.New("verify request failed :expect only one prof, but got more than one")
	}
	if len(originProf.Result) == 0 {
		return errors.New("verify request failed :prof is nil")
	}
	if len(request.RPCAdds[originChainId].ListenIP) == 0 {
		return errors.New("verify request failed :expect at least one origin Listen IP, but got zero")
	}

	log.Printf("[start verify request] [%s] orgin, listen ip: [%s], consensus ip: [%s]", request.CrossId, request.RPCAdds[originChainId].ListenIP, request.RPCAdds[originChainId].ConsensusIPs)

	rpcReply := &pb.RPCReply{
		ChainId: originChainId,
		Success: true,
		Prof:    originProf,
	}

	success := v.VerifyRPCReply(request.RPCAdds[originChainId], rpcReply)

	if !success {
		return errors.New("[verify request failed]")
	}

	log.Printf("[%s] verify orgin request [%s] success\n", request.CrossId, originChainId)
	return nil
}

// VerifyRPCReply judge the reply is legal or not according to majority rule
func (v *Verifier) VerifyRPCReply(rpcAddr *pb.ChainRPCAdd, reply *pb.RPCReply) bool {
	// plus 1 mean Listen as one consensus node
	numCall := len(rpcAddr.ConsensusIPs) + 1
	respChan := make(chan bool, numCall)

	numAdmit := 0

	go getProfFrom(rpcAddr.ListenIP, reply, respChan)

	for _, ip := range rpcAddr.ConsensusIPs {
		go getProfFrom(ip, reply, respChan)
	}

	for resp := range respChan {
		if resp {
			numAdmit++
			if numAdmit >= numCall/2 {
				break
			}
		}
	}

	if numAdmit >= numCall/2 {
		log.Printf("[verifyRPCReply verify success] chainId: [%s] key: [%s], [admit number: %d] >= [1/2*numcall: %d]\n", reply.ChainId, reply.Prof.Key, numAdmit, numCall/2)
		return true
	} else {
		log.Printf("[verifyRPCReply verify failed] chainId: [%s] key: [%s], because [admit number: %d] < [1/2*numcall: %d]\n", reply.ChainId, reply.Prof.Key, numAdmit, numCall/2)
		return false
	}
}

// getProfFrom invoke Service.Prove at `ip`, and judge resp and reply equal or not
func getProfFrom(ip string, reply *pb.RPCReply, respChan chan bool) {
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Println(err)
		respChan <- false
		return
	}

	client := pb.NewServiceClient(conn)
	if err != nil {
		log.Printf("[%s] [%s] fail new rpc client [%s]\n", reply.ChainId, reply.Prof.Key, err)
		respChan <- false
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	resp, err := client.Prove(ctx, reply.Prof)
	if err != nil {
		log.Printf("[%s] [%s] rpc call failed [%s]\n", reply.ChainId, reply.Prof.Key, err)
		respChan <- false
		return
	}

	if resp.Result != reply.Prof.Result {
		log.Printf("[getProfFrom veify failed] chainId:[%s]  ip: [%s]  txId: [%s] , result not equal, expected: [%s], got: [%s]\n", reply.ChainId, ip, reply.Prof.Key, reply.Prof.Result, resp.Result)
		respChan <- false
		return
	}
	// success verify
	respChan <- true
}
