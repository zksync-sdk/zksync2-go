package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// Fee represents the transaction fee parameters.
type Fee struct {
	GasLimit *hexutil.Big `json:"gas_limit"` // Maximum amount of gas allowed for the transaction.
	// Maximum amount of gas the user is willing to pay for a single byte of pubdata.
	GasPerPubdataLimit   *hexutil.Big `json:"gas_per_pubdata_limit"`
	MaxFeePerGas         *hexutil.Big `json:"max_fee_per_gas"`          // EIP-1559 fee cap per gas.
	MaxPriorityFeePerGas *hexutil.Big `json:"max_priority_fee_per_gas"` // EIP-1559 tip per gas.
}
