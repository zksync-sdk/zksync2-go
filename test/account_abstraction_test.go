package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"testing"
)

func TestIntegration_ApprovalPaymaster(t *testing.T) {
	AirdropAmount := big.NewInt(10)
	MinimalAllowance := big.NewInt(1)
	MintAmount := big.NewInt(7)

	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	// ====== Deploy Token ======
	tokenAbi, err := TokenMetaData.GetAbi()
	assert.NoError(t, err, "GetAbi should not return an error")

	tokenConstructor, err := tokenAbi.Pack("", "Crown", "Crown", uint8(18))
	assert.NoError(t, err, "Pack should not return an error")

	tokenDeployHash, err := wallet.DeployWithCreate(nil, accounts.CreateTransaction{
		Bytecode: common.FromHex(TokenMetaData.Bin),
		Calldata: tokenConstructor,
	})
	assert.NoError(t, err, "Deploy should not return an error")

	tokenDeployReceipt, err := client.WaitMined(context.Background(), tokenDeployHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	tokenAddress := tokenDeployReceipt.ContractAddress
	assert.NotNil(t, tokenAddress, "Contract should be deployed")

	// ===== Mint tokens to wallet =====
	token, err := NewToken(tokenAddress, client)
	assert.NoError(t, err, "NewToken should not return an error")

	opts, err := bind.NewKeyedTransactorWithChainID(wallet.Signer().PrivateKey(), wallet.Signer().Domain().ChainId)
	assert.NoError(t, err, "NewKeyedTransactorWithChainID should not return an error")

	mint, err := token.Mint(opts, wallet.Address(), AirdropAmount)
	assert.NoError(t, err, "Mint should not return an error")

	_, err = client.WaitMined(context.Background(), mint.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")

	// ===== Deploy Paymaster =====
	_, paymasterAbi, bytecode, err := utils.ReadStandardJson("./testdata/Paymaster.json")
	assert.NoError(t, err, "ReadStandardJson should not return an error")

	paymasterConstructor, err := paymasterAbi.Pack("", tokenAddress)
	assert.NoError(t, err, "Pack should not return an error")

	paymasterDeployHash, err := wallet.Deploy(nil, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: paymasterConstructor,
	})
	assert.NoError(t, err, "DeployWithCreate should not return an error")

	paymasterDeployReceipt, err := client.WaitMined(context.Background(), paymasterDeployHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	paymasterAddress := paymasterDeployReceipt.ContractAddress
	assert.NotNil(t, paymasterAddress, "Contract should be deployed")

	// ===== Transfer some ETH to paymaster, so it can pay fee with ETH =====
	transferTx, err := wallet.Transfer(nil, accounts.TransferTransaction{
		To:     paymasterAddress,
		Amount: big.NewInt(2_000_000_000_000_000_000),
		Token:  utils.LegacyEthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	_, err = client.WaitMined(context.Background(), transferTx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")

	// Read token and ETH balances from user and paymaster accounts
	balanceBefore, err := wallet.Balance(context.Background(), utils.LegacyEthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	tokenBalanceBefore, err := token.BalanceOf(nil, wallet.Address())
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterBalanceBefore, err := client.BalanceAt(context.Background(), paymasterAddress, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	paymasterTokenBalanceBefore, err := token.BalanceOf(nil, paymasterAddress)
	assert.NoError(t, err, "BalanceOf should not return an error")

	calldata, err := tokenAbi.Pack("mint", wallet.Address(), MintAmount)
	assert.NoError(t, err, "Pack should not return an error")

	paymasterParams, err := utils.GetPaymasterParams(
		paymasterAddress,
		&types.ApprovalBasedPaymasterInput{
			Token:            tokenAddress,
			MinimalAllowance: MinimalAllowance,
			InnerInput:       []byte{},
		})
	assert.NoError(t, err, "GetPaymasterParams should not return an error")

	hash, err := wallet.SendTransaction(context.Background(), &accounts.Transaction{
		To:   &tokenAddress,
		Data: calldata,
		Meta: &types.Eip712Meta{
			PaymasterParams: paymasterParams,
		},
	})
	assert.NoError(t, err, "SendTransaction should not return an error")

	_, err = client.WaitMined(context.Background(), hash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	balanceAfter, err := wallet.Balance(context.Background(), utils.LegacyEthAddress, nil)
	assert.NoError(t, err, "Balance should not return an error")

	tokenBalanceAfter, err := token.BalanceOf(nil, wallet.Address())
	assert.NoError(t, err, "BalanceOf should not return an error")

	paymasterBalanceAfter, err := client.BalanceAt(context.Background(), paymasterAddress, nil)
	assert.NoError(t, err, "BalanceAt should not return an error")

	paymasterTokenBalanceAfter, err := token.BalanceOf(nil, paymasterAddress)
	assert.NoError(t, err, "BalanceOf should not return an error")

	assert.True(t, paymasterTokenBalanceBefore.Cmp(big.NewInt(0)) == 0, "Paymaster token balance before transaction should be 0")
	assert.True(t, tokenBalanceBefore.Cmp(AirdropAmount) == 0, "Wallet token balance before transaction should be equal to airdrop amount")
	assert.True(t, paymasterBalanceBefore.Cmp(paymasterBalanceAfter) > 0, "Paymaster balance after transaction should be decreased")
	assert.True(t, paymasterTokenBalanceAfter.Cmp(MinimalAllowance) == 0, "Paymaster token balance after transaction should be minimal allowance amount")
	assert.True(t, balanceBefore.Cmp(balanceAfter) == 0, "Balance should be the same")

	tmp := new(big.Int).Add(tokenBalanceBefore, MintAmount)
	assert.True(t, tokenBalanceAfter.Cmp(tmp.Sub(tmp, MinimalAllowance)) == 0, "Wallet token balance should be increased by difference between mint amount and allowance amount")
}
