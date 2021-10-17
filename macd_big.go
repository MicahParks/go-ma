package ma

import (
	"math/big"
)

type MACDBig struct {
	Long  *EMABig
	Short *EMABig
}

func NewMACDBig(long, short *EMABig) (macd MACDBig) {
	return MACDBig{
		Long:  long,
		Short: short,
	}
}

func (macd MACDBig) Calculate(next *big.Float) (result *big.Float) {
	return new(big.Float).Sub(macd.Short.Calculate(next), macd.Long.Calculate(next))
}
