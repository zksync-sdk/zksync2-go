package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// Deprecated: Will be removed in the future releases.
type BlockNumber string // Enums or hex value

var (
	BlockNumberEarliest  BlockNumber = "earliest"
	BlockNumberLatest    BlockNumber = "latest"
	BlockNumberPending   BlockNumber = "pending"
	BlockNumberCommitted BlockNumber = "committed"
	BlockNumberFinalized BlockNumber = "finalized"
)

// Block represents a block.
type Block struct {
	Header           *types.Header
	Uncles           []*types.Header
	Transactions     []*TransactionResponse
	Withdrawals      types.Withdrawals
	Hash             common.Hash
	Size             *big.Int
	TotalDifficulty  *big.Int
	SealFields       []interface{}
	ReceivedAt       time.Time
	ReceivedFrom     interface{}
	L1BatchNumber    *big.Int
	L1BatchTimestamp *big.Int
}

// BatchDetails contains batch information.
type BatchDetails struct {
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	CommitTxHash   *common.Hash `json:"commitTxHash"`
	CommittedAt    time.Time    `json:"committedAt"`
	ExecuteTxHash  *common.Hash `json:"executeTxHash"`
	ExecutedAt     time.Time    `json:"executedAt"`
	L1GasPrice     uint64       `json:"l1GasPrice"`
	L1TxCount      uint         `json:"l1TxCount"`
	L2FairGasPrice uint         `json:"l2FairGasPrice"`
	L2TxCount      uint         `json:"l2TxCount"`
	Number         uint         `json:"number"`
	ProveTxHash    *common.Hash `json:"proveTxHash"`
	ProvenAt       time.Time    `json:"provenAt"`
	RootHash       *common.Hash `json:"rootHash"`
	Status         string       `json:"status"`
	Timestamp      uint         `json:"timestamp"`
}

// BlockDetails contains block details.
type BlockDetails struct {
	Number                    uint         `json:"number"`
	L1BatchNumber             uint         `json:"l1BatchNumber"`
	Timestamp                 uint         `json:"timestamp"`
	L1TxCount                 uint         `json:"l1TxCount"`
	L2TxCount                 uint         `json:"l2TxCount"`
	RootHash                  common.Hash  `json:"rootHash"`
	Status                    string       `json:"status"`
	CommitTxHash              *common.Hash `json:"commitTxHash"`
	CommittedAt               time.Time    `json:"committedAt"`
	ProveTxHash               *common.Hash `json:"proveTxHash"`
	ProvenAt                  time.Time    `json:"provenAt"`
	ExecuteTxHash             *common.Hash `json:"executeTxHash"`
	ExecutedAt                time.Time    `json:"executedAt"`
	L1GasPrice                *big.Int     `json:"l1GasPrice"`
	L2FairGasPrice            *big.Int     `json:"l2FairGasPrice"`
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	OperatorAddress common.Address `json:"operatorAddress"`
	ProtocolVersion string         `json:"protocolVersion"`
}

type RawBlockTransaction struct {
	CommonData struct {
		L1 struct {
			CanonicalTxHash    common.Hash    `json:"canonicalTxHash"`
			DeadlineBlock      *big.Int       `json:"deadlineBlock"`
			EthBlock           *big.Int       `json:"ethBlock"`
			EthHash            common.Hash    `json:"ethHash"`
			FullFee            hexutil.Big    `json:"fullFee"`
			GasLimit           hexutil.Big    `json:"gasLimit"`
			GasPerPubdataLimit hexutil.Big    `json:"gasPerPubdataLimit"`
			Layer2TipFee       hexutil.Big    `json:"layer2TipFee"`
			MaxFeePerGas       hexutil.Big    `json:"maxFeePerGas"`
			OpProcessingType   string         `json:"opProcessingType"`
			PriorityQueueType  string         `json:"priorityQueueType"`
			RefundRecipient    common.Address `json:"refundRecipient"`
			Sender             common.Address `json:"sender"`
			SerialId           *big.Int       `json:"serialId"`
			ToMint             string         `json:"toMint"`
		} `json:"L1"`
	} `json:"common_data"`
	Execute struct {
		Calldata        hexutil.Bytes   `json:"calldata"`
		ContractAddress common.Address  `json:"contractAddress"`
		FactoryDeps     []hexutil.Bytes `json:"factoryDeps"`
		Value           hexutil.Big     `json:"value"`
	} `json:"execute"`
	ReceivedTimestampMs uint64 `json:"received_timestamp_ms"`
}
