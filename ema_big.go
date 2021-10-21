package ma

import (
	"math/big"
)

var (
	bigOne = big.NewFloat(1)
)

type EMABig struct {
	constant         *big.Float
	prev             *big.Float
	oneMinusConstant *big.Float
}

func NewEMABig(periods uint, sma, smoothing *big.Float) (ema *EMABig) {
	if smoothing == nil || smoothing.Cmp(big.NewFloat(0)) == 0 {
		smoothing = big.NewFloat(DefaultEMASmoothing)
	}

	ema = &EMABig{
		constant: new(big.Float).Quo(smoothing, new(big.Float).Add(bigOne, big.NewFloat(float64(periods)))),
		prev:     sma,
	}
	ema.oneMinusConstant = new(big.Float).Sub(bigOne, ema.constant)

	return ema
}

func (ema *EMABig) Calculate(next *big.Float) (result *big.Float) {
	ema.prev = new(big.Float).Add(new(big.Float).Mul(next, ema.constant), new(big.Float).Mul(ema.prev, ema.oneMinusConstant))

	return new(big.Float).Copy(ema.prev)
}
