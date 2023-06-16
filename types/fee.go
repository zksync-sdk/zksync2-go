package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type Fee struct {
	GasLimit             *hexutil.Big `json:"gas_limit"`
	GasPerPubdataLimit   *hexutil.Big `json:"gas_per_pubdata_limit"`
	MaxFeePerGas         *hexutil.Big `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas *hexutil.Big `json:"max_priority_fee_per_gas"`
}
