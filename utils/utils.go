package utils

import (
	"encoding/json"
	"fmt"
)

const (
	// SpvTxVerifyMethod spv证明方法
	SpvTxVerifyMethod = "verify_tx"

	// SyncBlockHeaderMethod 同步区块头方法
	SyncBlockHeaderMethod = "sync_block_header"

	// GetBlockHeaderMethod 获取区块头方法
	GetBlockHeaderMethod = "get_block_header"
)

// GetSpvContractName 获取spv交易证明的参数
func GetSpvContractName(gatewayId, chainId string) string {
	return "spv" + gatewayId + chainId
}

// GetBlockHeaderParam 获取区块头的参数
func GetBlockHeaderParam(blockHeight int64) string {
	blockHeightByte := []byte(fmt.Sprintf("%d", blockHeight))
	param := make(map[string][]byte)
	param["block_height"] = blockHeightByte
	paramByte, _ := json.Marshal(param)
	return string(paramByte)
}

// GetSyncBlockHeaderParameter 获取同步区块头的参数
func GetSyncBlockHeaderParameter(blockHeight uint64, blockHeader []byte) string {
	res := make(map[string][]byte)
	res["block_height"] = []byte(fmt.Sprintf("%d", blockHeight))
	res["block_header"] = blockHeader
	resJson, _ := json.Marshal(res)
	return string(resJson)
}
