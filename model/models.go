package model

type Stock struct {
	Company string  `json:"company"`
	Ticker  string  `json:"ticker"`
	Price   float64 `json:"price"`
	PERatio float64 `json:"peRatio"`
	ID      int     `json:"id"`
}
