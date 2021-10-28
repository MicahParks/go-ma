package ma

import (
	"math/big"
)

// MACDBig represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type MACDBig struct {
	Long  *EMABig // Typical 26 periods.
	Short *EMABig // Typical 12 periods.
}

type MACDResultsBig struct {
	Long   *big.Float
	Result *big.Float
	Short  *big.Float
}

// NewMACDBig creates a new MACD data structure and returns the initial result.
//
// TODO Return the initial result.
func NewMACDBig(long, short *EMABig) (macd MACDBig) {
	return MACDBig{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
func (macd MACDBig) Calculate(next *big.Float) (results MACDResultsBig) {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return MACDResultsBig{
		Long:   long,
		Result: new(big.Float).Sub(short, long),
		Short:  short,
	}
}
