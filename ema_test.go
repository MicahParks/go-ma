package ma_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ma"
)

var (
	bigTestSmoothing = big.NewFloat(ma.DefaultEMASmoothing)
)

func BenchmarkEMABig_Calculate(b *testing.B) {
	_, sma := ma.NewSMABig(bigPrices[:testPeriod])

	ema := ma.NewEMABig(testPeriod, sma, bigTestSmoothing)

	for _, p := range bigPrices[testPeriod:] {
		ema.Calculate(p)
	}
}

func BenchmarkEMAFloat_Calculate(b *testing.B) {
	_, sma := ma.NewSMAFloat(prices[:testPeriod])

	ema := ma.NewEMAFloat(testPeriod, sma, ma.DefaultEMASmoothing)

	for _, p := range prices[testPeriod:] {
		ema.Calculate(p)
	}
}
