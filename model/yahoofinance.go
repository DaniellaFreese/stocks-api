package model

type Quote struct {
	Name      string  `json:"longName"`
	Symbol    string  `json:"symbol"`
	Open      float64 `json:"open"`
	High      float64 `json:"regularMarketDayHigh"`
	Low       float64 `json:"regularMarketDayLow"`
	Vol       float64 `json:"regularMarketVolume"`
	PERatio   float64 `json:"trailingPE"`
	PBRatio   float64 `json:"priceToBook"`
	MktCap    float64 `json:"marketCap"`
	FiftyTwoH float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoL float64 `json:"fiftyTwoWeekLow"`
	EPS       float64 `json:"epsTrailingTwelveMonths"`
}

type Result struct {
	Key []Quote `json:"result"`
}

type Response struct {
	Key Result `json:"quoteResponse"`
}
