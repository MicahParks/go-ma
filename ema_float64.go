package ma

const (

	// TODO
	DefaultSmoothing = 2
)

// TODO
type EMA struct {
	prev     float64
	constant float64
}

// TODO
func NewEMAFloat(periods uint, smoothing float64) (ema *EMA) {

	ema = &EMA{
		constant: smoothing / (1 + float64(periods)),
	}

	return ema
}

// TODO
func (e *EMA) Calculate(next float64) (result float64) {

	e.prev = next*e.constant + e.prev*(1-e.constant)

	return e.prev
}
