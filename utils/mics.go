package utils

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

func NewBigZero() *hexutil.Big {
	return (*hexutil.Big)(new(big.Int))
}

func NewBig(n int64) *hexutil.Big {
	return (*hexutil.Big)(big.NewInt(n))
}

func NewCallMsg(call ethereum.CallMsg) *types.CallMsg {
	return &types.CallMsg{
		From:      call.From,
		To:        call.To,
		Gas:       call.Gas,
		GasPrice:  call.GasPrice,
		GasFeeCap: call.GasFeeCap,
		GasTipCap: call.GasTipCap,
		Value:     call.Value,
		Data:      call.Data,
		Meta: &types.Eip712Meta{
			GasPerPubdata: NewBig(DefaultGasPerPubdataLimit.Int64()),
		},
	}
}
