package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Log struct {
	types.Log
	L1BatchNumber *hexutil.Big `json:"l1BatchNumber"`
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

type L2ToL1Log struct {
	BlockNumber      *hexutil.Big   `json:"blockNumber"`
	BlockHash        common.Hash    `json:"blockHash"`
	L1BatchNumber    *hexutil.Big   `json:"l1BatchNumber"`
	TransactionIndex *hexutil.Uint  `json:"transactionIndex"`
	ShardId          *hexutil.Uint  `json:"shardId"`
	IsService        bool           `json:"isService"`
	Sender           common.Address `json:"sender"`
	Key              string         `json:"key"`
	Value            string         `json:"value"`
	TxHash           common.Hash    `json:"transactionHash"`
	Index            *hexutil.Uint  `json:"logIndex"`
}
