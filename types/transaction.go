package types

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/zksync-sdk/zksync2-go/eip712"
	"math/big"
)

// EIP712TxType represents an EIP-712 transaction type.
const EIP712TxType = `0x71`

// Transaction provides support for ZKsync Era-specific features
// such as account abstraction and paymasters.
// Smart contracts must be deployed with this transaction type.
type Transaction struct {
	Nonce     *big.Int        `json:"nonce"`     // Nonce to use for the transaction execution.
	GasTipCap *big.Int        `json:"gasTipCap"` // EIP-1559 tip per gas.
	GasFeeCap *big.Int        `json:"gasFeeCap"` // EIP-1559 fee cap per gas.
	Gas       *big.Int        `json:"gas"`       // Gas limit to set for the transaction execution.
	To        *common.Address `json:"to"`        // The address of the recipient.
	Value     *big.Int        `json:"value"`     // Funds to transfer along the transaction (nil = 0 = no funds).
	Data      hexutil.Bytes   `json:"data"`      // Input data, usually an ABI-encoded contract method invocation.

	ChainID *big.Int        `json:"chainID"` // Chain ID of the network.
	From    *common.Address `json:"from"`    // The address of the sender.

	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int `json:"gasPerPubdata"`
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes `json:"customSignature"`
	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
	// it should contain the bytecode of the contract being deployed.
	// If the contract is a factory contract, i.e. it can deploy other contracts,
	// the array should also contain the bytecodes of the contracts which it can deploy.
	FactoryDeps []hexutil.Bytes `json:"factoryDeps"`
	// PaymasterParams contains parameters for configuring the custom paymaster
	// for the transaction.
	PaymasterParams *PaymasterParams `json:"paymasterParams"`
}

func (tx *Transaction) Encode(sig []byte) ([]byte, error) {
	// use custom struct to get right RLP sequence and types to use default rlp encoder
	zkSyncTxRLP := struct {
		Nonce                uint64
		MaxPriorityFeePerGas *big.Int
		MaxFeePerGas         *big.Int
		GasLimit             *big.Int
		To                   *common.Address `rlp:"nil"` // nil means contract creation
		Value                *big.Int
		Data                 hexutil.Bytes
		// zkSync part
		ChainId1 *big.Int // legacy
		Empty1   string   // legacy
		Empty2   string   // legacy
		ChainId2 *big.Int
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
		ChainId1:             tx.ChainID,
		ChainId2:             tx.ChainID,
		From:                 tx.From,
		GasPerPubdata:        tx.GasPerPubdata,
		FactoryDeps:          tx.FactoryDeps,
		CustomSignature:      tx.CustomSignature,
		PaymasterParams:      tx.PaymasterParams,
	}
	if len(zkSyncTxRLP.CustomSignature) == 0 {
		zkSyncTxRLP.CustomSignature = sig
	}

	res, err := rlp.EncodeToBytes(zkSyncTxRLP)
	if err != nil {
		return nil, fmt.Errorf("failed to encode RLP bytes: %w", err)
	}
	return append([]byte{0x71}, res...), nil
}

func (tx *Transaction) Decode(input []byte) error {
	type zkSyncTxRLP struct {
		Nonce                uint64
		MaxPriorityFeePerGas *big.Int
		MaxFeePerGas         *big.Int
		GasLimit             *big.Int
		To                   *common.Address `rlp:"nil"` // nil means contract creation
		Value                *big.Int
		Data                 hexutil.Bytes
		// zkSync part
		ChainId1 *big.Int // legacy
		Empty1   string   // legacy
		Empty2   string   // legacy
		ChainId2 *big.Int
		From     *common.Address
		// Meta fields   *Meta
		GasPerPubdata   *big.Int
		FactoryDeps     []hexutil.Bytes
		CustomSignature hexutil.Bytes
		PaymasterParams *PaymasterParams `rlp:"nil"`
	}
	var decodedTx zkSyncTxRLP
	err := rlp.DecodeBytes(input[1:], &decodedTx)
	if err != nil {
		return err
	}

	tx.Nonce = new(big.Int).SetUint64(decodedTx.Nonce)
	tx.GasTipCap = decodedTx.MaxPriorityFeePerGas
	tx.GasFeeCap = decodedTx.MaxFeePerGas
	tx.Gas = decodedTx.GasLimit
	tx.To = decodedTx.To
	tx.Value = decodedTx.Value
	tx.Data = decodedTx.Data
	tx.ChainID = decodedTx.ChainId2
	tx.From = decodedTx.From
	tx.GasPerPubdata = decodedTx.GasPerPubdata
	tx.CustomSignature = decodedTx.CustomSignature
	tx.FactoryDeps = decodedTx.FactoryDeps
	tx.PaymasterParams = decodedTx.PaymasterParams
	return nil
}

