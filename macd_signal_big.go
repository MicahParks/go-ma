package ma

import (
	"math/big"
)

// TODO
type MACDSignalBig struct {
	macd      MACDBig
	signalEMA *EMABig
	prevBuy   bool
}

type MACDSignalResultsBig struct {
	BuySignal *bool
	MACD      MACDResultsBig
	SignalEMA *big.Float
}

func NewMACDSignalBig(macd MACDBig, signalEMA *EMABig, next *big.Float) (*MACDSignalBig, MACDSignalResultsBig) {
	macdFloatSignal := &MACDSignalBig{
		macd:      macd,
		signalEMA: signalEMA,
	}

	macdResult := macd.Calculate(next)
	results := MACDSignalResultsBig{
		MACD:      macdResult,
		SignalEMA: signalEMA.Calculate(macdResult.Result),
	}

	macdFloatSignal.prevBuy = results.MACD.Result.Cmp(results.SignalEMA) == 1

	return macdFloatSignal, results
}

// Calculate TODO
func (m *MACDSignalBig) Calculate(next *big.Float) MACDSignalResultsBig {
	macd := m.macd.Calculate(next)
	signalEMA := m.signalEMA.Calculate(macd.Result)

	results := MACDSignalResultsBig{
		MACD:      macd,
		SignalEMA: signalEMA,
	}

	targetCmp := 1
	if m.prevBuy {
		targetCmp = -1
	}
	if results.MACD.Result.Cmp(results.SignalEMA) == targetCmp {
		buy := !m.prevBuy
		m.prevBuy = buy
		results.BuySignal = &buy
	}

	return results
}

func DefaultMACDSignalBig(initial []*big.Float) *MACDSignalBig {
	if RequiredSamplesForDefaultMACDSignal > len(initial) {
		return nil
	}

	_, shortSMA := NewSMABig(initial[:DefaultShortMACDPeriod])
	shortEMA := NewEMABig(DefaultShortMACDPeriod, shortSMA, nil)

	var mostRecentShortEMA *big.Float
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		mostRecentShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewSMABig(initial[:DefaultLongMACDPeriod])
	longEMA := NewEMABig(DefaultLongMACDPeriod, longSMA, nil)

	firstMACDResult := new(big.Float).Sub(mostRecentShortEMA, longSMA)

	macd := NewMACDBig(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], nil)

	signal, _ := NewMACDSignalBig(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
