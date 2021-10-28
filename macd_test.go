package ma_test

import (
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	for _, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func BenchmarkMACDFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	for _, p := range prices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func TestMACDBig_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	var res float64
	for i, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		res, _ = macd.Calculate(p).Result.Float64()
		if res != results[i] {
			t.FailNow()
		}
	}
}

func TestMACDFloat_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	for i, p := range prices[ma.DefaultLongMACDPeriod:] {
		if macd.Calculate(p).Result != results[i] {
			t.FailNow()
		}
	}
}
