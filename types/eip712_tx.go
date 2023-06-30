package types

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

// EIP712TxType represents an EIP-712 transaction type.
const EIP712TxType = `0x71`

// Transaction712 represents an EIP-712 compliant transaction.
// It shares similarities with regular transactions but also includes zkSync-specific features such as account
// abstraction and paymasters.
// Smart contracts must be deployed with support for the EIP-712 transaction type.
type Transaction712 struct {
	Nonce      *big.Int         // Nonce to use for the transaction execution.
	GasTipCap  *big.Int         // EIP-1559 tip per gas.
	GasFeeCap  *big.Int         // EIP-1559 fee cap per gas.
	Gas        *big.Int         // Gas limit to set for the transaction execution.
	To         *common.Address  // The address of the recipient.
	Value      *big.Int         // Funds to transfer along the transaction (nil = 0 = no funds).
	Data       hexutil.Bytes    // Input data, usually an ABI-encoded contract method invocation.
	AccessList types.AccessList // EIP-2930 access list.

	ChainID *big.Int        // Chain ID of the network.
	From    *common.Address // The address of the sender.
	Meta    *Eip712Meta     // EIP-712 metadata.
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
		// Meta fields   *Meta
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

func (tx *Transaction712) EIP712Type() string {
	return "Transaction"
}

func (tx *Transaction712) EIP712Types() []apitypes.Type {
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

func (tx *Transaction712) EIP712Message() (apitypes.TypedDataMessage, error) {
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
	factoryDepsHashes, err := tx.getFactoryDepsHashes()
	if err != nil {
		return nil, err
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
		"factoryDeps":            factoryDepsHashes,
		"paymasterInput":         paymasterInput,
	}, nil
}

func (tx *Transaction712) getFactoryDepsHashes() ([]interface{}, error) {
	if tx.Meta == nil || len(tx.Meta.FactoryDeps) == 0 {
		return []interface{}{}, nil
	}
	res := make([]interface{}, len(tx.Meta.FactoryDeps))
	for i, d := range tx.Meta.FactoryDeps {
		h, err := hashBytecode(d)
		if err != nil {
			return nil, fmt.Errorf("failed to get hash of some bytecode in FactoryDeps")
		}
		res[i] = h
	}
	return res, nil
}

// Eip712Meta L2-specific transaction metadata.
type Eip712Meta struct {
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *hexutil.Big `json:"gasPerPubdata,omitempty"`
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes `json:"customSignature,omitempty"`
	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
	// it should contain the bytecode of the contract being deployed.
	// If the contract is a factory contract, i.e. it can deploy other contracts,
	// the array should also contain the bytecodes of the contracts which it can deploy.
	FactoryDeps []hexutil.Bytes `json:"factoryDeps"`
	// PaymasterParams contains parameters for configuring the custom paymaster
	// for the transaction.
	PaymasterParams *PaymasterParams `json:"paymasterParams,omitempty"`
}

func (m *Eip712Meta) MarshalJSON() ([]byte, error) {
	type Alias Eip712Meta
	fdb := make([][]uint, len(m.FactoryDeps))
	for i, v := range m.FactoryDeps {
		fdb[i] = make([]uint, len(v))
		for j, b := range v {
			fdb[i][j] = uint(b)
		}
	}
	return json.Marshal(&struct {
		FactoryDeps [][]uint `json:"factoryDeps"`
		*Alias
	}{
		FactoryDeps: fdb,
		Alias:       (*Alias)(m),
	})
}

// PaymasterParams contains parameters for configuring the custom paymaster for the transaction.
type PaymasterParams struct {
	Paymaster      common.Address `json:"paymaster"`      // address of the paymaster
	PaymasterInput []byte         `json:"paymasterInput"` // encoded input
}

func (p *PaymasterParams) MarshalJSON() ([]byte, error) {
	type PaymasterParams struct {
		Paymaster      common.Address `json:"paymaster"`
		PaymasterInput []int          `json:"paymasterInput"`
	}
	var input []int
	for _, b := range p.PaymasterInput {
		input = append(input, int(b))
	}
	params := PaymasterParams{
		Paymaster:      p.Paymaster,
		PaymasterInput: input,
	}

	return json.Marshal(params)
}

func hashBytecode(bytecode []byte) ([]byte, error) {
	if len(bytecode)%32 != 0 {
		return nil, errors.New("bytecode length in bytes must be divisible by 32")
	}
	bytecodeHash := sha256.Sum256(bytecode)
	// get real length of bytecode, which is presented as 32-byte words
	length := big.NewInt(int64(len(bytecode) / 32))
	if length.BitLen() > 16 {
		return nil, errors.New("bytecode length must be less than 2^16 bytes")
	}
	// replace first 2 bytes of hash with version
	version := []byte{1, 0}
	copy(bytecodeHash[0:2], version)
	// replace second 2 bytes of hash with bytecode length
	length2b := make([]byte, 2)
	length2b = length.FillBytes(length2b) // 0-padded in 2 bytes
	copy(bytecodeHash[2:4], length2b)
	return bytecodeHash[:], nil
}
