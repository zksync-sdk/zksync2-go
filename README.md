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
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/zksync-sdk/zksync2-go"
)

...

// first, init Ethereum Signer, from your mnemonic
es, err := zksync2.NewEthSignerFromMnemonic("<mnemonic words>", zksync2.ZkSyncChainIdMainnet)

// or from raw PrivateKey bytes
es, err = zksync2.NewEthSignerFromRawPrivateKey(pkBytes, zksync2.ZkSyncChainIdMainnet)

// also, init ZkSync Provider, specify ZkSync2 RPC URL (e.g. testnet)
zp, err := zksync2.NewDefaultProvider("https://zksync2-testnet.zksync.dev")

// then init Wallet, passing just created Ethereum Signer and ZkSync Provider   
w, err := zksync2.NewWallet(es, zp)

// init default RPC clients to Ethereum node (Goerli network in case of ZkSync2 testnet)
ethRpc, err := rpc.Dial("https://goerli.infura.io/v3/<your_infura_node_id>")

// and use it to create Ethereum Provider by Wallet 
ep, err := w.CreateEthereumProvider(ethRpc)
```

Now, using this three instances - ZkSync Provider, Ethereum Provider and Wallet, 
you can perform all basic actions with ZkSync network, just explore their public methods and package helpers.

Few examples below:

### Deposit
```go
tx, err := ep.Deposit(
    zksync2.CreateETH(),
    big.NewInt(1_000_000_000_000_000), 
    common.HexToAddress("<target_L2_address>"), 
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("L1 Tx hash", tx.Hash())
```
Also, you can get hash of corresponding L2 transaction:
```go
l1Receipt, err := ep.GetClient().TransactionReceipt(context.Background(), l1Hash)
if err != nil {
    panic(err)
}
l2Hash, err := ep.GetL2HashFromPriorityOp(l1Receipt)
if err != nil {
    panic(err)
}
fmt.Println("L2 Tx hash", l2Hash)
```
And claim back failed deposit:
```go
cfdHash, err := w.ClaimFailedDeposit(l2Hash, ep)
if err != nil {
    panic(err)
}
fmt.Println("ClaimFailedDeposit hash", cfdHash)
```

### Transfer
```go
hash, err := w.Transfer(
    common.HexToAddress("<target_L2_address>"), 
    big.NewInt(1000000000000),
    nil, 
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Tx hash", hash)
```

### Withdraw
```go
// first, you must initiate Withdraw 
wHash, err := w.Withdraw(
    common.HexToAddress("<target_L1_address>"), 
    big.NewInt(1000000000000), 
    nil, 
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Withdraw Tx hash", wHash)

// now you must wait until this Tx is being finalized
_, err = w.GetProvider().WaitFinalized(context.Background(), wHash)
if err != nil {
    panic(err)
}

// then, you need to call FinalizeWithdraw 
fwHash, err := w.FinalizeWithdraw(wHash, 0)
if err != nil {
    panic(err)
}
fmt.Println("FinalizeWithdraw Tx hash", fwHash)
```
Also, you can check its status with:
```go
iswf, err := w.IsWithdrawFinalized(wHash, 0)
if err != nil {
    panic(err)
}
fmt.Println("Is Withdraw finalized?", iswf)
```

### Deploy smart contract (basic case)
```go
hash, err := w.Deploy(bytecode, nil, nil, nil, nil)
if err != nil {
    panic(err)
}
fmt.Println("Tx hash", hash)

// use helper to get (compute) address of deployed SC
addr, err := zksync2.ComputeL2Create2Address(
	common.HexToAddress("<deployer_L2_address>"), 
	bytecode, 
	nil, 
	nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Deployed address", addr.String())
```

### Execute smart contract (basic case)
```go
// you need to encode calldata of executing function and its parameters
// or use ABI.Pack() method for loaded ABI of your contract ("github.com/ethereum/go-ethereum/accounts/abi")
calldata := crypto.Keccak256([]byte("get()"))[:4]
hash, err := w.Execute(
    common.HexToAddress("<contract_address>"),
    calldata,
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Tx hash", hash)
```

### Deploy smart account 
```go
hash, err := w.DeployAccount(bytecode, constructor, nil, nil)
if err != nil {
	log.Panic(err)
}

if err != nil {
    log.Panic(err)
}
fmt.Println("Transaction: ", hash)

// Wait unit transaction is finalized
_, err = zp.WaitMined(context.Background(), hash)
if err != nil {
    log.Panic(err)
}

// Get address of deployed smart contract
contractAddress, err := utils.ComputeL2Create2Address(
    w.GetAddress(),
    bytecode,
    constructor,
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Account address: ", contractAddress.String())

```

### Get balance of ZkSync account
```go
bal, err := w.GetBalance()
if err != nil {
    panic(err)
}
fmt.Println("Balance", bal)
```

### Get enhanced Transaction data for ZkSync Txs
```go
tx, err := zp.GetTransaction(txHash)
if err != nil {
    panic(err)
}
// explore zksync2.TransactionResponse struct
```

### Get enhanced Transaction Receipt data for ZkSync Txs
```go
receipt, err := zp.GetTransactionReceipt(txHash)
if err != nil {
    panic(err)
}
// explore zksync2.TransactionReceipt struct
```
