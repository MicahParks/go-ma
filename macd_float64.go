package ma

// TODO
type MACD struct {
	Long  *EMAFloat // Typical 26 periods.
	Short *EMAFloat // Typical 12 periods.
}

// TODO
func NewMACDFloat(long, short *EMAFloat) (macd MACD) {
	return MACD{
		Long:  long,
		Short: short,
	}
}

// TODO
func (macd MACD) Calculate(next float64) (result float64) {
	return macd.Short.Calculate(next) - macd.Long.Calculate(next)
}

// TODO Make the "signal line" from the result of the MACD. Typically nine periods.
