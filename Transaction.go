package zksync2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

const (
	EIP712TxType = byte(0x71)
)

type Transaction struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      *BigInt        `json:"gas"`
	GasPrice *BigInt        `json:"gasPrice"`
	Value    *BigInt        `json:"value"`
	Data     hexutil.Bytes  `json:"data"`
	//
	Eip712Meta *Eip712Meta
}

func (tx *Transaction) txType() byte { return EIP712TxType }

func CreateFunctionCallTransaction(from, to common.Address, ergsLimit, ergsPrice, value *big.Int, feeToken common.Address, data hexutil.Bytes) *Transaction {
	return &Transaction{
		From:     from,
		To:       to,
		Gas:      ToBigInt(ergsLimit),
		GasPrice: ToBigInt(ergsPrice),
		Value:    ToBigInt(value),
		Data:     data,
		Eip712Meta: &Eip712Meta{
			FeeToken:       feeToken,
			ErgsPerPubdata: NewBigInt(0),
			ErgsPerStorage: NewBigInt(0),
			FactoryDeps:    nil,
			AAParams:       nil,
		},
	}
}

type Transaction712 struct {
	*types.Transaction
	ChainId *BigInt
	Meta    *Eip712Meta
}

func NewTransaction712(nonce *big.Int, to common.Address, value *big.Int, gasLimit, gasPrice *big.Int, data []byte,
	chainId *BigInt, meta *Eip712Meta) *Transaction712 {
	ltx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce.Uint64(),
		GasPrice: gasPrice,
		Gas:      gasLimit.Uint64(),
		To:       &to,
		Value:    value,
		Data:     data,
	})
	return &Transaction712{
		Transaction: ltx,
		ChainId:     chainId,
		Meta:        meta,
	}

}

type TransactionRequest struct {
	To             common.Address
	Nonce          *big.Int
	Value          *big.Int
	Data           []byte
	ErgsLimit      *big.Int
	FeeToken       common.Address
	ErgsPerPubdata *big.Int
	ErgsPrice      *big.Int
}

func (r *TransactionRequest) FromTx(tx *Transaction712) *TransactionRequest {
	r.To = *tx.To()
	r.Nonce = big.NewInt(int64(tx.Nonce()))
	r.Value = tx.Value()
	r.Data = tx.Data()
	r.ErgsLimit = big.NewInt(int64(tx.Gas()))
	r.FeeToken = tx.Meta.FeeToken
	r.ErgsPerPubdata = tx.Meta.ErgsPerPubdata.Int
	r.ErgsPrice = tx.GasPrice()
	return r
}
