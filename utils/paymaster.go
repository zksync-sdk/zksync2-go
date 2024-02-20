package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/contracts/paymasterflow"
	"github.com/zksync-sdk/zksync2-go/types"
	"log"
	"strings"
)

var paymasterFlowAbi abi.ABI

func init() {
	var err error
	paymasterFlowAbi, err = abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	if err != nil {
		log.Fatal("failed to load paymasterFlowAbi: %w", err)
	}
}

// GetApprovalBasedPaymasterInput returns encoded input for an approval-based paymaster.
func GetApprovalBasedPaymasterInput(paymasterInput types.ApprovalBasedPaymasterInput) ([]byte, error) {
	return paymasterFlowAbi.Pack("approvalBased",
		paymasterInput.Token,
		paymasterInput.MinimalAllowance,
		paymasterInput.InnerInput)
}

// GetGeneralPaymasterInput returns encoded input for a general-based paymaster.
func GetGeneralPaymasterInput(paymasterInput types.GeneralPaymasterInput) ([]byte, error) {
	return paymasterFlowAbi.Pack("general", paymasterInput.GetInput())
}

// GetPaymasterParams returns a correctly-formed paymaster parameters for common paymaster flows.
func GetPaymasterParams(paymasterAddress common.Address, paymasterInput types.PaymasterInput) (*types.PaymasterParams, error) {
	if paymasterInput.GetType() == "General" {
		generalPaymasterInput, ok := paymasterInput.(*types.GeneralPaymasterInput)
		if !ok {
			return &types.PaymasterParams{}, errors.New("cannot convert PaymasterInput to GeneralPaymasterInput type")
		}
		input, err := GetGeneralPaymasterInput(*generalPaymasterInput)
		if err != nil {
			return &types.PaymasterParams{}, err
		}
		return &types.PaymasterParams{
			Paymaster:      paymasterAddress,
			PaymasterInput: input,
		}, nil
	} else if paymasterInput.GetType() == "ApprovalBased" {
		approvalBasedPaymasterInput, ok := paymasterInput.(*types.ApprovalBasedPaymasterInput)
		if !ok {
			return &types.PaymasterParams{}, errors.New("cannot convert PaymasterInput to ApprovalBasedPaymasterInput type")
		}
		input, err := GetApprovalBasedPaymasterInput(*approvalBasedPaymasterInput)
		if err != nil {
			return &types.PaymasterParams{}, err
		}
		return &types.PaymasterParams{
			Paymaster:      paymasterAddress,
			PaymasterInput: input,
		}, nil
	} else {
		return &types.PaymasterParams{}, fmt.Errorf("cannot recognize given paymaster input type: %s", paymasterInput.GetType())
	}
}
