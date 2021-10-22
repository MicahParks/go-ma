package ma

import (
	"math/big"
)

type SMABig struct {
	cache    []*big.Float
	cacheLen uint
	index    uint
	periods  *big.Float
}

func NewSMABig(initial []*big.Float) (sma *SMABig, result *big.Float) {
	result = new(big.Float)

	sma = &SMABig{
		cacheLen: uint(len(initial)),
	}
	sma.cache = make([]*big.Float, sma.cacheLen)
	sma.periods = big.NewFloat(float64(sma.cacheLen))

	for i, p := range initial {
		if i != 0 {
			sma.cache[i] = p
		}
		result = result.Add(result, p)
	}
	result = result.Quo(result, sma.periods)

	return sma, result
}

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
