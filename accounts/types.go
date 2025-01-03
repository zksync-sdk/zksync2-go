package accounts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/bridgehub"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l1sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/zksynchyperchain"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
)

// Cache holds cached bridge addresses and contracts for optimized access.
// This cache uses a lazy-loading approach, meaning data is only fetched when needed.
type Cache struct {
	clientL1 *ethclient.Client
	clientL2 *clients.Client

	bridgeAddresses     *types.BridgeContracts
	mainContractAddress *common.Address
	bridgehubAddress    *common.Address

	l2BridgeContracts *types.L2BridgeContracts
	l1BridgeContracts *types.L1BridgeContracts
	mainContract      *zksynchyperchain.IZkSyncHyperchain
	bridgehub         *bridgehub.IBridgehub
}

// NewCache creates an instance of Cache with the provided clients.
// Contracts returned from Cache are associated with the provided clients.
func NewCache(clientL2 *clients.Client, clientL1 *ethclient.Client) *Cache {
	return &Cache{clientL2: clientL2, clientL1: clientL1}
}

// MainContractAddress returns the main contract address from cache.
// If not already cached, it fetches the addresses,
// caches it, and then returns the value from the cache.
func (c *Cache) MainContractAddress() (common.Address, error) {
	if c.mainContractAddress == nil {
		mainContractAddress, err := c.clientL2.MainContractAddress(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.mainContractAddress = &mainContractAddress
	}
	return *c.mainContractAddress, nil
}

// BridgehubAddress returns the Bridgehub contract address from cache.
// If not already cached, it fetches the addresses,
// caches it, and then returns the value from the cache.
func (c *Cache) BridgehubAddress() (common.Address, error) {
	if c.bridgehubAddress == nil {
		bridgehubContractAddress, err := c.clientL2.BridgehubContractAddress(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.bridgehubAddress = &bridgehubContractAddress
	}
	return *c.bridgehubAddress, nil
}

// L2DefaultBridgeAddress returns the L2 default bridge address from cache.
// If not already cached, it fetches the bridge addresses,
// caches them, and then returns the value from the cache.
func (c *Cache) L2DefaultBridgeAddress() (common.Address, error) {
	if c.bridgeAddresses == nil {
		bridgeAddresses, err := c.clientL2.BridgeContracts(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.bridgeAddresses = bridgeAddresses
	}
	return c.bridgeAddresses.L2Erc20Bridge, nil
}

// L1DefaultBridgeAddress returns the L1 default bridge address from cache.
// If not already cached, it fetches the bridge addresses,
// caches them, and then returns the value from the cache.
func (c *Cache) L1DefaultBridgeAddress() (common.Address, error) {
	if c.bridgeAddresses == nil {
		bridgeAddresses, err := c.clientL2.BridgeContracts(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.bridgeAddresses = bridgeAddresses
	}
	return c.bridgeAddresses.L1Erc20Bridge, nil
}

// L2SharedBridgeAddress returns the L2 shared bridge address from cache.
// If not already cached, it fetches the bridge addresses,
// caches them, and then returns the value from the cache.
func (c *Cache) L2SharedBridgeAddress() (common.Address, error) {
	if c.bridgeAddresses == nil {
		bridgeAddresses, err := c.clientL2.BridgeContracts(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.bridgeAddresses = bridgeAddresses
	}
	return c.bridgeAddresses.L2SharedBridge, nil
}

// L1SharedBridgeAddress returns the L1  shared bridge address from cache.
// If not already cached, it fetches the bridge addresses,
// caches them, and then returns the value from the cache.
func (c *Cache) L1SharedBridgeAddress() (common.Address, error) {
	if c.bridgeAddresses == nil {
		bridgeAddresses, err := c.clientL2.BridgeContracts(context.Background())
		if err != nil {
			return common.Address{}, err
		}
		c.bridgeAddresses = bridgeAddresses
	}
	return c.bridgeAddresses.L1SharedBridge, nil
}

// MainContract returns the main contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) MainContract() (*zksynchyperchain.IZkSyncHyperchain, error) {
	if c.bridgehub == nil {
		mainContractAddress, err := c.MainContractAddress()
		if err != nil {
			return nil, err
		}
		c.mainContract, err = zksynchyperchain.NewIZkSyncHyperchain(mainContractAddress, c.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IBridgehub: %w", err)
		}
	}
	return c.mainContract, nil
}

// Bridgehub returns the Bridgehub contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) Bridgehub() (*bridgehub.IBridgehub, error) {
	if c.bridgehub == nil {
		bridgehubAddress, err := c.BridgehubAddress()
		if err != nil {
			return nil, err
		}
		c.bridgehub, err = bridgehub.NewIBridgehub(bridgehubAddress, c.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IBridgehub: %w", err)
		}
	}
	return c.bridgehub, nil
}

// L2BridgeContracts returns the L2 bridge contracts from cache.
// It is not cached, fetches the contracts, cache them and
// returns the value from cache.
func (c *Cache) L2BridgeContracts() (*types.L2BridgeContracts, error) {
	if c.l2BridgeContracts == nil {
		defaultBridge, err := c.L2DefaultBridge()
		if err != nil {
			return nil, err
		}
		sharedBridge, err := c.L2SharedBridge()
		if err != nil {
			return nil, err
		}
		c.l2BridgeContracts = &types.L2BridgeContracts{Erc20: defaultBridge, Shared: sharedBridge}
	}
	return c.l2BridgeContracts, nil
}

// L1BridgeContracts returns the L1 bridge contracts from cache.
// It is not cached, fetches the contracts, cache them and
// returns the value from cache.
func (c *Cache) L1BridgeContracts() (*types.L1BridgeContracts, error) {
	if c.l2BridgeContracts == nil {
		defaultBridge, err := c.L1DefaultBridge()
		if err != nil {
			return nil, err
		}
		sharedBridge, err := c.L1SharedBridge()
		if err != nil {
			return nil, err
		}
		c.l1BridgeContracts = &types.L1BridgeContracts{Erc20: defaultBridge, Shared: sharedBridge}
	}
	return c.l1BridgeContracts, nil
}

// L2DefaultBridge returns the L2 default bridge contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) L2DefaultBridge() (*l2bridge.IL2Bridge, error) {
	if c.l2BridgeContracts == nil || c.l2BridgeContracts.Erc20 == nil {
		c.l2BridgeContracts = &types.L2BridgeContracts{}
		defaultL2BridgeAddress, err := c.L2DefaultBridgeAddress()
		if err != nil {
			return nil, err
		}
		c.l2BridgeContracts.Erc20, err = l2bridge.NewIL2Bridge(defaultL2BridgeAddress, c.clientL2)
		if err != nil {
			return nil, fmt.Errorf("failed to load IL2Bridge: %w", err)
		}
	}
	return c.l2BridgeContracts.Erc20, nil
}

// L1DefaultBridge returns the L1 default bridge contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) L1DefaultBridge() (*l1bridge.IL1Bridge, error) {
	if c.l1BridgeContracts == nil || c.l1BridgeContracts.Erc20 == nil {
		c.l1BridgeContracts = &types.L1BridgeContracts{}
		defaultL1BridgeAddress, err := c.L1DefaultBridgeAddress()
		if err != nil {
			return nil, err
		}
		c.l1BridgeContracts.Erc20, err = l1bridge.NewIL1Bridge(defaultL1BridgeAddress, c.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IL1Bridge: %w", err)
		}
	}
	return c.l1BridgeContracts.Erc20, nil
}

// L2SharedBridge returns the L2 shared bridge contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) L2SharedBridge() (*l2sharedbridge.IL2SharedBridge, error) {
	if c.l2BridgeContracts == nil || c.l2BridgeContracts.Shared == nil {
		c.l2BridgeContracts = &types.L2BridgeContracts{}
		sharedL2BridgeAddress, err := c.L2SharedBridgeAddress()
		if err != nil {
			return nil, err
		}
		c.l2BridgeContracts.Shared, err = l2sharedbridge.NewIL2SharedBridge(sharedL2BridgeAddress, c.clientL2)
		if err != nil {
			return nil, fmt.Errorf("failed to load IL2SharedBridge: %w", err)
		}
	}
	return c.l2BridgeContracts.Shared, nil
}

