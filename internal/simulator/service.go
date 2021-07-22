package simulator

import (
	"fmt"
	"time"
	"trading-simulator/internal/binance"
	"trading-simulator/internal/helper"
	"trading-simulator/internal/indicators"
)

func AddIndicators(candles []binance.Candle, indics []Indicator) ([]binance.Candle, error) {
	nativeFields := []string{"Close", "Open"}
	var err error
	backupCandles := candles
	candles = backupCandles
	for _, indicator := range indics {
		if helper.IsInStringList(indicator.Name, nativeFields) {
			continue
		}
		funcName := "Add" + indicator.Name
		candles, err = indicators.CallFuncByName(funcName, candles, indicator.Params)
		if err != nil {
			return nil, err
		}
	}
	return candles, nil
}

func GetComparisonsList(conditions Condition) []Comparison {
	var comps []Comparison
	tempCondition := conditions
	if len(tempCondition.Comparisons) > 0 {
		comps = append(comps, tempCondition.Comparisons...)
	}

	for _, v := range tempCondition.Conditions {
		comps = append(comps, GetComparisonsList(v)...)
	}
	return comps
}

func GetIndicatorsList(comparisons []Comparison) []Indicator {
	var indics []Indicator
	for _, comp := range comparisons {
		indics = append(indics, comp.Indicator1)
		indics = append(indics, comp.Indicator2)
	}
	return indics
}

func MakeTrades(condition Condition, candles []binance.Candle, wallet Wallet) (Wallet, []Trade) {
	var tradeList []Trade
	var trade Trade
	for _, candle := range candles {

		apply, err := ApplyCondition(condition, candle)
		if err != nil {
			fmt.Println(err)
		}
		if apply {
			trade = wallet.BuyCUR1(wallet.CUR2*0.1, candle.Close)
		} else {
			trade = wallet.SellAllCUR1(candle.Close)
		}

		// si le trade a bien été effectué
		if trade.CUR1 != 0 {
			trade.Date = time.Unix(0, candle.CloseTime*1000000)
			tradeList = append(tradeList, trade)
		}
	}

	//liquidation des actifs
	closeTrade := wallet.SellAllCUR1(candles[len(candles)-1].Close)
	closeTrade.Date = time.Unix(0, candles[len(candles)-1].CloseTime*1000000)

	tradeList = append(tradeList, closeTrade)

	return wallet, tradeList
}
