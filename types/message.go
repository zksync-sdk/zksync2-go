package types

import "github.com/ethereum/go-ethereum/common"

type L2ToL1MessageProof struct {
	Id    int           `json:"id"`
	Proof []common.Hash `json:"proof"`
	Root  common.Hash   `json:"root"`
}
