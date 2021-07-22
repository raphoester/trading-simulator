package indicators

import (
	"errors"
	"strconv"
	"trading-simulator/internal/binance"
)

func AddMovingAverage(candles []binance.Candle, params map[string]string) ([]binance.Candle, error) {
	unitCount, err := strconv.Atoi(params["UnitCount"])
	if err != nil {
		return nil, errors.New("unitcount parameter is not a valid int type")
	}
	for i := 0; i < len(candles); i++ {
		if candles[i].SMA == nil {
			candles[i].SMA = make(map[int]float64)
		}
		if i > unitCount {
			// ajout d'un champ moyenne à l'objet candle numéro i
			// la moyenne est calculée d'après les unitCount champs précédents.
			candles[i].SMA[unitCount] = CandlesClosePriceAvg(candles[i-unitCount : i])
		}
	}
	return candles, nil
}

func CandlesClosePriceAvg(candles []binance.Candle) float64 {
	var sum float64
	for _, v := range candles {
		sum += v.Close
	}
	// fmt.Println("sum : ", sum, "len : ", len(candles))
	return sum / float64(len(candles))
}
