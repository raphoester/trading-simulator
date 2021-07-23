package simulator

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"trading-simulator/internal/binance"
	"trading-simulator/internal/helper"
)

func SimulatorHandler(w http.ResponseWriter, r *http.Request) {
	var strategy Strategy
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	currentTime := helper.GetCurrentMilis()

	if err := decoder.Decode(&strategy); err != nil {
		http.Error(w, errors.New("invalid json data format : "+err.Error()).Error(), http.StatusBadRequest)
	}

	// si les dates de début ou de fin n'ont pas été définies
	if strategy.End == 0 {
		strategy.End = helper.GetCurrentMilis()
	}
	if strategy.Start == 0 {
		strategy.Start = strategy.End - 60000000
	}

	fmt.Println(strategy.End)
	fmt.Println(strategy.Start)

	if !helper.IsValidTimeUnit(strategy.TimeUnit) {
		http.Error(w, errors.New("invalid time unit given").Error(), http.StatusBadRequest)
	}
	if strategy.End > currentTime || strategy.Start > currentTime {
		http.Error(w, errors.New("unix time given is after current timestamp").Error(), http.StatusBadRequest)
	}

	pair := strategy.Currency1 + strategy.Currency2
	candles := binance.OHLCVHistory(strategy.Start, strategy.End, strategy.TimeUnit, pair)

	// récupérer la liste des comparaisons
	comparisons := GetComparisonsList(strategy.Condition)
	//récupérer la liste des indicateurs
	indicators := GetIndicatorsList(comparisons)
	// ajouter les indicateurs
	candles, err := AddIndicators(candles, indicators)
	if err != nil {
		http.Error(w, errors.New("unable to add indicators"+err.Error()).Error(), http.StatusBadRequest)
	}
	// effectuer les trades
	wallet := Wallet{
		CUR1: 0,
		CUR2: strategy.Capital,
	}

	wallet, trades := MakeTrades(strategy.Condition, candles, wallet)

	result := Result{
		Condition: strategy.Condition,
		Wallet:    wallet,
		Trades:    trades,
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}
