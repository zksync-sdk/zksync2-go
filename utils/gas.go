package utils

import "math/big"

var (
	// RequiredL1ToL2GasPerPubdataLimit It is possible to provide practically any gasPerPubdataByte for L1->L2 transactions,
	// since the cost per gas will be adjusted respectively. Use 800 as a relatively optimal value for now.
	RequiredL1ToL2GasPerPubdataLimit = big.NewInt(800)

	// This gas limit will be used for displaying the error messages when the users do not have enough fee.

	L1RecommendedMinErc20DepositGasLimit = big.NewInt(400_000)
	L1RecommendedMinEthDepositGasLimit   = big.NewInt(200_000)

	// MaxPriorityFeePerGas is fixed because L2 node does not support eth_maxPriorityFeePerGas method
	MaxPriorityFeePerGas = big.NewInt(1_000_000_000)

	// Currently, for some reason the SDK may return slightly smaller L1 gas limit than required for initiating L1->L2
	// transaction. We use a coefficient to ensure that the transaction will be accepted.

	L1FeeEstimationCoefNumerator   = big.NewInt(12)
	L1FeeEstimationCoefDenominator = big.NewInt(10)
)

func ScaleGasLimit(gasLimit *big.Int) *big.Int {
	gasLimit.Mul(L1FeeEstimationCoefNumerator, gasLimit)
	return gasLimit.Div(gasLimit, L1FeeEstimationCoefDenominator)
}
