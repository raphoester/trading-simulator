package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const queryPrefix = "https://api.binance.com/api/v3"

type ErrorResponse struct {
	Code int
	Msg  string
}

func OHLCVQuery(symbol string, interval string, start int64, end int64, limit int) []byte {

	url := fmt.Sprintf("%v/klines?symbol=%s&interval=%s&startTime=%d&endTime=%d&limit=%d", queryPrefix, symbol, interval, start, end, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var error ErrorResponse
	err = json.Unmarshal(body, &error)
	if err == nil {
		fmt.Println(error)
		log.Fatal()
	}
	return body
}
