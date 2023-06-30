package types

import "github.com/ethereum/go-ethereum/common"

// MessageProof represents a message proof.
type MessageProof struct {
	Id    int           `json:"id"`
	Proof []common.Hash `json:"proof"`
	Root  common.Hash   `json:"root"`
}
