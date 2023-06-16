package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

func CreateFunctionCallTransaction(from, to common.Address, gasPrice, gasLimit, value *big.Int, data hexutil.Bytes,
	customSignature hexutil.Bytes, paymasterParams *types.PaymasterParams) *types.Transaction {
	return &types.Transaction{
		From:     from,
		To:       to,
		Gas:      (*hexutil.Big)(gasLimit),
		GasPrice: (*hexutil.Big)(gasPrice),
		Value:    (*hexutil.Big)(value),
		Data:     data,
		Eip712Meta: &types.Eip712Meta{
			GasPerPubdata:   NewBig(160_000),
			CustomSignature: customSignature,
			PaymasterParams: paymasterParams,
		},
	}
}

func Create2ContractTransaction(from common.Address, gasPrice, gasLimit *big.Int,
	bytecode, calldata hexutil.Bytes, deps []hexutil.Bytes,
	customSignature hexutil.Bytes, paymasterParams *types.PaymasterParams) *types.Transaction {
	if len(deps) > 0 {
		deps = append(deps, bytecode)
	} else {
		deps = []hexutil.Bytes{bytecode}
	}
	return &types.Transaction{
		From:     from,
		To:       ContractDeployerAddress,
		Gas:      (*hexutil.Big)(gasLimit),
		GasPrice: (*hexutil.Big)(gasPrice),
		Value:    nil,
		Data:     calldata,
		Eip712Meta: &types.Eip712Meta{
			GasPerPubdata:   NewBig(160_000),
			CustomSignature: customSignature,
			FactoryDeps:     deps,
			PaymasterParams: paymasterParams,
		},
	}
}
