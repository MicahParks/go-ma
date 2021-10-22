package ma

const (

	// DefaultEMASmoothing is a common smoothing constant for the EMA algorithm.
	DefaultEMASmoothing = 2
)

// EMAFloat represents the state of an Exponential Moving Average.
type EMAFloat struct {
	constant         float64
	prev             float64
	oneMinusConstant float64
}

// NewEMAFloat creates a new EMA data structure (EMA) algorithm.
func NewEMAFloat(periods uint, sma, smoothing float64) (ema *EMAFloat) {
	if smoothing == 0 {
		smoothing = DefaultEMASmoothing
	}

	ema = &EMAFloat{
		constant: smoothing / (1 + float64(periods)),
		prev:     sma,
	}
	ema.oneMinusConstant = 1 - ema.constant

	return ema
}

// Calculate producers the next EMA result given the next input.
func (ema *EMAFloat) Calculate(next float64) (result float64) {
	ema.prev = next*ema.constant + ema.prev*ema.oneMinusConstant

	return ema.prev
}
