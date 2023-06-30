package types

import "github.com/ethereum/go-ethereum/common"

// Deprecated: Deprecated in favor of ethereum.FilterQuery
type FilterQuery struct {
	BlockHash *common.Hash
	FromBlock *BlockNumber
	ToBlock   *BlockNumber
	Addresses []common.Address
	Topics    [][]common.Hash
}
