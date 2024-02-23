# [0.4.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.3.2...v0.4.0) (2024-02-23)


### Bug Fixes

* **accounts:** `FullRequiredDepositFee` use default bridge address when missing `msg.BridgeAddress` ([840c653](https://github.com/zksync-sdk/zksync2-go/commit/840c6537f8c558b3bdb2861620e1e3f8f2a31351))
* **accounts:** `WalletL2.SignTransaction` populates transaction before signing ([13f8d7f](https://github.com/zksync-sdk/zksync2-go/commit/13f8d7f6f95b7f2dd4c572e533d54e040ed463fe))
* **accounts:** change default priority fee to 0 for zksync tx ([37dbd47](https://github.com/zksync-sdk/zksync2-go/commit/37dbd47a2cdf6c2b935d68e3fa11c8a84e2e58cd))
* **accounts:** resolve error in formating string ([6be5c6b](https://github.com/zksync-sdk/zksync2-go/commit/6be5c6bb3d1e7354cf773e00fc505abaf9ac1d58))
* **clients:** allow `requestExecute` transactions in block without signatrues ([1744ee7](https://github.com/zksync-sdk/zksync2-go/commit/1744ee769003ec13ceb5dbf594f5fca75cefdd77))
* **clients:** deprecate `client.ConfirmedTokens` ([33372c4](https://github.com/zksync-sdk/zksync2-go/commit/33372c43294d570e78efe4f27acef1a73017fb0b))
* **clients:** deprecate `client.TokenPrice` ([d5acec3](https://github.com/zksync-sdk/zksync2-go/commit/d5acec3ca932989cf36892163c42af0f0a6ecbab))
* **clients:** patch to work with `geth:1.13.13` version [#27](https://github.com/zksync-sdk/zksync2-go/issues/27) ([8ee9d08](https://github.com/zksync-sdk/zksync2-go/commit/8ee9d088c5b39a137f79549e22459042269555b2))
* resolve colision among dependencies [#29](https://github.com/zksync-sdk/zksync2-go/issues/29) ([dc8ad7a](https://github.com/zksync-sdk/zksync2-go/commit/dc8ad7a343d85f3675f20b88b9dc29b8ebf62f6b))
* **types:** aligned types with `zksync-core` v18 ([191cc74](https://github.com/zksync-sdk/zksync2-go/commit/191cc74d9b1b2eba1b9e2e0c41f8cb5c8c92f39c))
* **types:** extend bridge contracts with WETH bridges ([a21f294](https://github.com/zksync-sdk/zksync2-go/commit/a21f294ecf3c071016585f1d33e96489810a727b))
* **utils:** pass bytes to `GetGeneralPaymasterInput` function [#30](https://github.com/zksync-sdk/zksync2-go/issues/30) ([8e62183](https://github.com/zksync-sdk/zksync2-go/commit/8e621834994c137d524e598d0f3bb60990ca1ece))


### Features

* **accounts:** add `WalletL2.DeploymentNonce` for fetching deployment nonce ([8099778](https://github.com/zksync-sdk/zksync2-go/commit/809977882e6d9ee17d52ffcfccf5c8fa7497886a))
* **clients:** add `BaseClient.Proof` method for fetching storage proofs ([9e06517](https://github.com/zksync-sdk/zksync2-go/commit/9e06517b4fdb01d776bb0634fb8e0b2b29703fef))
* update contracts to align with version `v20.4.0` ([463a6f3](https://github.com/zksync-sdk/zksync2-go/commit/463a6f3c10215d193a70473ffd1041cb71e2b3ba))
* **utils:** add `NonceHolderAddress` variable ([0365a37](https://github.com/zksync-sdk/zksync2-go/commit/0365a37903389a5e3e032ac8ed4a95d82e21869f))

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
