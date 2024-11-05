package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/testneterc20token"
	"github.com/zksync-sdk/zksync2-go/utils"
	"log"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"
)

const TokenPath = "./testdata/tokens.json"

func readTokens() []TokenData {
	file, err := os.Open(TokenPath)
	if err != nil {
		log.Fatal("Could not find tokens.json")
	}

	var tokens []TokenData
	decoder := json.NewDecoder(file)
	if errDecode := decoder.Decode(&tokens); err != nil {
		log.Fatalf("Error decoding JSON: %s", errDecode)
	}
	return tokens
}

func deployMultisigAccount(wallet *accounts.Wallet, client *clients.Client) {
	_, abi, bytecode, err := utils.ReadStandardJson("./testdata/TwoUserMultisig.json")
	if err != nil {
		log.Fatal(err)
	}

	constructor, err := abi.Pack("", Address1, Address2)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := wallet.DeployAccount(nil, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: constructor,
		Salt:     Salt,
	})
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := client.WaitMined(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}
	multisigAccountAddress := receipt.ContractAddress
	if multisigAccountAddress != MultisigAccount {
		log.Fatal("multisig account addresses mismatch")
	}

	// transfer ETH to multisig account
	transferTxHash, err := wallet.Transfer(nil, accounts.TransferTransaction{
		To:     multisigAccountAddress,
		Amount: big.NewInt(2_000_000_000_000_000_000),
		Token:  utils.LegacyEthAddress,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), transferTxHash)
	if err != nil {
		log.Fatal(err)
	}

	// transfer token to multisig account
	transferTxHash, err = wallet.Transfer(nil, accounts.TransferTransaction{
		To:     multisigAccountAddress,
		Amount: big.NewInt(20),
		Token:  L2Dai,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), transferTxHash)
	if err != nil {
		log.Fatal(err)
	}

	// transfer approval token to multisig account
	transferTxHash, err = wallet.Transfer(nil, accounts.TransferTransaction{
		To:     multisigAccountAddress,
		Amount: big.NewInt(5),
		Token:  ApprovalToken,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), transferTxHash)
	if err != nil {
		log.Fatal(err)
	}

	if !IsEthBasedChain {
		// transfer base token to multisig account
		transferTxHash, err = wallet.Transfer(nil, accounts.TransferTransaction{
			To:     multisigAccountAddress,
			Amount: big.NewInt(2_000_000_000_000_000_000),
			Token:  utils.L2BaseTokenAddress,
		})
		if err != nil {
			log.Fatal(err)
		}

		_, err = client.WaitMined(context.Background(), transferTxHash)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func deployPaymasterAndToken(wallet *accounts.Wallet, client *clients.Client) {
	tokenAbi, err := TokenMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	constructor, err := tokenAbi.Pack("", "Crown", "Crown", uint8(18))
	if err != nil {
		log.Fatal(err)
	}

	hash, err := wallet.Deploy(nil, accounts.Create2Transaction{
		Bytecode: common.FromHex(TokenMetaData.Bin),
		Calldata: constructor,
		Salt:     Salt,
	})
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := client.WaitMined(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}
	tokenAddress := receipt.ContractAddress
	if tokenAddress != ApprovalToken {
		log.Fatal("token addresses mismatch")
	}

	// mint tokens to wallet
	token, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(wallet.Signer().PrivateKey(), wallet.Signer().ChainID())
	if err != nil {
		log.Fatal(err)
	}

	mint, err := token.Mint(opts, wallet.Address(), big.NewInt(50))
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), mint.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// deploy paymaster
	_, paymasterAbi, bytecode, err := utils.ReadStandardJson("./testdata/Paymaster.json")
	if err != nil {
		log.Fatal(err)
	}

	constructor, err = paymasterAbi.Pack("", tokenAddress)
	if err != nil {
		log.Fatal(err)
	}

	hash, err = wallet.DeployAccount(nil, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: constructor,
		Salt:     Salt,
	})
	if err != nil {
		log.Fatal(err)
	}

	receipt, err = client.WaitMined(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}
	paymasterAddress := receipt.ContractAddress
	if paymasterAddress != Paymaster {
		log.Fatal("paymaster addresses mismatch")
	}

	// transfer base token to paymaster so it could pay fee
	transferTxHash, err := wallet.Transfer(nil, accounts.TransferTransaction{
		To:     paymasterAddress,
		Amount: big.NewInt(2_000_000_000_000_000_000),
		Token:  utils.L2BaseTokenAddress,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), transferTxHash)
	if err != nil {
		log.Fatal(err)
	}

}

