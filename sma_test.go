package ma_test

import (
	"testing"

	"github.com/MicahParks/go-ma"
)

const (
	testPeriod = 10
)

func BenchmarkSMABig_Calculate(b *testing.B) {
	sma, _ := ma.NewBigSMA(bigPrices[:testPeriod])

	for _, p := range bigPrices[testPeriod:] {
		sma.Calculate(p)
	}
}

func BenchmarkSMAFloat_Calculate(b *testing.B) {
	sma, _ := ma.NewSMA(prices[:testPeriod])

	for i, p := range prices[testPeriod:] {
		if smaResults[i] != sma.Calculate(p) {
			b.FailNow()
		}
	}
}

func TestSMABig_Calculate(t *testing.T) {
	sma, _ := ma.NewBigSMA(bigPrices[:testPeriod])

	for _, p := range bigPrices[testPeriod:] {
		sma.Calculate(p)
	}
}

func TestSMA_Calculate(t *testing.T) {
	sma, _ := ma.NewSMA(prices[:testPeriod])

	for i, p := range prices[testPeriod:] {
		if smaResults[i] != sma.Calculate(p) {
			t.FailNow()
		}
	}
}