func (tx *Transaction) TypedData() (*apitypes.TypedData, error) {
	domain := eip712.ZkSyncEraEIP712Domain(tx.ChainID.Int64())
	message, err := tx.typedDataMessage()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &apitypes.TypedData{
		Types: apitypes.Types{
			"Transaction":       tx.types(),
			domain.EIP712Type(): domain.EIP712Types(),
		},
		PrimaryType: "Transaction",
		Domain:      domain.EIP712Domain(),
		Message:     message,
	}, nil
}

func (tx *Transaction) Hash() ([]byte, error) {
	typedData, err := tx.TypedData()
	if err != nil {
		return nil, fmt.Errorf("failed to get typed data: %w", err)
	}
	hash, err := tx.hashTypedData(*typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed data: %w", err)
	}
	return hash, nil
}

func (tx *Transaction) Copy() *Transaction {
	if tx == nil {
		return nil
	}

	cpy := &Transaction{
		Nonce:           new(big.Int),
		GasTipCap:       new(big.Int),
		GasFeeCap:       new(big.Int),
		Gas:             new(big.Int),
		Value:           new(big.Int),
		ChainID:         new(big.Int),
		To:              copyAddressPtr(tx.To),
		From:            copyAddressPtr(tx.From),
		Data:            common.CopyBytes(tx.Data),
		CustomSignature: common.CopyBytes(tx.CustomSignature),
		GasPerPubdata:   new(big.Int),
		FactoryDeps:     make([]hexutil.Bytes, len(tx.FactoryDeps)),
	}

	if tx.Nonce != nil {
		cpy.Nonce.Set(tx.Nonce)
	}
	if tx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.GasTipCap)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.Gas != nil {
		cpy.Gas.Set(tx.Gas)
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPerPubdata != nil {
		cpy.GasPerPubdata.Set(tx.GasPerPubdata)
	}

	for i, dep := range tx.FactoryDeps {
		cpy.FactoryDeps[i] = common.CopyBytes(dep)
	}

	if tx.PaymasterParams != nil {
		cpy.PaymasterParams = &PaymasterParams{
			Paymaster:      *copyAddressPtr(&tx.PaymasterParams.Paymaster),
			PaymasterInput: common.CopyBytes(tx.PaymasterParams.PaymasterInput),
		}
	}

	return cpy
}

func (tx *Transaction) hashTypedData(data apitypes.TypedData) ([]byte, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed data domain: %w", err)
	}
	dataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed message: %w", err)
	}
	prefixedData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domain), string(dataHash)))
	prefixedDataHash := crypto.Keccak256(prefixedData)
	return prefixedDataHash, nil
}

func (tx *Transaction) types() []apitypes.Type {
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

func (tx *Transaction) typedDataMessage() (apitypes.TypedDataMessage, error) {
	paymaster := big.NewInt(0)
	paymasterInput := hexutil.Bytes{}
	if tx.PaymasterParams != nil {
		paymaster = big.NewInt(0).SetBytes(tx.PaymasterParams.Paymaster.Bytes())
		paymasterInput = tx.PaymasterParams.PaymasterInput
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
		"gasPerPubdataByteLimit": tx.GasPerPubdata.String(),
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

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	fdb := make([][]uint, len(tx.FactoryDeps))
	for i, v := range tx.FactoryDeps {
		fdb[i] = make([]uint, len(v))
		for j, b := range v {
			fdb[i][j] = uint(b)
		}
	}
	return json.Marshal(&struct {
		*Alias
		FactoryDeps [][]uint `json:"factoryDeps"`
	}{
		Alias:       (*Alias)(tx),
		FactoryDeps: fdb,
	})
}

func (tx *Transaction) getFactoryDepsHashes() ([]interface{}, error) {
	if len(tx.FactoryDeps) == 0 {
		return []interface{}{}, nil
	}
	res := make([]interface{}, len(tx.FactoryDeps))
	for i, d := range tx.FactoryDeps {
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
	Paymaster      common.Address `json:"paymaster"`      // Address of the paymaster.
	PaymasterInput []byte         `json:"paymasterInput"` // Encoded input.
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

func copyAddressPtr(a *common.Address) *common.Address {
	if a == nil {
		return nil
	}
	cpy := *a
	return &cpy
}
