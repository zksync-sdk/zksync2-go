# 🚀 zksync2-go Golang SDK 🚀

[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE-MIT)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache%202.0-orange)](LICENSE-APACHE)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](https://www.contributor-covenant.org/)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-orange)](.github/CONTRIBUTING.md)
[![X (formerly Twitter) Follow](https://badgen.net/badge/twitter/@zksyncDevs/1DA1F2?icon&label)](https://x.com/zksyncDevs)

[![ZKsync Era Logo](logo.svg)](https://zksync.io/)

In order to provide easy access to all the features of ZKsync Era, the `zksync2-go` Golang SDK was created,
which is made in a way that has an interface very similar to those of [geth](https://geth.ethereum.org/). In
fact, `geth` is a dependency of our library and most of the objects exported by `zksync2-go` (
e.g. `Wallet`, `Client` etc.) inherit from the corresponding `geth` objects and override only the fields that need
to be changed.

While most of the existing SDKs should work out of the box, deploying smart contracts or using unique ZKsync Era features,
like account abstraction, requires providing additional fields to those that Ethereum transactions have by default.

The library is made in such a way that after replacing `geth` with `zksync2-go` most client apps will work out of
box.

🔗 For a detailed walkthrough, refer to the [official documentation](https://docs.zksync.io/sdk/go/getting-started).

## 📌 Overview

To begin, it is useful to have a basic understanding of the types of objects available and what they are responsible for, at a high level:

-   `Client` provides connection to the ZKsync Era blockchain, which allows querying the blockchain state, such as account, block or transaction details,
    querying event logs or evaluating read-only code using call. Additionally, the client facilitates writing to the blockchain by sending
    transactions.
-   `Wallet` wraps all operations that interact with an account. An account generally has a private key, which can be used to sign a variety of
    types of payloads. It provides easy usage of the most common features.

## 🛠 Prerequisites

-   `go: >= 1.21` ([installation guide](https://go.dev/doc/install))

## 📥 Installation & Setup

```bash
go get github.com/zksync-sdk/zksync2-go
```

## 📝 Examples

The complete examples with various use cases are available [here](https://github.com/zksync-sdk/zksync2-examples/tree/main/go).

### Connect to the ZKsync Era network:

```go
import (
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/zksync-sdk/zksync2-go/clients"
    "log"
    "os"
)

var (
    ZkSyncProvider   = "https://sepolia.era.zksync.dev" // zkSync Era testnet  
    EthereumProvider = "https://rpc.ankr.com/eth_sepolia" // Sepolia testnet
)
// Connect to ZKsync network
client, err := clients.Dial(ZkSyncProvider)
if err != nil {
    log.Panic(err)
}
defer client.Close()

// Connect to Ethereum network
ethClient, err := ethclient.Dial(EthereumProvider)
if err != nil {
    log.Panic(err)
}
defer ethClient.Close()
```

### Get the latest block number

```ts
blockNumber, err := client.BlockNumber(context.Background())
if err != nil {
    log.Panic(err)
}
fmt.Println("Block number: ", blockNumber)
```

### Get the latest block

```ts
block, err := client.BlockByNumber(context.Background(), nil)
if err != nil {
    log.Panic(err)
}
fmt.Printf("%+v\n", block)
```

### Create a wallet

```ts
privateKey := os.Getenv("PRIVATE_KEY")
w, err := accounts.NewWallet(common.Hex2Bytes(privateKey), &client, ethClient)
if err != nil {
    log.Panic(err)
}
```

### Check account balances

```go
balance, err := w.Balance(context.Background(), utils.EthAddress, nil) // balance on ZKsync Era network
if err != nil {
    log.Panic(err)
}
fmt.Println("Balance: ", balance)

balanceL1, err := w.BalanceL1(nil, utils.EthAddress) // balance on goerli network
if err != nil {
    log.Panic(err)
}
fmt.Println("Balance L1: ", balanceL1)
```

### Transfer funds

Transfer funds among accounts on L2 network.


```go
chainID, err := client.ChainID(context.Background())
if err != nil {
    log.Panic(err)
}
receiver, err := accounts.NewRandomWallet(chainID.Int64(), &client, ethClient)
if err != nil {
    log.Panic(err)
}

tx, err := w.Transfer(accounts.TransferTransaction{
    To:     receiver.Address(),
    Amount: big.NewInt(7_000_000_000_000_000_000),
    Token:  utils.EthAddress,
})
if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", tx.Hash())
```

### Deposit funds

Transfer funds from L1 to L2 network.

```ts
tx, err := w.Deposit(accounts.DepositTransaction{
    Token:  utils.EthAddress, 
    Amount: big.NewInt(2_000_000_000_000_000_000),
    To:     w.Address(),
})
if err != nil {
    log.Panic(err)
}
fmt.Println("L1 transaction: ", tx.Hash())
```

### Withdraw funds

Transfer funds from L2 to L1 network.

```ts
tx, err := w.Withdraw(accounts.WithdrawalTransaction{
    To:     w.Address(),
        Amount: big.NewInt(1_000_000_000_000_000_000),
        Token:  utils.EthAddress,
})
if err != nil {
    log.Panic(err)
}
fmt.Println("Withdraw transaction: ", tx.Hash())
```

## 🤖 Running tests

In order to run test you need to run [local-setup](https://github.com/matter-labs/local-setup) on your machine.
For running tests, use:

```shell
make run-tests
```


## 🤝 Contributing

We welcome contributions from the community! If you're interested in contributing to the `zksync2-go` Golang SDK,
please take a look at our [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for guidelines and details on the process.

Thank you for making `zksync2-go` Golang SDK better! 🙌
