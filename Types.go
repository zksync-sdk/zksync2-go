package zksync2

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

var (
	EthAddress              = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ContractDeployerAddress = common.HexToAddress("0x0000000000000000000000000000000000008006")
	NonceHolderAddress      = common.HexToAddress("0x0000000000000000000000000000000000008003")
	MessengerAddress        = common.HexToAddress("0x0000000000000000000000000000000000008008")
)

const (
	ZkSyncChainIdMainnet   int64 = 280
	ZkSyncChainIdLocalhost int64 = 42

	PriorityQueueTypeDeque      uint8 = 0
	PriorityQueueTypeHeapBuffer uint8 = 1
	PriorityQueueTypeHeap       uint8 = 2
)

type BridgeContracts struct {
	L1Erc20DefaultBridge common.Address `json:"l1Erc20DefaultBridge"`
	L2Erc20DefaultBridge common.Address `json:"l2Erc20DefaultBridge"`
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

type Fee struct {
	GasLimit             *hexutil.Big `json:"gas_limit"`
	GasPerPubdataLimit   *hexutil.Big `json:"gas_per_pubdata_limit"`
	MaxFeePerGas         *hexutil.Big `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas *hexutil.Big `json:"max_priority_fee_per_gas"`
}

type BlockNumber string // Enums or hex value

var (
	BlockNumberEarliest  BlockNumber = "earliest"
	BlockNumberLatest    BlockNumber = "latest"
	BlockNumberPending   BlockNumber = "pending"
	BlockNumberCommitted BlockNumber = "committed"
	BlockNumberFinalized BlockNumber = "finalized"
)

type BlockDetails struct {
	CommitTxHash  common.Hash `json:"commitTxHash"`
	CommittedAt   time.Time   `json:"committedAt"`
	ExecuteTxHash common.Hash `json:"executeTxHash"`
	ExecutedAt    time.Time   `json:"executedAt"`
	L1TxCount     int         `json:"l1TxCount"`
	L2TxCount     int         `json:"l2TxCount"`
	Number        int         `json:"number"`
	ProveTxHash   common.Hash `json:"proveTxHash"`
	ProvenAt      time.Time   `json:"provenAt"`
	RootHash      common.Hash `json:"rootHash"`
	Status        string      `json:"status"`
	Timestamp     uint        `json:"timestamp"`
}

func NewBigZero() *hexutil.Big {
	return (*hexutil.Big)(new(big.Int))
}

func NewBig(n int64) *hexutil.Big {
	return (*hexutil.Big)(big.NewInt(n))
}
