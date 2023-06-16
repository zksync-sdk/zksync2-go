package types

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// TODO check if ethereum.CallMsg can be inherited

type Transaction struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      *hexutil.Big   `json:"gas"`
	GasPrice *hexutil.Big   `json:"gasPrice"`
	Value    *hexutil.Big   `json:"value"`
	Data     hexutil.Bytes  `json:"data"`
	//
	Eip712Meta *Eip712Meta `json:"eip712Meta"`
}

func NewTransactionFromCallMsg(call ethereum.CallMsg) *Transaction {
	return &Transaction{
		From:     call.From,
		To:       *call.To,
		Gas:      (*hexutil.Big)(big.NewInt(int64(call.Gas))),
		GasPrice: (*hexutil.Big)(call.GasPrice),
		Value:    (*hexutil.Big)(call.Value),
		Data:     call.Data,
		Eip712Meta: &Eip712Meta{
			GasPerPubdata: (*hexutil.Big)(big.NewInt(160_000)),
		},
	}
}

type Tx struct {
	To            common.Address
	Amount        *big.Int
	Token         *Token
	BridgeAddress common.Address
}

type RequestExecuteTransaction struct {
	ContractL2               common.Address
	L2Value                  *big.Int
	Calldata                 []byte
	L2GasLimit               *big.Int
	L2GasPerPubdataByteLimit *big.Int
	FactoryDeps              [][]byte
	RefundRecipient          common.Address
	Auth                     *bind.TransactOpts
}

type DepositTransaction struct {
	Token             *Token
	Amount            *big.Int
	To                common.Address
	OperatorTip       *big.Int
	BridgeAddress     common.Address
	ApproveERC20      bool
	L2GasLimit        *big.Int
	GasPerPubdataByte *big.Int
	RefundRecipient   common.Address
	Auth              *bind.TransactOpts
	ApproveAuth       *bind.TransactOpts
	CustomBridgeData  []byte
}
