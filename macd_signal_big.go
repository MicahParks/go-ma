package ma

import (
	"math/big"
)

// BigMACDSignal represents an MACD and signal EMA pair.
type BigMACDSignal struct {
	macd      BigMACD
	signalEMA *BigEMA
	prevBuy   bool
}

// BigMACDSignalResults holds the results of an MACD and signal EMA pair's calculation.
//
// If the calculation triggered a buy signal BuySignal will not be `nil`. It will be a pointer to `true`, if the signal
// indicates a buy and a pointer to `false` if the signal indicates a sell.
type BigMACDSignalResults struct {
	BuySignal *bool
	MACD      BigMACDResults
	SignalEMA *big.Float
}

// NewBigMACDSignal creates a new MACD and signal EMA pair. It's calculations are done in tandem to produced buy/sell
// signals.
func NewBigMACDSignal(macd BigMACD, signalEMA *BigEMA, next *big.Float) (*BigMACDSignal, BigMACDSignalResults) {
	macdSignal := &BigMACDSignal{
		macd:      macd,
		signalEMA: signalEMA,
	}

	macdResult := macd.Calculate(next)
	results := BigMACDSignalResults{
		MACD:      macdResult,
		SignalEMA: signalEMA.Calculate(macdResult.Result),
	}

	macdSignal.prevBuy = results.MACD.Result.Cmp(results.SignalEMA) == 1

	return macdSignal, results
}

// Calculate computes the next MACD and signal EMA pair's results. It may also trigger a buy/sell signal.
func (m *BigMACDSignal) Calculate(next *big.Float) BigMACDSignalResults {
	macd := m.macd.Calculate(next)
	signalEMA := m.signalEMA.Calculate(macd.Result)

	results := BigMACDSignalResults{
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

// DefaultBigMACDSignal is a helper function to create an MACD and signal EMA pair from the default parameters given the
// initial input data points.
//
// There must be at least 35 data points in the initial slice. This accounts for the number of data points required to
// make the MACD (26) and the number of data points to make the signal EMA from the MACD (9).
//
// If the initial input slice does not have enough data points, the function will return `nil`.
//
// If the initial input slice does has too many data points, the MACD and signal EMA pair will be "caught" up to the
// last data point given, but no results will be accessible.
func DefaultBigMACDSignal(initial []*big.Float) *BigMACDSignal {
	if RequiredSamplesForDefaultMACDSignal > len(initial) {
		return nil
	}

	_, shortSMA := NewBigSMA(initial[:DefaultShortMACDPeriod])
	shortEMA := NewBigEMA(DefaultShortMACDPeriod, shortSMA, nil)

	var latestShortEMA *big.Float
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewBigSMA(initial[:DefaultLongMACDPeriod])
	longEMA := NewBigEMA(DefaultLongMACDPeriod, longSMA, nil)

	firstMACDResult := new(big.Float).Sub(latestShortEMA, longSMA)

	macd := NewBigMACD(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], nil)

	signal, _ := NewBigMACDSignal(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
