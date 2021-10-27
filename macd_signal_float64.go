package ma

// TODO
type MACDSignalFloat struct {
	MACD      MACDFloat
	SignalEMA *EMAFloat
	buy       bool
}

func NewMACDSignalFloat(macd MACDFloat, signalEMA *EMAFloat, next float64) (macdSignalFloat *MACDSignalFloat, macdResult, signalResult float64) {
	macdSignalFloat = &MACDSignalFloat{
		MACD:      macd,
		SignalEMA: signalEMA,
	}

	macdResult, _, _ = macd.Calculate(next)
	signalResult = signalEMA.Calculate(next)

	macdSignalFloat.buy = macdResult > signalResult

	return macdSignalFloat, macdResult, signalResult
}

// Calculate TODO
func (m *MACDSignalFloat) Calculate(next float64) (buySignal *bool, macdResult, signalResult float64) {
	macdResult, _, _ = m.MACD.Calculate(next)
	signalResult = m.SignalEMA.Calculate(next)

	buy := macdResult > signalResult

	if buy != m.buy {
		m.buy = buy
		buySignal = &buy
	}

	return buySignal, macdResult, signalResult
}
