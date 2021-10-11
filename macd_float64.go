package ma

// TODO
type MACD struct {
	Long  *EMAFloat
	Short *EMAFloat
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