// L1SharedBridge returns the L2 shared bridge contract from cache.
// It is not cached, fetches the contract, cache it and
// returns the value from cache.
func (c *Cache) L1SharedBridge() (*l1sharedbridge.IL1SharedBridge, error) {
	if c.l1BridgeContracts == nil || c.l1BridgeContracts.Shared == nil {
		c.l1BridgeContracts = &types.L1BridgeContracts{}
		sharedL1BridgeAddress, err := c.L1SharedBridgeAddress()
		if err != nil {
			return nil, err
		}
		c.l1BridgeContracts.Shared, err = l1sharedbridge.NewIL1SharedBridge(sharedL1BridgeAddress, c.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IL1SharedBridge: %w", err)
		}
	}
	return c.l1BridgeContracts.Shared, nil
}

// CallOpts is the collection of options to fine tune a contract call request from
// a specific associated account.
// Its primary purpose is to be transformed into bind.CallOpts, wherein the 'From'
// field represents the associated account.
type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// ToCallOpts transforms CallOpts to bind.CallOpts.
func (o *CallOpts) ToCallOpts(from common.Address) *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     o.Pending,
		From:        from,
		BlockNumber: o.BlockNumber,
		Context:     ensureContext(o.Context),
	}
}

// CallMsg contains parameters for contract call from a specific account associated
// with WalletL1, WalletL2 or Deployer.
// Its primary purpose is to be transformed into types.CallMsg, wherein the 'From'
// field represents the associated account.
type CallMsg struct {
	To        *common.Address // The address of the recipient.
	Gas       uint64          // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int        // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value     *big.Int        // Amount of wei sent along with the call.
	Data      []byte          // Input data, usually an ABI-encoded contract method invocation

	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes
	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
	// it should contain the bytecode of the contract being deployed.
	// If the contract is a factory contract, i.e. it can deploy other contracts,
	// the array should also contain the bytecodes of the contracts which it can deploy.
	FactoryDeps []hexutil.Bytes
	// PaymasterParams contains parameters for configuring the custom paymaster
	// for the transaction.
	PaymasterParams *types.PaymasterParams
}

