package utils

import (
	"fmt"
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

func CheckBaseCost(baseCost, value *big.Int) error {
	if baseCost.Cmp(value) > 0 {
		return fmt.Errorf(
			"the base cost of performing the priority operation is higher than the provided value parameter for the transaction: baseCost: %d, provided value: %d", baseCost, value)
	}
	return nil
}

/*
export async function checkBaseCost(
    baseCost: ethers.BigNumber,
    value: ethers.BigNumberish | Promise<ethers.BigNumberish>
) {
    if (baseCost.gt(await value)) {
        throw new Error(
            `The base cost of performing the priority operation is higher than the provided value parameter ` +
                `for the transaction: baseCost: ${baseCost}, provided value: ${value}`
        );
    }
}
*/
