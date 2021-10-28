package ma

const (

	// DefaultLongMACDPeriod is a common MACD period for the long input EMA algorithm.
	DefaultLongMACDPeriod = 26

	// DefaultShortMACDPeriod is a common MACD period for the short input EMA algorithm.
	DefaultShortMACDPeriod = 12

	// DefaultSignalEMAPeriod is a common MACD period for the signal EMA that will crossover the MACD line.
	DefaultSignalEMAPeriod = 9
)

// MACDFloat represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type MACDFloat struct {
	Long  *EMAFloat // Typical 26 periods.
	Short *EMAFloat // Typical 12 periods.
}

type MACDResultsFloat struct {
	Long   float64
	Result float64
	Short  float64
}

// NewMACDFloat creates a new MACD data structure and returns the initial result.
//
// TODO Return the initial result.
func NewMACDFloat(long, short *EMAFloat) (macd MACDFloat) {
	return MACDFloat{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
//
// TODO Return the result from short and long?
func (macd MACDFloat) Calculate(next float64) (results MACDResultsFloat) {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return MACDResultsFloat{
		Long:   long,
		Result: short - long,
		Short:  short,
	}
}

// TODO Make the "signal line" from the result of the MACD. Typically nine periods.
