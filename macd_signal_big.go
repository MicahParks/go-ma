package ma

import (
	"math/big"
)

// TODO
type MACDSignalBig struct {
	MACD      MACDBig
	SignalEMA *EMABig
	buy       bool
}

func NewMACDFloatSignal(macd MACDBig, signalEMA *EMABig, next *big.Float) (macdFloatSignal *MACDSignalBig, macdResult, signalResult *big.Float) {
	macdFloatSignal = &MACDSignalBig{
		MACD:      macd,
		SignalEMA: signalEMA,
	}

	macdResult, _, _ = macd.Calculate(next)
	signalResult = signalEMA.Calculate(next)

	macdFloatSignal.buy = macdResult.Cmp(signalResult) == 1

	return macdFloatSignal, macdResult, signalResult
}

// Calculate TODO
func (m *MACDSignalBig) Calculate(next *big.Float) (buySignal *bool, macdResult, signalResult *big.Float) {
	macdResult, _, _ = m.MACD.Calculate(next)
	signalResult = m.SignalEMA.Calculate(next)

	buy := macdResult.Cmp(signalResult) == 1

	if buy != m.buy {
		m.buy = buy
		buySignal = &buy
	}

	return buySignal, macdResult, signalResult
}
