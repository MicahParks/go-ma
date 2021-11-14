package ma

import (
	"math/big"
)

var (
	bigOne = big.NewFloat(1)
)

// BigEMA represents the state of an Exponential Moving Average (EMA) algorithm.
type BigEMA struct {
	constant         *big.Float
	prev             *big.Float
	oneMinusConstant *big.Float
}

// NewBigEMA creates a new EMA data structure.
func NewBigEMA(periods uint, sma, smoothing *big.Float) *BigEMA {
	if smoothing == nil || smoothing.Cmp(big.NewFloat(0)) == 0 {
		smoothing = big.NewFloat(DefaultEMASmoothing)
	}

	ema := &BigEMA{
		constant: new(big.Float).Quo(smoothing, new(big.Float).Add(bigOne, big.NewFloat(float64(periods)))),
		prev:     sma,
	}
	ema.oneMinusConstant = new(big.Float).Sub(bigOne, ema.constant)

	return ema
}

// Calculate produces the next EMA result given the next input.
func (ema *BigEMA) Calculate(next *big.Float) (result *big.Float) {
	ema.prev = new(big.Float).Add(new(big.Float).Mul(next, ema.constant), new(big.Float).Mul(ema.prev, ema.oneMinusConstant))

	return new(big.Float).Copy(ema.prev)
}
