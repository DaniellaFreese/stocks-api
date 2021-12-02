package api

import (
	"net/http"

	"github.com/DaniellaFreese/stocks-api/repository"
	"github.com/DaniellaFreese/stocks-api/restClient"
	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {

	provideController()

	mux := chi.NewRouter()
	mux.Get("/", homePage)
	mux.Get("/stock/{stockID}", stockID)
	mux.Get("/stock", stockList)
	mux.Put("/stock/{ticker}", addStock)
	mux.Get("/quote/{ticker}", quote)

	return mux
}

func provideController() {
	//handler
	controller := repository.InitController(repository.NewListRepo())
	newController(controller)
	//restclient
	restClient.NewController(controller)
}
