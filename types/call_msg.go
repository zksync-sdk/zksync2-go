package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// CallMsg contains parameters for contract call using L2 transaction.
type CallMsg struct {
	From      common.Address  // The sender of the 'transaction'.
	To        *common.Address // The destination contract (nil for contract creation).
	Gas       uint64          // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int        // Wei <-> gas exchange ratio
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value     *big.Int        // Amount of wei sent along with the call.
	Data      []byte          // Input data, usually an ABI-encoded contract method invocation.

	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes
	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
	// it should contain the bytecode of the contract being deployed.
	// If the contract is a factory contract, i.e. it can deploy other contracts,
	// the array should also contain the bytecodes of the contracts which it can deploy.
	FactoryDeps []hexutil.Bytes
	// PaymasterParams contains parameters for configuring the custom paymaster
	// for the transaction.
	PaymasterParams *PaymasterParams
}

func (m CallMsg) MarshalJSON() ([]byte, error) {
	arg := map[string]interface{}{
		"type": "0x71",
		"from": m.From,
		"to":   m.To,
	}
	if len(m.Data) > 0 {
		arg["data"] = hexutil.Bytes(m.Data)
	}
	if m.Value != nil {
		arg["value"] = (*hexutil.Big)(m.Value)
	}
	if m.Gas != 0 {
		arg["gas"] = hexutil.Uint64(m.Gas)
	}
	if m.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(m.GasPrice)
	}
	if m.GasTipCap != nil {
		arg["maxPriorityFeePerGas"] = (*hexutil.Big)(m.GasTipCap)
	}
	if m.GasFeeCap != nil {
		arg["maxFeePerGas"] = (*hexutil.Big)(m.GasFeeCap)
	}

	eip712Meta := map[string]interface{}{}

	if m.GasPerPubdata != nil {
		eip712Meta["gasPerPubdata"] = (*hexutil.Big)(m.GasPerPubdata)
	}
	if len(m.CustomSignature) > 0 {
		eip712Meta["customSignature"] = m.CustomSignature
	}
	if len(m.FactoryDeps) > 0 {
		// Convert FactoryDeps into [][]uint format
		fdb := make([][]uint, len(m.FactoryDeps))
		for i, v := range m.FactoryDeps {
			fdb[i] = make([]uint, len(v))
			for j, b := range v {
				fdb[i][j] = uint(b)
			}
		}
		eip712Meta["factoryDeps"] = fdb
	}
	if m.PaymasterParams != nil {
		eip712Meta["paymasterParams"] = m.PaymasterParams
	}

	// Add eip712Meta to the main argument map
	arg["eip712Meta"] = eip712Meta
	return json.Marshal(arg)
}
