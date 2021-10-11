package ma

const (

	// TODO
	DefaultSmoothing = 2
)

// TODO
type EMAFloat struct {
	prev             float64
	constant         float64
	oneMinusConstant float64
}

// TODO Verify
func NewEMAFloat(periods uint, sma, smoothing float64) (ema *EMAFloat) {

	ema = &EMAFloat{
		constant: smoothing / (1 + float64(periods)),
		prev:     sma,
	}
	ema.oneMinusConstant = 1 - ema.constant

	return ema
}

// TODO
func (e *EMAFloat) Calculate(next float64) (result float64) {

	e.prev = next*e.constant + e.prev*e.oneMinusConstant

	return e.prev
}
