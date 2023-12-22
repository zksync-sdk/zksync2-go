package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

const TokenPath = "./token.json"

type token struct {
	L1Address string `json:"l1Address"`
	L2Address string `json:"l2Address"`
	L1TxHash  string `json:"L1TxHash"`
	L2TxHash  string `json:"L2TxHash"`
}

func readToken() *token {
	file, err := os.Open(TokenPath)
	if err != nil {
		log.Printf("Coud not find token.json, creating new one")
		return nil
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			log.Fatalf("Error closing file: %s", err)
		}
	}(file)

	var tokenData token
	decoder := json.NewDecoder(file)
	if errDecode := decoder.Decode(&tokenData); err != nil {
		log.Fatalf("Error decoding JSON: %s", errDecode)
	}
	return &tokenData
}

func writeToken(token token) {
	file, err := os.Create(TokenPath)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			log.Fatalf("Error closing file: %s", err)
		}
	}(file)

	// Marshal data to JSON
	jsonData, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %s", err)
	}

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %s", err)
	}
}

func checkIfTokenNeedsToBeCreated(client clients.Client, ethClient *ethclient.Client) (bool, bool) {
	tokenData := readToken()
	if tokenData == nil {
		return true, true
	}
	_, err := ethClient.CodeAt(context.Background(), common.HexToAddress(tokenData.L1Address), nil)
	if err != nil {
		return false, false
	}
	_, err = client.CodeAt(context.Background(), common.HexToAddress(tokenData.L2Address), nil)
	if err != nil {
		return true, false
	}

	return true, true
}

func createTokenL1(ethClient *ethclient.Client, privateKey *ecdsa.PrivateKey, publicKey common.Address) common.Address {
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, tokenContract, err := DeployToken(auth, ethClient, "DAI", "DAI", 18)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitDeployed(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := tokenContract.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := tokenContract.Decimals(nil)
	if err != nil {
		log.Fatal(nil)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx, err = tokenContract.Mint(auth, publicKey, big.NewInt(100_000_000))
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatal("Wait mint: ", err)
	}

	balance, err := tokenContract.BalanceOf(nil, publicKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Token address: ", address)
	log.Println("Symbol:", symbol)
	log.Println("Decimals:", decimals)
	log.Println("Balance L1: ", balance)

	return address
}

func createTokenL2(wallet *accounts.Wallet, client clients.Client, ethClient *ethclient.Client, l1Token common.Address) (common.Address, common.Hash, common.Hash) {
	tx, err := wallet.Deposit(nil, accounts.DepositTransaction{
		Token:           l1Token,
		Amount:          big.NewInt(30),
		To:              wallet.Address(),
		ApproveERC20:    true,
		RefundRecipient: wallet.Address(),
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatal(err)
	}

	l1Receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
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

	tokenL2Address, err := client.L2TokenAddress(context.Background(), l1Token)
	if err != nil {
		log.Panic(err)
	}

	tokenL2Balance, err := wallet.Balance(context.Background(), tokenL2Address, nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Balance L2: ", tokenL2Balance)

	return tokenL2Address, tx.Hash(), l2Tx.Hash
}

func wait() {
	const maxAttempts = 30

	nodeURL := "http://localhost:3050"
	client, err := clients.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	for i := 0; i < maxAttempts; i++ {
		_, err := client.NetworkID(context.Background())
		if err == nil {
			log.Println("Node is ready to receive traffic.")
			return
		}

		log.Println("Node not ready yet. Retrying...")
		time.Sleep(20 * time.Second)
	}

	log.Fatal("Maximum retries exceeded.")
}

func TestMain(m *testing.M) {
	wait()

	client, err := clients.Dial(ZkSyncEraProvider)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	ethClient, err := ethclient.Dial(EthereumProvider)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKey := crypto.PubkeyToAddress(*publicKeyECDSA)

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	if err != nil {
		log.Fatal(err)
	}

	shouldCreateL1Token, shouldCreateL2Token := checkIfTokenNeedsToBeCreated(client, ethClient)
	if shouldCreateL1Token {
		l1TokenAddress := createTokenL1(ethClient, privateKey, publicKey)
		l2TokenAddress, l1Tx, l2Tx := createTokenL2(wallet, client, ethClient, l1TokenAddress)
		writeToken(token{
			L1Address: l1TokenAddress.Hex(),
			L2Address: l2TokenAddress.Hex(),
			L1TxHash:  l1Tx.Hex(),
			L2TxHash:  l2Tx.Hex(),
		})

	} else if !shouldCreateL1Token && shouldCreateL2Token {
		tokenData := readToken()
		l2TokenAddress, l1Tx, l2Tx := createTokenL2(wallet, client, ethClient, common.HexToAddress(tokenData.L1Address))
		tokenData.L2Address = l2TokenAddress.Hex()
		tokenData.L1TxHash = l1Tx.Hex()
		tokenData.L1TxHash = l2Tx.Hex()

		writeToken(*tokenData)
	} else {
		fmt.Println("Token has been already created.")
	}

	os.Exit(m.Run())
}
