package ma

import (
	"math/big"
)

// BigMACD represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type BigMACD struct {
	Long  *BigEMA
	Short *BigEMA
}

// BigMACDResults holds the results fo an MACD calculation.
type BigMACDResults struct {
	Long   *big.Float
	Result *big.Float
	Short  *big.Float
}

// NewBigMACD creates a new MACD data structure and returns the initial result.
func NewBigMACD(long, short *BigEMA) BigMACD {
	return BigMACD{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
func (macd BigMACD) Calculate(next *big.Float) BigMACDResults {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return BigMACDResults{
		Long:   long,
		Result: new(big.Float).Sub(short, long),
		Short:  short,
	}
}

// SignalEMA creates a signal EMA for the current MACD.
//
// The first MACD result *must* be saved in order to create the signal EMA. Then, the next period samples required for
// the creation of the signal EMA must be given. The period length of the EMA is `1 + len(next)`.
func (macd BigMACD) SignalEMA(firstMACDResult *big.Float, next []*big.Float, smoothing *big.Float) (signalEMA *BigEMA, signalResult *big.Float, macdResults []BigMACDResults) {
	macdBigs := make([]*big.Float, len(next))
	macdResults = make([]BigMACDResults, len(next))

	for i, p := range next {
		result := macd.Calculate(p)
		macdBigs[i] = result.Result
		macdResults[i] = result
	}

	_, sma := NewBigSMA(append([]*big.Float{firstMACDResult}, macdBigs...))

	return NewBigEMA(uint(len(next)+1), sma, smoothing), sma, macdResults
}
