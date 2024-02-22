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
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

const TokenPath = "./testdata/tokens.json"

func readTokens() []TokenData {
	file, err := os.Open(TokenPath)
	if err != nil {
		log.Printf("Could not find tokens.json")
		return nil
	}

	var tokens []TokenData
	decoder := json.NewDecoder(file)
	if errDecode := decoder.Decode(&tokens); err != nil {
		log.Fatalf("Error decoding JSON: %s", errDecode)
	}
	return tokens
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
		log.Fatal(err)
	}

	tokenL2Balance, err := wallet.Balance(context.Background(), tokenL2Address, nil)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	defer client.Close()

	ethClient, err := ethclient.Dial(EthereumProvider)
	if err != nil {
		log.Fatal(err)
	}

	wallet, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
	if err != nil {
		log.Fatal(err)
	}

	L1Tokens = readTokens()
	L1Dai = L1Tokens[0].Address

	L2Dai, L1DepositTx, L2DepositTx = createTokenL2(wallet, client, ethClient, L1Dai)

	os.Exit(m.Run())
}
