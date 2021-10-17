package ma

const (

	// TODO
	DefaultSmoothing = 2
)

// TODO
type EMAFloat struct {
	constant         float64
	prev             float64
	oneMinusConstant float64
}

// TODO Verify
func NewEMAFloat(periods uint, sma, smoothing float64) (ema *EMAFloat) {
	if smoothing == 0 {
		smoothing = DefaultSmoothing
	}

	ema = &EMAFloat{
		constant: smoothing / (1 + float64(periods)),
		prev:     sma,
	}
	ema.oneMinusConstant = 1 - ema.constant

	return ema
}

// TODO
func (ema *EMAFloat) Calculate(next float64) (result float64) {
	ema.prev = next*ema.constant + ema.prev*ema.oneMinusConstant

	return ema.prev
}
