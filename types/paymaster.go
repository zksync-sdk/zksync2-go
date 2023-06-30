package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// PaymasterInput is an interface that represents input data for a paymaster.
// Paymasters can implement this interface to provide specific data required
// for payment processing.
type PaymasterInput interface {
	// GetType returns the type of the paymaster input. It should provide a
	// unique identifier for the type of input data.
	GetType() string
	// GetInput returns the actual input data in the form of a byte slice. The
	// data format may vary depending on the type and implementation of the
	// paymaster.
	GetInput() []byte
}

// ApprovalBasedPaymasterInput contains approval-based paymaster input.
// It should be used if the user is required to set a certain allowance to a token for the paymaster to operate.
type ApprovalBasedPaymasterInput struct {
	Token common.Address // ERC20 token used to pay transaction fee.
	// Minimal allowance of Token towards the paymaster from the account that pays the fee with the token.
	MinimalAllowance *big.Int
	InnerInput       []byte // Additional payload that can be sent to the paymaster to implement any logic.
}

func (a *ApprovalBasedPaymasterInput) GetType() string {
	return "ApprovalBased"
}

func (a *ApprovalBasedPaymasterInput) GetInput() []byte {
	return a.InnerInput
}

// GeneralPaymasterInput contains general paymaster input.
// It should be used if no prior actions are required from the user for the paymaster to operate.
type GeneralPaymasterInput []byte

func (g *GeneralPaymasterInput) GetType() string {
	return "General"
}

func (g *GeneralPaymasterInput) GetInput() []byte {
	return *g
}
