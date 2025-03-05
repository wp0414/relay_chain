package fabric

import (
	"encoding/json"
	"errors"
	"fmt"
	fab_proto "github.com/golang/protobuf/proto"
	fabcommon "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"relay-chain/module/listener"
	"relay-chain/module/plugin"
	"relay-chain/module/server"
	"relay-chain/module/types/pb"
)

var _ plugin.ChainPlugin = &Fabric{}

type Fabric struct {
	ChainId       string
	ChannelClient *channel.Client
	LegerClient   *ledger.Client
	EventClient   *event.Client
}

func (fa *Fabric) GetChainId() string {
	return fa.ChainId
}

func (fa *Fabric) InitClient(configYmlPath string) error {
	fabricSDK, err := fabsdk.New(config.FromFile(configYmlPath))
	if err != nil {
		return err
	}
	channelCtx := fabricSDK.ChannelContext(fa.ChainId, fabsdk.WithUser("User1"))
	fabClient, err := channel.New(channelCtx)
	if err != nil {
		return err
	}
	ledgerClient, err := ledger.New(channelCtx)
	if err != nil {
		return err
	}
	// event.WithBlockEvents() is necessary. If it is nil, we could not get event payload(it is nil).
	eventClient, err := event.New(channelCtx, event.WithBlockEvents())

	fa.ChannelClient = fabClient
	fa.LegerClient = ledgerClient
	fa.EventClient = eventClient

	return nil
}

func (fa *Fabric) InvokeContract(contractName, method string, kvs map[string]string) (channel.Response, error) {
	req := channel.Request{
		ChaincodeID: contractName,
		Fcn:         method,
		Args:        fa.Parse(kvs),
	}
	response, err := fa.ChannelClient.Execute(req)
	if err != nil {
		return channel.Response{}, err
	}
	log.Printf("[invoke] [%s]-[%s] success, response:[%#v]\n", contractName, method, response)
	return response, err
}

func (fa *Fabric) Execute(request *pb.Request) error {
	crossTx, ok := request.Transactions[fa.ChainId]
	if !ok {
		return errors.New("fail find crossTx")
	}
	exeTx := crossTx.ExecuteTransaction
	resp, err := fa.InvokeContract(exeTx.ContractName, exeTx.Method, exeTx.Params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", exeTx.ContractName, exeTx.Method, err))
	}
	if request.Profs == nil {
		request.Profs = make(map[string]*pb.Prof)
	}
	request.Profs[fa.ChainId] = &pb.Prof{Key: string(resp.TransactionID), Result: string(resp.Payload)}
	return nil
}

func (fa *Fabric) Parse(params map[string]string) [][]byte {
	var kvs [][]byte
	for _, value := range params {
		kvs = append(kvs, []byte(value))
	}
	return kvs
}

func (fa *Fabric) Commit(response *pb.Response) error {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = server.DefaultServer.DB.PutResponse(response.CrossId, string(responseBytes))
	if err != nil {
		// one node fail is acceptable.
		log.Printf("put response failed,err: [%s]", err)
	}

	params := map[string]string{
		plugin.CrossId:       response.CrossId,
		plugin.Status:        plugin.SuccessStatus,
		plugin.ResponseParam: string(responseBytes),
	}
	_, err = fa.InvokeContract(plugin.CrossTransactionContractName, plugin.SetCrossStatusResponse, params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", plugin.CrossTransactionContractName, plugin.SetCrossStatusResponse, err))
	}
	return nil
}

func (fa *Fabric) Rollback(request *pb.Request) error {
	err := server.DefaultServer.DB.PutResponse(request.CrossId, plugin.RollbackResponse)
	if err != nil {
		log.Printf("put Rollback response failed, err: [%s]", err)
	}

	crossTx, ok := request.Transactions[fa.ChainId]
	if !ok {
		return errors.New("fail find crossTx")
	}
	rollTx := crossTx.RollbackTransaction
	_, err = fa.InvokeContract(rollTx.ContractName, rollTx.Method, rollTx.Params)
	if err != nil {
		return errors.New(fmt.Sprintf("fail call contract [%s]-[%s]:[%s]", rollTx.ContractName, rollTx.Method, err))
	}
	return nil
}

func (fa *Fabric) GetTxByTxId(txId string) *pb.Prof {
	transaction, err := fa.LegerClient.QueryTransaction(fab.TransactionID(txId))
	if err != nil {
		log.Println(err)
		return nil
	}
	resp, err := fa.payloadToContractResult(transaction.TransactionEnvelope.Payload)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &pb.Prof{
		Key:    txId,
		Result: string(resp.Payload),
	}
}

// payloadToContractResult from blob unmarshal fabric contract result.
// See https://www.cnblogs.com/jiftle/p/15302746.html and https://www.shuzhiduo.com/A/qVdeRmPgzP/
func (fa *Fabric) payloadToContractResult(blob []byte) (*peer.Response, error) {
	var payload fabcommon.Payload
	err := fab_proto.Unmarshal(blob, &payload)
	if err != nil {
		return nil, err
	}

	var tx peer.Transaction
	err = fab_proto.Unmarshal(payload.Data, &tx)
	if err != nil {
		return nil, err
	}

	if len(tx.Actions) < 1 {
		return nil, fmt.Errorf("tx has no action")
	}

	var action peer.ChaincodeActionPayload
	err = fab_proto.Unmarshal(tx.Actions[0].Payload, &action)
	if err != nil {
		return nil, err
	}

	/*	// >>>>> this can fetch invoked method and params in contract
		propPayload := peer.ChaincodeProposalPayload{}

		err = proto.Unmarshal(action.ChaincodeProposalPayload, &propPayload)
		if err != nil {
			return err
		}
		invokeSpec := peer.ChaincodeInvocationSpec{}
		err = proto.Unmarshal(propPayload.Input, &invokeSpec)
		if err != nil {
			return err
		}
		// <<<<< */

	responsePayload := peer.ProposalResponsePayload{}

	err = fab_proto.Unmarshal(action.Action.ProposalResponsePayload, &responsePayload)
	if err != nil {
		return nil, err
	}

	chainCodeAction := peer.ChaincodeAction{}
	err = fab_proto.Unmarshal(responsePayload.Extension, &chainCodeAction)
	if err != nil {
		return nil, err
	}
	return chainCodeAction.Response, nil
}

func (fa *Fabric) SubscribeCrossEvent() {
	registration, events, err := fa.EventClient.RegisterChaincodeEvent(plugin.CrossTransactionContractName, plugin.CrossStartTopic)
	defer fa.EventClient.Unregister(registration)
	if err != nil {
		log.Fatalln("failed to register chaincode event")
	}
	for {
		e, ok := <-events
		if !ok {
			log.Fatalln("listen chan is close!")
		}
		if e == nil {
			log.Println("[listen] [waring] nil event")
			continue
		}

		log.Printf("[event received] [block heigt %d]: %s - [txid: %s] [payload: %s]\n", e.BlockNumber, e.EventName, e.TxID, e.Payload)
		go listener.BeginCrossChain(e.Payload)
	}
}