// ToCallMsg transforms CallMsg to types.CallMsg.
func (m *CallMsg) ToCallMsg(from common.Address) types.CallMsg {
	return types.CallMsg{
		From:            from,
		To:              m.To,
		Gas:             m.Gas,
		GasPrice:        m.GasPrice,
		GasFeeCap:       m.GasFeeCap,
		GasTipCap:       m.GasTipCap,
		Value:           m.Value,
		Data:            m.Data,
		GasPerPubdata:   m.GasPerPubdata,
		CustomSignature: m.CustomSignature,
		FactoryDeps:     m.FactoryDeps,
		PaymasterParams: m.PaymasterParams,
	}
}

// WithdrawalCallMsg contains the common data required to execute a withdrawal call on L1 from L2.
// This execution is initiated by the account associated with WalletL2.
type WithdrawalCallMsg struct {
	To            common.Address  // The address of the recipient on L1.
	Amount        *big.Int        // The amount of the token to transfer.
	Token         common.Address  // The address of the token. ETH by default.
	BridgeAddress *common.Address // The address of the bridge contract to be used.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	PaymasterParams *types.PaymasterParams // The paymaster parameters.
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
}

// ToWithdrawalCallMsg transforms WithdrawalCallMsg to clients.WithdrawalCallMsg.
func (m *WithdrawalCallMsg) ToWithdrawalCallMsg(from common.Address) clients.WithdrawalCallMsg {
	return clients.WithdrawalCallMsg{
		To:              m.To,
		Amount:          m.Amount,
		Token:           m.Token,
		BridgeAddress:   m.BridgeAddress,
		From:            from,
		Gas:             m.Gas,
		GasPrice:        m.GasPrice,
		GasFeeCap:       m.GasFeeCap,
		GasTipCap:       m.GasTipCap,
		PaymasterParams: m.PaymasterParams,
		GasPerPubdata:   m.GasPerPubdata,
	}
}

// TransferCallMsg contains the common data required to execute a transfer call on L2.
// This execution is initiated by the account associated with WalletL2.
type TransferCallMsg struct {
	To     common.Address // The address of the recipient.
	Amount *big.Int       // The amount of the token to transfer.
	Token  common.Address // The address of the token. ETH by default.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	PaymasterParams *types.PaymasterParams // The paymaster parameters.
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
}

// ToTransferCallMsg transforms TransferCallMsg to clients.TransferCallMsg.
func (m *TransferCallMsg) ToTransferCallMsg(from common.Address) clients.TransferCallMsg {
	return clients.TransferCallMsg{
		To:              m.To,
		Amount:          m.Amount,
		Token:           m.Token,
		From:            from,
		Gas:             m.Gas,
		GasPrice:        m.GasPrice,
		GasFeeCap:       m.GasFeeCap,
		GasTipCap:       m.GasTipCap,
		PaymasterParams: m.PaymasterParams,
		GasPerPubdata:   m.GasPerPubdata,
	}
}

// DepositCallMsg contains the common data required to execute a deposit call on L2 from L1.
// This execution is initiated by the account associated with WalletL1.
type DepositCallMsg struct {
	To     common.Address // The address that will receive the deposited tokens on L2.
	Token  common.Address // The address of the token to deposit.
	Amount *big.Int       // The amount of the token to be deposited.

	// If the ETH value passed with the transaction is not explicitly stated Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	BridgeAddress *common.Address // The address of the bridge contract to be used.

	L2GasLimit *big.Int // Maximum amount of L2 gas that transaction can consume during execution on L2.

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	CustomBridgeData []byte // Additional data that can be sent to a bridge.

	Value     *big.Int // The amount of wei sent along with the call.
	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.
}

// ToDepositTransaction transforms DepositCallMsg to DepositTransaction.
func (m *DepositCallMsg) ToDepositTransaction() DepositTransaction {
	return DepositTransaction{
		Token:             m.Token,
		Amount:            m.Amount,
		To:                m.To,
		OperatorTip:       m.OperatorTip,
		BridgeAddress:     m.BridgeAddress,
		L2GasLimit:        m.L2GasLimit,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
		CustomBridgeData:  m.CustomBridgeData,
	}
}

