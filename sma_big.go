package ma

import (
	"math/big"
)

// SMABig represents the state of a Simple Moving Average (SMA) algorithm.
type SMABig struct {
	cache    []*big.Float
	cacheLen uint
	index    uint
	periods  *big.Float
}

// NewSMABig creates a new SMA data structure and the initial result. The period used will be the length of the initial
// input slice.
func NewSMABig(initial []*big.Float) (sma *SMABig, result *big.Float) {
	sma = &SMABig{
		cacheLen: uint(len(initial)),
	}
	sma.cache = make([]*big.Float, sma.cacheLen)
	sma.periods = big.NewFloat(float64(sma.cacheLen))

	result = big.NewFloat(0)
	for i, p := range initial {
		if i != 0 {
			sma.cache[i] = p
		}
		result = result.Add(result, p)
	}
	result = result.Quo(result, sma.periods)

	return sma, result
}

// Calculate produces the next SMA result given the next input.
func (sma *SMABig) Calculate(next *big.Float) (result *big.Float) {
	sma.cache[sma.index] = next
	sma.index++
	if sma.index == sma.cacheLen {
		sma.index = 0
	}

	result = big.NewFloat(0)
	for _, p := range sma.cache {
		result = result.Add(result, p)
	}
	result = result.Quo(result, sma.periods)

	return result
}
