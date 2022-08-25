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
	Eip712Meta *Eip712Meta
}

func CreateFunctionCallTransaction(from, to common.Address, ergsLimit, ergsPrice, value *big.Int, data hexutil.Bytes) *Transaction {
	return &Transaction{
		From:     from,
		To:       to,
		Gas:      (*hexutil.Big)(ergsLimit),
		GasPrice: (*hexutil.Big)(ergsPrice),
		Value:    (*hexutil.Big)(value),
		Data:     data,
		Eip712Meta: &Eip712Meta{
			ErgsPerPubdata:  NewBigZero(),
			FactoryDeps:     nil,
			PaymasterParams: nil,
		},
	}
}

func Create2ContractTransaction(from common.Address, ergsLimit, ergsPrice *big.Int, bytecode, calldata hexutil.Bytes) *Transaction {
	return &Transaction{
		From:     from,
		To:       ContractDeployerAddress,
		Gas:      (*hexutil.Big)(ergsLimit),
		GasPrice: (*hexutil.Big)(ergsPrice),
		Value:    nil,
		Data:     calldata,
		Eip712Meta: &Eip712Meta{
			ErgsPerPubdata:  NewBigZero(),
			FactoryDeps:     []hexutil.Bytes{bytecode},
			PaymasterParams: nil,
		},
	}
}
