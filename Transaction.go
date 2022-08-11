package zksync2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

const (
	EIP712TxType = `0x71`
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

func CreateFunctionCallTransaction(from, to common.Address, ergsLimit, ergsPrice, value *big.Int, feeToken common.Address, data hexutil.Bytes) *Transaction {
	return &Transaction{
		From:     from,
		To:       to,
		Gas:      (*hexutil.Big)(ergsLimit),
		GasPrice: (*hexutil.Big)(ergsPrice),
		Value:    (*hexutil.Big)(value),
		Data:     data,
		Eip712Meta: &Eip712Meta{
			FeeToken:       feeToken,
			ErgsPerPubdata: NewBigZero(),
			ErgsPerStorage: NewBigZero(),
			FactoryDeps:    nil,
			AAParams:       nil,
		},
	}
}

type Transaction712 struct {
	// Legacy Tx part
	types.LegacyTx
	// zkSync part
	ChainID *big.Int
	Meta    *Eip712Meta
}

func NewTransaction712(nonce *big.Int, to common.Address, value *big.Int, gasLimit, gasPrice *big.Int, data []byte,
	chainID *big.Int, meta *Eip712Meta) *Transaction712 {
	return &Transaction712{
		LegacyTx: types.LegacyTx{
			Nonce:    nonce.Uint64(),
			GasPrice: gasPrice,
			Gas:      gasLimit.Uint64(),
			To:       &to,
			Value:    value,
			Data:     data,
		},
		ChainID: chainID,
		Meta:    meta,
	}
}

func (tx *Transaction712) RLPValues() ([]byte, error) {
	// TODO first, set signature's V, R, S

	return rlp.EncodeToBytes(tx)
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
	r.To = *tx.To
	r.Nonce = big.NewInt(int64(tx.Nonce))
	r.Value = tx.Value
	r.Data = tx.Data
	r.ErgsLimit = big.NewInt(int64(tx.Gas))
	r.FeeToken = tx.Meta.FeeToken
	r.ErgsPerPubdata = tx.Meta.ErgsPerPubdata.ToInt()
	r.ErgsPrice = tx.GasPrice
	return r
}

func (r *TransactionRequest) GetEIP712Type() string {
	return "TransactionRequest"
}

func (r *TransactionRequest) GetEIP712Types() []apitypes.Type {
	return []apitypes.Type{
		{Name: "txType", Type: "uint8"},
		{Name: "to", Type: "address"},
		{Name: "value", Type: "uint256"},
		{Name: "data", Type: "bytes"},
		{Name: "feeToken", Type: "address"},
		{Name: "ergsLimit", Type: "uint256"},
		{Name: "ergsPerPubdataByteLimit", Type: "uint256"},
		{Name: "ergsPrice", Type: "uint256"},
		{Name: "nonce", Type: "uint256"},
	}
}

func (r *TransactionRequest) GetEIP712Message() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{
		"txType":                  EIP712TxType,
		"to":                      r.To.String(),
		"value":                   r.Value.String(),
		"data":                    r.Data,
		"feeToken":                r.FeeToken.String(),
		"ergsLimit":               r.ErgsLimit.String(),
		"ergsPerPubdataByteLimit": r.ErgsPerPubdata.String(),
		"ergsPrice":               r.ErgsPrice.String(),
		"nonce":                   r.Nonce.String(),
	}
}
