package indicators

import (
	"errors"
	"strconv"
	"trading-simulator/internal/binance"
)

func AddExponentialMovingAverage(candles []binance.Candle, params map[string]string) ([]binance.Candle, error) {
	unitCount, err := strconv.Atoi(params["UnitCount"])
	if err != nil {
		return nil, errors.New("unitcount parameter is not a valid int type")
	}

	step := 5
	pond := float64(2 / (float64(unitCount) + 1))
	// création d'une seed à partir de la SMA
	seed := float64(CandlesClosePriceAvg(candles[:5]))

	//initialisation de l'EMA sur le nombre d'unités requises
	for i := step; i < step+unitCount; i++ {
		seed = ExponentialMovingAverage(candles[i].Close, seed, pond)
		// fmt.Println("seed :", seed)

	}

	if candles[step+unitCount].EMA == nil {
		candles[step+unitCount].EMA = make(map[int]float64)
	}
	candles[step+unitCount].EMA[unitCount] = seed

	for i := step + unitCount + 1; i < len(candles); i++ {
		if candles[i].EMA == nil {
			candles[i].EMA = make(map[int]float64)
		}
		candles[i].EMA[unitCount] = ExponentialMovingAverage(candles[i].Close, candles[i-1].EMA[unitCount], pond)
	}

	return candles, nil
}

func ExponentialMovingAverage(price, last, pond float64) float64 {
	// fmt.Println(price, last, pond)
	return (pond * price) + (1-pond)*last
}
