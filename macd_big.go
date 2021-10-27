package ma

import (
	"math/big"
)

// MACDBig represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type MACDBig struct {
	Long  *EMABig // Typical 26 periods.
	Short *EMABig // Typical 12 periods.
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
func (macd MACDBig) Calculate(next *big.Float) (result, short, long *big.Float) {
	short = macd.Short.Calculate(next)
	long = macd.Long.Calculate(next)
	return new(big.Float).Sub(short, long), short, long
}
