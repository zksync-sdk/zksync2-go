package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"testing"
)

func TestIntegration_NewSmartAccount(t *testing.T) {
	account := accounts.NewSmartAccount(
		Address1,
		PrivateKey1,
		&accounts.SignPayloadWithECDSA,
		&accounts.PopulateTransactionECDSA,
		nil)
	assert.NotNil(t, account, "Account should not be nil")
}

func TestIntegrationSmartAccount_Connect(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, nil)

	account = account.Connect(client)
	assert.NotNil(t, account, "Connect should return non-nil value")
}

func TestIntegrationSmartAccount_Address(t *testing.T) {
	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, nil)
	assert.Equal(t, Address1, account.Address())
}

func TestIntegrationSmartAccount_BalanceETH(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balance, err := account.Balance(nil, utils.LegacyEthAddress)

	assert.NoError(t, err, "Balance should not return an error")
	assert.True(t, balance.Cmp(big.NewInt(0)) >= 0, "Balance should return a non-negative number")
}

func TestIntegrationSmartAccount_BalanceToken(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balance, err := account.Balance(nil, L2Dai)

	assert.NoError(t, err, "Balance should not return an error")
	assert.True(t, balance.Cmp(big.NewInt(0)) >= 0, "Balance should return a non-negative number")
}

func TestIntegration_EthBasedChain_SmartAccount_AllBalances(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balances, err := account.AllBalances(context.Background())

	assert.NoError(t, err, "AllBalances should not return an error")
	assert.Len(t, balances, 2, "Should have ETH and DAI balance")
}

func TestIntegration_NonEthBasedChain_SmartAccount_AllBalances(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balances, err := account.AllBalances(context.Background())

	assert.NoError(t, err, "AllBalances should not return an error")
	assert.Len(t, balances, 3, "Should have ETH and DAI balance")
}

func TestIntegrationSmartAccount_Nonce(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	nonce, err := account.Nonce(context.Background(), nil)

	assert.NoError(t, err, "Nonce should not return an error")
	assert.True(t, nonce >= 0, "Nonce should be non-negative number")
}

func TestIntegrationSmartAccount_DeploymentNonce(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	deploymentNonce, err := account.DeploymentNonce(nil)

	assert.NoError(t, err, "DeploymentNonce should not return an error")
	assert.True(t, deploymentNonce.Cmp(big.NewInt(0)) >= 0, "Deployment nonce should be non-negative number")
}

