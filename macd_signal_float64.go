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

// Calculate TODO
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

func DefaultMACDSignalFloat(initial []float64) *MACDSignalFloat {
	if DefaultShortMACDPeriod+DefaultLongMACDPeriod+DefaultSignalEMAPeriod > len(initial) {
		return nil
	}

	_, shortSMA := NewSMAFloat(initial[:DefaultShortMACDPeriod])
	shortEMA := NewEMAFloat(DefaultShortMACDPeriod, shortSMA, 0)

	var mostRecentShortEMA float64
	for _, p := range initial[DefaultShortMACDPeriod:DefaultLongMACDPeriod] {
		mostRecentShortEMA = shortEMA.Calculate(p)
	}

	_, longSMA := NewSMAFloat(initial[:DefaultLongMACDPeriod])
	longEMA := NewEMAFloat(DefaultLongMACDPeriod, longSMA, 0)

	firstMACDResult := mostRecentShortEMA - longSMA

	macd := NewMACDFloat(longEMA, shortEMA)

	signalEMA, _, _ := macd.SignalEMA(firstMACDResult, initial[DefaultLongMACDPeriod:DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1], 0)

	signal, _ := NewMACDSignalFloat(macd, signalEMA, initial[DefaultLongMACDPeriod+DefaultSignalEMAPeriod-1])

	for i := DefaultLongMACDPeriod + DefaultSignalEMAPeriod; i < len(initial); i++ {
		signal.Calculate(initial[i])
	}

	return signal
}
