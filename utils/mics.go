package utils

import (
	"fmt"
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
		CallMsg: call,
		Meta: &types.Eip712Meta{
			GasPerPubdata: NewBig(DefaultGasPerPubdataLimit.Int64()),
		},
	}
}

// Deprecated: Will be removed in the future releases.
func ToFilterArg(q types.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = q.FromBlock
		}
		arg["toBlock"] = q.ToBlock
	}
	return arg, nil
}
