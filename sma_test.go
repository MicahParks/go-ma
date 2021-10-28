package ma_test

import (
	"testing"

	"github.com/MicahParks/go-ma"
)

const (
	testPeriod = 10
)

func BenchmarkSMABig_Calculate(b *testing.B) {
	sma, _ := ma.NewSMABig(bigPrices[:testPeriod])

	for _, p := range bigPrices[testPeriod:] {
		sma.Calculate(p)
	}
}

func BenchmarkSMAFloat_Calculate(b *testing.B) {
	sma, _ := ma.NewSMAFloat(prices[:testPeriod])

	for i, p := range prices[testPeriod:] {
		if smaResults[i] != sma.Calculate(p) {
			b.FailNow()
		}
	}
}

func TestSMABig_Calculate(t *testing.T) {
	sma, _ := ma.NewSMABig(bigPrices[:testPeriod])

	for _, p := range bigPrices[testPeriod:] {
		sma.Calculate(p)
	}
}

func TestSMAFloat_Calculate(t *testing.T) {
	sma, _ := ma.NewSMAFloat(prices[:testPeriod])

	for i, p := range prices[testPeriod:] {
		if smaResults[i] != sma.Calculate(p) {
			t.FailNow()
		}
	}
}
