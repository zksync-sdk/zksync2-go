package zksync2

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

var (
	EthAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

const (
	PriorityQueueTypeDeque      uint8 = 0
	PriorityQueueTypeHeapBuffer uint8 = 1
	PriorityQueueTypeHeap       uint8 = 2
)

type BridgeContracts struct {
	L1Erc20DefaultBridge common.Address `json:"l1Erc20DefaultBridge"`
	L1EthDefaultBridge   common.Address `json:"l1EthDefaultBridge"`
	L2Erc20DefaultBridge common.Address `json:"l2Erc20DefaultBridge"`
	L2EthDefaultBridge   common.Address `json:"l2EthDefaultBridge"`
}

type Token struct {
	L1Address common.Address `json:"l1Address"`
	L2Address common.Address `json:"l2Address"`
	Name      string         `json:"name"`
	Symbol    string         `json:"symbol"`
	Decimals  uint           `json:"decimals"`
}

func (t *Token) IsETH() bool {
	return bytes.Equal(t.L2Address.Bytes(), common.Address{}.Bytes()) && t.Symbol == "ETH"
}

func CreateETH() *Token {
	return &Token{
		L1Address: common.Address{},
		L2Address: common.Address{},
		Name:      `ETH`,
		Symbol:    `ETH`,
		Decimals:  18,
	}
}

type L2ToL1MessageProof struct {
	Id    int           `json:"id"`
	Proof []common.Hash `json:"proof"`
	Root  common.Hash   `json:"root"`
}

type BlockNumber string // Enums or hex value

var (
	BlockNumberEarliest  BlockNumber = "earliest"
	BlockNumberLatest    BlockNumber = "latest"
	BlockNumberPending   BlockNumber = "pending"
	BlockNumberCommitted BlockNumber = "committed"
	BlockNumberFinalized BlockNumber = "finalized"
)

type BigInt struct {
	*big.Int
}

// MarshalJSON overwrite to encode value as hex string
func (x *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(hexutil.EncodeBig(x.Int))
}

func NewBigInt(x int64) *BigInt {
	return &BigInt{big.NewInt(x)}
}

func ToBigInt(x *big.Int) *BigInt {
	return &BigInt{x}
}
