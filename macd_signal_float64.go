package ma

const (

	// RequiredSamplesForDefaultMACDSignal is the required number of period samples for an MACD signal using the default
	// arguments.
	RequiredSamplesForDefaultMACDSignal = DefaultLongMACDPeriod + DefaultSignalEMAPeriod
)

// MACDSignal represents an MACD and signal EMA pair.
type MACDSignal struct {
	macd      MACD
	signalEMA *EMA
	prevBuy   bool
}

// MACDSignalResults holds the results of an MACD and signal EMA pair's calculation.
//
// If the calculation triggered a buy signal BuySignal will not be `nil`. It will be a pointer to `true`, if the signal
// indicates a buy and a pointer to `false` if the signal indicates a sell.
type MACDSignalResults struct {
	BuySignal *bool
	MACD      MACDResults
	SignalEMA float64
}

// NewMACDSignal creates a new MACD and signal EMA pair. It's calculations are done in tandem to produced buy/sell
// signals.
func NewMACDSignal(macd MACD, signalEMA *EMA, next float64) (*MACDSignal, MACDSignalResults) {
	macdSignal := &MACDSignal{
		macd:      macd,
		signalEMA: signalEMA,
	}

	macdResult := macd.Calculate(next)
	results := MACDSignalResults{
		MACD:      macdResult,
		SignalEMA: signalEMA.Calculate(macdResult.Result),
	}

	macdSignal.prevBuy = results.MACD.Result > results.SignalEMA

	return macdSignal, results
}

// Calculate computes the next MACD and signal EMA pair's results. It may also trigger a buy/sell signal.
func (m *MACDSignal) Calculate(next float64) MACDSignalResults {
	macd := m.macd.Calculate(next)
	signalEMA := m.signalEMA.Calculate(macd.Result)

	results := MACDSignalResults{
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

// DefaultMACDSignal is a helper function to create an MACD and signal EMA pair from the default parameters given the
// initial input data points.
//
// There must be at least 35 data points in the initial slice. This accounts for the number of data points required to
// make the MACD (26) and the number of data points to make the signal EMA from the MACD (9).
//
// If the initial input slice does not have enough data points, the function will return `nil`.
//
// If the initial input slice does has too many data points, the MACD and signal EMA pair will be "caught" up to the
// last data point given, but no results will be accessible.
func DefaultMACDSignal(initial []float64) *MACDSignal {
	if RequiredSamplesForDefaultMACDSignal > len(initial) {
		return nil
	}

	_, shortSMA := NewSMA(initial[:DefaultShortMACDPeriod])
	shortEMA := NewEMA(DefaultShortMACDPeriod, shortSMA, 0)

	var latestShortEMA float64
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		latestShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewSMA(initial[:DefaultLongMACDPeriod])
	longEMA := NewEMA(DefaultLongMACDPeriod, longSMA, 0)

	firstMACDResult := latestShortEMA - longSMA

	macd := NewMACD(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], 0)

	signal, _ := NewMACDSignal(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
