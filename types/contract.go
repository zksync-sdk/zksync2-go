package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l1sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2sharedbridge"
)

// BridgeContracts represents the addresses of default bridge contracts for both L1 and L2.
type BridgeContracts struct {
	L1Erc20DefaultBridge common.Address `json:"l1Erc20DefaultBridge"`  // Default L1Bridge contract address.
	L2Erc20DefaultBridge common.Address `json:"l2Erc20DefaultBridge"`  // Default L2Bridge contract address.
	L1WethBridge         common.Address `json:"l1WethBridge"`          // WETH L1Bridge contract address.
	L2WethBridge         common.Address `json:"l2WethBridge"`          // WETH L2Bridge contract address.
	L1SharedBridge       common.Address `json:"l1SharedDefaultBridge"` // Default L1SharedBridge contract address.
	L2SharedBridge       common.Address `json:"l2SharedDefaultBridge"` // Default L2SharedBridge contract address.
}

// L1BridgeContracts represents the L1 bridge contracts.
type L1BridgeContracts struct {
	Erc20  *l1bridge.IL1Bridge             // Default L1Bridge contract.
	Shared *l1sharedbridge.IL1SharedBridge // L1SharedBridge contract.
}

// L2BridgeContracts represents the L2 bridge contracts.
type L2BridgeContracts struct {
	Erc20  *l2bridge.IL2Bridge             // Default L2Bridge contract.
	Shared *l2sharedbridge.IL2SharedBridge // Shared L2Bridge contract.
}

// AccountAbstractionVersion represents an enumeration of account abstraction versions.
type AccountAbstractionVersion uint8

const (
	None     AccountAbstractionVersion = iota // Used for contracts that are not accounts.
	Version1                                  // Used for contracts that are accounts.
)

// AccountNonceOrdering represents an enumeration of account nonce ordering formats.
type AccountNonceOrdering uint8

const (
	// Sequential nonces should be ordered in the same way as in externally owned accounts (EOAs).
	// This means, for instance, that the operator will always wait for a transaction with nonce `X`
	// before processing a transaction with nonce `X+1`.
	Sequential AccountNonceOrdering = iota
	Arbitrary                       // Nonces can be ordered in arbitrary order.
)

// ContractAccountInfo represent contract account information containing details on the supported account
// abstraction version and nonce ordering format.
type ContractAccountInfo struct {
	SupportedAAVersion AccountAbstractionVersion // The supported account abstraction version.
	NonceOrdering      AccountNonceOrdering      // The nonce ordering format.
}
