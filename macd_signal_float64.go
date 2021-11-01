package ma

const (

	// RequiredSamplesForDefaultMACDSignal is the required number of period samples for an MACD signal using the default
	// arguments.
	RequiredSamplesForDefaultMACDSignal = DefaultLongMACDPeriod + DefaultSignalEMAPeriod
)

// MACDSignalFloat represents an MACD and signal EMA pair.
type MACDSignalFloat struct {
	macd      MACDFloat
	signalEMA *EMAFloat
	prevBuy   bool
}

// MACDSignalResultsFloat holds the results of an MACD and signal EMA pair's calculation.
//
// If the calculation triggered a buy signal BuySignal will not be `nil`. It will be a pointer to `true`, if the signal
// indicates a buy and a pointer to `false` if the signal indicates a sell.
type MACDSignalResultsFloat struct {
	BuySignal *bool
	MACD      MACDResultsFloat
	SignalEMA float64
}

// NewMACDSignalFloat creates a new MACD and signal EMA pair. It's calculations are done in tandem to produced buy/sell
// signals.
func NewMACDSignalFloat(macd MACDFloat, signalEMA *EMAFloat, next float64) (*MACDSignalFloat, MACDSignalResultsFloat) {
	macdSignalFloat := &MACDSignalFloat{
		macd:      macd,
		signalEMA: signalEMA,
	}

	macdResult := macd.Calculate(next)
	results := MACDSignalResultsFloat{
		MACD:      macdResult,
		SignalEMA: signalEMA.Calculate(macdResult.Result),
	}

	macdSignalFloat.prevBuy = results.MACD.Result > results.SignalEMA

	return macdSignalFloat, results
}

// Calculate computes the next MACD and signal EMA pair's results. It may also trigger a buy/sell signal.
func (m *MACDSignalFloat) Calculate(next float64) MACDSignalResultsFloat {
	macd := m.macd.Calculate(next)
	signalEMA := m.signalEMA.Calculate(macd.Result)

	results := MACDSignalResultsFloat{
		MACD:      macd,
		SignalEMA: signalEMA,
	}

	if results.MACD.Result > results.SignalEMA != m.prevBuy {
		buy := !m.prevBuy
		m.prevBuy = buy
		results.BuySignal = &buy
	}

	return results
}

// DefaultMACDSignalFloat is a helper function to create an MACD and signal EMA pair from the default parameters given the
// initial input data points.
//
// There must be at least 35 data points in the initial slice. This accounts for the number of data points required to
// make the MACD (26) and the number of data points to make the signal EMA from the MACD (9).
//
// If the initial input slice does not have enough data points, the function will return `nil`.
//
// If the initial input slice does has too many data points, the MACD and signal EMA pair will be "caught" up to the
// last data point given, but no results will be accessible.
func DefaultMACDSignalFloat(initial []float64) *MACDSignalFloat {
	if RequiredSamplesForDefaultMACDSignal > len(initial) {
		return nil
	}

	_, shortSMA := NewSMAFloat(initial[:DefaultShortMACDPeriod])
	shortEMA := NewEMAFloat(DefaultShortMACDPeriod, shortSMA, 0)

	var latestShortEMA float64
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewSMAFloat(initial[:DefaultLongMACDPeriod])
	longEMA := NewEMAFloat(DefaultLongMACDPeriod, longSMA, 0)

	firstMACDResult := latestShortEMA - longSMA

	macd := NewMACDFloat(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], 0)

	signal, _ := NewMACDSignalFloat(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
