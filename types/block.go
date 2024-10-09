package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
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
	L1BatchNumber    *big.Int // The batch number on L1.
	L1BatchTimestamp *big.Int // The timestamp of the batch on L1.
}

// BatchDetails contains batch information.
type BatchDetails struct {
	// Hashes of the base system contracts involved in the batch.
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	CommitTxHash   *common.Hash `json:"commitTxHash"`   // The transaction hash of the commit operation on L1.
	CommittedAt    time.Time    `json:"committedAt"`    // The timestamp when the block was committed on L1.
	ExecuteTxHash  *common.Hash `json:"executeTxHash"`  // The transaction hash of the execution on L1.
	ExecutedAt     time.Time    `json:"executedAt"`     // The timestamp when the block execution was completed on L1.
	L1GasPrice     *big.Int     `json:"l1GasPrice"`     // L1 gas price at the time of the block's execution.
	L1TxCount      uint64       `json:"l1TxCount"`      // The number of L1 transactions included in the batch.
	L2FairGasPrice *big.Int     `json:"l2FairGasPrice"` // Fair gas price on L2 at the time of the block's execution.
	L2TxCount      uint64       `json:"l2TxCount"`      // The number of L2 transactions associated with this batch.
	Number         *big.Int     `json:"number"`         // L1 batch number.
	ProveTxHash    *common.Hash `json:"proveTxHash"`    // The transaction hash of the proof submission on L1.
	ProvenAt       time.Time    `json:"provenAt"`       // The timestamp when the proof was submitted on L1.
	RootHash       *common.Hash `json:"rootHash"`       // Root hash of the state after processing the batch.
	Status         string       `json:"status"`         // Current status of the batch (e.g., verified).
	Timestamp      uint64       `json:"timestamp"`      // Unix timestamp when the batch was processed.
}

// BlockDetails contains block details.
type BlockDetails struct {
	Number         *big.Int     `json:"number"`         // The number of the block.
	L1BatchNumber  *big.Int     `json:"l1BatchNumber"`  // Corresponding L1 batch number.
	Timestamp      uint64       `json:"timestamp"`      // Unix timestamp when the block was committed.
	L1TxCount      uint64       `json:"l1TxCount"`      // The number of L1 transactions included in the block.
	L2TxCount      uint64       `json:"l2TxCount"`      // The number of L2 transactions included in the block.
	RootHash       common.Hash  `json:"rootHash"`       // Root hash of the block's state after execution.
	Status         string       `json:"status"`         // Current status of the block (e.g., verified, executed).
	CommitTxHash   *common.Hash `json:"commitTxHash"`   // The transaction hash of the commit operation on L1.
	CommittedAt    time.Time    `json:"committedAt"`    // The timestamp when the block was committed on L1.
	ProveTxHash    *common.Hash `json:"proveTxHash"`    // The transaction hash of the proof submission on L1.
	ProvenAt       time.Time    `json:"provenAt"`       // The timestamp when the proof was submitted on L1.
	ExecuteTxHash  *common.Hash `json:"executeTxHash"`  // The transaction hash of the execution on L1.
	ExecutedAt     time.Time    `json:"executedAt"`     // The timestamp when the block execution was completed on L1.
	L1GasPrice     *big.Int     `json:"l1GasPrice"`     // L1 gas price at the time of the block's execution.
	L2FairGasPrice *big.Int     `json:"l2FairGasPrice"` // Fair gas price on L2 at the time of the block's execution.
	// A collection of hashes for the base system contracts involved in the block.
	BaseSystemContractsHashes struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"baseSystemContractsHashes"`
	OperatorAddress common.Address `json:"operatorAddress"` // Address of the operator who committed the block.
	ProtocolVersion string         `json:"protocolVersion"` // Version of the ZKsync protocol the block was committed under.
}

// RawBlockTransaction represents a raw block transaction.
type RawBlockTransaction struct {
	// General information about the L2 transaction.
	CommonData struct {
		L2 struct {
			Nonce uint64 `json:"nonce"`
			Fee   struct {
				GasLimit             hexutil.Big `json:"gas_limit"`
				MaxFeePerGas         hexutil.Big `json:"max_fee_per_gas"`
				MaxPriorityFeePerGas hexutil.Big `json:"max_priority_fee_per_gas"`
				GasPerPubdataLimit   hexutil.Big `json:"gas_per_pubdata_limit"`
			} `json:"fee"`
			InitiatorAddress common.Address `json:"initiatorAddress"`
			Signature        []byte         `json:"signature"`
			TransactionType  string         `json:"transactionType"`
			Input            struct {
				Hash common.Hash `json:"hash"`
				Data []uint64    `json:"data"`
			} `json:"input"`
			PaymasterParams struct {
				Paymaster      common.Address `json:"paymaster"`
				PaymasterInput []byte         `json:"paymasterInput"`
			} `json:"paymasterParams"`
		} `json:"L2"`
	} `json:"common_data"`
	// Details regarding the execution of the transaction.
	Execute struct {
		ContractAddress common.Address  `json:"contractAddress"`
		Calldata        hexutil.Bytes   `json:"calldata"`
		Value           hexutil.Big     `json:"value"`
		FactoryDeps     []hexutil.Bytes `json:"factoryDeps"`
	} `json:"execute"`
	// Timestamp when the transaction was received, in milliseconds.
	ReceivedTimestampMs uint64 `json:"received_timestamp_ms"`
	// Raw bytes of the transaction as a hexadecimal string.
	RawBytes hexutil.Bytes `json:"raw_bytes"`
}
