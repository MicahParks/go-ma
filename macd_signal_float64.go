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
func NewMACDSignalFloat(macd MACDFloat, signalEMA *EMAFloat, next float64) (macdSignalFloat *MACDSignalFloat, results MACDSignalResultsFloat) {
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
	macd := m.macd.Calculate(next)
	signalEMA := m.signalEMA.Calculate(macd.Result)

	results = MACDSignalResultsFloat{
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
