package ma_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDSignalBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMABig(bigPrices[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMABig(ma.DefaultSignalEMAPeriod, signalSMA, nil)
	for _, p := range bigPrices[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	signal, _ := ma.NewMACDSignalBig(macd, signalEMA, bigPrices[ma.DefaultLongMACDPeriod+1])

	for _, p := range bigPrices[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func BenchmarkMACDSignalFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMAFloat(prices[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMAFloat(ma.DefaultSignalEMAPeriod, signalSMA, 0)
	for _, p := range prices[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	signal, _ := ma.NewMACDSignalFloat(macd, signalEMA, prices[ma.DefaultLongMACDPeriod+1])

	for _, p := range prices[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func TestMACDSignalBig_Calculate(t *testing.T) {
	var latestShortEMA *big.Float
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	firstMACDResult := new(big.Float).Sub(latestShortEMA, longSMA)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, bigPrices[ma.DefaultLongMACDPeriod:ma.RequiredSamplesForDefaultMACDSignal-1], nil)

	signal, firstResults := ma.NewMACDSignalBig(macd, signalEMA, bigPrices[ma.RequiredSamplesForDefaultMACDSignal-1])

	if firstResults.BuySignal != signalResults[0] {
		if firstResults.BuySignal == nil || signalResults[0] == nil || *firstResults.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestMACDSignalFloat_Calculate(t *testing.T) {
	var latestShortEMA float64
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	firstMACDResult := latestShortEMA - longSMA

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, prices[ma.DefaultLongMACDPeriod:ma.RequiredSamplesForDefaultMACDSignal-1], 0)

	signal, firstResult := ma.NewMACDSignalFloat(macd, signalEMA, prices[ma.RequiredSamplesForDefaultMACDSignal-1])

	if firstResult.BuySignal != signalResults[0] {
		if firstResult.BuySignal == nil || signalResults[0] == nil || *firstResult.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalFloat(t *testing.T) {
	signal, firstResult := ma.DefaultMACDSignalFloat(prices[:ma.RequiredSamplesForDefaultMACDSignal])

	if firstResult.BuySignal != signalResults[0] {
		if firstResult.BuySignal == nil || signalResults[0] == nil || *firstResult.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalBig(t *testing.T) {
	signal, firstResult := ma.DefaultMACDSignalBig(bigPrices[:ma.RequiredSamplesForDefaultMACDSignal])

	if firstResult.BuySignal != signalResults[0] {
		if firstResult.BuySignal == nil || signalResults[0] == nil || *firstResult.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalFloatNil(t *testing.T) {
	signal, _ := ma.DefaultMACDSignalFloat(nil)
	if signal != nil {
		t.FailNow()
	}
}

func TestDefaultMACDSignalBigNil(t *testing.T) {
	signal, _ := ma.DefaultMACDSignalBig(nil)
	if signal != nil {
		t.FailNow()
	}
}

func TestDefaultMACDSignalCatchUpFloat(t *testing.T) {
	const catchUp = 5

	signal, firstResult := ma.DefaultMACDSignalFloat(prices[:ma.RequiredSamplesForDefaultMACDSignal+catchUp])

	if firstResult.BuySignal != signalResults[0] {
		if firstResult.BuySignal == nil || signalResults[0] == nil || *firstResult.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range prices[ma.RequiredSamplesForDefaultMACDSignal+catchUp:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1+catchUp]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}

func TestDefaultMACDSignalCatchUpBig(t *testing.T) {
	const catchUp = 5

	signal, firstResult := ma.DefaultMACDSignalBig(bigPrices[:ma.RequiredSamplesForDefaultMACDSignal+catchUp])

	if firstResult.BuySignal != signalResults[0] {
		if firstResult.BuySignal == nil || signalResults[0] == nil || *firstResult.BuySignal != *signalResults[0] {
			t.FailNow()
		}
	}

	for i, p := range bigPrices[ma.RequiredSamplesForDefaultMACDSignal+catchUp:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1+catchUp]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}