// ToRequestExecuteCallMsg transforms DepositCallMsg to RequestExecuteCallMsg.
func (m *DepositCallMsg) ToRequestExecuteCallMsg() RequestExecuteCallMsg {
	return RequestExecuteCallMsg{
		ContractAddress:   m.To,
		L2GasLimit:        m.L2GasLimit,
		L2Value:           m.Amount,
		OperatorTip:       m.OperatorTip,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
		Value:             m.Value,
		Gas:               m.Gas,
		GasPrice:          m.GasPrice,
		GasFeeCap:         m.GasFeeCap,
		GasTipCap:         m.GasTipCap,
	}
}

// ToCallMsg transforms DepositCallMsg to ethereum.CallMsg.
func (m *DepositCallMsg) ToCallMsg(from, l1Bridge common.Address) (ethereum.CallMsg, error) {
	l1BridgeAbi, err := l1bridge.IL1BridgeMetaData.GetAbi()
	if err != nil {
		return ethereum.CallMsg{}, fmt.Errorf("failed to load IL1Bridge ABI: %w", err)
	}

	calldata, err := l1BridgeAbi.Pack("deposit", m.To, m.Token, m.Amount, m.L2GasLimit, m.GasPerPubdataByte, m.RefundRecipient)
	if err != nil {
		return ethereum.CallMsg{}, err
	}

	bridge := l1Bridge
	if m.BridgeAddress != nil {
		bridge = *m.BridgeAddress
	}

	return ethereum.CallMsg{
		From:      from,
		To:        &bridge,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     m.Value,
		Data:      calldata,
	}, nil
}

// ToTransactOpts transforms DepositCallMsg to TransactOptsL1.
func (m *DepositCallMsg) ToTransactOpts() TransactOptsL1 {
	return TransactOptsL1{
		Value:     m.Value,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		GasLimit:  m.Gas,
	}
}

// PopulateEmptyFields populates required empty fields with default values.
func (m *DepositCallMsg) PopulateEmptyFields(from common.Address) {
	if m.To == (common.Address{}) {
		m.To = from
	}
	if m.OperatorTip == nil {
		m.OperatorTip = big.NewInt(0)
	}
	if m.GasPerPubdataByte == nil {
		m.GasPerPubdataByte = utils.RequiredL1ToL2GasPerPubdataLimit
	}
	if m.Token == (common.Address{}) {
		m.Token = utils.LegacyEthAddress
	}
}

// RequestExecuteCallMsg contains the common data required to execute a call for a request execution of an L2
// transaction from L1. This execution is initiated by the account associated with WalletL1.
type RequestExecuteCallMsg struct {
	ContractAddress common.Address // The L2 receiver address.
	Calldata        []byte         // The input of the L2 transaction.
	L2GasLimit      *big.Int       // Maximum amount of L2 gas that transaction can consume during execution on L2.
	MintValue       *big.Int       // The amount of base token that needs to be minted on non-ETH-based L2.
	L2Value         *big.Int       // `msg.value` of L2 transaction.
	FactoryDeps     [][]byte       // An array of L2 bytecodes that will be marked as known on L2.

	// If the ETH value passed with the transaction is not explicitly stated Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	Value     *big.Int // The amount of wei sent along with the call.
	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.
}

// ToRequestExecuteTransaction transforms RequestExecuteCallMsg to RequestExecuteTransaction.
func (m *RequestExecuteCallMsg) ToRequestExecuteTransaction() RequestExecuteTransaction {
	return RequestExecuteTransaction{
		ContractAddress:   m.ContractAddress,
		Calldata:          m.Calldata,
		L2GasLimit:        m.L2GasLimit,
		MintValue:         m.MintValue,
		L2Value:           m.L2Value,
		FactoryDeps:       m.FactoryDeps,
		OperatorTip:       m.OperatorTip,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
	}
}

// ToCallMsgWithChainID transforms RequestExecuteCallMsg to ethereum.CallMsg.
func (m *RequestExecuteCallMsg) ToCallMsgWithChainID(from common.Address, chainID *big.Int) (ethereum.CallMsg, error) {
	bridgehubAbi, err := bridgehub.IBridgehubMetaData.GetAbi()
	if err != nil {
		return ethereum.CallMsg{}, fmt.Errorf("failed to load IZkSync ABI: %w", err)
	}
	requestExecuteCalldata, err := bridgehubAbi.Pack("requestL2TransactionDirect", bridgehub.L2TransactionRequestDirect{
		ChainId:                  chainID,
		MintValue:                m.MintValue,
		L2Contract:               m.ContractAddress,
		L2Value:                  m.L2Value,
		L2Calldata:               m.Calldata,
		L2GasLimit:               m.L2GasLimit,
		L2GasPerPubdataByteLimit: m.GasPerPubdataByte,
		FactoryDeps:              m.FactoryDeps,
		RefundRecipient:          m.RefundRecipient,
	})
	if err != nil {
		return ethereum.CallMsg{}, err
	}

	return ethereum.CallMsg{
		From:      from,
		To:        &m.ContractAddress,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     m.Value,
		Data:      requestExecuteCalldata,
	}, nil
}

