package repository

import "github.com/DaniellaFreese/stocks-api/model"

type Repository interface {
	GetStock(id int) (*model.Stock, error)
	GetStocks() (*[]model.Stock, error)
	//fix this to return the list of stocks
	AddStock(stock *model.Stock) (*[]model.Stock, error)
	DeleteStock(id int) (*[]model.Stock, error)
}

type Controller struct {
	Service Repository
}

func InitController(listRepo *ListRepo) *Controller {
	return &Controller{
		Service: listRepo,
	}
}
