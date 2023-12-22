package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"testing"
)

func TestIntegrationWallet_MainContract(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	mainContract, err := wallet.MainContract(context.Background())

	assert.NoError(t, err, "MainContract should not return an error")
	assert.NotNil(t, mainContract, "MainContract should return a non-nil address")
}

func TestIntegrationWallet_L1BridgeContracts(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	contracts, err := wallet.L1BridgeContracts(context.Background())

	assert.NoError(t, err, "L1BridgeContracts should not return an error")
	assert.NotNil(t, contracts, "L1BridgeContracts should return a non-nil bridge contracts")
}

func TestIntegrationWallet_BalanceL1(t *testing.T) {
	tokenData := readToken()

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	ethBalance, err := wallet.BalanceL1(nil, utils.EthAddress)

	assert.NoError(t, err, "BalanceL1 should not return an error")
	assert.NotNil(t, ethBalance, "BalanceL1 should return a non-nil ETH balance")

	tokenBalance, err := wallet.BalanceL1(nil, common.HexToAddress(tokenData.L1Address))

	assert.NoError(t, err, "BalanceL1 should not return an error")
	assert.NotNil(t, tokenBalance, "BlockByNumber should return a non-nil token balance")
}

func TestIntegrationWallet_AllowanceL1(t *testing.T) {
	tokenData := readToken()

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	bridgeContracts, err := client.BridgeContracts(context.Background())
	assert.NoError(t, err, "BridgeContracts should not return an error")

	ethBalance, err := wallet.AllowanceL1(nil, common.HexToAddress(tokenData.L1Address), bridgeContracts.L1Erc20DefaultBridge)

	assert.NoError(t, err, "AllowanceL1 should not return an error")
	assert.NotNil(t, ethBalance, "AllowanceL1 should return a non-nil token allowance")
}

func TestIntegrationWallet_L2TokenAddress(t *testing.T) {
	tokenData := readToken()

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2Address, err := wallet.L2TokenAddress(context.Background(), common.HexToAddress(tokenData.L1Address))

	assert.NoError(t, err, "L2TokenAddress should not return an error")
	assert.Equal(t, common.HexToAddress(tokenData.L2Address), l2Address, "L2 token addresses should be the same")
}

func TestIntegrationWallet_ApproveERC20(t *testing.T) {
	tokenData := readToken()
	l1TokenAddress := common.HexToAddress(tokenData.L1Address)
	approveAmount := big.NewInt(1)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	bridgeContracts, err := client.BridgeContracts(context.Background())
	assert.NoError(t, err, "BridgeContracts should not return an error")

	allowanceBefore, err := wallet.AllowanceL1(nil, l1TokenAddress, bridgeContracts.L1Erc20DefaultBridge)
	assert.NoError(t, err, "AllowanceL1 should not return an error")

	tx, err := wallet.ApproveERC20(nil, common.HexToAddress(tokenData.L1Address), approveAmount, bridgeContracts.L1Erc20DefaultBridge)
	assert.NoError(t, err, "ApproveERC20 should not return an error")

	txReceipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")

	allowanceAfter, err := wallet.AllowanceL1(nil, l1TokenAddress, bridgeContracts.L1Erc20DefaultBridge)
	assert.NoError(t, err, "AllowanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(allowanceAfter, allowanceBefore).Int64() == 1, "Allowance must be increased")
}

func TestIntegrationWallet_BaseCost(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	baseCost, err := wallet.BaseCost(nil, big.NewInt(100_000), big.NewInt(800), big.NewInt(500_000))

	assert.NoError(t, err, "BaseCost should not return an error")
	assert.True(t, baseCost.Cmp(big.NewInt(0)) > 0, "BaseCost should return a positive number")
}

func TestIntegrationWallet_DepositETH(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeDeposit, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceBeforeDeposit, err := wallet.BalanceL1(nil, utils.EthAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	tx, err := wallet.Deposit(nil, accounts.DepositTransaction{
		To:              wallet.Address(),
		Token:           utils.EthAddress,
		Amount:          amount,
		RefundRecipient: wallet.Address(),
	})
	assert.NoError(t, err, "Deposit should not return an error")

	l1Receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")

	l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
	assert.NoError(t, err, "L2TransactionFromPriorityOp should not return an error")

	l2Receipt, err := client.WaitMined(context.Background(), l2Tx.Hash)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, l2Receipt.BlockHash, "Transaction should be mined")

	l2BalanceAfterDeposit, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceAfterDeposit, err := wallet.BalanceL1(nil, utils.EthAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceAfterDeposit, l2BalanceBeforeDeposit).Cmp(amount) >= 0, "Balance on L2 should be increased")
	assert.True(t, new(big.Int).Sub(l1BalanceBeforeDeposit, l1BalanceAfterDeposit).Cmp(amount) >= 0, "Balance on L1 should be decreased")
}

