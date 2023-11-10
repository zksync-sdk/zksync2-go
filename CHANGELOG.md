## [0.3.2](https://github.com/zksync-sdk/zksync2-go/compare/v0.3.1...v0.3.2) (2023-11-10)


### Bug Fixes

* **accounts,types:** resolve errors in error messages ([167cf5b](https://github.com/zksync-sdk/zksync2-go/commit/167cf5bb6064fbac637f31a99bcef118c063d251))

## [0.3.1](https://github.com/zksync-sdk/zksync2-go/compare/v0.3.0...v0.3.1) (2023-08-21)


### Bug Fixes

* **accounts:** ensure the `Context` if it's `nil` in `WalletL2.Balance` ([e0d6064](https://github.com/zksync-sdk/zksync2-go/commit/e0d6064b02ae0b7a74a36d9e6c269d1ea703399b))
* **accounts:** use correct `Context` and block number in `WalletL1.BalanceL1` ([2af4513](https://github.com/zksync-sdk/zksync2-go/commit/2af451316a71e9d9d848319bb9a9c4da95bc5332)), closes [#22](https://github.com/zksync-sdk/zksync2-go/issues/22)

## [0.3.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.2.0...v0.3.0) (2023-06-30)

### Features

* refactor the design to be more geth friendly ([1e141c0](https://github.com/zksync-sdk/zksync2-go/commit/1e141c05b0d5209e3e5b440ec44a688d10432686))

### Bug Fixes

* update readme with account abstraction examples ([f321951](https://github.com/zksync-sdk/zksync2-go/commit/f3219511ef4be8d6b343ff6df6d441db4918d6e0))


## [0.2.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.1.1...v0.2.0) (2023-06-16)  

### Features

* repack the source code ([82dae4b](https://github.com/zksync-sdk/zksync2-go/commit/82dae4bb2b0643bb94643ce8800ab4a4975fa35c))
* handle error graceful ([8b6c93f](https://github.com/zksync-sdk/zksync2-go/commit/8b6c93f046efb6e3e7e2ee36b8d50021d04686b0))
* handle error when call EstimateAndSend ([a17dec6](https://github.com/zksync-sdk/zksync2-go/commit/a17dec6a5cd18f484976fe60af0b82a5309c044e))
* add account abstraction ([bd81fc3](https://github.com/zksync-sdk/zksync2-go/commit/bd81fc32e804371cd9d0a377c7f9ec28647fffdf))
* **provider:** add zksync rpc methods ([2f0081b](https://github.com/zksync-sdk/zksync2-go/commit/2f0081b2cc9c0a35fa84247ebbb926c40ddc992c))
