package ma

import (
	"math/big"
)

// TODO
type MACDSignalBig struct {
	MACD      MACDBig
	SignalEMA *EMABig
	prevBuy   bool
}

type MACDSignalResultsBig struct {
	BuySignal *bool
	MACD      MACDResultsBig
	Signal    *big.Float
}

func NewMACDSignalBig(macd MACDBig, signalEMA *EMABig, next *big.Float) (macdFloatSignal *MACDSignalBig, results MACDSignalResultsBig) {
	macdFloatSignal = &MACDSignalBig{
		MACD:      macd,
		SignalEMA: signalEMA,
	}

	results = MACDSignalResultsBig{
		MACD:   macd.Calculate(next),
		Signal: signalEMA.Calculate(next),
	}

	macdFloatSignal.prevBuy = results.MACD.Result.Cmp(results.Signal) == 1

	return macdFloatSignal, results
}

// Calculate TODO
func (m *MACDSignalBig) Calculate(next *big.Float) (results MACDSignalResultsBig) {
	results = MACDSignalResultsBig{
		MACD:   m.MACD.Calculate(next),
		Signal: m.SignalEMA.Calculate(next),
	}

	targetCmp := 1
	if m.prevBuy {
		targetCmp = -1
	}
	if results.MACD.Result.Cmp(results.Signal) == targetCmp {
		buy := !m.prevBuy
		m.prevBuy = buy
		results.BuySignal = &buy
	}

	return results
}
