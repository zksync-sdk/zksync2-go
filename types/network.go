package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// Deprecated: Deprecated in favor of Receipt
type TransactionReceipt struct {
	types.Receipt
	// extend
	From              common.Address `json:"from"`
	To                common.Address `json:"to"`
	EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
	L1BatchNumber     *hexutil.Big   `json:"l1BatchNumber"`
	L1BatchTxIndex    *hexutil.Big   `json:"l1BatchTxIndex"`
	Logs              []*Log         `json:"logs"`
	L2ToL1Logs        []*L2ToL1Log   `json:"l2ToL1Logs"`
}

func (r *TransactionReceipt) MarshalJSON() ([]byte, error) {
	// get json of embedded types.Receipt with its custom marshaller
	rj, err := json.Marshal(r.Receipt)
	if err != nil {
		return nil, err
	}
	// decode back to abstract struct
	var buf map[string]interface{}
	if err := json.Unmarshal(rj, &buf); err != nil {
		return nil, err
	}
	// mixin our fields
	buf["from"] = r.From
	buf["to"] = r.To
	buf["effectiveGasPrice"] = r.EffectiveGasPrice
	buf["l1BatchNumber"] = r.L1BatchNumber
	buf["l1BatchTxIndex"] = r.L1BatchTxIndex
	buf["logs"] = r.Logs
	buf["l2ToL1Logs"] = r.L2ToL1Logs
	// encode to json again all together
	return json.Marshal(&buf)
}

func (r *TransactionReceipt) UnmarshalJSON(input []byte) error {
	if err := r.Receipt.UnmarshalJSON(input); err != nil {
		return err
	}
	type TransactionReceipt struct {
		From              common.Address `json:"from"`
		To                common.Address `json:"to"`
		EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
		L1BatchNumber     *hexutil.Big   `json:"l1BatchNumber"`
		L1BatchTxIndex    *hexutil.Big   `json:"l1BatchTxIndex"`
		Logs              []*Log         `json:"logs"`
		L2ToL1Logs        []*L2ToL1Log   `json:"l2ToL1Logs"`
	}
	var dec TransactionReceipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	r.From = dec.From
	r.To = dec.To
	r.EffectiveGasPrice = dec.EffectiveGasPrice
	r.L1BatchNumber = dec.L1BatchNumber
	r.L1BatchTxIndex = dec.L1BatchTxIndex
	r.Logs = dec.Logs
	r.L2ToL1Logs = dec.L2ToL1Logs
	return nil
}

// Receipt represents the results of a transaction.
type Receipt struct {
	types.Receipt

	From              common.Address `json:"from"`
	To                common.Address `json:"to"`
	EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
	L1BatchNumber     *hexutil.Big   `json:"l1BatchNumber"`
	L1BatchTxIndex    *hexutil.Big   `json:"l1BatchTxIndex"`
	Logs              []*Log         `json:"logs"`
	L2ToL1Logs        []*L2ToL1Log   `json:"l2ToL1Logs"`
}

func (r *Receipt) MarshalJSON() ([]byte, error) {
	// get json of embedded types.Receipt with its custom marshaller
	rj, err := json.Marshal(r.Receipt)
	if err != nil {
		return nil, err
	}
	// decode back to abstract struct
	var buf map[string]interface{}
	if err := json.Unmarshal(rj, &buf); err != nil {
		return nil, err
	}
	// mixin our fields
	buf["from"] = r.From
	buf["to"] = r.To
	buf["effectiveGasPrice"] = r.EffectiveGasPrice
	buf["l1BatchNumber"] = r.L1BatchNumber
	buf["l1BatchTxIndex"] = r.L1BatchTxIndex
	buf["logs"] = r.Logs
	buf["l2ToL1Logs"] = r.L2ToL1Logs
	// encode to json again all together
	return json.Marshal(&buf)
}

func (r *Receipt) UnmarshalJSON(input []byte) error {
	if err := r.Receipt.UnmarshalJSON(input); err != nil {
		return err
	}
	type TransactionReceipt struct {
		From              common.Address `json:"from"`
		To                common.Address `json:"to"`
		EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
		L1BatchNumber     *hexutil.Big   `json:"l1BatchNumber"`
		L1BatchTxIndex    *hexutil.Big   `json:"l1BatchTxIndex"`
		Logs              []*Log         `json:"logs"`
		L2ToL1Logs        []*L2ToL1Log   `json:"l2ToL1Logs"`
	}
	var dec TransactionReceipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	r.From = dec.From
	r.To = dec.To
	r.EffectiveGasPrice = dec.EffectiveGasPrice
	r.L1BatchNumber = dec.L1BatchNumber
	r.L1BatchTxIndex = dec.L1BatchTxIndex
	r.Logs = dec.Logs
	r.L2ToL1Logs = dec.L2ToL1Logs
	return nil
}

// TransactionResponse includes all properties of a transaction
// as well as several properties that are useful once it has been mined.
type TransactionResponse struct {
	BlockHash            *common.Hash   `json:"blockHash"`
	BlockNumber          *hexutil.Big   `json:"blockNumber"`
	ChainID              hexutil.Big    `json:"chainId"`
	From                 common.Address `json:"from"`
	Gas                  hexutil.Uint64 `json:"gas"`
	GasPrice             hexutil.Big    `json:"gasPrice"`
	Hash                 common.Hash    `json:"hash"`
	Data                 hexutil.Bytes  `json:"input"`
	L1BatchNumber        hexutil.Big    `json:"l1BatchNumber"`
	L1BatchTxIndex       hexutil.Big    `json:"l1BatchTxIndex"`
	MaxFeePerGas         hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas hexutil.Big    `json:"maxPriorityFeePerGas"`
	Nonce                hexutil.Uint64 `json:"nonce"`
	V                    *hexutil.Big   `json:"v"`
	R                    *hexutil.Big   `json:"r"`
	S                    *hexutil.Big   `json:"s"`
	To                   common.Address `json:"to"`
	TransactionIndex     hexutil.Uint   `json:"transactionIndex"`
	Type                 hexutil.Uint64 `json:"type"`
	Value                hexutil.Big    `json:"value"`
}

// TransactionDetails contains transaction details.
type TransactionDetails struct {
	EthCommitTxHash  common.Hash    `json:"ethCommitTxHash"`
	EthExecuteTxHash common.Hash    `json:"ethExecuteTxHash"`
	EthProveTxHash   common.Hash    `json:"ethProveTxHash"`
	Fee              hexutil.Big    `json:"fee"`
	InitiatorAddress common.Address `json:"initiatorAddress"`
	IsL1Originated   bool           `json:"isL1Originated"`
	ReceivedAt       time.Time      `json:"receivedAt"`
	Status           string         `json:"status"`
}
