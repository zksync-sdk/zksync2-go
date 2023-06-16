package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PaymasterInput interface {
	GetType() string
	GetInput() []byte
}

type ApprovalBasedPaymasterInput struct {
	Token            common.Address
	MinimalAllowance *big.Int
	InnerInput       []byte
}

func (a *ApprovalBasedPaymasterInput) GetType() string {
	return "ApprovalBased"
}

func (a *ApprovalBasedPaymasterInput) GetInput() []byte {
	return a.InnerInput
}

type GeneralPaymasterInput []byte

func (g *GeneralPaymasterInput) GetType() string {
	return "General"
}

func (g *GeneralPaymasterInput) GetInput() []byte {
	return *g
}
