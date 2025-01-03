package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Fee represents the transaction fee parameters.
type Fee struct {
	GasLimit *hexutil.Big `json:"gas_limit"` // Maximum amount of gas allowed for the transaction.
	// Maximum amount of gas the user is willing to pay for a single byte of pubdata.
	GasPerPubdataLimit   *hexutil.Big `json:"gas_per_pubdata_limit"`
	MaxFeePerGas         *hexutil.Big `json:"max_fee_per_gas"`          // EIP-1559 fee cap per gas.
	MaxPriorityFeePerGas *hexutil.Big `json:"max_priority_fee_per_gas"` // EIP-1559 tip per gas.
}

// FeeParams represents the fee parameters configuration.
type FeeParams struct {
	// Fee parameter configuration for the current version of the ZKsync protocol.
	V2 struct {
		// Settings related to transaction fee computation.
		Config struct {
			MinimalL2GasPrice   *big.Int `json:"minimal_l2_gas_price"`  // Minimal gas price on L2.
			ComputeOverheadPart *big.Int `json:"compute_overhead_part"` // Compute overhead part in fee calculation.
			PubdataOverheadPart *big.Int `json:"pubdata_overhead_part"` // Public data overhead part in fee calculation.
			BatchOverheadL1Gas  *big.Int `json:"batch_overhead_l1_gas"` // Overhead in L1 gas for a batch of transactions.
			MaxGasPerBatch      *big.Int `json:"max_gas_per_batch"`     // Maximum gas allowed per batch.
			MaxPubdataPerBatch  *big.Int `json:"max_pubdata_per_batch"` // Maximum amount of public data allowed per batch.
		} `json:"config"`
		// Represents the BaseToken<->ETH conversion ratio.
		ConversionRation struct {
			Denominator *big.Int `json:"denominator"` // Represents the denominator part of the conversion ratio.
			Numerator   *big.Int `json:"numerator"`   // Represents the numerator part of the conversion ratio.
		} `json:"conversion_ratio"`
		L1GasPrice     *big.Int `json:"l1_gas_price"`     // Current L1 gas price.
		L1PubdataPrice *big.Int `json:"l1_pubdata_price"` // Price of storing public data on L1.
	} `json:"V2"`
}
