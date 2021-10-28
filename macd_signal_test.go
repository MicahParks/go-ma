package ma_test

import (
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDSignalBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPricesForSignal[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMABig(ma.DefaultSignalEMAPeriod, signalSMA, nil)
	for _, p := range bigPricesForSignal[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	signal, _ := ma.NewMACDSignalBig(macd, signalEMA, bigPricesForSignal[ma.DefaultLongMACDPeriod+1])

	for _, p := range bigPricesForSignal[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func BenchmarkMACDSignalFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range pricesForSignal[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMAFloat(ma.DefaultSignalEMAPeriod, signalSMA, 0)
	for _, p := range pricesForSignal[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	signal, _ := ma.NewMACDSignalFloat(macd, signalEMA, pricesForSignal[ma.DefaultLongMACDPeriod+1])

	for _, p := range pricesForSignal[ma.DefaultLongMACDPeriod+2:] {
		signal.Calculate(p)
	}
}

func TestMACDSignalBig_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPricesForSignal[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMABig(ma.DefaultSignalEMAPeriod, signalSMA, nil)
	for _, p := range bigPricesForSignal[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPricesForSignal[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	signal, firstResult := ma.NewMACDSignalBig(macd, signalEMA, bigPricesForSignal[ma.DefaultLongMACDPeriod+1])

	if firstResult.BuySignal != signalResults[0] {
		t.FailNow()
	}

	for i, p := range bigPricesForSignal[ma.DefaultLongMACDPeriod+2:] {
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
	_, shortSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range pricesForSignal[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, signalSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultSignalEMAPeriod])
	signalEMA := ma.NewEMAFloat(ma.DefaultSignalEMAPeriod, signalSMA, 0)
	for _, p := range pricesForSignal[ma.DefaultSignalEMAPeriod:ma.DefaultLongMACDPeriod] {
		signalEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(pricesForSignal[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	signal, firstResult := ma.NewMACDSignalFloat(macd, signalEMA, pricesForSignal[ma.DefaultLongMACDPeriod+1])

	if firstResult.BuySignal != signalResults[0] {
		t.FailNow()
	}

	for i, p := range pricesForSignal[ma.DefaultLongMACDPeriod+2:] {
		actual := signal.Calculate(p).BuySignal
		expected := signalResults[i+1]
		if actual != expected {
			if actual == nil || expected == nil || *actual != *expected {
				t.FailNow()
			}
		}
	}
}
