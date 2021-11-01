package main

import (
	"log"
	"os"

	"github.com/MicahParks/go-ma"
)

func main() {
	// Create a logger.
	logger := log.New(os.Stdout, "", 0)

	// Gather some data.
	//
	// For production systems, it'd be best to gather data asynchronously.
	prices := testData()

	// Create the MACD and signal EMA pair.
	signal := ma.DefaultMACDSignalFloat(prices[:ma.RequiredSamplesForDefaultMACDSignal])

	// Iterate through the rest of the data and print the results.
	var results ma.MACDSignalResultsFloat
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
}

func testData() (prices []float64) {
	return []float64{
		459.99,
		448.85,
		446.06,
		450.81,
		442.8,
		448.97,
		444.57,
		441.4,
		430.47,
		420.05,
		431.14,
		425.66,
		430.58,
		431.72,
		437.87,
		428.43,
		428.35,
		432.5,
		443.66,
		455.72,
		454.49,
		452.08,
		452.73,
		461.91,
		463.58,
		461.14,
		452.08,
		442.66,
		428.91,
		429.79,
		431.99,
		427.72,
		423.2,
		426.21,
		426.98,
		435.69,
		434.33,
		429.8,
		419.85,
		426.24,
		402.8,
		392.05,
		390.53,
		398.67,
		406.13,
		405.46,
		408.38,
		417.2,
		430.12,
		442.78,
		439.29,
		445.52,
		449.98,
		460.71,
		458.66,
		463.84,
		456.77,
		452.97,
		454.74,
		443.86,
		428.85,
		434.58,
		433.26,
		442.93,
		439.66,
		441.35,
	}
}
