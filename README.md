zkSync 2 Golang SDK
===


## Requirements

Go version >= 1.17

Also, Go modules was used

## Installation

Add this package to your project: 

```shell
$ go get github.com/zksync-sdk/zksync2-go
```

## Using examples

The complete examples with various use cases are available [here](https://github.com/zksync-sdk/zksync2-examples/tree/main/go).

### How to init main components

```go
import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/zksync-sdk/zksync2-go/accounts"
    "github.com/zksync-sdk/zksync2-go/clients"
    "github.com/zksync-sdk/zksync2-go/utils"
    "log"
    "math/big"
    "os"
)

var (
    PrivateKey       = os.Getenv("PRIVATE_KEY")
    ZkSyncProvider   = "https://testnet.era.zksync.dev"
    EthereumProvider = "https://rpc.ankr.com/eth_goerli"
)
// Connect to zkSync network
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

// Create wallet
w, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, ethClient)
if err != nil {
    log.Panic(err)
}
```

Now, using this three instances - zkSync client, ethereum client and wallet, 
you can perform all basic actions with ZkSync network.

### Get balance of ZkSync account
```go
balance, err := w.Balance(context.Background(), utils.EthAddress, nil)
if err != nil {
	log.Panic(err)
}
fmt.Println("Balance: ", balance)
```

### Deposit

#### Deposit ETH token
```go
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

#### Deposit ERC20 token
```go
TokenL1Address   := common.HexToAddress("0xf10A110E59a22b444c669C83b02f0E6d945b2b69")

tx, err := w.Deposit(accounts.DepositTransaction{
    Token:           TokenL1Address,
    Amount:          big.NewInt(5),
    To:              w.Address(),
    ApproveERC20:    true,
    RefundRecipient: w.Address(),
})
if err != nil {
    log.Panic(err)
}
fmt.Println("L1 deposit transaction: ", tx.Hash())
```

Also, you can get hash of corresponding L2 transaction:
```go
// Wait for deposit transaction to be finalized on L1 network
_, err = bind.WaitMined(context.Background(), ethClient, tx)
if err != nil {
    log.Panic(err)
}

// Get transaction receipt for deposit transaction on L1 network
l1Receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
if err != nil {
    log.Panic(err)
}

l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
if err != nil {
    log.Panic(err)
}
fmt.Println("L2 transaction", l2Tx.Hash)
```
And claim back failed deposit:
```go
cfdTx, err := w.ClaimFailedDeposit(nil, l2Tx.Hash)
if err != nil {
    fmt.Println(err) // this should trigger if deposit succeed
}
fmt.Println("ClaimFailedDeposit hash: ", cfdTx.Hash)
```

### Transfer

#### Transfer ETH token
```go
tx, err := w.Transfer(accounts.TransferTransaction{
    To:     common.HexToAddress(PublicKey2),
    Amount: big.NewInt(7_000_000_000_000_000_000),
    Token:  utils.EthAddress,
})
if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", tx.Hash())
```

#### Transfer ERC20 token
```go
TokenL2Address := common.HexToAddress("0x16D1b10bC0BEa66B47bEe0f4c72543fE79192f0f")
Receiver := common.HexToAddress("0xa61464658AfeAf65CccaaFD3a512b69A83B77618") 

tx, err := w.Transfer(accounts.TransferTransaction{
    To:     common.HexToAddress(Receiver),
    Amount: big.NewInt(1),
    Token:  TokenL2Address,
})
if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", tx.Hash())
```

### Withdraw

#### Withdraw ETH token
```go
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

#### Withdraw ERC20 token
```go
TokenL2Address   = common.HexToAddress("0x16D1b10bC0BEa66B47bEe0f4c72543fE79192f0f")

tx, err := w.Withdraw(accounts.WithdrawalTransaction{
		To:     w.Address(),
		Amount: big.NewInt(1),
		Token:  TokenL2Address,
})
if err != nil {
	log.Panic(err)
}
fmt.Println("Withdraw transaction: ", tx.Hash())
```

After the withdrawal is finished on L2 network, `FinalizeWithdraw` is required to
be executed on L1 network.
```go
// Wait until transaction is finalized
_, err = client.WaitFinalized(context.Background(), tx.Hash())
if err != nil {
    log.Panic(err)
}

// Perform finalize withdrawal
finalizeWithdrawTx, err := w.FinalizeWithdraw(nil, tx.Hash(), 0)
if err != nil {
    log.Panic(err)
}

fmt.Println("Finalize withdraw transaction: ", finalizeWithdrawTx.Hash())
```
Also, you can check its status with:
```go
isFinalized, err := w.IsWithdrawFinalized(nil, tx.Hash(), 0)
if err != nil {
    log.Panic(err)
}
fmt.Println("Is Withdraw finalized?", isFinalized)
```

### Deploy smart contract
```go
// Read smart contract bytecode
bytecode, err := os.ReadFile("Incrementer.zbin")
if err != nil {
    log.Panic(err)
}

// Get ABI
abi, err := incrementer.IncrementerMetaData.GetAbi()
if err != nil {
    log.Panic(err)
}

// Encode constructor arguments
constructor, err := abi.Pack("", big.NewInt(2))
if err != nil {
    log.Panicf("error while encoding constructor arguments: %s", err)
}

// Deploy paymaster contract
hash, err := w.DeployAccountWithCreate(accounts.CreateTransaction{
    Bytecode: bytecode,
    Calldata: constructor,
})
if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", hash)

// Wait unit transaction is finalized
_, err = client.WaitMined(context.Background(), hash)
if err != nil {
    log.Panic(err)
}

receipt, err := client.TransactionReceipt(context.Background(), hash)
if err != nil {
    log.Panic(err)
}
contractAddress := receipt.ContractAddress
fmt.Println("Paymaster address", contractAddress.String())
```

### Execute smart contract

#### Static execution using bindings
```go
incrementerContract, err := incrementer.NewIncrementer(contractAddress, client)
if err != nil {
	log.Panic(err)
}

// Execute Get method 
value, err := incrementerContract.Get(nil)
if err != nil {
	log.Panic(err)
}
fmt.Println("Value before Increment method execution: ", value)

// Start configuring transaction parameters
opts, err := bind.NewKeyedTransactorWithChainID(w.Signer().PrivateKey(), w.Signer().Domain().ChainId)
if err != nil {
	log.Panic(err)
}

// Execute Set method from storage smart contract with configured transaction parameters
tx, err := incrementerContract.Increment(opts)
if err != nil {
	log.Panic(err)
}
// Wait for transaction to be finalized
_, err = client.WaitMined(context.Background(), tx.Hash())
if err != nil {
	log.Panic(err)
}
```
#### Dynamic execution using EIP-712 transaction
```go
abiIncrementer, err := abi.JSON(strings.NewReader("<path to abi>")); 
if err != nil {
    log.Panic(err)
}
// Encode mint function from token contract
calldata, err := abiIncrementer.Pack("increment")
if err != nil {
	log.Panic(err)
}

hash, err := w.SendTransaction(context.Background(), &accounts.Transaction{
	To:   &TokenAddress,
	Data: calldata,
})
fmt.Println("Transaction: ", hash)
```
**In order to utilize Account Abstraction and Paymaster features, EIP-712 transactions must be used!**

### Deploy smart account 
```go
// Read paymaster contract from standard json
_, paymasterAbi, bytecode, err := utils.ReadStandardJson("Paymaster.json")
if err != nil {
    log.Panic(err)
}

// Encode paymaster constructor
constructor, err := paymasterAbi.Pack("", common.HexToAddress(TokenAddress))
if err != nil {
    log.Panic(err)
}

// Deploy paymaster contract
hash, err := w.DeployAccount(accounts.Create2Transaction{Bytecode: bytecode, Calldata: constructor})
if err != nil {
    log.Panic(err)
}
if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", hash)

// Wait unit transaction is finalized
_, err = client.WaitMined(context.Background(), hash)
if err != nil {
    log.Panic(err)
}

// Get address of deployed smart contract
contractAddress, err := utils.ComputeL2Create2Address(w.Address(), bytecode, constructor, nil)
if err != nil {
    log.Panic(err)
}
fmt.Println("Paymaster address: ", contractAddress.String())

```