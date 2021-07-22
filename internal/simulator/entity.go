package simulator

import "time"

type Indicator struct {
	Name   string
	Params map[string]string
}

type Comparison struct {
	Operator   string
	Indicator1 Indicator
	Indicator2 Indicator
}

type Condition struct {
	Keyword string

	Conditions  []Condition
	Comparisons []Comparison
}

type Strategy struct {
	Currency1 string
	Currency2 string
	Capital   float64
	TimeUnit  string
	Start     int64
	End       int64

	Condition Condition
}

type Result struct {
	Condition Condition
	Wallet    Wallet
	Trades    []Trade
}

type Trade struct {
	BuyOrSell bool // buy : 1, sell : 0
	Rate      float64
	CUR1      float64
	CUR2      float64
	Date      time.Time
}
