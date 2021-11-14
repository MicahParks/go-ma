package ma

const (

	// DefaultEMASmoothing is a common smoothing constant for the EMA algorithm.
	DefaultEMASmoothing = 2
)

// EMA represents the state of an Exponential Moving Average (EMA).
type EMA struct {
	constant         float64
	prev             float64
	oneMinusConstant float64
}

// NewEMA creates a new EMA data structure.
func NewEMA(periods uint, sma, smoothing float64) *EMA {
	if smoothing == 0 {
		smoothing = DefaultEMASmoothing
	}

	ema := &EMA{
		constant: smoothing / (1 + float64(periods)),
		prev:     sma,
	}
	ema.oneMinusConstant = 1 - ema.constant

	return ema
}

// Calculate produces the next EMA result given the next input.
func (ema *EMA) Calculate(next float64) (result float64) {
	ema.prev = next*ema.constant + ema.prev*ema.oneMinusConstant

	return ema.prev
}
