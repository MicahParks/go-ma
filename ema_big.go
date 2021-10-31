package ma

import (
	"math/big"
)

var (
	bigOne = big.NewFloat(1)
)

// EMABig represents the state of an Exponential Moving Average (EMA) algorithm.
type EMABig struct {
	constant         *big.Float
	prev             *big.Float
	oneMinusConstant *big.Float
}

// NewEMABig creates a new EMA data structure.
func NewEMABig(periods uint, sma, smoothing *big.Float) *EMABig {
	if smoothing == nil || smoothing.Cmp(big.NewFloat(0)) == 0 {
		smoothing = big.NewFloat(DefaultEMASmoothing)
	}

	ema := &EMABig{
		constant: new(big.Float).Quo(smoothing, new(big.Float).Add(bigOne, big.NewFloat(float64(periods)))),
		prev:     sma,
	}
	ema.oneMinusConstant = new(big.Float).Sub(bigOne, ema.constant)

	return ema
}

// Calculate producers the next EMA result given the next input.
func (ema *EMABig) Calculate(next *big.Float) (result *big.Float) {
	ema.prev = new(big.Float).Add(new(big.Float).Mul(next, ema.constant), new(big.Float).Mul(ema.prev, ema.oneMinusConstant))

	return new(big.Float).Copy(ema.prev)
}
