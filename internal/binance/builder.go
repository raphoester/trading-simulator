package binance

import (
	"trading-simulator/internal/helper"
)

func OHLCVHistory(start int64, end int64, interval string, symbol string) []Candle {
	intervalInMS := helper.IntervalToMs(interval)
	resultsCount, err := helper.TimeLapseToIntervalCount(start, end, intervalInMS)
	if err != nil {
		return nil
	}
	const maxResultsByRequest = 1000
	if resultsCount <= maxResultsByRequest {
		result := OHLCVQuery(symbol, interval, start, end, maxResultsByRequest)
		return JsonToCandlesArray(result)
	} else {
		callsCount64 := resultsCount / maxResultsByRequest
		if resultsCount%maxResultsByRequest != 0 {
			callsCount64 += 1
		}
		callsCount := int(callsCount64)
		timeLapseByRequest := intervalInMS * maxResultsByRequest

		var ret []Candle
		for i := 0; i < callsCount; i++ {
			result := OHLCVQuery(symbol, interval, start+(int64(i)*timeLapseByRequest), start+(int64(i+1)*timeLapseByRequest), maxResultsByRequest)
			candles := JsonToCandlesArray(result)
			ret = append(ret, candles...)
		}
		return ret
	}
}
