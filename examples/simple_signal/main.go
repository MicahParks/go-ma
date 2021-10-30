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

	// Get the first result of the short SMA. The first result of the SMA is the same as the first result of the EMA.
	//
	// The length of the initial slice is the number of periods used for calculating the SMA.
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])

	// Create the short EMA data structure.
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)

	// Save the last value of the short EMA for the first MACD calculation.
	var mostRecentShortEMA float64

	// Catch up the short EMA to the period where the long EMA plus the signal EMA will be at.
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		mostRecentShortEMA = shortEMA.Calculate(p)
	}

	// Create the long EMA.
	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	// The first result returned from calculating the MACD will be the second possible MACD result. To get the first
	// possible MACD result, use the most recent short and long EMA values. For the long EMA value, this will be
	// equivalent to the most recent long SMA value.
	firstMACDResult := mostRecentShortEMA - longSMA

	// Create the MACD from the short and long EMAs.
	macd := ma.NewMACDFloat(longEMA, shortEMA)

	// Create the signal EMA.
	signalEMA, signalResult, macdResults := macd.SignalEMA(firstMACDResult, prices[ma.DefaultLongMACDPeriod:ma.DefaultLongMACDPeriod+ma.DefaultSignalEMAPeriod-1], 0)
	logger.Printf("Period index: %d\n  Buy Signal: %v\n  MACD: %.5f\n  Signal EMA: %.5f", ma.DefaultLongMACDPeriod+ma.DefaultSignalEMAPeriod-1, nil, macdResults[len(macdResults)-1].Result, signalResult)

	// Create the signal from the MACD and signal EMA.
	signal, results := ma.NewMACDSignalFloat(macd, signalEMA, prices[ma.DefaultLongMACDPeriod+ma.DefaultSignalEMAPeriod-1]) // TODO Wrong index, probably.
	logger.Printf("Period index: %d\n  Buy Signal: %v\n  MACD: %.5f\n  Signal EMA: %.5f", ma.DefaultLongMACDPeriod+ma.DefaultSignalEMAPeriod-1, results.BuySignal, results.MACD.Result, results.SignalEMA)

	// Use the remaining data to generate the signal results for each period.
	for i := ma.DefaultLongMACDPeriod + ma.DefaultSignalEMAPeriod; i < len(prices); i++ {
		results = signal.Calculate(prices[i])
		if results.BuySignal != nil {
			println("boi")
		}
		logger.Printf("Period index: %d\n  Buy Signal: %v\n  MACD: %.5f\n  Signal EMA: %.5f", i, results.BuySignal, results.MACD.Result, results.SignalEMA)
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