// ToTransactOpts transforms RequestExecuteCallMsg to TransactOptsL1.
func (m *RequestExecuteCallMsg) ToTransactOpts() TransactOptsL1 {
	return TransactOptsL1{
		Value:     m.Value,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		GasLimit:  m.Gas,
	}
}

// TransactOptsL1 contains common data required to create a valid transaction on L1
// using the account associated with WalletL1.
// Its primary purpose is to be transformed into bind.TransactOpts, wherein the 'From' and 'Signer'
// fields are linked to the associated account.
type TransactOptsL1 struct {
	Nonce     *big.Int        // Nonce to use for the transaction execution (nil = use pending state).
	Value     *big.Int        // Funds to transfer along the transaction (nil = 0 = no funds).
	GasPrice  *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle).
	GasFeeCap *big.Int        // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasTipCap *big.Int        // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasLimit  uint64          // Gas limit to set for the transaction execution (0 = estimate).
	Context   context.Context // Network context to support cancellation and timeouts (nil = no timeout).
}

// ToTransactOpts transforms TransactOptsL1 to bind.TransactOpts.
func (t *TransactOptsL1) ToTransactOpts(from common.Address, signer bind.SignerFn) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:      from,
		Nonce:     t.Nonce,
		Signer:    signer,
		Value:     t.Value,
		GasPrice:  t.GasPrice,
		GasFeeCap: t.GasFeeCap,
		GasTipCap: t.GasTipCap,
		GasLimit:  t.GasLimit,
		Context:   t.Context,
	}
}

// TransactOpts contains common data required to create a valid transaction on L2
// using the account associated with WalletL2 and Deployer.
// Its primary purpose is to enrich bind.TransactOpts with ZKsync features,
// wherein the 'From' and 'Signer' fields are linked to the associated account.
type TransactOpts struct {
	Nonce     *big.Int        // Nonce to use for the transaction execution (nil = use pending state).
	Value     *big.Int        // Funds to transfer along the transaction (nil = 0 = no funds).
	GasPrice  *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle).
	GasFeeCap *big.Int        // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasTipCap *big.Int        // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasLimit  uint64          // Gas limit to set for the transaction execution (0 = estimate).
	Context   context.Context // Network context to support cancellation and timeouts (nil = no timeout).

	PaymasterParams *types.PaymasterParams // The paymaster parameters.
	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int
}

// Transaction is similar to types.Transaction but does not include the From field. This design is intended for use
// with entities which already have an associated account. The From field is bound to
// their associated account, and thus, it is not included in this type.
type Transaction struct {
	To        *common.Address // The address of the recipient.
	Data      hexutil.Bytes   // Input data, usually an ABI-encoded contract method invocation.
	Value     *big.Int        // Funds to transfer along the transaction (nil = 0 = no funds).
	Nonce     *big.Int        // Nonce to use for the transaction execution.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	Gas       uint64          // Gas limit to set for the transaction execution.

	ChainID *big.Int // Chain ID of the network.

	// GasPerPubdata denotes the maximum amount of gas the user is willing
	// to pay for a single byte of pubdata.
	GasPerPubdata *big.Int `json:"gasPerPubdata,omitempty"`
	// CustomSignature is used for the cases in which the signer's account
	// is not an EOA.
	CustomSignature hexutil.Bytes `json:"customSignature,omitempty"`
	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
	// it should contain the bytecode of the contract being deployed.
	// If the contract is a factory contract, i.e. it can deploy other contracts,
	// the array should also contain the bytecodes of the contracts which it can deploy.
	FactoryDeps []hexutil.Bytes `json:"factoryDeps"`
	// PaymasterParams contains parameters for configuring the custom paymaster
	// for the transaction.
	PaymasterParams *types.PaymasterParams `json:"paymasterParams,omitempty"`
}

// ToTransaction transforms Transaction to types.Transaction.
func (t *Transaction) ToTransaction(from common.Address) *types.Transaction {
	return &types.Transaction{
		Nonce:           t.Nonce,
		GasTipCap:       t.GasTipCap,
		GasFeeCap:       t.GasFeeCap,
		Gas:             new(big.Int).SetUint64(t.Gas),
		To:              t.To,
		Value:           t.Value,
		Data:            t.Data,
		ChainID:         t.ChainID,
		From:            &from,
		GasPerPubdata:   t.GasPerPubdata,
		CustomSignature: t.CustomSignature,
		FactoryDeps:     t.FactoryDeps,
		PaymasterParams: t.PaymasterParams,
	}
}

