package main

import (
	"log"
	"os"

	"github.com/MicahParks/go-ma"
)

// This example is functionally equivalent to the default signal, but it can be modified for customization.
// These customizations can include period lengths and EMA smoothing constants.
// Be careful to correctly index any slices during creations and read all function/method documentation.
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
	_, shortSMA := ma.NewSMA(prices[:ma.DefaultShortMACDPeriod])

	// Create the short EMA data structure.
	shortEMA := ma.NewEMA(ma.DefaultShortMACDPeriod, shortSMA, 0)

	// Save the last value of the short EMA for the first MACD calculation.
	var latestShortEMA float64

	// Catch up the short EMA to the period where the long EMA plus the signal EMA will be at.
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	// Create the long EMA.
	_, longSMA := ma.NewSMA(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMA(ma.DefaultLongMACDPeriod, longSMA, 0)

	// The first result returned from calculating the MACD will be the second possible MACD result. To get the first
	// possible MACD result, use the most recent short and long EMA values. For the long EMA value, this will be
	// equivalent to the most recent long SMA value.
	firstMACDResult := latestShortEMA - longSMA

	// Create the MACD from the short and long EMAs.
	macd := ma.NewMACD(longEMA, shortEMA)

	// Create the signal EMA.
	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, prices[ma.DefaultLongMACDPeriod:ma.RequiredSamplesForDefaultMACDSignal-1], 0)

	// Create the signal from the MACD and signal EMA.
	signal, results := ma.NewMACDSignal(macd, signalEMA, prices[ma.RequiredSamplesForDefaultMACDSignal-1])
	buySignal := "Do nothing." // The first result's buy signal will never buy a buy/sell.
	logger.Printf("Period index: %d\n  Buy Signal: %s\n  MACD: %.5f\n  Signal EMA: %.5f", ma.RequiredSamplesForDefaultMACDSignal-1, buySignal, results.MACD.Result, results.SignalEMA)

	// Use the remaining data to generate the signal results for each period.
	for i := ma.RequiredSamplesForDefaultMACDSignal; i < len(prices); i++ {
		results = signal.Calculate(prices[i])

		// Interpret the buy signal.
		if results.BuySignal != nil {
			if *results.BuySignal {
				buySignal = "Buy, buy, buy!"
			} else {
				buySignal = "Sell, sell, sell!"
			}
		} else {
			buySignal = "Do nothing."
		}

		logger.Printf("Period index: %d\n  Buy Signal: %s\n  MACD: %.5f\n  Signal EMA: %.5f", i, buySignal, results.MACD.Result, results.SignalEMA)
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
