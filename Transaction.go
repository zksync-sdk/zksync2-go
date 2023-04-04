package zksync2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Transaction struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      *hexutil.Big   `json:"gas"`
	GasPrice *hexutil.Big   `json:"gasPrice"`
	Value    *hexutil.Big   `json:"value"`
	Data     hexutil.Bytes  `json:"data"`
	//
	Eip712Meta *Eip712Meta `json:"eip712Meta"`
}

func CreateFunctionCallTransaction(from, to common.Address, gasPrice, gasLimit, value *big.Int, data hexutil.Bytes,
	customSignature hexutil.Bytes, paymasterParams *PaymasterParams) *Transaction {
	return &Transaction{
		From:     from,
		To:       to,
		Gas:      (*hexutil.Big)(gasLimit),
		GasPrice: (*hexutil.Big)(gasPrice),
		Value:    (*hexutil.Big)(value),
		Data:     data,
		Eip712Meta: &Eip712Meta{
			GasPerPubdata:   NewBig(160000),
			CustomSignature: customSignature,
			FactoryDeps:     nil,
			PaymasterParams: paymasterParams,
		},
	}
}

func Create2ContractTransaction(from common.Address, gasPrice, gasLimit *big.Int,
	bytecode, calldata hexutil.Bytes, deps []hexutil.Bytes,
	customSignature hexutil.Bytes, paymasterParams *PaymasterParams) *Transaction {
	if len(deps) > 0 {
		deps = append(deps, bytecode)
	} else {
		deps = []hexutil.Bytes{bytecode}
	}
	return &Transaction{
		From:     from,
		To:       ContractDeployerAddress,
		Gas:      (*hexutil.Big)(gasLimit),
		GasPrice: (*hexutil.Big)(gasPrice),
		Value:    nil,
		Data:     calldata,
		Eip712Meta: &Eip712Meta{
			GasPerPubdata:   NewBig(160000),
			CustomSignature: customSignature,
			FactoryDeps:     deps,
			PaymasterParams: paymasterParams,
		},
	}
}
