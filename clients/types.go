package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/contracts/l2assetrouter"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2nativetokenvault"
	"github.com/zksync-sdk/zksync2-go/types"
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

	PaymasterParams *types.PaymasterParams // The paymaster parameters.
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes
}

// ToL1CallMsg transforms TransferCallMsg to ethereum.CallMsg.
func (m *TransferCallMsg) ToL1CallMsg() (*ethereum.CallMsg, error) {
	var (
		value *big.Int
		data  []byte
		to    *common.Address
	)

	if m.Token == utils.L2BaseTokenAddress {
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

// ToCallMsg transforms TransferCallMsg to types.CallMsg.
func (m *TransferCallMsg) ToCallMsg() (*types.CallMsg, error) {
	var (
		value         *big.Int
		data          []byte
		to            *common.Address
		gasPerPubdata *big.Int
	)

	if m.Token == utils.L2BaseTokenAddress {
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

	if m.GasPerPubdata == nil {
		gasPerPubdata = utils.DefaultGasPerPubdataLimit
	}

	return &types.CallMsg{
		From:            m.From,
		To:              to,
		Gas:             m.Gas,
		GasPrice:        m.GasPrice,
		GasFeeCap:       m.GasFeeCap,
		GasTipCap:       m.GasTipCap,
		Value:           value,
		Data:            data,
		GasPerPubdata:   gasPerPubdata,
		PaymasterParams: m.PaymasterParams,
		CustomSignature: m.CustomSignature,
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

	PaymasterParams *types.PaymasterParams // The paymaster parameters.
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes
}

// ToL1CallMsg transforms WithdrawalCallMsg to ethereum.CallMsg.
func (m *WithdrawalCallMsg) ToL1CallMsg(ctx context.Context, client *Client, defaultL2Bridge *common.Address) (*ethereum.CallMsg, error) {
	if m.Token == utils.L2BaseTokenAddress {
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
	}

	ntv, err := l2nativetokenvault.NewIL2NativeTokenVault(utils.L2NativeTokenVaultAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to init l2NativeTokenVault: %w", err)
	}
	assetId, err := ntv.AssetId(&bind.CallOpts{Context: ctx}, m.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset id from l2NativeTokenVault: %w", err)
	}
	originChainId, err := ntv.OriginChainId(&bind.CallOpts{Context: ctx}, assetId)
	if err != nil {
		return nil, fmt.Errorf("failed to get origin chain id from l2NativeTokenVault: %w", err)
	}
	l1ChainId, err := client.L1ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get L1 chain id: %w", err)
	}
	isTokenL1Native := originChainId == l1ChainId || m.Token == utils.EthAddressInContracts
	bridgeAddress := m.BridgeAddress
	if bridgeAddress != nil && isTokenL1Native {
		// If the legacy L2SharedBridge is deployed we use it for l1 native tokens.
		bridgeAddress = defaultL2Bridge
	} else if bridgeAddress != nil && !isTokenL1Native {
		bridgeAddress = &utils.L2AssetRouterAddress
	}

	// For non L1 native tokens we need to use the AssetRouter.
	// For L1 native tokens we can use the legacy withdraw method.
	if !isTokenL1Native {
		chainId, errChainId := client.ChainID(ctx)
		if errChainId != nil {
			return nil, fmt.Errorf("failed to get chain id: %w", err)
		}
		nativeAssetId, errNativeAssetId := utils.NativeTokenVaultAssetId(chainId, m.Token)
		if errNativeAssetId != nil {
			return nil, fmt.Errorf("failed to encode asset id: %w", err)
		}
		nativeAssetData, errNativeAssetData := utils.NativeTokenVaultTransferData(m.Amount, m.To, m.Token)
		if errNativeAssetData != nil {
			return nil, fmt.Errorf("failed to encode asset data: %w", err)
		}
		bridgeAbi, errBridgeAbi := l2assetrouter.IL2AssetRouterMetaData.GetAbi()
		if errBridgeAbi != nil {
			return nil, fmt.Errorf("failed to load l2AssetRouterAbi: %w", err)
		}
		data, errData := bridgeAbi.Pack("withdraw", nativeAssetId, nativeAssetData)
		if errData != nil {
			return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
		}

		return &ethereum.CallMsg{
			From:      m.From,
			To:        &utils.L2AssetRouterAddress,
			Gas:       m.Gas,
			GasPrice:  m.GasPrice,
			GasFeeCap: m.GasFeeCap,
			GasTipCap: m.GasTipCap,
			Value:     big.NewInt(0),
			Data:      data,
		}, nil
	}

	l2BridgeAbi, err := l2bridge.IL2BridgeMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to load l2BridgeAbi: %w", err)
	}
	data, err := l2BridgeAbi.Pack("withdraw", m.To, m.Token, m.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	return &ethereum.CallMsg{
		From:      m.From,
		To:        bridgeAddress,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     big.NewInt(0),
		Data:      data,
	}, nil
}

// ToCallMsg transforms WithdrawalCallMsg to types.CallMsg.
func (m *WithdrawalCallMsg) ToCallMsg(ctx context.Context, client *Client, defaultL2Bridge *common.Address) (*types.CallMsg, error) {
	msg, err := m.ToL1CallMsg(ctx, client, defaultL2Bridge)
	if err != nil {
		return nil, err
	}

	gasPerPubdata := m.GasPerPubdata
	if gasPerPubdata == nil {
		gasPerPubdata = utils.DefaultGasPerPubdataLimit
	}

	return &types.CallMsg{
		From:            msg.From,
		To:              msg.To,
		Gas:             msg.Gas,
		GasPrice:        msg.GasPrice,
		GasFeeCap:       msg.GasFeeCap,
		GasTipCap:       msg.GasTipCap,
		Value:           msg.Value,
		Data:            msg.Data,
		GasPerPubdata:   gasPerPubdata,
		PaymasterParams: m.PaymasterParams,
		CustomSignature: m.CustomSignature,
	}, nil
}

type blockMarshaling struct {
	ParentHash    common.Hash         `json:"parentHash"       gencodec:"required"`
	UncleHash     common.Hash         `json:"sha3Uncles"       gencodec:"required"`
	Coinbase      common.Address      `json:"miner"`
	Root          common.Hash         `json:"stateRoot"        gencodec:"required"`
	TxHash        common.Hash         `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash   common.Hash         `json:"receiptsRoot"     gencodec:"required"`
	Bloom         ethTypes.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty    *hexutil.Big        `json:"difficulty"       gencodec:"required"`
	Number        *hexutil.Big        `json:"number"           gencodec:"required"`
	GasLimit      hexutil.Uint64      `json:"gasLimit"         gencodec:"required"`
	GasUsed       hexutil.Uint64      `json:"gasUsed"          gencodec:"required"`
	Time          hexutil.Uint64      `json:"timestamp"        gencodec:"required"`
	Extra         hexutil.Bytes       `json:"extraData"        gencodec:"required"`
	MixDigest     common.Hash         `json:"mixHash"`
	Nonce         ethTypes.BlockNonce `json:"nonce"`
	BaseFee       *hexutil.Big        `json:"baseFeePerGas" rlp:"optional"`
	ExcessDataGas *hexutil.Big        `json:"excessDataGas" rlp:"optional"`

	Uncles           []*common.Hash `json:"uncles"`
	Hash             *common.Hash   `json:"hash"`
	L1BatchNumber    *hexutil.Big   `json:"l1BatchNumber"`
	L1BatchTimestamp *hexutil.Big   `json:"l1BatchTimestamp"`
	TotalDifficulty  *hexutil.Big   `json:"totalDifficulty"`
	Size             *hexutil.Big   `json:"size"`
	SealFields       []interface{}  `json:"sealFields"`

	Transactions []*types.TransactionResponse `json:"transactions"`
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
