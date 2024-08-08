package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// Receipt represents the results of a transaction.
type Receipt struct {
	types.Receipt

	From              common.Address `json:"from"`
	To                common.Address `json:"to"`
	EffectiveGasPrice *hexutil.Big   `json:"effectiveGasPrice"`
	L1BatchNumber     *hexutil.Big   `json:"l1BatchNumber"`  // The batch number on the L1 network.
	L1BatchTxIndex    *hexutil.Big   `json:"l1BatchTxIndex"` // The transaction index within the batch on the L1 network.
	Logs              []*Log         `json:"logs"`
	L2ToL1Logs        []*L2ToL1Log   `json:"l2ToL1Logs"` // The logs of L2 to L1 messages.
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
	L1BatchNumber        hexutil.Big    `json:"l1BatchNumber"`  // The batch number on the L1 network.
	L1BatchTxIndex       hexutil.Big    `json:"l1BatchTxIndex"` // The transaction index within the batch on the L1 network.
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
	EthCommitTxHash  *common.Hash   `json:"ethCommitTxHash"`  // The transaction hash of the commit operation.
	EthExecuteTxHash *common.Hash   `json:"ethExecuteTxHash"` // The transaction hash of the execution.
	EthProveTxHash   *common.Hash   `json:"ethProveTxHash"`   // The transaction hash of the proof submission.
	Fee              hexutil.Big    `json:"fee"`              // The transaction fee.
	GasPerPubdata    hexutil.Big    `json:"gasPerPubdata"`    // Gas amount per unit of public data for this transaction.
	InitiatorAddress common.Address `json:"initiatorAddress"` // Address of the transaction initiator.
	IsL1Originated   bool           `json:"isL1Originated"`   // Indicates whether the transaction originated on Layer 1.
	ReceivedAt       time.Time      `json:"receivedAt"`       // Timestamp when the transaction was received.
	Status           string         `json:"status"`           // Current status of the transaction (e.g., verified).
}

// ProtocolVersion represents the protocol version.
type ProtocolVersion struct {
	VersionId uint `json:"version_id"` // Protocol version ID.
	Timestamp uint `json:"timestamp"`  // Unix timestamp of the version's activation.
	// Contains the hashes of various verification keys used in the protocol.
	VerificationKeysHashes struct {
		Params struct {
			RecursionNodeLevelVkHash    common.Hash `json:"recursion_node_level_vk_hash"`
			RecursionLeafLevelVkHash    common.Hash `json:"recursion_leaf_level_vk_hash"`
			RecursionCircuitsSetVksHash common.Hash `json:"recursion_circuits_set_vks_hash"`
		} `json:"params"`
		RecursionSchedulerLevelVkHash common.Hash `json:"recursion_scheduler_level_vk_hash"`
	} `json:"verification_keys_hashes"`
	// Hashes of the base system contracts.
	BaseSystemContracts struct {
		Bootloader common.Hash `json:"bootloader"`
		DefaultAa  common.Hash `json:"default_aa"`
	} `json:"base_system_contracts"`
	// Hash of the transaction used for the system upgrade, if any.
	L2SystemUpgradeTxHash *common.Hash `json:"l2_system_upgrade_tx_hash"`
}

// TransactionWithDetailedOutput represents the transaction with detailed output.
type TransactionWithDetailedOutput struct {
	TransactionHash common.Hash `json:"transactionHash"` // The transaction hash.
	// Storage slots.
	StorageLogs []struct {
		Address      common.Address `json:"address"`
		Key          string         `json:"key"`
		WrittenValue string         `json:"writtenValue"`
	} `json:"storageLogs"`
	// Generated events.
	Events []struct {
		Address             common.Address `json:"address"`
		Topics              []common.Hash  `json:"topics"`
		Data                hexutil.Bytes  `json:"data"`
		BlockHash           *common.Hash   `json:"blockHash"`
		BlockNumber         *hexutil.Big   `json:"blockNumber"`
		L1BatchNumber       *hexutil.Big   `json:"l1BatchNumber"`
		TransactionHash     common.Hash    `json:"transactionHash"`
		TransactionIndex    hexutil.Uint   `json:"transactionIndex"`
		LogIndex            *hexutil.Uint  `json:"logIndex"`
		TransactionLogIndex *hexutil.Uint  `json:"transactionLogIndex"`
		LogType             *hexutil.Bytes `json:"logType"`
		Removed             bool           `json:"removed"`
	} `json:"events"`
}
