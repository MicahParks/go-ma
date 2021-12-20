[![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/go-ma.svg)](https://pkg.go.dev/github.com/MicahParks/go-ma) [![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/go-ma)](https://goreportcard.com/report/github.com/MicahParks/go-ma)
# go-ma
The Simple Moving Average (SMA), Exponential Moving (EMA), and Moving Average Convergence Divergence (MACD) technical
analysis algorithms implemented in Golang.

If there are any other moving average algorithms you'd like implemented, please feel free to open an issue or a PR.
Example: Weighted Moving Average (WMA)

```
import "github.com/MicahParks/go-ma"
```

# Usage
Please see the `examples` directory for full customizable examples.

## Create an MACD and signal EMA data structure
This will produce points on the MACD and signal EMA lines, as well as buy/sell signals.
```go
// Create a logger.
logger := log.New(os.Stdout, "", 0)

// Gather some data.
//
// For production systems, it'd be best to gather data asynchronously.
prices := testData()

// Create the MACD and signal EMA pair.
signal := ma.DefaultMACDSignal(prices[:ma.RequiredSamplesForDefaultMACDSignal])

// Iterate through the rest of the data and print the results.
var results ma.MACDSignalResults
for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal:] {
	results = signal.Calculate(p)

	// Interpret the buy signal.
	var buySignal string
	if results.BuySignal != nil {
		if *results.BuySignal {
			buySignal = "Buy, buy, buy!"
		} else {
			buySignal = "Sell, sell, sell!"
		}
	} else {
		buySignal = "Do nothing."
	}

	logger.Printf("Price index: %d\n  MACD: %.5f\n  Signal EMA: %.5f\n  Buy signal: %s", i+ma.RequiredSamplesForDefaultMACDSignal, results.MACD.Result, results.SignalEMA, buySignal)
}
```

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/MicahParks/go-ma
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkEMABig_Calculate-6             1000000000               0.0000272 ns/op
BenchmarkEMA_Calculate-6           1000000000               0.0000009 ns/op
BenchmarkMACDSignalBig_Calculate-6      1000000000               0.0004847 ns/op
BenchmarkMACDSignal_Calculate-6    1000000000               0.0000064 ns/op
BenchmarkMACDBig_Calculate-6            1000000000               0.0000389 ns/op
BenchmarkMACD_Calculate-6          1000000000               0.0000019 ns/op
BenchmarkSMABig_Calculate-6             1000000000               0.0000476 ns/op
BenchmarkSMA_Calculate-6           1000000000               0.0000009 ns/op
PASS
ok      github.com/MicahParks/go-ma     0.014s
```

# Resources
I built and tested this package based on the resources here:
* [Investopedia](https://www.investopedia.com/terms/m/macd.asp)
* [Invest Excel](https://investexcel.net/how-to-calculate-macd-in-excel/)
