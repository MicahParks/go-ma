package ma

import (
	"math/big"
)

// MACDBig represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type MACDBig struct {
	Long  *EMABig
	Short *EMABig
}

// MACDResultsBig holds the results fo an MACD calculation.
type MACDResultsBig struct {
	Long   *big.Float
	Result *big.Float
	Short  *big.Float
}

// NewMACDBig creates a new MACD data structure and returns the initial result.
func NewMACDBig(long, short *EMABig) MACDBig {
	return MACDBig{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
func (macd MACDBig) Calculate(next *big.Float) MACDResultsBig {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return MACDResultsBig{
		Long:   long,
		Result: new(big.Float).Sub(short, long),
		Short:  short,
	}
}

// SignalEMA creates a signal EMA for the current MACD.
//
// The first MACD result *must* be saved in order to create the signal EMA. Then, the next period samples required for
// the creation of the signal EMA must be given. The period length of the EMA is `1 + len(next)`.
func (macd MACDBig) SignalEMA(firstMACDResult *big.Float, next []*big.Float, smoothing *big.Float) (signalEMA *EMABig, signalResult *big.Float, macdResults []MACDResultsBig) {
	macdBigs := make([]*big.Float, len(next))
	macdResults = make([]MACDResultsBig, len(next))

	for i, p := range next {
		result := macd.Calculate(p)
		macdBigs[i] = result.Result
		macdResults[i] = result
	}

	_, sma := NewSMABig(append([]*big.Float{firstMACDResult}, macdBigs...))

	return NewEMABig(uint(len(next)+1), sma, smoothing), sma, macdResults
}
