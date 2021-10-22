package ma

// SMAFloat represents the state of a Simple Moving Average (SMA) algorithm.
type SMAFloat struct {
	cache    []float64
	cacheLen uint
	index    uint
	periods  float64
}

// NewSMAFloat creates a new SMA data structure and the initial result. The period used will be the length of the
// initial input slice.
func NewSMAFloat(initial []float64) (sma *SMAFloat, result float64) {

	sma = &SMAFloat{
		cacheLen: uint(len(initial)),
	}
	sma.cache = make([]float64, sma.cacheLen)
	sma.periods = float64(sma.cacheLen)

	for i, p := range initial {
		if i != 0 {
			sma.cache[i] = p
		}
		result += p
	}
	result /= sma.periods

	return sma, result
}

// Calculate produces the next SMA result given the next input.
func (sma *SMAFloat) Calculate(next float64) (result float64) {
	sma.cache[sma.index] = next
	sma.index++
	if sma.index == sma.cacheLen {
		sma.index = 0
	}

	for _, p := range sma.cache {
		result += p
	}
	result /= sma.periods

	return result
}
