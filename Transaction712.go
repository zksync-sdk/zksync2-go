package zksync2

import (
	"errors"
	"fmt"
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

type Transaction712 struct {
	// most like to DynamicFeeTx
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        *big.Int
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       hexutil.Bytes
	AccessList types.AccessList
	// Signature values
	V, R, S *big.Int
	// zkSync part
	ChainID *big.Int
	From    *common.Address `rlp:"-"` // ?? empty ??
	// Meta fields   *Eip712Meta
	ErgsPerPubdata  *hexutil.Big
	FactoryDeps     []hexutil.Bytes
	CustomSignature hexutil.Bytes
	PaymasterParams *PaymasterParams
}

func NewTransaction712(chainID *big.Int, nonce *big.Int, gasLimit *big.Int, to common.Address, value *big.Int,
	data []byte, maxPriorityFeePerGas, maxFeePerGas *big.Int, from common.Address, meta *Eip712Meta) *Transaction712 {
	return &Transaction712{
		Nonce:           nonce.Uint64(),
		GasTipCap:       maxPriorityFeePerGas,
		GasFeeCap:       maxFeePerGas,
		Gas:             gasLimit,
		To:              &to,
		Value:           value,
		Data:            data,
		ChainID:         chainID,
		From:            &from,
		ErgsPerPubdata:  meta.ErgsPerPubdata,
		FactoryDeps:     meta.FactoryDeps,
		CustomSignature: meta.CustomSignature,
		PaymasterParams: meta.PaymasterParams,
	}
}

func (tx *Transaction712) RLPValues(sig []byte) ([]byte, error) {
	if len(sig) == 65 {
		// set signature's V, R, S values
		tx.V = big.NewInt(0).SetBytes(sig[64:])
		tx.R = big.NewInt(0).SetBytes(sig[0:32])
		tx.S = big.NewInt(0).SetBytes(sig[32:64])
	} else if len(sig) > 0 {
		return nil, errors.New("invalid length of signature")
	}
	res, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to encode RLP bytes: %w", err)
	}
	return append([]byte{0x71}, res...), nil
}

func (tx *Transaction712) GetEIP712Type() string {
	return "Transaction"
}

func (tx *Transaction712) GetEIP712Types() []apitypes.Type {
	return []apitypes.Type{
		{Name: "txType", Type: "uint256"},
		{Name: "from", Type: "uint256"},
		{Name: "to", Type: "uint256"},
		{Name: "ergsLimit", Type: "uint256"},
		{Name: "ergsPerPubdataByteLimit", Type: "uint256"},
		{Name: "maxFeePerErg", Type: "uint256"},
		{Name: "maxPriorityFeePerErg", Type: "uint256"},
		{Name: "paymaster", Type: "uint256"},
		{Name: "nonce", Type: "uint256"},
		{Name: "value", Type: "uint256"},
		{Name: "data", Type: "bytes"},
		{Name: "factoryDeps", Type: "bytes32[]"},
		{Name: "paymasterInput", Type: "bytes"},
	}
}

func (tx *Transaction712) GetEIP712Message() apitypes.TypedDataMessage {
	paymaster := big.NewInt(0)
	paymasterInput := hexutil.Bytes{}
	if tx.PaymasterParams != nil {
		paymaster = big.NewInt(0).SetBytes(tx.PaymasterParams.Paymaster.Bytes())
		paymasterInput = tx.PaymasterParams.PaymasterInput
	}
	return apitypes.TypedDataMessage{
		"txType":                  EIP712TxType,
		"from":                    big.NewInt(0).SetBytes(tx.From.Bytes()).String(),
		"to":                      big.NewInt(0).SetBytes(tx.To.Bytes()).String(),
		"ergsLimit":               tx.Gas.String(),
		"ergsPerPubdataByteLimit": tx.ErgsPerPubdata.String(),
		"maxFeePerErg":            tx.GasFeeCap.String(),
		"maxPriorityFeePerErg":    tx.GasTipCap.String(),
		"paymaster":               paymaster.String(),
		"nonce":                   big.NewInt(0).SetUint64(tx.Nonce).String(),
		"value":                   tx.Value.String(),
		"data":                    tx.Data,
		"factoryDeps":             tx.getFactoryDepsHashes(),
		"paymasterInput":          paymasterInput,
	}
}

func (tx *Transaction712) getFactoryDepsHashes() []interface{} {
	if len(tx.FactoryDeps) == 0 {
		return []interface{}{}
	}
	res := make([]interface{}, len(tx.FactoryDeps))
	for i, d := range tx.FactoryDeps {
		h, err := HashBytecode(d)
		if err != nil {
			panic(fmt.Sprintf("failed to get hash of some bytecode in FactoryDeps"))
		}
		res[i] = h
	}
	return res
}
