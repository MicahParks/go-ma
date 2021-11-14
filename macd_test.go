package ma_test

import (
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewBigSMA(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewBigEMA(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewBigSMA(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewBigEMA(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewBigMACD(longEMA, shortEMA)

	for _, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func BenchmarkMACDFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMA(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMA(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMA(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMA(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACD(longEMA, shortEMA)

	for _, p := range prices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func TestMACDBig_Calculate(t *testing.T) {
	_, shortSMA := ma.NewBigSMA(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewBigEMA(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewBigSMA(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewBigEMA(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewBigMACD(longEMA, shortEMA)

	var res float64
	for i, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		res, _ = macd.Calculate(p).Result.Float64()
		if res != results[i] {
			t.FailNow()
		}
	}
}

func TestMACD_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMA(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMA(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMA(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMA(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACD(longEMA, shortEMA)

	for i, p := range prices[ma.DefaultLongMACDPeriod:] {
		if macd.Calculate(p).Result != results[i] {
			t.FailNow()
		}
	}
}
