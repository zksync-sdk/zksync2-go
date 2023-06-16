package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
)

type BridgeContracts struct {
	L1Erc20DefaultBridge common.Address `json:"l1Erc20DefaultBridge"`
	L2Erc20DefaultBridge common.Address `json:"l2Erc20DefaultBridge"`
}

type L1BridgeContracts struct {
	Erc20 *l1bridge.IL1Bridge
}

type L2BridgeContracts struct {
	Erc20 *l2bridge.IL2Bridge
}
