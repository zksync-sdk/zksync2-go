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
