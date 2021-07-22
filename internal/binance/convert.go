package binance

import (
	"encoding/json"
	"log"
	"strings"
)

func JsonToCandlesArray(jsonData []byte) []Candle {
	jsonString := string(jsonData)
	jsonString = strings.ReplaceAll(jsonString, "\"", "")
	var numericData [][]float64

	err := json.Unmarshal([]byte(jsonString), &numericData)
	if err != nil {
		log.Fatal(err)
	}
	return ArrayToCandle(numericData)
}

func ArrayToCandle(arr [][]float64) []Candle {
	var ret []Candle
	for _, v := range arr {
		candle := Candle{
			OpenTime:    int64(v[0]),
			Open:        v[1],
			High:        v[2],
			Low:         v[3],
			Close:       v[4],
			Volume:      v[5],
			CloseTime:   int64(v[6]),
			QAV:         v[7],
			TradesCount: int(v[8]),
			TBAV:        v[9],
			TQAV:        v[10],
			Ignore:      v[11],
		}
		ret = append(ret, candle)
	}
	return ret
}
