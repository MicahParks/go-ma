package ma_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ma"
)

func BenchmarkMACDBig_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	for _, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func BenchmarkMACDFloat_Calculate(b *testing.B) {
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	for _, p := range prices[ma.DefaultLongMACDPeriod:] {
		macd.Calculate(p)
	}
}

func TestMACDBig_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMABig(bigPrices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMABig(ma.DefaultShortMACDPeriod, shortSMA, nil)
	for _, p := range bigPrices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMABig(bigPrices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMABig(ma.DefaultLongMACDPeriod, longSMA, nil)

	macd := ma.NewMACDBig(longEMA, shortEMA)

	var res float64
	var bigRes *big.Float
	for i, p := range bigPrices[ma.DefaultLongMACDPeriod:] {
		bigRes, _, _ = macd.Calculate(p)
		res, _ = bigRes.Float64()
		if res != results[i] {
			t.FailNow()
		}
	}
}

func TestMACDFloat_Calculate(t *testing.T) {
	_, shortSMA := ma.NewSMAFloat(prices[:ma.DefaultShortMACDPeriod])
	shortEMA := ma.NewEMAFloat(ma.DefaultShortMACDPeriod, shortSMA, 0)
	for _, p := range prices[ma.DefaultShortMACDPeriod:ma.DefaultLongMACDPeriod] {
		shortEMA.Calculate(p)
	}

	_, longSMA := ma.NewSMAFloat(prices[:ma.DefaultLongMACDPeriod])
	longEMA := ma.NewEMAFloat(ma.DefaultLongMACDPeriod, longSMA, 0)

	macd := ma.NewMACDFloat(longEMA, shortEMA)

	var res float64
	for i, p := range prices[ma.DefaultLongMACDPeriod:] {
		res, _, _ = macd.Calculate(p)
		if res != results[i] {
			t.FailNow()
		}
	}
}

func floatToBig(s []float64) (b []*big.Float) {
	l := len(s)
	b = make([]*big.Float, l)
	for i := 0; i < l; i++ {
		b[i] = big.NewFloat(s[i])
	}
	return b
}

var (
	prices = []float64{
		459.99,
		448.85,
		446.06,
		450.81,
		442.8,
		448.97,
		444.57,
		441.4,
		430.47,
		420.05,
		431.14,
		425.66,
		430.58,
		431.72,
		437.87,
		428.43,
		428.35,
		432.5,
		443.66,
		455.72,
		454.49,
		452.08,
		452.73,
		461.91,
		463.58,
		461.14,
		452.08,
		442.66,
		428.91,
		429.79,
		431.99,
		427.72,
		423.2,
		426.21,
		426.98,
		435.69,
		434.33,
		429.8,
		419.85,
		426.24,
		402.8,
		392.05,
		390.53,
		398.67,
		406.13,
		405.46,
		408.38,
		417.2,
		430.12,
		442.78,
		439.29,
		445.52,
		449.98,
		460.71,
		458.66,
		463.84,
		456.77,
		452.97,
		454.74,
		443.86,
		428.85,
		434.58,
		433.26,
		442.93,
		439.66,
		441.35,
	}
	bigPrices = floatToBig(prices)
	results   = []float64{
		7.70337838145673003964475356042385101318359375,
		6.41607475695587936570518650114536285400390625,
		4.23751978326481548720039427280426025390625000,
		2.55258332486573635833337903022766113281250000,
		1.37888571985365615546470507979393005371093750,
		0.10298149119910249282838776707649230957031250,
		-1.25840195280312627801322378218173980712890625,
		-2.07055819009491415272350423038005828857421875,
		-2.62184232825688923185225576162338256835937500,
		-2.32906674045494810343370772898197174072265625,
		-2.18163211479298979611485265195369720458984375,
		-2.40262627286642782564740628004074096679687500,
		-3.34212168135286447068210691213607788085937500,
		-3.53036313607753982068970799446105957031250000,
		-5.50747124862732562178280204534530639648437500,
		-7.85127422856083967417362146079540252685546875,
		-9.71936745524880052471417002379894256591796875,
		-10.42286650800718916798359714448451995849609375,
		-10.26016215893747585141682066023349761962890625,
		-10.06920961013429405284114181995391845703125000,
		-9.57191961195360363490181043744087219238281250,
		-8.36963349244507526236702688038349151611328125,
		-6.30163572368542190815787762403488159179687500,
		-3.59968150914517082128440961241722106933593750,
		-1.72014836089687150888494215905666351318359375,
		0.26900323229938294389285147190093994140625000,
		2.18017324711348692289902828633785247802734375,
		4.50863780861044460834818892180919647216796875,
		6.11802015384472497316892258822917938232421875,
		7.72243059351438887460972182452678680419921875,
		8.32745380871403995115542784333229064941406250,
		8.40344118462587630347115918993949890136718750,
		8.50840632319352607737528160214424133300781250,
		7.62576184402922763183596543967723846435546875,
		5.64994908292868558419286273419857025146484375,
		4.49465476488205695204669609665870666503906250,
		3.43298936168440604888019151985645294189453125,
		3.33347385363293824411812238395214080810546875,
		2.95666285611525836429791525006294250488281250,
		2.76256121582525793201057240366935729980468750,
	}
)