func TestIntegrationWallet_DepositToken(t *testing.T) {
	tokenData := readToken()
	amount := big.NewInt(5)
	l1TokenAddress := common.HexToAddress(tokenData.L1Address)
	l2TokenAddress := common.HexToAddress(tokenData.L2Address)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeDeposit, err := wallet.Balance(context.Background(), l2TokenAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceBeforeDeposit, err := wallet.BalanceL1(nil, l1TokenAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	tx, err := wallet.Deposit(nil, accounts.DepositTransaction{
		To:              wallet.Address(),
		Token:           l1TokenAddress,
		Amount:          amount,
		ApproveERC20:    true,
		RefundRecipient: wallet.Address(),
	})
	assert.NoError(t, err, "Deposit should not return an error")

	l1Receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")

	l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
	assert.NoError(t, err, "L2TransactionFromPriorityOp should not return an error")

	l2Receipt, err := client.WaitMined(context.Background(), l2Tx.Hash)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, l2Receipt.BlockHash, "Transaction should be mined")

	l2BalanceAfterDeposit, err := wallet.Balance(context.Background(), l2TokenAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceAfterDeposit, err := wallet.BalanceL1(nil, l1TokenAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceAfterDeposit, l2BalanceBeforeDeposit).Cmp(amount) >= 0, "Balance on L2 should be increased")
	assert.True(t, new(big.Int).Sub(l1BalanceBeforeDeposit, l1BalanceAfterDeposit).Cmp(amount) >= 0, "Balance on L1 should be decreased")
}

func TestIntegrationWallet_EstimateGasDeposit(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_FullRequiredDepositFee(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_IsWithdrawFinalized(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

// WRITE NEGATIVE TEST
func TestIntegrationWallet_ClaimFailedDeposit(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_RequestExecute(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_EstimateGasRequestExecute(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_Address(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_Signer(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_Balance(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_AllBalances(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_L2BridgeContracts(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_Withdraw(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_EstimateGasWithdraw(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_Transfer(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_EstimateGasTransfer(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_CallContract(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_PopulateTransaction(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_SignTransaction(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_SendTransaction(t *testing.T) {
	tokenData := readToken()

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", common.HexToAddress(Receiver), big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	tokenAddress := common.HexToAddress(tokenData.L2Address)
	txHash, err := wallet.SendTransaction(context.Background(), &accounts.Transaction{
		To:   &tokenAddress,
		Data: approveTokenCalldata,
	})
	assert.NoError(t, err, "SendTransaction should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationWallet_Deploy(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_DeployWithCreate(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_DeployAccount(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationWallet_DeployAccountWithCreate(t *testing.T) {
	//client, err := clients.Dial(ZkSyncEraProvider)
	//defer client.Close()
	//assert.NoError(t, err, "clients.Dial should not return an error")
	//
	//block, err := client.BlockByNumber(context.Background(), nil)
	//
	//assert.NoError(t, err, "BlockByNumber should not return an error")
	//assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}
