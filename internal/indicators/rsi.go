package indicators

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"trading-simulator/internal/binance"
)

func AddRelativeStrengthIndex(candles []binance.Candle, params map[string]string) ([]binance.Candle, error) {
	for i := 0; i < len(candles); i++ {

		period, err := strconv.Atoi(params["UnitCount"])
		if err != nil {
			return nil, errors.New("unitcount parameter is not a valid int type")
		}
		//  trouver la liste des 5 dernières hausses et des 5 dernières baisses
		stop := false
		var upList []float64
		var downList []float64

		initCount := 0
		for i := 1; !stop; i++ {
			diff := candles[i-1].Close - candles[i].Close
			if diff > 0 {
				upList = append(upList, diff)
			} else {
				downList = append(downList, diff)
			}
			if len(upList) >= 5 && len(downList) >= 5 {
				initCount = i
				stop = true
			}
		}

		fmt.Println(initCount)

		// trouver la SMA à partir des listes
		upSMA := SMAFromFloatList(upList[:len(upList)-5])
		downSMA := SMAFromFloatList(downList[:len(downList)-5])

		if candles[period].RSI == nil {
			candles[period].RSI = make(map[int]float64)
		}
		//initialisation de la SMMA selon le nombre de jours requis
		upSMMA := upSMA
		downSMMA := downSMA

		for i := initCount + 1; i < initCount+1+period; i++ {
			diff := candles[i-1].Close - candles[i].Close
			if diff > 0 {
				upSMMA = ExponentialMovingAverage(diff, upSMMA, float64(1/period))
			} else if diff < 0 {
				downSMMA = ExponentialMovingAverage(math.Abs(diff), downSMMA, float64(1/period))
			} else {
				downSMMA = ExponentialMovingAverage(0, downSMMA, float64(1/period))
				upSMMA = ExponentialMovingAverage(0, upSMMA, float64(1/period))
			}
		}

		ending := initCount + 1 + period
		// calculer la SSMA à partir de la fin de l'initialisation

		for i = ending; i < len(candles); i++ {
			diff := candles[i-1].Close - candles[i].Close
			if diff > 0 {
				upSMMA = ExponentialMovingAverage(diff, upSMMA, float64(1/period))
			} else if diff < 0 {
				downSMMA = ExponentialMovingAverage(math.Abs(diff), downSMMA, float64(1/period))
			} else {
				downSMMA = ExponentialMovingAverage(0, downSMMA, float64(1/period))
				upSMMA = ExponentialMovingAverage(0, upSMMA, float64(1/period))
			}

			if candles[i].RSI == nil {
				candles[i].RSI = make(map[int]float64)
			}
			candles[i].RSI[period] = RSI(upSMMA, downSMMA)

		}

	}

	return candles, nil
}

func RSI(U float64, D float64) float64 {
	return 100 - (100 / (1 + (U / D)))
}

func SMAFromFloatList(list []float64) float64 {
	var sum float64
	for _, v := range list {
		sum += v
	}
	return sum / float64(len(list))
}
