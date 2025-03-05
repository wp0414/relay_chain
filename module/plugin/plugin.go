package plugin

import (
	"relay-chain/module/types/pb"
)

const (
	CrossTransactionContractName = "Cross"
	// topics
	CrossStartTopic    = "CrossStart"
	CrossResponseTopic = "CrossResponse"

	// filed
	Status   = "Status"
	Payload  = "Payload"
	Response = "Response"

	// methods
	StartCross             = "StartCross"
	SetChainId             = "SetChainId"
	ShowChainId            = "ShowChainId"
	SetCross               = "SetCross"
	ShowCross              = "ShowCross"
	SetCrossStatusResponse = "SetCrossStatusResponse"

	// params(map key)
	CrossId       = "CrossId"
	PayloadParam  = "Payload"
	ResponseParam = "Response"
	Type          = "Type"
	StatusParam   = "Status"

	// status
	SuccessStatus = "Success"
	ErrorStatus   = "Fail"

	CrossChainProve = "Service.Prove"

	RollbackResponse = "fail cross"
)

type ChainPlugin interface {

	// GetChainId return the id of this chain
	GetChainId() string

	// InitClient initiate the chain client using the specific config path(yaml file)
	InitClient(configYmlPath string) error

	// Execute invoke business contract on this chain
	Execute(request *pb.Request) error

	// Commit  response on this chain to Cross contract
	Commit(response *pb.Response) error

	// Rollback request on this chain to Cross contract
	Rollback(request *pb.Request) error

	// GetTxByTxId return the Prof according to txId on this chain
	GetTxByTxId(txId string) *pb.Prof

	// SubscribeCrossEvent subscribe cross event on this chain
	SubscribeCrossEvent()
}
