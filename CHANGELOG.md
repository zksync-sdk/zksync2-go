# [1.0.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.7.0...v1.0.0) (2024-11-06)


### Bug Fixes

* **accounts:** implement lazy chaching in wallets ([60f84ab](https://github.com/zksync-sdk/zksync2-go/commit/60f84abc8964932d0d015c1e7a25ad5233cbcb93))


### Features

* **accounts:** add `Cache` which provides lazy caching for bridge contracts ([542647f](https://github.com/zksync-sdk/zksync2-go/commit/542647fe0526002e4cad6f33ebfdafa23a99e52b))
* **accounts:** add `PaymasterParams` and `GasPerPubdata` to `TransferCallMsg` ([8cf5594](https://github.com/zksync-sdk/zksync2-go/commit/8cf5594ce1e9e65d4023fd64fda044ebdc8f20ec))
* **accounts:** add `PaymasterParams` and `GasPerPubdata` to `WithdrawalCallMsg` ([d294db0](https://github.com/zksync-sdk/zksync2-go/commit/d294db0cb0479bced356fca84220b37bcc4f30e8))
* **accounts:** add `TransactionOpts` with ZKsync features ([f4321a3](https://github.com/zksync-sdk/zksync2-go/commit/f4321a393da4d8995ad715a5628dc5ecf3e3cd12))
* **accounts:** change `SmartAccount.Balance` to utilize `CallOpts` ([e4155d6](https://github.com/zksync-sdk/zksync2-go/commit/e4155d644bee419ff80f3a57afd8d704d0a6aaf6))
* **accounts:** change `WalletL2.Balance` to utilize `CallOpts` ([1d28994](https://github.com/zksync-sdk/zksync2-go/commit/1d28994a2cfb7a03f982984f708bde0aec93d6eb))
* **accounts:** remove `PaymasterParams` from transfer and withdrawal tx ([efc0f60](https://github.com/zksync-sdk/zksync2-go/commit/efc0f6067e5cb08f5b474583c5b125f30d452b46))
* **accounts:** rename `DepositTransaction` fields ([c070cfe](https://github.com/zksync-sdk/zksync2-go/commit/c070cfeb0a73f548a04f693c5ec6b9f0703e45a6))
* **accounts:** rename `NewBaseSignerFromMnemonic` method ([1b2bf80](https://github.com/zksync-sdk/zksync2-go/commit/1b2bf802ff14da77461520ee80ec5045cb708669))
* **accounts:** rename `Transaction` method ([fb1bbe9](https://github.com/zksync-sdk/zksync2-go/commit/fb1bbe9c43db39d578914c1ee761d43301d1b05c))
* **accounts:** rename `TransactionOpts` to `TransactionOptsL1` ([ba3ad1e](https://github.com/zksync-sdk/zksync2-go/commit/ba3ad1ec83204915da993e3d5ea0d6f1caa8d523))
* **accounts:** rename `WalletL1.ApproveERC20` to `WalletL1.ApproveToken` ([80dd564](https://github.com/zksync-sdk/zksync2-go/commit/80dd564b524a95aaee355f7a2f4cc5f75457bdf2))
* **accounts:** rename `WalletL1.EstimateCustomBridgeDepositL2Gas` ([f310974](https://github.com/zksync-sdk/zksync2-go/commit/f31097484092fbd45277ada1df1972cbfda4766b))
* **accounts:** rename `WalletL1.EstimateDefaultBridgeDepositL2Gas` ([bbc0b48](https://github.com/zksync-sdk/zksync2-go/commit/bbc0b488b265394f3fb0a222257981a4db51dc33))
* **accounts:** use `TransactionOptsL1` in `WalletL1` ([2757f4b](https://github.com/zksync-sdk/zksync2-go/commit/2757f4bb2a8ee7dc2fa177f796a5c17013d90897))
* **accounts:** use `TransactOpts` in `Deployer` ([5491f63](https://github.com/zksync-sdk/zksync2-go/commit/5491f639b93b103516da0fe8538a3256b84f0317))
* **accounts:** use zksync transaction for transfer and withdrawal ([99762f5](https://github.com/zksync-sdk/zksync2-go/commit/99762f56c95cbeda0278dd3965ece760bc3c4324))
* **clients:** add `CustomSignature` and `GasPerPubdata` to `TransferCallMsq` ([df00687](https://github.com/zksync-sdk/zksync2-go/commit/df006870c5eb50b27e6468d1465ee4cd7122e4d7))
* **clients:** add `CustomSignature` and `GasPerPubdata` to `WithdrawalCallMsq` ([4cd0dcf](https://github.com/zksync-sdk/zksync2-go/commit/4cd0dcfb971c00abeb0e71b7a17820a2d9e94cf2))
* **clients:** rename `TransferCallMsg` methods ([487b902](https://github.com/zksync-sdk/zksync2-go/commit/487b9022113da38912a067cb28c8c58421902c93))
* **clients:** rename `WithdrawalCallMsg` methods ([bd63b3d](https://github.com/zksync-sdk/zksync2-go/commit/bd63b3dd1840dde38245ae7c7af241204a58a5b8))
* **eip712:** change the API of `Domain` struct ([670010a](https://github.com/zksync-sdk/zksync2-go/commit/670010aa040ab89600501ece8703ce11b9ca9df1))
* **eip712:** remove `TypedData` interface ([9f69f2e](https://github.com/zksync-sdk/zksync2-go/commit/9f69f2eee913d2427457328956618dcc2921e067))
* redesign `Signer` and replace `BaseSigner` with `ECDSASigner` ([a38b8bc](https://github.com/zksync-sdk/zksync2-go/commit/a38b8bc5d956782d41ddf0eda8ab7646adf401f2))
* remove `AllowList` field from transaction and message types ([571fb96](https://github.com/zksync-sdk/zksync2-go/commit/571fb9636b04ba83238477b28d0c2038c18e4826))
* remove `Deployer`, `AdapterL1` and `AdapterL1` interfaces ([dfcc937](https://github.com/zksync-sdk/zksync2-go/commit/dfcc937b5fb6e20da13e5dd9ebffeb0461eca2db))
* remove `EIP712Meta` struct and inline its properties in `CallMsg` ([4e48d91](https://github.com/zksync-sdk/zksync2-go/commit/4e48d9187254c569c6acc0959793c26c5136cfd7))
* remove deprecated code ([269f787](https://github.com/zksync-sdk/zksync2-go/commit/269f78797dc86bcb52d983ce80113a948fa20cff))
* rename `BaseClient` to `Client` ([e91f3bc](https://github.com/zksync-sdk/zksync2-go/commit/e91f3bc918718af6e60cd45f5f70161b7833108b))
* **types:** improve the design of L2 transaction ([c2aa297](https://github.com/zksync-sdk/zksync2-go/commit/c2aa2975af618992db355c1c1355ddddfe1399e8))
* **types:** rename `EIP712TxType` to `TxType` ([f009818](https://github.com/zksync-sdk/zksync2-go/commit/f0098188c09751fa1f433745e3ae262c8e0483ac))
* **types:** rename the fields for default ERC20 bridges in `BridgeContracts` ([5fb3925](https://github.com/zksync-sdk/zksync2-go/commit/5fb3925f0882a835cf5b0fd199134874dfd3d817))
* **types:** update `FeeParams` field types from `uint64` to `big.Int` ([02d02a2](https://github.com/zksync-sdk/zksync2-go/commit/02d02a2398d5dc66d21eaec72cd3c916b3548327))
* **types:** use `big.Int` for block and batch number in `BlockDetails` ([8fb632b](https://github.com/zksync-sdk/zksync2-go/commit/8fb632b15b00fc9a62fdfb3387e429f2a4b34432))
* **types:** use `big.Int` for gas prices and block numbers in `BatchDetails` ([02a6e8e](https://github.com/zksync-sdk/zksync2-go/commit/02a6e8ea450171f20b9db51c866952e90819be67))


### BREAKING CHANGES

* **accounts:** Rename `NewBaseSignerFromMnemonic` method to
`NewECDSASignerFromMenmonic`.
* **types:** Change `FeeParams` field types from `uint64` to
`big.Int`.
* **accounts:** Rename `Transaction.ToTransaction712` to
`Transaction.ToTransaction`.
* **accounts:** Renamed `DepositTransaction` fields:
`ApproveERC20` to `ApproveToken`, `ApproveBaseERC20` to
`ApproveBaseToken`.
* **accounts:** Rename `WalletL1.ApproveERC20` to `WalletL1.ApproveToken`.
* **accounts:** The `SmartAccount.Balance` method utilizes `CallOpts`.
* **accounts:** The `WalletL2.Balance` method utilizes `CallOpts`.
* **accounts:** Remove `PaymasterParams` from `TransferTransaction`
and `WithdrawalTransaction` since it can be utilze using `TransactionOpts`.
* **accounts:** Use `TransactionOptsL1` in `WalletL1` when
interacting with contract API.
* **accounts:** The `TransactionOpts` is renamed to
`TransactionOptsL1` since it's changed to be used only for
L1 interaction.
* **accounts:** `WalletL2` and `Wallet` for `transfer` and
`withdraw` methods use zksync transaction in order to provide
utilization of zksync features.
* **clients:** Rename `WithdrawalCallMsg.ToCallMsg` to
`WithdrawalCallMsg.ToL1CallMsg` and `WithdrawalCallMsg.ToZkCallMsg`
to `WithdrawalCallMsg.ToCallMsg`.
* **clients:** Rename `TransferCallMsg.ToCallMsg` to
`TransferCallMsg.ToL1CallMsg` and `TransferCallMsg.ToZkCallMsg` to
`TransferCallMsg.ToCallMsg`.
* **eip712:** The `eip712.Domain` methods have been renamed
to be more convenient.
* **eip712:** Removed `eip712.TypedData` interface.
* **types:** Rename properties `L1Erc20DefaultBridge` to `L1Erc20Bridge`
and `L2Erc20DefaultBridge` to `L2Erc20Bridge` in `BridgeContracts`.
* **types:** Use `big.Int` type in `BlockDetails` for `Number`
and `L1BatchNumber`.
* **types:** Use `big.Int` type in `BatchDetails` for `Number`,
`L2FairGasPrice` and `L1GasPrice`.
* The `EIP712Meta` struct has been removed, and its properties
are now directly inlined into the `CallMsg` struct.
* **types:** Rename constant `EIP712TxType` to `TxType`.
* The `Signer` interface has been redesigned to be minimal and more
extensible. The `BaseSigner` has been replaced by `ECDSASigner`. Additionally,
`Wallet`, `WalletL1`, and `WalletL2` now use `ECDSASigner` in their constructors.
* **types:** `Transaction712` has been replaced by `Transaction`, an improved version
that offers a more convenient API and greater extensibility.
* **accounts:** Rename `WalletL1.EstimateCustomBridgeDepositL2Gas` to
`WalletL1.EstimateDepositL2GasFromCustomBridge`. Added the missing
`l2Value` parameter.
* **accounts:** Rename `WalletL1.EstimateDefaultBridgeDepositL2Gas` to
`WalletL1.EstimateDepositL2GasFromDefaultBridge`.
* Remove `Deployer`, `AdapterL1` and `AdapterL2`
interfaces. Rename `BaseDeployer` to `Deployer`.
* Remove `Client` interface and rename `BaseClient`
to `Client`.
* Remove `AllowList` field from transaction and message types.
* Remove deprecated code.

# [0.7.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.6.0...v0.7.0) (2024-07-27)


### Features

* **clients:** add `BaseClient.BalanceAtByTag` method ([abff602](https://github.com/zksync-sdk/zksync2-go/commit/abff602206fdab12bddff57710b63da3e5be7691))
* **clients:** add `BaseClient.BlockByTag` method ([d86c806](https://github.com/zksync-sdk/zksync2-go/commit/d86c806ccb5054c9c1305b0b66cf676fe0bb68f7))
* **clients:** add `BaseClient.CallContractByTag` method ([1d64eab](https://github.com/zksync-sdk/zksync2-go/commit/1d64eab4fe672021b999d85b4bf47a1146161611))
* **clients:** add `BaseClient.CodeAtByTag` method ([b8ab9e4](https://github.com/zksync-sdk/zksync2-go/commit/b8ab9e4acb14f397b019b2e52e5ab4bbce2f5ca3))
* **clients:** add `BaseClient.HeaderByTag` method ([12cddcd](https://github.com/zksync-sdk/zksync2-go/commit/12cddcda8fe0434cb45a3533ff88bdf69c8c9489))
* **clients:** add `BaseClient.IsMessageSignatureCorrect` ([e2d1312](https://github.com/zksync-sdk/zksync2-go/commit/e2d131291b0431af9b48bedcbfdd75d095fbe030))
* **clients:** add `BaseClient.IsTypedDataSignatureCorrect` ([f95918a](https://github.com/zksync-sdk/zksync2-go/commit/f95918a412db44035932cfba6bcb3172dd064d37))
* **clients:** add `BaseClient.NonceAtByTag` method ([1dc6a59](https://github.com/zksync-sdk/zksync2-go/commit/1dc6a59cc213cdf9ba9d34dc237a31da3769b565))
* **clients:** add `BaseClient.StorageAtByTag` method ([24c3361](https://github.com/zksync-sdk/zksync2-go/commit/24c3361e5e447e5a2136cc87cc291e6c4ff76958))
* **clients:** add `BaseClient.TransactionCountByTag` method ([b01170e](https://github.com/zksync-sdk/zksync2-go/commit/b01170eb3a76b4e014beb2e7439917263d4028bd))
* **utils:** add `HashedL2ToL1Msg` method ([98e7825](https://github.com/zksync-sdk/zksync2-go/commit/98e7825b9cc63c7798a245bc83abf51d61950fbb))
* **utils:** add `IsMessageSignatureCorrect` ([001cfef](https://github.com/zksync-sdk/zksync2-go/commit/001cfef3705ef5f7d71c6de0d4871c13e198264a))
* **utils:** add `IsTypedDataSignatureCorrect` function ([09d9b5d](https://github.com/zksync-sdk/zksync2-go/commit/09d9b5d3916a62aa8563d246c68f0019c00eacec))

# [0.6.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.5.0...v0.6.0) (2024-06-24)


### Features

* align types and RPC endpoints with version `v24.7.0` of a node ([a411218](https://github.com/zksync-sdk/zksync2-go/commit/a41121803b02d455e3040d62b177db422ce167ce))
* **clients,types:** add support for `zks_getFeeParams` RPC method ([b6bce81](https://github.com/zksync-sdk/zksync2-go/commit/b6bce81215662ea691619a7b22c578678717eb9c))
* **clients,types:** add support for `zks_getProtocolVersion` RPC method ([72f2247](https://github.com/zksync-sdk/zksync2-go/commit/72f22473ebe4543094ba5c232b2918212cad66e2))
* **clients,types:** add support for `zks_sendRawTransactionWithDetailedOutput` RPC method ([fa9c571](https://github.com/zksync-sdk/zksync2-go/commit/fa9c571722ad92f099f8ce54e5ba7c326efd27a8))
* **clients:** add `BaseClient.IsL2BridgeLegacy` method ([8209dc3](https://github.com/zksync-sdk/zksync2-go/commit/8209dc3e380fbd93ce6584f7f85f45239ce0dd15))
* **clients:** add `BaseClient.L2TokenAddressFromCustomBridge()` method ([ec03104](https://github.com/zksync-sdk/zksync2-go/commit/ec031047d572ae404c468b2f232c64dc079e2c22))
* provider support for custom shared bridge ([cfdd543](https://github.com/zksync-sdk/zksync2-go/commit/cfdd54382261bb065991ed437d9281546409ec3b))


### BREAKING CHANGES

* Types and RPC endpoints are aligned with version
`v24.7.0` of a node.

# [0.5.0](https://github.com/zksync-sdk/zksync2-go/compare/v0.4.0...v0.5.0) (2024-06-09)


### Bug Fixes

* **accounts:** use correct token when trasfering and withdrawing ETH ([573779b](https://github.com/zksync-sdk/zksync2-go/commit/573779b57436fdc5a0e5777dadec4859f903f99a))


### Features

* add `SmartAccount` to provide better support for AA ([65a5686](https://github.com/zksync-sdk/zksync2-go/commit/65a56869e2c74ba48045b2d8b56223a11e65b41a))
* add Bridgehub support ([82d6474](https://github.com/zksync-sdk/zksync2-go/commit/82d647411892496838457daa97e7584a68a33bb5))
* **clients:** add `BaseClient.RawBlockTransactions` and `BaseClient.BytecodeByHash` ([0a41362](https://github.com/zksync-sdk/zksync2-go/commit/0a41362300e9d53051f276a34423bf736e1b8f36))
* **types:** add `Transaction712.Decode()` for decoding serialized transactions ([9b6dbc1](https://github.com/zksync-sdk/zksync2-go/commit/9b6dbc128394fa6214d052dc5a0dfa5b0ada51bd)), closes [#49](https://github.com/zksync-sdk/zksync2-go/issues/49)

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
