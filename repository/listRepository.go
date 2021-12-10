package repository

import "github.com/DaniellaFreese/stocks-api/model"

//in real world should be a DB connection object
type ListRepo struct {
	stocks map[int]model.Stock
}

//Create a new List Repo,in real world would initialize db connection
func NewListRepo() *ListRepo {
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
	repo := ListRepo{
		map[int]model.Stock{
			s1.ID: s1,
			s2.ID: s2,
			s3.ID: s3,
		},
	}
	return &repo
}

func (l ListRepo) GetStock(id int) (*model.Stock, error) {
	stock := l.stocks[id]
	return &stock, nil
}

func (l ListRepo) GetStocks() (*[]model.Stock, error) {
	stocks := []model.Stock{}

	for _, v := range l.stocks {
		stocks = append(stocks, v)
	}

	return &stocks, nil
}

//make an api call to get the stock details then add it to the map for now will mock it out
func (l ListRepo) AddStock(stock *model.Stock) (*[]model.Stock, error) {

	id := len(l.stocks) + 1
	(*stock).ID = id
	l.stocks[stock.ID] = (*stock)

	return l.GetStocks()

}

//then add it to the map for now will mock it out
func (l ListRepo) DeleteStock(id int) (*[]model.Stock, error) {
	delete(l.stocks, id)

	return l.GetStocks()
}
