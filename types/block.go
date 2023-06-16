package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

type BlockNumber string // Enums or hex value

var (
	BlockNumberEarliest  BlockNumber = "earliest"
	BlockNumberLatest    BlockNumber = "latest"
	BlockNumberPending   BlockNumber = "pending"
	BlockNumberCommitted BlockNumber = "committed"
	BlockNumberFinalized BlockNumber = "finalized"
)

type Block struct {
	types.Block
	L1BatchNumber    *hexutil.Big `json:"l1BatchNumber"`
	L1BatchTimestamp *hexutil.Big `json:"l1BatchTimestamp"`
}

type BatchDetails struct {
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	CommitTxHash   common.Hash `json:"commitTxHash"`
	CommittedAt    time.Time   `json:"committedAt"`
	ExecuteTxHash  common.Hash `json:"executeTxHash"`
	ExecutedAt     time.Time   `json:"executedAt"`
	L1GasPrice     uint64      `json:"l1GasPrice"`
	L1TxCount      uint        `json:"l1TxCount"`
	L2FairGasPrice uint        `json:"l2FairGasPrice"`
	L2TxCount      uint        `json:"l2TxCount"`
	Number         uint        `json:"number"`
	ProveTxHash    common.Hash `json:"proveTxHash"`
	ProvenAt       time.Time   `json:"provenAt"`
	RootHash       common.Hash `json:"rootHash"`
	Status         string      `json:"status"`
	Timestamp      uint        `json:"timestamp"`
}

type BlockDetails struct {
	CommitTxHash  common.Hash `json:"commitTxHash"`
	CommittedAt   time.Time   `json:"committedAt"`
	ExecuteTxHash common.Hash `json:"executeTxHash"`
	ExecutedAt    time.Time   `json:"executedAt"`
	L1TxCount     uint        `json:"l1TxCount"`
	L2TxCount     uint        `json:"l2TxCount"`
	Number        uint        `json:"number"`
	ProveTxHash   common.Hash `json:"proveTxHash"`
	ProvenAt      time.Time   `json:"provenAt"`
	RootHash      common.Hash `json:"rootHash"`
	Status        string      `json:"status"`
	Timestamp     uint        `json:"timestamp"`
}

type BlockRange struct {
	Beginning *big.Int `json:"beginning"`
	End       *big.Int `json:"end"`
}

func (r *BlockRange) UnmarshalJSON(input []byte) error {
	var data [2]*hexutil.Big
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}
	r.Beginning = data[0].ToInt()
	r.End = data[1].ToInt()
	return nil
}
