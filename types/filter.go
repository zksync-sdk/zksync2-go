package types

import "github.com/ethereum/go-ethereum/common"

type FilterQuery struct {
	BlockHash *common.Hash
	FromBlock *BlockNumber
	ToBlock   *BlockNumber
	Addresses []common.Address
	Topics    [][]common.Hash
}
