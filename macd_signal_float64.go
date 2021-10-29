package ma

// TODO
type MACDSignalFloat struct {
	macd      MACDFloat
	signalEMA *EMAFloat
	prevBuy   bool
}

type MACDSignalResultsFloat struct {
	BuySignal *bool
	MACD      MACDResultsFloat
	SignalEMA float64
}

// TODO Make the input a data structure.
func NewMACDSignalFloat(macd MACDFloat, signalEMAPeriod uint, signalEMASmoothing float64, next []float64) (macdSignalFloat *MACDSignalFloat) {
	_, sma := NewSMAFloat(next)
	signalEMA := NewEMAFloat(signalEMAPeriod, sma, signalEMASmoothing)

	macdSignalFloat = &MACDSignalFloat{
		macd:      macd,
		signalEMA: signalEMA,
	}

	results = MACDSignalResultsFloat{
		MACD:      macd.Calculate(next),
		SignalEMA: signalEMA.Calculate(next),
	}

	macdSignalFloat.prevBuy = results.MACD.Result > results.SignalEMA

	return macdSignalFloat, results
}

// Calculate TODO
func (m *MACDSignalFloat) Calculate(next float64) (results MACDSignalResultsFloat) {
	results = MACDSignalResultsFloat{
		MACD:      m.macd.Calculate(next),
		SignalEMA: m.signalEMA.Calculate(next),
	}

	if results.MACD.Result > results.SignalEMA != m.prevBuy {
		buy := !m.prevBuy
		m.prevBuy = buy
		results.BuySignal = &buy
	}

	return results
}
