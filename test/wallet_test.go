package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"os"
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

	tokenBalance, err := wallet.BalanceL1(nil, L1Dai)

	assert.NoError(t, err, "BalanceL1 should not return an error")
	assert.NotNil(t, tokenBalance, "BlockByNumber should return a non-nil token balance")
}

func TestIntegrationWallet_AllowanceL1(t *testing.T) {
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

	ethBalance, err := wallet.AllowanceL1(nil, L1Dai, bridgeContracts.L1Erc20DefaultBridge)

	assert.NoError(t, err, "AllowanceL1 should not return an error")
	assert.NotNil(t, ethBalance, "AllowanceL1 should return a non-nil token allowance")
}

func TestIntegrationWallet_L2TokenAddress(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2Address, err := wallet.L2TokenAddress(context.Background(), L1Dai)

	assert.NoError(t, err, "L2TokenAddress should not return an error")
	assert.Equal(t, L2Dai, l2Address, "L2 token addresses should be the same")
}

func TestIntegrationWallet_ApproveERC20(t *testing.T) {
	l1TokenAddress := L1Dai
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

	tx, err := wallet.ApproveERC20(nil, L1Dai, approveAmount, bridgeContracts.L1Erc20DefaultBridge)
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
	amount := big.NewInt(5)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeDeposit, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceBeforeDeposit, err := wallet.BalanceL1(nil, L1Dai)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	tx, err := wallet.Deposit(nil, accounts.DepositTransaction{
		To:              wallet.Address(),
		Token:           L1Dai,
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

	l2BalanceAfterDeposit, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceAfterDeposit, err := wallet.BalanceL1(nil, L1Dai)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceAfterDeposit, l2BalanceBeforeDeposit).Cmp(amount) >= 0, "Balance on L2 should be increased")
	assert.True(t, new(big.Int).Sub(l1BalanceBeforeDeposit, l1BalanceAfterDeposit).Cmp(amount) >= 0, "Balance on L1 should be decreased")
}

func TestIntegrationWallet_FullRequiredDepositFeeETH(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	feeData := &accounts.FullDepositFee{
		MaxFeePerGas:         big.NewInt(1_500_000_010),
		MaxPriorityFeePerGas: big.NewInt(1_500_000_000),
		BaseCost:             big.NewInt(289_770_500_000_000),
		L1GasLimit:           big.NewInt(22_792),
		L2GasLimit:           big.NewInt(579_541),
	}

	depositFee, err := wallet.FullRequiredDepositFee(context.Background(), accounts.DepositCallMsg{
		Token: utils.EthAddress,
		To:    Receiver,
	})

	assert.NoError(t, err, "FullRequiredDepositFee should not return an error")
	assert.Equal(t, feeData, depositFee, "Fees should be the same")
}

func TestIntegrationWallet_FullRequiredDepositFeeNotEnoughBalance(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	chainId, err := client.ChainID(context.Background())
	assert.NoError(t, err, "ChainID should not return an error")

	wallet, err := accounts.NewRandomWallet(chainId.Int64(), &client, ethClient)
	assert.NoError(t, err, "NewRandomWallet should not return an error")

	_, err = wallet.FullRequiredDepositFee(context.Background(), accounts.DepositCallMsg{
		Token: utils.EthAddress,
		To:    Receiver,
	})

	assert.Error(t, err, "Should throw error when there is not enough balance")
	assert.Contains(t, err.Error(), "not enough balance for deposit")
}

func TestIntegrationWallet_FullRequiredDepositFeeTokenNotEnoughAllowance(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	_, err = wallet.FullRequiredDepositFee(context.Background(), accounts.DepositCallMsg{
		Token: L1Dai,
		To:    Address,
	})

	assert.Error(t, err, "Should throw error when there is not enough allowance")
	assert.EqualError(t, err, "not enough allowance to cover the deposit")
}

