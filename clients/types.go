package clients

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
)

// TransferCallMsg contains parameters for transfer call.
type TransferCallMsg struct {
	To     common.Address // The address of the recipient.
	Amount *big.Int       // The amount of the token to transfer.
	Token  common.Address // The address of the token. ETH by default.
	From   common.Address // The address of the sender.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.

	PaymasterParams *zkTypes.PaymasterParams // The paymaster parameters.
}

func (m *TransferCallMsg) ToCallMsg() (*ethereum.CallMsg, error) {
	var (
		value *big.Int
		data  []byte
		to    *common.Address
	)

	if m.Token == utils.LegacyEthAddress {
		value = m.Amount
		to = &m.To
	} else {
		value = big.NewInt(0)
		to = &m.Token

		erc20abi, err := erc20.IERC20MetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to load erc20abi: %w", err)
		}
		data, err = erc20abi.Pack("transfer", m.To, m.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to pack transfer function: %w", err)
		}
	}

	return &ethereum.CallMsg{
		From:      m.From,
		To:        to,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     value,
		Data:      data,
	}, nil
}

func (m *TransferCallMsg) ToZkCallMsg() (*zkTypes.CallMsg, error) {
	var (
		value *big.Int
		data  []byte
		to    *common.Address
		meta  *zkTypes.Eip712Meta
	)

	if m.Token == utils.LegacyEthAddress {
		value = m.Amount
		to = &m.To
	} else {
		value = big.NewInt(0)
		to = &m.Token

		erc20abi, err := erc20.IERC20MetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to load erc20abi: %w", err)
		}
		data, err = erc20abi.Pack("transfer", m.To, m.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to pack transfer function: %w", err)
		}
	}
	if m.PaymasterParams != nil {
		meta = &zkTypes.Eip712Meta{
			GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			PaymasterParams: m.PaymasterParams,
		}
	}

	return &zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:      m.From,
			To:        to,
			Gas:       m.Gas,
			GasPrice:  m.GasPrice,
			GasFeeCap: m.GasFeeCap,
			GasTipCap: m.GasTipCap,
			Value:     value,
			Data:      data,
		},
		Meta: meta,
	}, nil
}

// WithdrawalCallMsg contains parameters for withdrawal call.
type WithdrawalCallMsg struct {
	To            common.Address  // The address of the recipient on L1.
	Amount        *big.Int        // The amount of the token to transfer.
	Token         common.Address  // The address of the token. ETH by default.
	BridgeAddress *common.Address // The address of the bridge contract to be used.
	From          common.Address  // The address of the sender.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.

	PaymasterParams *zkTypes.PaymasterParams // The paymaster parameters.
}

func (m *WithdrawalCallMsg) ToCallMsg(defaultL2Bridge *common.Address) (*ethereum.CallMsg, error) {
	if m.Token == utils.LegacyEthAddress || m.Token == utils.L2BaseTokenAddress || m.Token == utils.EthAddressInContracts {
		ethTokenAbi, err := ethtoken.IEthTokenMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
		}

		data, errPack := ethTokenAbi.Pack("withdraw", m.To)
		if errPack != nil {
			return nil, fmt.Errorf("failed to pack withdraw function: %w", errPack)
		}
		return &ethereum.CallMsg{
			From:      m.From,
			To:        &utils.L2BaseTokenAddress,
			Gas:       m.Gas,
			GasPrice:  m.GasPrice,
			GasFeeCap: m.GasFeeCap,
			GasTipCap: m.GasTipCap,
			Value:     m.Amount,
			Data:      data,
		}, nil
	} else {
		l2BridgeAbi, err := l2bridge.IL2BridgeMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to load l2BridgeAbi: %w", err)
		}
		data, err := l2BridgeAbi.Pack("withdraw", m.To, m.Token, m.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
		}
		bridge := m.BridgeAddress
		if defaultL2Bridge != nil {
			bridge = defaultL2Bridge
		}

		return &ethereum.CallMsg{
			From:      m.From,
			To:        bridge,
			Gas:       m.Gas,
			GasPrice:  m.GasPrice,
			GasFeeCap: m.GasFeeCap,
			GasTipCap: m.GasTipCap,
			Value:     big.NewInt(0),
			Data:      data,
		}, nil
	}
}

func (m *WithdrawalCallMsg) ToZkCallMsg(defaultL2Bridge *common.Address) (*zkTypes.CallMsg, error) {
	var meta *zkTypes.Eip712Meta

	msg, err := m.ToCallMsg(defaultL2Bridge)
	if err != nil {
		return nil, err
	}
	if m.PaymasterParams != nil {
		meta = &zkTypes.Eip712Meta{
			GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			PaymasterParams: m.PaymasterParams,
		}
	}

	return &zkTypes.CallMsg{
		CallMsg: *msg,
		Meta:    meta,
	}, nil
}

type blockMarshaling struct {
	ParentHash    common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash     common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase      common.Address   `json:"miner"`
	Root          common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash        common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash   common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom         types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty    *hexutil.Big     `json:"difficulty"       gencodec:"required"`
	Number        *hexutil.Big     `json:"number"           gencodec:"required"`
	GasLimit      hexutil.Uint64   `json:"gasLimit"         gencodec:"required"`
	GasUsed       hexutil.Uint64   `json:"gasUsed"          gencodec:"required"`
	Time          hexutil.Uint64   `json:"timestamp"        gencodec:"required"`
	Extra         hexutil.Bytes    `json:"extraData"        gencodec:"required"`
	MixDigest     common.Hash      `json:"mixHash"`
	Nonce         types.BlockNonce `json:"nonce"`
	BaseFee       *hexutil.Big     `json:"baseFeePerGas" rlp:"optional"`
	ExcessDataGas *hexutil.Big     `json:"excessDataGas" rlp:"optional"`

	Uncles           []*common.Hash `json:"uncles"`
	Hash             *common.Hash   `json:"hash"`
	L1BatchNumber    *hexutil.Big   `json:"l1BatchNumber"`
	L1BatchTimestamp *hexutil.Big   `json:"l1BatchTimestamp"`
	TotalDifficulty  *hexutil.Big   `json:"totalDifficulty"`
	Size             *hexutil.Big   `json:"size"`
	SealFields       []interface{}  `json:"sealFields"`

	Transactions []*zkTypes.TransactionResponse `json:"transactions"`
}

// BlockRange represents a range of blocks with the starting and ending block numbers.
type BlockRange struct {
	Beginning *big.Int `json:"beginning"` // Starting block number of the range.
	End       *big.Int `json:"end"`       // Ending block number of the range.
}

func (r *BlockRange) UnmarshalJSON(input []byte) error {
	var data [2]*hexutil.Big
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}
	r.Beginning = data[0].ToInt()
	r.End = data[1].ToInt()
	return nil
}