func mintTokenOnL1(wallet *accounts.Wallet, ethClient *ethclient.Client, l1Token common.Address) {
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(chainID)
	}
	opts, optsErr := bind.NewKeyedTransactorWithChainID(wallet.Signer().PrivateKey(), chainID)
	if optsErr != nil {
		log.Fatal(optsErr)
	}

	amount, ok := new(big.Int).SetString("20000000000000000000000", 10)
	if !ok {
		log.Fatal("failed to convert string to big.Int")
	}

	if l1Token != utils.EthAddressInContracts {
		token, tokenErr := testneterc20token.NewITestnetERC20Token(l1Token, ethClient)
		if tokenErr != nil {
			log.Fatal(tokenErr)
		}

		mint, mintErr := token.Mint(opts, wallet.Address(), amount)
		if mintErr != nil {
			log.Fatal(mintErr)
		}

		_, err = bind.WaitMined(context.Background(), ethClient, mint)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func sendTokenToL2(wallet *accounts.Wallet, client *clients.Client, ethClient *ethclient.Client, l1Token common.Address) (common.Address, common.Hash, common.Hash) {
	amount, ok := new(big.Int).SetString("10000000000000000000000", 10)
	if !ok {
		log.Fatal("failed to convert string to big.Int")
	}

	l1Tx, err := wallet.Deposit(nil, accounts.DepositTransaction{
		Token:            l1Token,
		Amount:           amount,
		To:               wallet.Address(),
		ApproveERC20:     true,
		ApproveBaseERC20: true,
		RefundRecipient:  wallet.Address(),
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), ethClient, l1Tx)
	if err != nil {
		log.Fatal(err)
	}

	l1Receipt, err := ethClient.TransactionReceipt(context.Background(), l1Tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.WaitMined(context.Background(), l2Tx.Hash)
	if err != nil {
		log.Fatal(err)
	}

	l2TokenAddress, err := client.L2TokenAddress(context.Background(), l1Token)
	if err != nil {
		log.Fatal(err)
	}

	return l2TokenAddress, l1Tx.Hash(), l2Tx.Hash
}

func wait() {
	const maxAttempts = 30

	client, err := clients.Dial(L2ChainURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	for i := 0; i < maxAttempts; i++ {
		_, err = client.NetworkID(context.Background())
		if err == nil {
			log.Println("Node is ready to receive traffic.")
			return
		}

		log.Println("Node not ready yet. Retrying...")
		time.Sleep(20 * time.Second)
	}

	log.Fatal("Maximum retries exceeded.")
}

func prepare() {
	L1Tokens = readTokens()
	L1Dai = L1Tokens[0].Address

	client, err := clients.Dial(L2ChainURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ethClient, err := ethclient.Dial(L1ChainURL)
	if err != nil {
		log.Fatal(err)
	}

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey1), client, ethClient)
	if err != nil {
		log.Fatal(err)
	}

	baseToken, err := wallet.BaseToken(nil)
	if err != nil {
		log.Fatal(err)
	}
	l1BaseTokenBalance, err := wallet.BalanceL1(nil, baseToken)
	if err != nil {
		log.Fatal(err)
	}
	l2BaseTokenBalance, err := wallet.Balance(nil, utils.L2BaseTokenAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wallet address: ", wallet.Address())
	fmt.Println("Base token L1: ", baseToken)

	fmt.Println("L1 base token balance before mint: ", l1BaseTokenBalance)
	fmt.Println("L2 base token balance before mint: ", l2BaseTokenBalance)

	mintTokenOnL1(wallet, ethClient, baseToken)
	sendTokenToL2(wallet, client, ethClient, baseToken)

	l1BaseTokenBalance, err = wallet.BalanceL1(nil, baseToken)
	if err != nil {
		log.Fatal(err)
	}
	l2BaseTokenBalance, err = wallet.Balance(nil, utils.L2BaseTokenAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("L1 base token balance after mint: ", l1BaseTokenBalance)
	fmt.Println("L2 base token balance after mint: ", l2BaseTokenBalance)

	if baseToken != utils.EthAddressInContracts {
		l2EthAddress, l2EthAddressErr := wallet.L2TokenAddress(context.Background(), utils.EthAddressInContracts)
		if l2EthAddressErr != nil {
			log.Fatal(l2EthAddressErr)
		}
		l1EthTokenBalance, l1EthTokenBalanceErr := wallet.BalanceL1(nil, utils.LegacyEthAddress)
		if l1EthTokenBalanceErr != nil {
			log.Fatal(l1EthTokenBalanceErr)
		}
		l2EthTokenBalance, l2EthTokenBalanceErr := wallet.Balance(nil, l2EthAddress)
		if l2EthTokenBalanceErr != nil && l2EthTokenBalanceErr.Error() != "no contract code at given address" {
			log.Fatal(l2EthTokenBalanceErr)
		} else {
			l2EthTokenBalance = big.NewInt(0)
		}

		fmt.Println("ETH L1: ", utils.EthAddressInContracts)
		fmt.Println("ETH L2: ", l2EthAddress)

		fmt.Println("L1 ETH balance before mint: ", l1EthTokenBalance)
		fmt.Println("L2 ETH balance before mint: ", l2EthTokenBalance)

		mintTokenOnL1(wallet, ethClient, utils.EthAddressInContracts)
		sendTokenToL2(wallet, client, ethClient, utils.EthAddressInContracts)

		l1EthTokenBalance, l1EthTokenBalanceErr = wallet.BalanceL1(nil, utils.LegacyEthAddress)
		if l1EthTokenBalanceErr != nil {
			log.Fatal(l1EthTokenBalanceErr)
		}
		l2EthTokenBalance, l2EthTokenBalanceErr = wallet.Balance(nil, l2EthAddress)
		if l2EthTokenBalanceErr != nil {
			log.Fatal(l2EthTokenBalanceErr)
		}

		fmt.Println("L1 ETH balance after mint: ", l1EthTokenBalance)
		fmt.Println("L2 ETH balance after mint: ", l2EthTokenBalance)
	}

	mintTokenOnL1(wallet, ethClient, L1Dai)
	L2Dai, L1DepositTx, L2DepositTx = sendTokenToL2(wallet, client, ethClient, L1Dai)

	l1DaiTokenBalance, err := wallet.BalanceL1(nil, L1Dai)
	if err != nil {
		log.Fatal(err)
	}
	l2DaiTokenBalance, err := wallet.Balance(nil, L2Dai)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAI L1: ", L1Dai)
	fmt.Println("DAI L2: ", L2Dai)
	fmt.Println("L1 DAI balance: ", l1DaiTokenBalance)
	fmt.Println("L2 DAI balance: ", l2DaiTokenBalance)

	if approvalTokenBytecode, bytecodeErr := client.CodeAt(context.Background(), ApprovalToken, nil); len(approvalTokenBytecode) == 0 {
		deployPaymasterAndToken(wallet, client)
	} else if bytecodeErr != nil {
		log.Fatal(bytecodeErr)
	}

	if multisigAccountBytecode, bytecodeErr := client.CodeAt(context.Background(), MultisigAccount, nil); len(multisigAccountBytecode) == 0 {
		deployMultisigAccount(wallet, client)
	} else if bytecodeErr != nil {
		log.Fatal(bytecodeErr)
	}
}

func TestMain(m *testing.M) {
	var err error
	IsEthBasedChain, err = strconv.ParseBool(os.Getenv("ETH_BASED_CHAIN"))
	if err != nil {
		IsEthBasedChain = true
	}

	if !IsEthBasedChain {
		L2ChainURL = "http://localhost:15200"
	}

	wait()
	prepare()
	os.Exit(m.Run())
}
