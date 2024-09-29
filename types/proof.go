package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// LogProof represents a log proof for an L2 to L1 transaction.
type LogProof struct {
	Id    int32         `json:"id"`    // Identifier of the log within the transaction.
	Proof []common.Hash `json:"proof"` // Each element represents a piece of the proof for the specified log.
	Root  common.Hash   `json:"root"`  // Root hash of the proof, anchoring it to a specific state in the blockchain.
}

// MessageProof represents a log proof for an L2 to L1 transaction.
type MessageProof struct {
	Id    int32         `json:"id"`    // Identifier of the log within the transaction.
	Proof []common.Hash `json:"proof"` // Each element represents a piece of the proof for the specified log.
	Root  common.Hash   `json:"root"`  // Root hash of the proof, anchoring it to a specific state in the blockchain.
}

// StorageProof Merkle proofs for one or more storage values at the specified account.
type StorageProof struct {
	Address string `json:"address"` // Account address associated with the storage proofs.
	// Array of objects, each representing a storage proof for the requested keys.
	Proofs []struct {
		Key string `json:"key"` // Storage key for which the proof is provided.
		// An array of 32-byte hashes that constitute the Merkle path from the leaf node
		// (representing the storage key-value pair) to the root of the Merkle tree.
		Proof []string `json:"proof"`
		Value string   `json:"value"` // Value stored in the specified storage key at the time of the specified l1BatchNumber.
		// A 1-based index representing the position of the tree entry within the Merkle tree.
		// This index is used to help reconstruct the Merkle path during verification.
		Index int32 `json:"index"`
	} `json:"storageProof"`
}

// PriorityOpConfirmation represents confirmation data that is part of L2->L1 message
type PriorityOpConfirmation struct {
	L1BatchNumber     *big.Int
	L2MessageIndex    int32
	L2TxNumberInBlock *big.Int
	Proof             []common.Hash
}
