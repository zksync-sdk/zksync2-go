package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
)

// BridgeContracts represents the addresses of default bridge contracts for both L1 and L2.
type BridgeContracts struct {
	L1Erc20DefaultBridge common.Address `json:"l1Erc20DefaultBridge"` // Default L1Bridge contract address.
	L2Erc20DefaultBridge common.Address `json:"l2Erc20DefaultBridge"` // Default L2Bridge contract address.
	L1WethBridge         common.Address `json:"l1WethBridge"`         //  WETH L1Bridge contract address
	L2WethBridge         common.Address `json:"l2WethBridge"`         //  WETH L2Bridge contract address
}

// L1BridgeContracts represents the L1 bridge contracts.
type L1BridgeContracts struct {
	Erc20 *l1bridge.IL1Bridge // Default L1Bridge contract.
}

// L2BridgeContracts represents the L2 bridge contracts.
type L2BridgeContracts struct {
	Erc20 *l2bridge.IL2Bridge // Default L2Bridge contract.
}

// AccountAbstractionVersion represents an enumeration of account abstraction versions.
type AccountAbstractionVersion uint8

const (
	None AccountAbstractionVersion = iota
	Version1
)

// AccountNonceOrdering represents an enumeration of account nonce ordering formats.
type AccountNonceOrdering uint8

const (
	Sequential AccountNonceOrdering = iota
	Arbitrary
)

type ContractAccountInfo struct {
	SupportedAAVersion AccountAbstractionVersion
	NonceOrdering      AccountNonceOrdering
}
