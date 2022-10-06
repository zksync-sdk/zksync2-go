ZkSync 2 Golang SDK
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

// init default RPC client to Ethereum node (Goerli network in case of ZkSync2 testnet)
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
    big.NewInt(1000000000000000), 
    common.HexToAddress("<target_L2_address>"), 
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Tx hash", tx.Hash())
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
hash, err := w.Withdraw(
    common.HexToAddress("<target_L1_address>"), 
    big.NewInt(1000000000000), 
    nil, 
    nil,
)
if err != nil {
    panic(err)
}
fmt.Println("Tx hash", hash)
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

### Get balance of ZkSync account
```go
bal, err := w.GetBalance()
if err != nil {
    panic(err)
}
fmt.Println("Balance", bal)
```