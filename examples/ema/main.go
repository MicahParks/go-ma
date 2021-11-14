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

	// Determine the number of periods for the calculating the EMA.
	const periods = 10

	// Get the first result of the SMA. The first result of the SMA is the same as the first result of the EMA.
	//
	// The length of the initial slice is the number of periods used for calculating the SMA.
	_, sma := ma.NewSMA(prices[:periods])
	logger.Printf("Period index: %d\nPrice: %.2f\nEMA: %.2f", periods-1, prices[periods-1], sma)

	// Create the EMA data structure.
	ema := ma.NewEMA(periods, sma, 0)

	// Use the remaining data to generate the SMA for each period.
	var result float64
	for i := periods; i < len(prices); i++ {
		result = ema.Calculate(prices[i])
		logger.Printf("Period index: %d\nPrice: %.2f\nEMA: %.2f", i, prices[i], result)
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