func TestIntegration_EthBasedChain_SmartAccount_WithdrawEth(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     account.Address(),
		Amount: amount,
		Token:  utils.LegacyEthAddress,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_WithdrawEth(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	l2EthAddress, err := client.L2TokenAddress(context.Background(), utils.EthAddressInContracts)
	assert.NoError(t, err, "L2TokenAddress should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     account.Address(),
		Amount: amount,
		Token:  utils.LegacyEthAddress, // or l2EthAddress
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegration_EthBasedChain_SmartAccount_WithdrawEthUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     account.Address(),
			Amount: amount,
			Token:  utils.LegacyEthAddress,
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_WithdrawEthUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	l2EthAddress, err := client.L2TokenAddress(context.Background(), utils.EthAddressInContracts)
	assert.NoError(t, err, "L2TokenAddress should not return an error")

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     account.Address(),
			Amount: amount,
			Token:  utils.LegacyEthAddress, // or l2EthAddress
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_WithdrawBaseToken(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     account.Address(),
		Amount: amount,
		Token:  utils.L2BaseTokenAddress,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_WithdrawBaseTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     account.Address(),
			Amount: amount,
			Token:  utils.L2BaseTokenAddress,
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegrationSmartAccount_WithdrawToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     account.Address(),
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegrationSmartAccount_WithdrawTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(5)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     account.Address(),
			Amount: amount,
			Token:  L2Dai,
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegration_EthBasedChain_SmartAccount_TransferEth(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balanceBeforeTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	txHash, err := account.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  utils.LegacyEthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_TransferEth(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	sender := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)
	receiver := accounts.NewECDSASmartAccount(Address2, PrivateKey2, client)

	l2EthAddress, err := client.L2TokenAddress(context.Background(), utils.EthAddressInContracts)
	assert.NoError(t, err, "L2TokenAddress should not return an error")

	balanceBeforeTransferSender, err := sender.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	txHash, err := sender.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  utils.LegacyEthAddress, // l2EthAddress
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := sender.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := receiver.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegration_EthBasedChain_SmartAccount_TransferEthUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := account.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  utils.LegacyEthAddress,
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_TransferEthUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	sender := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)
	receiver := accounts.NewECDSASmartAccount(Address2, PrivateKey2, client)

	l2EthAddress, err := client.L2TokenAddress(context.Background(), utils.EthAddressInContracts)
	assert.NoError(t, err, "L2TokenAddress should not return an error")

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := sender.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := receiver.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := sender.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  utils.LegacyEthAddress, // or l2EthAddress
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := sender.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := receiver.Balance(nil, l2EthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_TransferBaseToken(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	balanceBeforeTransferSender, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	txHash, err := account.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  utils.L2BaseTokenAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegration_NonEthBasedChain_SmartAccount_TransferBaseTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := account.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  utils.L2BaseTokenAddress,
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, utils.L2BaseTokenAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := client.BalanceAt(context.Background(), Address2, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegrationSmartAccount_TransferToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	token, err := erc20.NewIERC20(L2Dai, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := token.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	txHash, err := account.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := token.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegrationSmartAccount_TransferTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(5)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	tokenContract, err := erc20.NewIERC20(L2Dai, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := tokenContract.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := account.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  L2Dai,
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, Address1)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := tokenContract.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegrationSmartAccount_PopulateTransaction(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	tx := &types.Transaction{
		To:    &Address2,
		Value: big.NewInt(7_000_000_000),
		From:  &Address1,
	}

	err = account.PopulateTransaction(context.Background(), tx)

	assert.NoError(t, err, "PopulateTransaction should not return an error")
	assert.Equal(t, &Address2, tx.To, "To addresses must be the same")
	assert.Equal(t, &Address1, tx.From, "From addresses must be the same")
	assert.Equal(t, big.NewInt(7_000_000_000), tx.Value, "Values must be the same")
	assert.NotNil(t, tx.ChainID, "Chain IDs must not be nil")
	assert.NotNil(t, tx.Data, "Data must not be nil")
	assert.NotNil(t, tx.Gas, "Gas must not be nil")
	assert.NotNil(t, tx.GasFeeCap, "GasFeeCap must not be nil")
	assert.NotNil(t, tx.GasTipCap, "GasTipCap must not be nil")
	assert.NotNil(t, tx.GasPerPubdata, "GasPerPubdata must not be nil")
}

func TestIntegrationSmartAccount_SignTransaction(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	signedTx, err := account.SignTransaction(context.Background(), &types.Transaction{
		To:    &Address2,
		Value: big.NewInt(1_000_000_000_000_000_000), // 1ETH
	})

	assert.NoError(t, err, "SignTransaction should not return an error")
	assert.NotNil(t, signedTx, "Transaction should not be nil")
}

func TestIntegrationSmartAccount_SignMessage(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	signature, err := account.SignMessage(context.Background(), []byte("Hello World!"))

	assert.NoError(t, err, "SignMessage should not return an error")
	assert.NotNil(t, signature, "Signature should not be nil")
}

func TestIntegrationSmartAccount_SignTypedData(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	signature, err := account.SignTypedData(context.Background(), apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    "Example",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(270),
		},
		Types: apitypes.Types{
			"Person": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "age", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "Person",
		Message: apitypes.TypedDataMessage{
			"name": "John",
			"age":  hexutil.EncodeUint64(30),
		},
	})
	assert.NoError(t, err, "SignTypedData should not return an error")
	assert.NotNil(t, signature, "Signature should not be nil")
}

func TestIntegrationSmartAccount_SendTransaction(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewECDSASmartAccount(Address1, PrivateKey1, client)

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Address2, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	txHash, err := account.SendTransaction(context.Background(), &types.Transaction{
		To:   &L2Dai,
		Data: approveTokenCalldata,
	})
	assert.NoError(t, err, "SendTransaction should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationMultisigSmartAccount_Withdraw(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     wallet.Address(),
		Amount: amount,
		Token:  utils.LegacyEthAddress,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegrationMultisigSmartAccount_WithdrawUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     wallet.Address(),
			Amount: amount,
			Token:  utils.LegacyEthAddress,
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegrationMultisigSmartAccount_WithdrawToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	withdrawHash, err := account.Withdraw(nil, accounts.WithdrawalTransaction{
		To:     wallet.Address(),
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
}

func TestIntegrationMultisigSmartAccount_WithdrawTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(5)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	assert.NoError(t, err, "NewWallet should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	l2BalanceBeforeWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeWithdrawal, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	withdrawHash, err := account.Withdraw(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.WithdrawalTransaction{
			To:     wallet.Address(),
			Amount: amount,
			Token:  L2Dai,
		})
	assert.NoError(t, err, "Withdraw should not return an error")

	withdrawReceipt, err := client.WaitFinalized(context.Background(), withdrawHash)
	assert.NoError(t, err, "client.WaitFinalized should not return an error")
	assert.NotNil(t, withdrawReceipt.BlockHash, "Withdraw transaction should be mined")

	finalizeWithdrawTx, err := wallet.FinalizeWithdraw(nil, withdrawHash, 0)
	assert.NoError(t, err, "FinalizeWithdraw should not return an error")

	finalizeWithdrawReceipt, err := bind.WaitMined(context.Background(), ethClient, finalizeWithdrawTx)
	assert.NoError(t, err, "bind.WaitMined should not return an error")
	assert.NotNil(t, finalizeWithdrawReceipt.BlockHash, "Finalize withdraw transaction should be mined")

	l2BalanceAfterWithdrawal, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterWithdrawal, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterWithdrawalPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterWithdrawalPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeWithdrawalPaymaster.Cmp(balanceAfterWithdrawalPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterWithdrawalPaymaster, approvalTokenBalanceBeforeWithdrawalPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(l2BalanceBeforeWithdrawal, l2BalanceAfterWithdrawal).Cmp(amount) >= 0, "Balance on L2 should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeWithdrawal, minimalAllowance).Cmp(approvalTokenBalanceAfterWithdrawal) == 0, "Sender approval token balance should be decreased")
}

func TestIntegrationMultisigSmartAccount_TransferEth(t *testing.T) {
	amount := big.NewInt(7_000_000_000)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	sender := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)
	receiver := accounts.NewECDSASmartAccount(Address2, PrivateKey2, client)

	balanceBeforeTransferSender, err := sender.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := receiver.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	txHash, err := sender.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  utils.LegacyEthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := sender.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := receiver.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegrationMultisigSmartAccount_TransferEthUsingPaymaster(t *testing.T) {
	amount := big.NewInt(7_000_000_000)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)
	receiver := accounts.NewECDSASmartAccount(Address2, PrivateKey2, client)

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := receiver.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := account.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  utils.LegacyEthAddress,
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := receiver.Balance(nil, utils.LegacyEthAddress)
	assert.NoError(t, err, "BalanceAt should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Receiver balance should be increased")
}

func TestIntegrationMultisigSmartAccount_TransferToken(t *testing.T) {
	amount := big.NewInt(5)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	token, err := erc20.NewIERC20(L2Dai, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	balanceBeforeTransferReceiver, err := token.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	txHash, err := account.Transfer(nil, accounts.TransferTransaction{
		To:     Address2,
		Amount: amount,
		Token:  L2Dai,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	balanceAfterTransferReceiver, err := token.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}

func TestIntegrationMultisigSmartAccount_TransferTokenUsingPaymaster(t *testing.T) {
	amount := big.NewInt(5)
	minimalAllowance := big.NewInt(1)

	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	ethClient, err := ethclient.Dial(L1ChainURL)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	account := accounts.NewMultisigECDSASmartAccount(MultisigAccount, []string{PrivateKey1, PrivateKey2}, client)

	tokenContract, err := erc20.NewIERC20(L2Dai, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	approvalToken, err := erc20.NewIERC20(ApprovalToken, client)
	assert.NoError(t, err, "NewIERC20 should not return an error")

	balanceBeforeTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceBeforeTransferSender, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferReceiver, err := tokenContract.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceBeforeTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceBeforeTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		Paymaster,
		&types.ApprovalBasedPaymasterInput{
			Token:            ApprovalToken,
			MinimalAllowance: minimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	txHash, err := account.Transfer(
		&accounts.TransactOpts{PaymasterParams: paymasterParams},
		accounts.TransferTransaction{
			To:     Address2,
			Amount: amount,
			Token:  L2Dai,
		})
	assert.NoError(t, err, "Transfer should not return an error")

	receipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")
	assert.NotNil(t, receipt.BlockHash, "Transaction should be mined")

	balanceAfterTransferSender, err := account.Balance(nil, L2Dai)
	assert.NoError(t, err, "Balance should not return an error")

	approvalTokenBalanceAfterTransferSender, err := approvalToken.BalanceOf(nil, MultisigAccount)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferReceiver, err := tokenContract.BalanceOf(nil, Address2)
	assert.NoError(t, err, "BalanceOf should not return an error")

	balanceAfterTransferPaymaster, err := client.BalanceAt(context.Background(), Paymaster, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	approvalTokenBalanceAfterTransferPaymaster, err := approvalToken.BalanceOf(nil, Paymaster)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, balanceBeforeTransferPaymaster.Cmp(balanceAfterTransferPaymaster) >= 0, "Paymaster balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceAfterTransferPaymaster, approvalTokenBalanceBeforeTransferPaymaster).Cmp(minimalAllowance) == 0, "Paymaster approval token balance should be increased")

	assert.True(t, new(big.Int).Sub(balanceBeforeTransferSender, balanceAfterTransferSender).Cmp(amount) >= 0, "Sender balance should be decreased")
	assert.True(t, new(big.Int).Sub(approvalTokenBalanceBeforeTransferSender, minimalAllowance).Cmp(approvalTokenBalanceAfterTransferSender) == 0, "Sender approval token balance should be decreased")

	assert.True(t, new(big.Int).Sub(balanceAfterTransferReceiver, balanceBeforeTransferReceiver).Cmp(amount) >= 0, "Address2 balance should be increased")
}