// ToCallMsg transforms Transaction to types.CallMsg.
func (t *Transaction) ToCallMsg(from common.Address) types.CallMsg {
	return types.CallMsg{
		From:            from,
		To:              t.To,
		Gas:             t.Gas,
		GasFeeCap:       t.GasFeeCap,
		GasTipCap:       t.GasTipCap,
		Value:           t.Value,
		Data:            t.Data,
		GasPerPubdata:   t.GasPerPubdata,
		CustomSignature: t.CustomSignature,
		FactoryDeps:     t.FactoryDeps,
		PaymasterParams: t.PaymasterParams,
	}
}

// TransferTransaction represents a transfer transaction on L2
// initiated by the account associated with WalletL2.
type TransferTransaction struct {
	To     common.Address // The address of the recipient.
	Amount *big.Int       // The amount of the token to transfer.
	Token  common.Address // The address of the token. ETH by default.
}

// ToTransaction transforms TransferTransaction to Transaction.
func (t *TransferTransaction) ToTransaction(opts *TransactOpts) *Transaction {
	return &Transaction{
		To:              &t.To,
		Value:           t.Amount,
		Nonce:           opts.Nonce,
		GasFeeCap:       opts.GasFeeCap,
		GasTipCap:       opts.GasTipCap,
		Gas:             opts.GasLimit,
		GasPerPubdata:   opts.GasPerPubdata,
		PaymasterParams: opts.PaymasterParams,
	}
}

// ToTransferCallMsg transforms TransferTransaction to clients.TransferCallMsg.
func (t *TransferTransaction) ToTransferCallMsg(from common.Address, opts *TransactOpts) clients.TransferCallMsg {
	return clients.TransferCallMsg{
		To:              t.To,
		Amount:          t.Amount,
		Token:           t.Token,
		From:            from,
		Gas:             opts.GasLimit,
		GasPrice:        opts.GasPrice,
		GasFeeCap:       opts.GasFeeCap,
		GasTipCap:       opts.GasTipCap,
		GasPerPubdata:   opts.GasPerPubdata,
		PaymasterParams: opts.PaymasterParams,
	}
}

// WithdrawalTransaction represents a withdrawal transaction on L1 from L2
// initiated by the account associated with WalletL2.
type WithdrawalTransaction struct {
	To            common.Address  // The address that will receive the withdrawn tokens on L1.
	Token         common.Address  // The address of the token to withdraw.
	Amount        *big.Int        // The amount of the token to withdraw.
	BridgeAddress *common.Address // The address of the bridge contract to be used.
}

// ToWithdrawalCallMsg transforms WithdrawalTransaction to clients.WithdrawalCallMsg.
func (t *WithdrawalTransaction) ToWithdrawalCallMsg(from common.Address, opts *TransactOpts) *clients.WithdrawalCallMsg {
	return &clients.WithdrawalCallMsg{
		To:              t.To,
		Amount:          t.Amount,
		Token:           t.Token,
		BridgeAddress:   t.BridgeAddress,
		From:            from,
		Gas:             opts.GasLimit,
		GasPrice:        opts.GasPrice,
		GasFeeCap:       opts.GasFeeCap,
		GasTipCap:       opts.GasTipCap,
		GasPerPubdata:   opts.GasPerPubdata,
		PaymasterParams: opts.PaymasterParams,
	}
}

// RequestExecuteTransaction represents a request execution of L2 transaction from L1
// initiated by the account associated with WalletL1.
type RequestExecuteTransaction struct {
	ContractAddress common.Address // The L2 receiver address.
	Calldata        []byte         // The input of the L2 transaction.
	L2GasLimit      *big.Int       // Maximum amount of L2 gas that transaction can consume during execution on L2.
	MintValue       *big.Int       // The amount of base token that needs to be minted on non-ETH-based L2.
	L2Value         *big.Int       // `msg.value` of L2 transaction.
	FactoryDeps     [][]byte       // An array of L2 bytecodes that will be marked as known on L2.

	// If the ETH value passed with the transaction is not explicitly stated Auth.Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address
}

// ToRequestExecuteCallMsg transforms RequestExecuteTransaction to RequestExecuteCallMsg.
func (t *RequestExecuteTransaction) ToRequestExecuteCallMsg(opts *TransactOptsL1) RequestExecuteCallMsg {
	return RequestExecuteCallMsg{
		ContractAddress:   t.ContractAddress,
		Calldata:          t.Calldata,
		L2GasLimit:        t.L2GasLimit,
		MintValue:         t.MintValue,
		L2Value:           t.L2Value,
		FactoryDeps:       t.FactoryDeps,
		OperatorTip:       t.OperatorTip,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
		Value:             opts.Value,
		GasPrice:          opts.GasPrice,
		GasFeeCap:         opts.GasFeeCap,
		GasTipCap:         opts.GasTipCap,
	}
}

