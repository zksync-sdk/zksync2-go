package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/zksync-sdk/zksync2-go"
	"github.com/zksync-sdk/zksync2-go/contracts/IPaymasterFlow"
	"log"
	"strings"
)

var paymasterFlowAbi abi.ABI

func init() {
	var err error
	paymasterFlowAbi, err = abi.JSON(strings.NewReader(IPaymasterFlow.IPaymasterFlowMetaData.ABI))
	if err != nil {
		log.Fatal("failed to load paymasterFlowAbi: %w", err)
	}
}

func GetApprovalBasedPaymasterInput(paymasterInput zksync2.ApprovalBasedPaymasterInput) ([]byte, error) {
	return paymasterFlowAbi.Pack("approvalBased",
		paymasterInput.Token,
		paymasterInput.MinimalAllowance,
		paymasterInput.InnerInput)
}

func GetGeneralPaymasterInput(paymasterInput zksync2.GeneralPaymasterInput) ([]byte, error) {
	return paymasterFlowAbi.Pack("general", paymasterInput)
}
