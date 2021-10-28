[![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/go-ma.svg)](https://pkg.go.dev/github.com/MicahParks/go-ma) [![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/go-ma)](https://goreportcard.com/report/github.com/MicahParks
# go-ma
The Simple Moving Average (SMA), Exponential Moving (EMA), and Moving Average Convergence Divergence (MACD) algorithms
implemented in Golang.

If there are any other moving average algorithms you'd like implemented, please feel free to open an issue or a PR.
Example: Weighted Moving Average (WMA)

# Usage

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/MicahParks/go-ma
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkEMABig_Calculate-6             1000000000               0.0000272 ns/op
BenchmarkEMAFloat_Calculate-6           1000000000               0.0000009 ns/op
BenchmarkMACDSignalBig_Calculate-6      1000000000               0.0004847 ns/op
BenchmarkMACDSignalFloat_Calculate-6    1000000000               0.0000064 ns/op
BenchmarkMACDBig_Calculate-6            1000000000               0.0000389 ns/op
BenchmarkMACDFloat_Calculate-6          1000000000               0.0000019 ns/op
BenchmarkSMABig_Calculate-6             1000000000               0.0000476 ns/op
BenchmarkSMAFloat_Calculate-6           1000000000               0.0000009 ns/op
PASS
ok      github.com/MicahParks/go-ma     0.014s
```
