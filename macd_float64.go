package ma

const (

	// DefaultLongMACDPeriod is a common MACD period for the long input EMA algorithm.
	DefaultLongMACDPeriod = 26

	// DefaultShortMACDPeriod is a common MACD period for the short input EMA algorithm.
	DefaultShortMACDPeriod = 12

	// DefaultSignalEMAPeriod is a common MACD period for the signal EMA that will crossover the MACD line.
	DefaultSignalEMAPeriod = 9
)

// MACD represents the state of a Moving Average Convergence Divergence (MACD) algorithm.
type MACD struct {
	Long  *EMA
	Short *EMA
}

// MACDResults holds the results fo an MACD calculation.
type MACDResults struct {
	Long   float64
	Result float64
	Short  float64
}

// NewMACD creates a new MACD data structure and returns the initial result.
func NewMACD(long, short *EMA) MACD {
	return MACD{
		Long:  long,
		Short: short,
	}
}

// Calculate produces the next MACD result given the next input.
func (macd MACD) Calculate(next float64) MACDResults {
	short := macd.Short.Calculate(next)
	long := macd.Long.Calculate(next)
	return MACDResults{
		Long:   long,
		Result: short - long,
		Short:  short,
	}
}

// SignalEMA creates a signal EMA for the current MACD.
//
// The first MACD result *must* be saved in order to create the signal EMA. Then, the next period samples required for
// the creation of the signal EMA must be given. The period length of the EMA is `1 + len(next)`.
func (macd MACD) SignalEMA(firstMACDResult float64, next []float64, smoothing float64) (signalEMA *EMA, signalResult float64, macdResults []MACDResults) {
	macdFloats := make([]float64, len(next))
	macdResults = make([]MACDResults, len(next))

	for i, p := range next {
		result := macd.Calculate(p)
		macdFloats[i] = result.Result
		macdResults[i] = result
	}

	_, sma := NewSMA(append([]float64{firstMACDResult}, macdFloats...))

	return NewEMA(uint(len(next)+1), sma, smoothing), sma, macdResults
}
