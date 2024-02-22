package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
	"testing"
)

func TestGetPaymasterParamsGeneral(t *testing.T) {
	expected := &types.PaymasterParams{
		Paymaster:      common.HexToAddress("0x0a67078A35745947A37A552174aFe724D8180c25"),
		PaymasterInput: common.Hex2Bytes("8c5a344500000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"),
	}

	params, err := GetPaymasterParams(
		common.HexToAddress("0x0a67078A35745947A37A552174aFe724D8180c25"),
		&types.GeneralPaymasterInput{},
	)

	assert.NoError(t, err, "GetPaymasterParams should not return error")
	assert.Equal(t, expected, params, "Parameter should be the same")
}

func TestGetPaymasterParamsApprovalBased(t *testing.T) {
	expected := &types.PaymasterParams{
		Paymaster:      common.HexToAddress("0x0a67078A35745947A37A552174aFe724D8180c25"),
		PaymasterInput: common.Hex2Bytes("949431dc00000000000000000000000065c899b5fb8eb9ae4da51d67e1fc417c7cb7e964000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
	}

	params, err := GetPaymasterParams(
		common.HexToAddress("0x0a67078A35745947A37A552174aFe724D8180c25"),
		&types.ApprovalBasedPaymasterInput{
			Token:            common.HexToAddress("0x65C899B5fb8Eb9ae4da51D67E1fc417c7CB7e964"),
			MinimalAllowance: big.NewInt(1),
			InnerInput:       []byte{},
		},
	)

	assert.NoError(t, err, "GetPaymasterParams should not return error")
	assert.Equal(t, expected, params, "Parameter should be the same")
}
