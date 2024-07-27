package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CallMsg contains parameters for contract call using EIP-712 transaction.
type CallMsg struct {
	ethereum.CallMsg
	Meta *Eip712Meta // EIP-712 metadata.
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