func TestIntegrationWallet_FullRequiredDepositFeeToken(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	feeData := &accounts.FullDepositFee{
		MaxFeePerGas:         big.NewInt(1_500_000_010),
		MaxPriorityFeePerGas: big.NewInt(1_500_000_000),
		BaseCost:             big.NewInt(288992000000000),
		L1GasLimit:           big.NewInt(210_981),
		L2GasLimit:           big.NewInt(577984),
	}

	bridgeContracts, err := client.BridgeContracts(context.Background())
	assert.NoError(t, err, "BridgeContracts should not return an error")

	approveTx, err := wallet.ApproveERC20(nil, L1Dai, big.NewInt(5), bridgeContracts.L1Erc20DefaultBridge)
	assert.NoError(t, err, "ApproveERC20 should not return an error")

	_, err = bind.WaitMined(context.Background(), ethClient, approveTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")

	depositFee, err := wallet.FullRequiredDepositFee(context.Background(), accounts.DepositCallMsg{
		Token: L1Dai,
		To:    Address,
	})

	assert.NoError(t, err, "FullRequiredDepositFee should not return an error")
	assert.Equal(t, feeData, depositFee, "Fees should be the same")
}

func TestIntegrationWallet_RequestExecute(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeTx, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceBeforeTx, err := wallet.BalanceL1(nil, utils.EthAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	mainContract, err := client.MainContractAddress(context.Background())
	assert.NoError(t, err, "client.MainContractAddress should not return an error")

	tx, err := wallet.RequestExecute(nil, accounts.RequestExecuteTransaction{
		ContractAddress: mainContract,
		L2GasLimit:      big.NewInt(900_000),
		L2Value:         amount,
	})
	assert.NoError(t, err, "Request execute should not return an error")

	l1Receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")

	l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
	assert.NoError(t, err, "L2TransactionFromPriorityOp should not return an error")

	l2Receipt, err := client.WaitMined(context.Background(), l2Tx.Hash)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, l2Receipt.BlockHash, "Transaction should be mined")

	l2BalanceAfterTx, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	l1BalanceAfterTx, err := wallet.BalanceL1(nil, utils.EthAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceAfterTx, l2BalanceBeforeTx).Cmp(amount) >= 0, "Balance on L2 should be increased")
	assert.True(t, new(big.Int).Sub(l1BalanceBeforeTx, l1BalanceAfterTx).Cmp(amount) >= 0, "Balance on L1 should be decreased")
}

func TestIntegrationWallet_BalanceETH(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	balance, err := wallet.Balance(context.Background(), utils.EthAddress, nil)

	assert.NoError(t, err, "Balance should not return an error")
	assert.True(t, balance.Cmp(big.NewInt(0)) >= 0, "Balance should return a non-negative number")
}

func TestIntegrationWallet_BalanceToken(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	balance, err := wallet.Balance(context.Background(), L2Dai, nil)

	assert.NoError(t, err, "Balance should not return an error")
	assert.True(t, balance.Cmp(big.NewInt(0)) >= 0, "Balance should return a non-negative number")
}

func TestIntegrationWallet_AllBalances(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	balances, err := wallet.AllBalances(context.Background())

	assert.NoError(t, err, "AllBalances should not return an error")
	assert.Len(t, balances, 2, "Should have ETH and DAI balance")
}

func TestIntegrationWallet_L2BridgeContracts(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	bridges, err := wallet.L2BridgeContracts(context.Background())

	assert.NoError(t, err, "L2BridgeContracts should not return an error")
	assert.NotNil(t, bridges, "L2BridgeContracts should return bridges")
}

func TestIntegrationWallet_DeploymentNonce(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWalletL2(common.Hex2Bytes(PrivateKey), &client)
	assert.NoError(t, err, "NewWallet should not return an error")

	deploymentNonce, err := wallet.DeploymentNonce(nil)

	assert.NoError(t, err, "DeploymentNonce should not return an error")
	assert.True(t, deploymentNonce.Cmp(big.NewInt(0)) >= 0, "Deployment nonce should be non-negative number")
}

