package repository

import "github.com/DaniellaFreese/stocks-api/model"

type ListRepository struct {
	stocks map[int]model.Stock
}

func NewDatabase() *ListRepository {
	s1 := model.Stock{
		Company: "Microsoft",
		Ticker:  "MSFT",
		Price:   337.41,
		PERatio: 37,
		ID:      1,
	}
	s2 := model.Stock{
		Company: "Rivian",
		Ticker:  "RVN",
		Price:   112,
		PERatio: 45,
		ID:      2,
	}
	s3 := model.Stock{
		Company: "JP Morgan",
		Ticker:  "JPM",
		Price:   160,
		PERatio: 10,
		ID:      3,
	}
	repo := ListRepository{
		map[int]model.Stock{
			s1.ID: s1,
			s2.ID: s2,
			s3.ID: s3,
		},
	}
	return &repo
}

func (l ListRepository) GetStock(id int) (*model.Stock, error) {
	stock := l.stocks[id]
	return &stock, nil
}

func (l ListRepository) GetStocks() (*[]model.Stock, error) {
	stocks := []model.Stock{}

	for _, v := range l.stocks {
		stocks = append(stocks, v)
	}

	return &stocks, nil
}
