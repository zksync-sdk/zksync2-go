package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// CallMsg contains parameters for contract call using EIP-712 transaction.
type CallMsg struct {
	From      common.Address  // The sender of the 'transaction'.
	To        *common.Address // The destination contract (nil for contract creation).
	Gas       uint64          // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int        // Wei <-> gas exchange ratio
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value     *big.Int        // Amount of wei sent along with the call.
	Data      []byte          // Input data, usually an ABI-encoded contract method invocation.
	Meta      *Eip712Meta     // EIP-712 metadata.
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
	if m.Meta != nil {
		arg["eip712Meta"] = m.Meta
	}
	return json.Marshal(arg)
}
