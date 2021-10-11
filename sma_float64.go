package ma

// TODO
type SMAFloat struct {
	cache    []float64
	cacheLen uint
	index    uint
	periods  float64
}

// TODO
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

// TODO
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
