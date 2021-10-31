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
func NewMACDFloat(long, short *EMAFloat) MACDFloat {
	return MACDFloat{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
func (macd MACDFloat) Calculate(next float64) MACDResultsFloat {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return MACDResultsFloat{
		Long:   long,
		Result: short - long,
		Short:  short,
	}
}

// SignalEMA TODO
func (macd MACDFloat) SignalEMA(firstMACDResult float64, next []float64, smoothing float64) (signalEMA *EMAFloat, signalResult float64, macdResults []MACDResultsFloat) {
	macdFloats := make([]float64, len(next))

	for i, p := range next {
		result := macd.Calculate(p)
		macdFloats[i] = result.Result
		macdResults = append(macdResults, result)
	}

	_, sma := NewSMAFloat(append([]float64{firstMACDResult}, macdFloats...))

	// TODO Remove other redundant default checks.
	// if smoothing == 0 {
	// 	smoothing = DefaultEMASmoothing
	// }

	return NewEMAFloat(uint(len(next)+1), sma, smoothing), sma, macdResults
}
