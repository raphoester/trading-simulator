package indicators

import "trading-simulator/internal/binance"

func CallFuncByName(name string, candles []binance.Candle, params map[string]string) ([]binance.Candle, error) {
	var err error
	err = nil
	switch name {
	case "AddSMA":
		candles, err = AddMovingAverage(candles, params)
	case "AddEMA":
		candles, err = AddExponentialMovingAverage(candles, params)
	}

	return candles, err
}
