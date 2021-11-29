package repository

import "github.com/DaniellaFreese/stocks-api/model"

type Repository interface {
	GetStock(id int) (*model.Stock, error)
	GetStocks() (*[]model.Stock, error)
	AddStock(ticker string) error
	DeleteStock(id string) error
}

type Controller struct {
	Service Repository
}

func InitController(listRepo *ListRepo) *Controller {
	return &Controller{
		Service: listRepo,
	}
}
