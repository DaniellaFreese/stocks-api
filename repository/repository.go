package repository

import "github.com/DaniellaFreese/stocks-api/model"

//domain layer
type StockRepository interface {
	GetStock(id string) (*model.Stock, error)
	GetStocks() (*[]model.Stock, error)
	AddStock() error
	DeleteStock(id string) error
}
