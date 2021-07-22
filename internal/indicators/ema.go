package indicators

import (
	"errors"
	"fmt"
	"strconv"
	"trading-simulator/internal/binance"
)

func AddExponentialMovingAverage(candles []binance.Candle, params map[string]string) ([]binance.Candle, error) {
	unitCount, err := strconv.Atoi(params["UnitCount"])
	if err != nil {
		return nil, errors.New("unitcount parameter is not a valid int type")
	}
	for i := 0; i < len(candles); i++ {
		if candles[i].EMA == nil {
			candles[i].EMA = make(map[int]float64)
		}
		if i > unitCount {
			//gnégnégné
			fmt.Println("ajout des EMA pas encore implémenté")
			// candles[i].EMA[unitCount] = CandlesClosePriceAvg(candles[i-unitCount : i])
		}
	}
	return candles, nil
}
