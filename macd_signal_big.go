package ma

import (
	"math/big"
)

// MACDSignalBig represents an MACD and signal EMA pair.
type MACDSignalBig struct {
	macd      MACDBig
	signalEMA *EMABig
	prevBuy   bool
}

// MACDSignalResultsBig holds the results of an MACD and signal EMA pair's calculation.
//
// If the calculation triggered a buy signal BuySignal will not be `nil`. It will be a pointer to `true`, if the signal
// indicates a buy and a pointer to `false` if the signal indicates a sell.
type MACDSignalResultsBig struct {
	BuySignal *bool
	MACD      MACDResultsBig
	SignalEMA *big.Float
}

// NewMACDSignalBig creates a new MACD and signal EMA pair. It's calculations are done in tandem to produced buy/sell
// signals.
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

// Calculate computes the next MACD and signal EMA pair's results. It may also trigger a buy/sell signal.
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

// DefaultMACDSignalBig is a helper function to create an MACD and signal EMA pair from the default parameters given the
// initial input data points.
//
// There must be at least 35 data points in the initial slice. This accounts for the number of data points required to
// make the MACD (26) and the number of data points to make the signal EMA from the MACD (9).
//
// If the initial input slice does not have enough data points, the function will return `nil`.
//
// If the initial input slice does has too many data points, the MACD and signal EMA pair will be "caught" up to the
// last data point given, but no results will be accessible.
func DefaultMACDSignalBig(initial []*big.Float) *MACDSignalBig {
	if RequiredSamplesForDefaultMACDSignal > len(initial) {
		return nil
	}

	_, shortSMA := NewSMABig(initial[:DefaultShortMACDPeriod])
	shortEMA := NewEMABig(DefaultShortMACDPeriod, shortSMA, nil)

	var latestShortEMA *big.Float
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewSMABig(initial[:DefaultLongMACDPeriod])
	longEMA := NewEMABig(DefaultLongMACDPeriod, longSMA, nil)

	firstMACDResult := new(big.Float).Sub(latestShortEMA, longSMA)

	macd := NewMACDBig(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], nil)

	signal, _ := NewMACDSignalBig(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
