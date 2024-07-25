package utils

import (
	"fmt"
	"math/big"
)

var (
	// RequiredL1ToL2GasPerPubdataLimit It is possible to provide practically any gasPerPubdataByte for L1->L2 transactions,
	// since the cost per gas will be adjusted respectively. Use 800 as a relatively optimal value for now.
	RequiredL1ToL2GasPerPubdataLimit = big.NewInt(800)
	// DefaultGasPerPubdataLimit The large L2 gas per pubdata to sign. This gas is enough to ensure that
	// any reasonable limit will be accepted. Note, that the operator is NOT required to
	// use the honest value of gas per pubdata, and it can use any value up to the one signed by the user.
	// In the future releases, we will provide a way to estimate the current gasPerPubdata.
	DefaultGasPerPubdataLimit = big.NewInt(50_000)

	// MaxPriorityFeePerGas is fixed because L2 node does not support eth_maxPriorityFeePerGas method
	MaxPriorityFeePerGas = big.NewInt(1_000_000_000)

	// L1RecommendedMinErc20DepositGasLimit This gas limit will be used for displaying the error messages when the users do not have enough fee.
	L1RecommendedMinErc20DepositGasLimit = big.NewInt(400_000)
	// L1RecommendedMinEthDepositGasLimit This gas limit will be used for displaying the error messages when the users do not have enough fee.
	L1RecommendedMinEthDepositGasLimit = big.NewInt(200_000)
)

// ScaleGasLimit scales the provided gas limit using a coefficient to ensure acceptance of L1->L2 transactions.
func ScaleGasLimit(gasLimit *big.Int) *big.Int {
	// Currently, for some reason the SDK may return slightly smaller L1 gas limit than required for initiating L1->L2
	// transaction. We use a coefficient to ensure that the transaction will be accepted.
	L1FeeEstimationCoefNumerator := big.NewInt(12)
	L1FeeEstimationCoefDenominator := big.NewInt(10)
	gasLimit.Mul(L1FeeEstimationCoefNumerator, gasLimit)
	return gasLimit.Div(gasLimit, L1FeeEstimationCoefDenominator)
}

// CheckBaseCost checks if the provided base cost is greater than the provided value.
// If it is, return an error indicating that there are not enough funds.
func CheckBaseCost(baseCost, value *big.Int) error {
	if baseCost.Cmp(value) > 0 {
		return fmt.Errorf(
			"the base cost of performing the priority operation is higher than the provided value parameter for the transaction: baseCost: %d, provided value: %d", baseCost, value)
	}
	return nil
}
