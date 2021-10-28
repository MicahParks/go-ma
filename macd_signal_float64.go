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
	Signal    float64
}

func NewMACDSignalFloat(macd MACDFloat, signalEMA *EMAFloat, next float64) (macdSignalFloat *MACDSignalFloat, results MACDSignalResultsFloat) {
	macdSignalFloat = &MACDSignalFloat{
		macd:      macd,
		signalEMA: signalEMA,
	}

	results = MACDSignalResultsFloat{
		MACD:   macd.Calculate(next),
		Signal: signalEMA.Calculate(next),
	}

	macdSignalFloat.prevBuy = results.MACD.Result > results.Signal

	return macdSignalFloat, results
}

// Calculate TODO
func (m *MACDSignalFloat) Calculate(next float64) (results MACDSignalResultsFloat) {
	results = MACDSignalResultsFloat{
		MACD:   m.macd.Calculate(next),
		Signal: m.signalEMA.Calculate(next),
	}

	if results.MACD.Result > results.Signal != m.prevBuy {
		buy := !m.prevBuy
		m.prevBuy = buy
		results.BuySignal = &buy
	}

	return results
}
