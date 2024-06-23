package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// Log represents a log entry.
type Log struct {
	types.Log
	L1BatchNumber *hexutil.Big `json:"l1BatchNumber"` // The batch number on L1.
}

func (l *Log) MarshalJSON() ([]byte, error) {
	// get json of embedded types.Log with its custom marshaller
	lj, err := json.Marshal(l.Log)
	if err != nil {
		return nil, err
	}
	// decode back to abstract struct
	var buf map[string]interface{}
	if err := json.Unmarshal(lj, &buf); err != nil {
		return nil, err
	}
	// mixin our fields
	buf["l1BatchNumber"] = l.L1BatchNumber
	// encode to json again all together
	return json.Marshal(&buf)
}

func (l *Log) UnmarshalJSON(input []byte) error {
	if err := l.Log.UnmarshalJSON(input); err != nil {
		return err
	}
	type Log struct {
		L1BatchNumber *hexutil.Big `json:"l1BatchNumber"`
	}
	var dec Log
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	l.L1BatchNumber = dec.L1BatchNumber
	return nil
}

// L2ToL1Log represents a layer 2 to layer 1 transaction log.
type L2ToL1Log struct {
	BlockNumber         *hexutil.Big  `json:"blockNumber"`         // The block number.
	BlockHash           common.Hash   `json:"blockHash"`           // The block hash.
	L1BatchNumber       *hexutil.Big  `json:"l1BatchNumber"`       // The batch number on L1.
	TransactionIndex    *hexutil.Uint `json:"transactionIndex"`    // The L2 transaction number in a block, in which the log was sent.
	TransactionLogIndex *hexutil.Uint `json:"transactionLogIndex"` // The transaction log index.
	TxIndexInL1Batch    *hexutil.Uint `json:"txIndexInL1Batch"`    // The transaction index in L1 batch.
	ShardId             *hexutil.Uint `json:"shardId"`             // The shard identifier, 0 - rollup, 1 - porter.
	// A boolean flag that is part of the log along with `key`, `value`, and `sender` address.
	// This field is required formally but does not have any special meaning.
	IsService bool           `json:"isService"`
	Sender    common.Address `json:"sender"`          // The L2 address which sent the log.
	Key       string         `json:"key"`             // The 32 bytes of information that was sent in the log.
	Value     string         `json:"value"`           // The 32 bytes of information that was sent in the log.
	TxHash    common.Hash    `json:"transactionHash"` // The transaction hash.
	Index     *hexutil.Uint  `json:"logIndex"`        // The log index.
}