func TestIntegrationWallet_Withdraw(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeWithdrawal, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawTx, err := wallet.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     wallet.Address(),
		Amount: amount,
		Token:  utils.EthAddress,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawTx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	isWithdrawFinalized, err := wallet.IsWithdrawFinalized(nil, withdrawTx.Hash(), 0)
	assert.NoError(t, err, "IsWithdrawFinalized should not return an error")
	assert.False(t, isWithdrawFinalized, "Withdraw transaction should not be finalized")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawTx.Hash(), 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegrationWallet_WithdrawToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	l2BalanceBeforeWithdrawal, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawTx, err := wallet.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     wallet.Address(),
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawTx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	isWithdrawFinalized, err := wallet.IsWithdrawFinalized(nil, withdrawTx.Hash(), 0)
	assert.NoError(t, err, "IsWithdrawFinalized should not return an error")
	assert.False(t, isWithdrawFinalized, "Withdraw transaction should not be finalized")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawTx.Hash(), 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegrationWallet_Transfer(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	balanceBeforeTransferSender, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Receiver, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	tx, err := wallet.Transfer(nil, accounts.TransferTransaction{
		To:     Receiver,
		Amount: amount,
		Token:  utils.EthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), tx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := wallet.BalanceL1(nil, utils.EthAddress)
	assert.NoError(t, err, "BalanceL1 should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegrationWallet_TransferToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	tokenContract, err := erc20.NewIERC20(L2Dai, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := tokenContract.BalanceOf(nil, Receiver)
	assert.NoError(t, err, "BalanceOf should not return an error")

	tx, err := wallet.Transfer(nil, accounts.TransferTransaction{
		To:     Receiver,
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), tx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := wallet.Balance(context.Background(), L2Dai, nil)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := tokenContract.BalanceOf(nil, Receiver)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegrationWallet_PopulateTransaction(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	nonce, err := wallet.Nonce(context.Background(), nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tx := &types.Transaction712{
		Nonce:     new(big.Int).SetUint64(nonce),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(250_000_000),
		Gas:       big.NewInt(154_379),
		To:        &Receiver,
		Value:     big.NewInt(7_000_000_000),
		ChainID:   big.NewInt(270),
		From:      &Address,
		Data:      hexutil.Bytes{},
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(50_000),
		},
	}

	populatedTx, err := wallet.PopulateTransaction(context.Background(), accounts.Transaction{
		To:    &Receiver,
		Value: big.NewInt(7_000_000_000),
	})

	assert.NoError(t, err, "PopulateTransaction should not return an error")
	assert.Equal(t, tx, populatedTx, "Transactions should be the same")
}

func TestIntegrationWallet_SignTransaction(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	signedTx, err := wallet.SignTransaction(&types.Transaction712{
		To:    &Receiver,
		Value: big.NewInt(1_000_000_000_000_000_000), // 1ETH
	})

	assert.NoError(t, err, "SignTransaction should not return an error")
	assert.NotNil(t, signedTx, "Transactions should be nil")
}

func TestIntegrationWallet_SendTransaction(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Receiver, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	txHash, err := wallet.SendTransaction(context.Background(), &accounts.Transaction{
		To:   &L2Dai,
		Data: approveTokenCalldata,
	})
	assert.NoError(t, err, "SendTransaction should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationWallet_DeployWithCreate(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	bytecode, err := os.ReadFile("./testdata/Storage.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	hash, err := wallet.DeployWithCreate(nil, accounts.CreateTransaction{Bytecode: bytecode})
	assert.NoError(t, err, "DeployWithCreate should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployWithCreateConstructor(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	bytecode, err := os.ReadFile("./testdata/Incrementer.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	abi, err := IncrementerMetaData.GetAbi()
	assert.NoError(t, err, "GetAbi should not return an error")

	constructor, err := abi.Pack("", big.NewInt(2))
	assert.NoError(t, err, "Pack should not return an error")

	hash, err := wallet.DeployWithCreate(nil, accounts.CreateTransaction{
		Bytecode: bytecode,
		Calldata: constructor,
	})
	assert.NoError(t, err, "DeployWithCreate should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployWithCreateDeps(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	demoBytecode, err := os.ReadFile("./testdata/Demo.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	fooBytecode, err := os.ReadFile("./testdata/Foo.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	hash, err := wallet.DeployWithCreate(nil, accounts.CreateTransaction{
		Bytecode:     demoBytecode,
		Dependencies: [][]byte{fooBytecode},
	})

	assert.NoError(t, err, "DeployWithCreate should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployWithCreateAccount(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	_, paymasterAbi, bytecode, err := utils.ReadStandardJson("./testdata/Paymaster.json")
	assert.NoError(t, err, "ReadStandardJson should not return an error")

	constructor, err := paymasterAbi.Pack("", L2Dai)
	assert.NoError(t, err, "Pack should not return an error")

	hash, err := wallet.DeployAccountWithCreate(nil, accounts.CreateTransaction{
		Bytecode: bytecode,
		Calldata: constructor,
	})
	assert.NoError(t, err, "DeployWithCreate should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_Deploy(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	bytecode, err := os.ReadFile("./testdata/Storage.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	hash, err := wallet.Deploy(nil, accounts.Create2Transaction{Bytecode: bytecode})
	assert.NoError(t, err, "Deploy should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployConstructor(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	bytecode, err := os.ReadFile("./testdata/Incrementer.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	abi, err := IncrementerMetaData.GetAbi()
	assert.NoError(t, err, "GetAbi should not return an error")

	constructor, err := abi.Pack("", big.NewInt(2))
	assert.NoError(t, err, "Pack should not return an error")

	hash, err := wallet.Deploy(nil, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: constructor,
	})
	assert.NoError(t, err, "Deploy should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployDeps(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	demoBytecode, err := os.ReadFile("./testdata/Demo.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	fooBytecode, err := os.ReadFile("./testdata/Foo.zbin")
	assert.NoError(t, err, "ReadFile should not return an error")

	hash, err := wallet.Deploy(nil, accounts.Create2Transaction{
		Bytecode:     demoBytecode,
		Dependencies: [][]byte{fooBytecode},
	})

	assert.NoError(t, err, "Deploy should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}

func TestIntegrationWallet_DeployAccount(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	_, paymasterAbi, bytecode, err := utils.ReadStandardJson("./testdata/Paymaster.json")
	assert.NoError(t, err, "ReadStandardJson should not return an error")

	constructor, err := paymasterAbi.Pack("", L2Dai)
	assert.NoError(t, err, "Pack should not return an error")

	hash, err := wallet.Deploy(nil, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: constructor,
	})
	assert.NoError(t, err, "Deploy should not return an error")

	receipt, err := client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	contractAddress := receipt.ContractAddress
	assert.NotNil(t, contractAddress, "Contract should be deployed")
}
