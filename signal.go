package ma

// TODO
type MACDFloatSignal struct {
	MACD       MACDFloat
	SignalEMA  *EMAFloat
	crossovers chan bool
	buy        bool
}

func NewMACDFloatSignal(macd MACDFloat, signalEMA *EMAFloat, next float64) (macdFloatSignal *MACDFloatSignal) {
	macdFloatSignal = &MACDFloatSignal{
		MACD:       macd,
		SignalEMA:  signalEMA,
		crossovers: make(chan bool),
	}

	macdFloatSignal.buy = macd.Calculate(next) > signalEMA.Calculate(next)

	return macdFloatSignal
}

// TODO
func (m *MACDFloatSignal) Consume(next float64) {
	buy := m.MACD.Calculate(next) > m.SignalEMA.Calculate(next)

	if buy != m.buy {
		m.crossovers <- buy
		m.buy = buy
	}
}

// TODO
func (m *MACDFloatSignal) Crossovers() <-chan bool {
	return m.crossovers
}
