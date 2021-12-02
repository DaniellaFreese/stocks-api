package repository

import (
	"fmt"

	"github.com/DaniellaFreese/stocks-api/model"
	"github.com/spf13/viper"
)

type Repository interface {
	GetStock(id int) (*model.Stock, error)
	GetStocks() (*[]model.Stock, error)
	AddStock(stock *model.Stock) (*[]model.Stock, error)
	DeleteStock(id int) (*[]model.Stock, error)
}

func setupEnv() *viper.Viper {
	viper.SetConfigName("localconfig")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w\n", err))
	}
	return viper.GetViper()
}

type Controller struct {
	Service Repository
	EnvVars *viper.Viper
}

func InitController(listRepo *ListRepo) *Controller {
	return &Controller{
		Service: listRepo,
		EnvVars: setupEnv(),
	}
}
