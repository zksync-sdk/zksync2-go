package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// MessageProof represents a message proof.
type MessageProof struct {
	Id    int           `json:"id"`
	Proof []common.Hash `json:"proof"`
	Root  common.Hash   `json:"root"`
}

// StorageProof Merkle proofs for one or more storage values at the specified account
type StorageProof struct {
	Address string `json:"address"`
	Proofs  []struct {
		Key   string   `json:"key"`
		Proof []string `json:"proof"`
		Value string   `json:"value"`
		Index int      `json:"index"`
	} `json:"storageProof"`
}

type PriorityOpConfirmation struct {
	L1BatchNumber     *big.Int
	L2MessageIndex    int
	L2TxNumberInBlock *big.Int
	Proof             []common.Hash
}