// ToCallMsg transforms RequestExecuteTransaction to types.CallMsg.
func (t *RequestExecuteTransaction) ToCallMsg(from common.Address, opts *TransactOptsL1) types.CallMsg {
	factoryDeps := make([]hexutil.Bytes, len(t.FactoryDeps))
	if len(t.FactoryDeps) > 0 {
		for i, d := range t.FactoryDeps {
			factoryDeps[i] = d
		}
	}
	return types.CallMsg{
		From:          from,
		To:            &t.ContractAddress,
		Gas:           opts.GasLimit,
		GasPrice:      opts.GasPrice,
		Value:         opts.Value,
		Data:          t.Calldata,
		GasPerPubdata: t.GasPerPubdataByte,
		FactoryDeps:   factoryDeps,
	}
}

// DepositTransaction represents a deposit transaction on L2 from L1
// initiated by the account associated with WalletL1.
type DepositTransaction struct {
	To     common.Address // The address of the token to deposit.
	Token  common.Address // The address of the token to deposit.
	Amount *big.Int       // The amount of the token to be deposited.

	// If the ETH value passed with the transaction is not explicitly stated Auth.Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	BridgeAddress *common.Address // The address of the bridge contract to be used.

	// Whether should the token approval be performed under the hood. Set this flag to true if you
	// bridge an ERC20 token and didn't call the ApproveToken function beforehand.
	ApproveToken bool

	// Whether should the base token approval be performed under the hood. Set this flag to true if you
	// bridge an ERC20 token and didn't call the ApproveToken function beforehand.
	ApproveBaseToken bool

	L2GasLimit *big.Int // Maximum amount of L2 gas that transaction can consume during execution on L2.

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	CustomBridgeData []byte // Additional data that can be sent to the bridge.

	ApproveAuth     *TransactOptsL1 // Authorization data for the token approval transaction.
	ApproveBaseAuth *TransactOptsL1 // Authorization data for the base token approval transaction.

}

// ToRequestExecuteTransaction transforms DepositTransaction to RequestExecuteTransaction.
func (t *DepositTransaction) ToRequestExecuteTransaction() *RequestExecuteTransaction {
	return &RequestExecuteTransaction{
		ContractAddress:   t.To,
		L2Value:           t.Amount,
		L2GasLimit:        t.L2GasLimit,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
	}
}

// ToDepositCallMsg transforms DepositTransaction to DepositCallMsg.
func (t *DepositTransaction) ToDepositCallMsg(opts *TransactOptsL1) DepositCallMsg {
	return DepositCallMsg{
		To:                t.To,
		Token:             t.Token,
		Amount:            t.Amount,
		OperatorTip:       t.OperatorTip,
		BridgeAddress:     t.BridgeAddress,
		L2GasLimit:        t.L2GasLimit,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
		CustomBridgeData:  t.CustomBridgeData,
		Value:             opts.Value,
		GasPrice:          opts.GasPrice,
		GasFeeCap:         opts.GasFeeCap,
		GasTipCap:         opts.GasTipCap,
	}
}

// PopulateEmptyFields populates required empty fields with default values.
func (t *DepositTransaction) PopulateEmptyFields(from common.Address) {
	if t.To == (common.Address{}) {
		t.To = from
	}
	if t.OperatorTip == nil {
		t.OperatorTip = big.NewInt(0)
	}
	if t.GasPerPubdataByte == nil {
		t.GasPerPubdataByte = utils.RequiredL1ToL2GasPerPubdataLimit
	}
	if t.Token == (common.Address{}) {
		t.Token = utils.LegacyEthAddress
	}
	if t.ApproveToken && t.ApproveAuth == nil {
		t.ApproveAuth = ensureTransactOptsL1(t.ApproveAuth)
	}
	if t.ApproveBaseToken && t.ApproveBaseAuth == nil {
		t.ApproveAuth = ensureTransactOptsL1(t.ApproveBaseAuth)
	}
}

// DeploymentType represents an enumeration of deployment types.
type DeploymentType string

const (
	DeployContract DeploymentType = "CONTRACT"
	DeployAccount  DeploymentType = "ACCOUNT"
)

// CreateTransaction represents the parameters for deploying a contract or account using the CREATE opcode.
// This transaction is initiated by the account associated with Deployer.
type CreateTransaction struct {
	Bytecode     []byte   // The bytecode of smart contract or smart account.
	Calldata     []byte   // The constructor calldata.
	Dependencies [][]byte // The bytecode of dependent smart contracts or smart accounts.
}

