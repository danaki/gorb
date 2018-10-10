# GORB (PoC)
Matching engine based on a limit order book written in Go

## Install

`go get ./...`

## TODO:
* [ ] Order cancelling
* [ ] Persistence (either with in-memory cache or not)
* [ ] Transactions? (https://github.com/claygod/transaction)
* [ ] 3rd party library for data structures (https://github.com/Workiva/go-datastructures or GODS)
* [ ] Freeze and refactor API + more tests with mocks
* [ ] REST API
* [ ] Memory leaks
* [ ] Benchmarks
* [ ] Concurency
* [ ] Histogram
* [ ] Fan events via websockets and/or message bus?
* [ ] Logging, Metrics (https://github.com/rcrowley/go-metrics) and last price (InfluxDB?)
* [ ] Fine-grained memory usage control + disk/flash swap for orders with low probability
* [ ] Memory pool for orders vs garbage collector
* [ ] Big numbers support
* [ ] Many tickers
