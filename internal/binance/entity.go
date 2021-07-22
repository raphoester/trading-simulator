package binance

type Candle struct {
	OpenTime    int64
	Open        float64
	High        float64
	Low         float64
	Close       float64
	Volume      float64
	CloseTime   int64
	QAV         float64
	TradesCount int
	TBAV        float64
	TQAV        float64
	Ignore      float64
	SMA         map[int]float64
	EMA         map[int]float64
}

type StreamCandle struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Candle    struct {
		StartTime      int64   `json:"t"`
		CloseTime      int64   `json:"T"`
		Symbol         string  `json:"s"`
		Interval       string  `json:"i"`
		FirstTradeID   uint64  `json:"f"`
		LastTradeID    uint64  `json:"L"`
		OpenPrice      float64 `json:"o,string"`
		ClosePrice     float64 `json:"c,string"`
		HighPrice      float64 `json:"h,string"`
		LowPrice       float64 `json:"l,string"`
		BaseAssetVolme float64 `json:"v,string"`
		TradesCount    int     `json:"n"`
		Closed         bool    `json:"x"`
		QAV            float64 `json:"q,string"`
		TBAV           float64 `json:"V,string"`
		TQAV           float64 `json:"Q,string"`
		Ignore         float64 `json:"B,string"`
	} `json:"k"`
}