// ToTransaction transforms CreateTransaction to Transaction.
func (t *CreateTransaction) ToTransaction(deploymentType DeploymentType, opts *TransactOpts) (*Transaction, error) {
	var (
		data []byte
		err  error
	)
	if deploymentType == DeployContract {
		data, err = utils.EncodeCreate(t.Bytecode, t.Calldata)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create call: %w", err)
		}
	} else {
		data, err = utils.EncodeCreateAccount(t.Bytecode, t.Calldata, types.Version1)
		if err != nil {
			return nil, fmt.Errorf("failed to encode createAccount call: %w", err)
		}

	}

	var factoryDeps []hexutil.Bytes
	if len(t.Dependencies) > 0 {
		factoryDeps = make([]hexutil.Bytes, len(t.Dependencies)+1)
		for i, d := range t.Dependencies {
			factoryDeps[i] = d
		}
		factoryDeps[len(t.Dependencies)] = t.Bytecode
	} else {
		factoryDeps = []hexutil.Bytes{t.Bytecode}
	}

	auth := opts
	if auth == nil {
		auth = &TransactOpts{Context: context.Background()}
	}

	gasPerPubdata := auth.GasPerPubdata
	if gasPerPubdata == nil {
		gasPerPubdata = utils.DefaultGasPerPubdataLimit
	}

	return &Transaction{
		To:              &utils.ContractDeployerAddress,
		Data:            data,
		Nonce:           auth.Nonce,
		GasFeeCap:       auth.GasFeeCap,
		GasTipCap:       auth.GasTipCap,
		Gas:             auth.GasLimit,
		GasPerPubdata:   gasPerPubdata,
		PaymasterParams: auth.PaymasterParams,
		FactoryDeps:     factoryDeps,
	}, nil
}

// Create2Transaction represents the parameters for deploying a contract or account using the CREATE2 opcode.
// This transaction is initiated by the account associated with Deployer.
type Create2Transaction struct {
	Bytecode     []byte   // The bytecode of smart contract or smart account.
	Calldata     []byte   // The constructor calldata.
	Salt         []byte   // The create2 salt.
	Dependencies [][]byte // The bytecode of dependent smart contracts or smart accounts.
}

// ToTransaction transforms Create2Transaction to Transaction.
func (t *Create2Transaction) ToTransaction(deploymentType DeploymentType, opts *TransactOpts) (*Transaction, error) {
	var (
		data []byte
		err  error
	)

	if deploymentType == DeployContract {
		data, err = utils.EncodeCreate2(t.Bytecode, t.Calldata, t.Salt)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create2 call: %w", err)
		}
	} else {
		data, err = utils.EncodeCreate2Account(t.Bytecode, t.Calldata, t.Salt, types.Version1)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create2Account call: %w", err)
		}
	}

	var factoryDeps []hexutil.Bytes
	if len(t.Dependencies) > 0 {
		factoryDeps = make([]hexutil.Bytes, len(t.Dependencies)+1)
		for i, d := range t.Dependencies {
			factoryDeps[i] = d
		}
		factoryDeps[len(t.Dependencies)] = t.Bytecode
	} else {
		factoryDeps = []hexutil.Bytes{t.Bytecode}
	}

	auth := opts
	if auth == nil {
		auth = &TransactOpts{Context: context.Background()}
	}

	gasPerPubdata := auth.GasPerPubdata
	if gasPerPubdata == nil {
		gasPerPubdata = utils.DefaultGasPerPubdataLimit
	}

	return &Transaction{
		To:              &utils.ContractDeployerAddress,
		Data:            data,
		Nonce:           auth.Nonce,
		GasFeeCap:       auth.GasFeeCap,
		GasTipCap:       auth.GasTipCap,
		Gas:             auth.GasLimit,
		GasPerPubdata:   gasPerPubdata,
		PaymasterParams: auth.PaymasterParams,
		FactoryDeps:     factoryDeps,
	}, nil
}

// FullDepositFee represents the total ETH fee required for performing the deposit on
// both L1 and L2 networks.
type FullDepositFee struct {
	MaxFeePerGas, // MaxFeePerGas of the L1 transaction if 1559 transaction is used.
	MaxPriorityFeePerGas, // MaxPriorityFeePerGas of the L1 transaction if 1559 transaction is used.
	GasPrice, // Gas price of the L1 transaction if legacy transaction is used.
	BaseCost, // Base cost of the L2 transaction.
	L1GasLimit, // Gas limit of the L1 transaction.
	L2GasLimit *big.Int // Gas limit of the L2 transaction.
}

// AllowanceParams contains the parameters required for approval of an ERC20 token.
type AllowanceParams struct {
	Token     common.Address // Token address
	Allowance *big.Int       // Allowance amount
}

// PayloadSigner signs various types of payloads, optionally using a some kind of secret.
// Returns the serialized signature.
// The clientL2 is used to fetch data from the network if it is required for signing.
type PayloadSigner func(ctx context.Context, payload []byte, secret interface{}, client *clients.Client) ([]byte, error)

// TransactionBuilder populates missing fields in a tx with default values.
// The clientL2 is used to fetch data from the network if it is required.
type TransactionBuilder func(ctx context.Context, tx *types.Transaction, secret interface{}, client *clients.Client) error
