## [1.0.1](https://github.com/danijelTxFusion/zksync2-go/compare/v1.0.0...v1.0.1) (2023-08-16)


### Bug Fixes

* **t.txt:** add some text ([e2cc7bf](https://github.com/danijelTxFusion/zksync2-go/commit/e2cc7bfeee0b713436487e158ac7e5eccddb04be))

# [1.0.0](https://github.com/danijelTxFusion/zksync2-go/compare/v0.4.2...v1.0.0) (2023-08-16)


### Features

* refactor the client ([4a02716](https://github.com/danijelTxFusion/zksync2-go/commit/4a02716c1a4ea6aa3e8380eaae6160b33f53fc31))


### BREAKING CHANGES

* `Client` interface
`EstimateGas` method that is part of `Client` interface now receives
additional parameter to override gas price and gas limit.
