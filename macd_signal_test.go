package ma_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDSignalBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewBigSMA(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewBigEMA(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewBigSMA(bigPrices[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewBigEMA(ma.DefaultSignalEMAPeriod, signalSMA, nil)
	for _, p := range bigPrices[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewBigSMA(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewBigEMA(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewBigMACD(longEMA, shortEMA)

	signal, _ := ma.NewBigMACDSignal(macd, signalEMA, bigPrices[ma.DefaultLongMACDPeriod+1])

	for _, p := range bigPrices[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func BenchmarkMACDSignalFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMA(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMA(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMA(prices[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMA(ma.DefaultSignalEMAPeriod, signalSMA, 0)
	for _, p := range prices[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMA(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMA(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACD(longEMA, shortEMA)

	signal, _ := ma.NewMACDSignal(macd, signalEMA, prices[ma.DefaultLongMACDPeriod+1])

	for _, p := range prices[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func TestMACDSignalBig_Calculate(t *testing.T) {
	var latestShortEMA *big.Float
	_, shortSMA := ma.NewBigSMA(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewBigEMA(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewBigSMA(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewBigEMA(ma.DefaultLongMACDPeriod, longSMA, nil)

	firstMACDResult := new(big.Float).Sub(latestShortEMA, longSMA)

	macd := ma.NewBigMACD(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, bigPrices[ma.DefaultLongMACDPeriod:ma.RequiredSamplesForDefaultMACDSignal-1], nil)

	signal, _ := ma.NewBigMACDSignal(macd, signalEMA, bigPrices[ma.RequiredSamplesForDefaultMACDSignal-1])

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestMACDSignal_Calculate(t *testing.T) {
	var latestShortEMA float64
	_, shortSMA := ma.NewSMA(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMA(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMA(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMA(ma.DefaultLongMACDPeriod, longSMA, 0)

	firstMACDResult := latestShortEMA - longSMA

	macd := ma.NewMACD(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, prices[ma.DefaultLongMACDPeriod:ma.RequiredSamplesForDefaultMACDSignal-1], 0)

	signal, _ := ma.NewMACDSignal(macd, signalEMA, prices[ma.RequiredSamplesForDefaultMACDSignal-1])

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignal(t *testing.T) {
	signal := ma.DefaultMACDSignal(prices[:ma.RequiredSamplesForDefaultMACDSignal])

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalBig(t *testing.T) {
	signal := ma.DefaultBigMACDSignal(bigPrices[:ma.RequiredSamplesForDefaultMACDSignal])

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalNil(t *testing.T) {
	signal := ma.DefaultMACDSignal(nil)
	if signal != nil {
		t.FailNow()
	}
}

func TestDefaultMACDSignalBigNil(t *testing.T) {
	signal := ma.DefaultBigMACDSignal(nil)
	if signal != nil {
		t.FailNow()
	}
}

func TestDefaultMACDSignalCatchUp(t *testing.T) {
	const catchUp = 5

	signal := ma.DefaultMACDSignal(prices[:ma.RequiredSamplesForDefaultMACDSignal+catchUp])

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal+catchUp:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+catchUp]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalCatchUpBig(t *testing.T) {
	const catchUp = 5

	signal := ma.DefaultBigMACDSignal(bigPrices[:ma.RequiredSamplesForDefaultMACDSignal+catchUp])

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal+catchUp:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+catchUp]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}
