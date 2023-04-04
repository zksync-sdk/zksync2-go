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
	Nonce      *big.Int
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        *big.Int
	To         *common.Address
	Value      *big.Int
	Data       hexutil.Bytes
	AccessList types.AccessList
	// zkSync part
	ChainID *big.Int
	From    *common.Address
	Meta    *Eip712Meta
}

func NewTransaction712(chainID *big.Int, nonce *big.Int, gasLimit *big.Int, to common.Address, value *big.Int,
	data []byte, maxPriorityFeePerGas, maxFeePerGas *big.Int, from common.Address, meta *Eip712Meta) *Transaction712 {
	return &Transaction712{
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
		Gas:       gasLimit,
		To:        &to,
		Value:     value,
		Data:      data,
		ChainID:   chainID,
		From:      &from,
		Meta:      meta,
	}
}

func (tx *Transaction712) RLPValues(sig []byte) ([]byte, error) {
	// use custom struct to get right RLP sequence and types to use default rlp encoder
	txRLP := struct {
		Nonce                uint64
		MaxPriorityFeePerGas *big.Int
		MaxFeePerGas         *big.Int
		GasLimit             *big.Int
		To                   *common.Address `rlp:"nil"` // nil means contract creation
		Value                *big.Int
		Data                 hexutil.Bytes
		// zkSync part
		ChainID1 *big.Int // legacy
		Empty1   string   // legacy
		Empty2   string   // legacy
		ChainID2 *big.Int
		From     *common.Address
		// Meta fields   *Eip712Meta
		GasPerPubdata   *big.Int
		FactoryDeps     []hexutil.Bytes
		CustomSignature hexutil.Bytes
		PaymasterParams *PaymasterParams
	}{
		Nonce:                tx.Nonce.Uint64(),
		MaxPriorityFeePerGas: tx.GasTipCap,
		MaxFeePerGas:         tx.GasFeeCap,
		GasLimit:             tx.Gas,
		To:                   tx.To,
		Value:                tx.Value,
		Data:                 tx.Data,
		ChainID1:             tx.ChainID,
		ChainID2:             tx.ChainID,
		From:                 tx.From,
		GasPerPubdata:        tx.Meta.GasPerPubdata.ToInt(),
		FactoryDeps:          tx.Meta.FactoryDeps,
		CustomSignature:      tx.Meta.CustomSignature,
		PaymasterParams:      tx.Meta.PaymasterParams,
	}
	if len(txRLP.CustomSignature) == 0 {
		if len(sig) == 65 {
			txRLP.CustomSignature = sig
		} else if len(sig) > 0 {
			return nil, errors.New("invalid length of signature")
		}
	}

	res, err := rlp.EncodeToBytes(txRLP)
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
		{Name: "gasLimit", Type: "uint256"},
		{Name: "gasPerPubdataByteLimit", Type: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256"},
		{Name: "maxPriorityFeePerGas", Type: "uint256"},
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
	if tx.Meta != nil && tx.Meta.PaymasterParams != nil {
		paymaster = big.NewInt(0).SetBytes(tx.Meta.PaymasterParams.Paymaster.Bytes())
		paymasterInput = tx.Meta.PaymasterParams.PaymasterInput
	}
	value := `0x0`
	if tx.Value != nil {
		value = tx.Value.String()
	}
	return apitypes.TypedDataMessage{
		"txType":                 EIP712TxType,
		"from":                   big.NewInt(0).SetBytes(tx.From.Bytes()).String(),
		"to":                     big.NewInt(0).SetBytes(tx.To.Bytes()).String(),
		"gasLimit":               tx.Gas.String(),
		"gasPerPubdataByteLimit": tx.Meta.GasPerPubdata.String(),
		"maxFeePerGas":           tx.GasFeeCap.String(),
		"maxPriorityFeePerGas":   tx.GasTipCap.String(),
		"paymaster":              paymaster.String(),
		"nonce":                  tx.Nonce.String(),
		"value":                  value,
		"data":                   tx.Data,
		"factoryDeps":            tx.getFactoryDepsHashes(),
		"paymasterInput":         paymasterInput,
	}
}

func (tx *Transaction712) getFactoryDepsHashes() []interface{} {
	if tx.Meta == nil || len(tx.Meta.FactoryDeps) == 0 {
		return []interface{}{}
	}
	res := make([]interface{}, len(tx.Meta.FactoryDeps))
	for i, d := range tx.Meta.FactoryDeps {
		h, err := HashBytecode(d)
		if err != nil {
			panic(fmt.Sprintf("failed to get hash of some bytecode in FactoryDeps"))
		}
		res[i] = h
	}
	return res
}
