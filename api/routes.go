package api

import (
	"net/http"

	"github.com/DaniellaFreese/stocks-api/repository"
	"github.com/DaniellaFreese/stocks-api/restClient"
	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {

	provideController()
	provideRC()

	mux := chi.NewRouter()
	mux.Get("/", homePage)
	mux.Get("/stock/{stockID}", stockID)
	mux.Get("/stock", stockList)
	mux.Put("/stock/{ticker}", addStock)
	mux.Delete("/stock/{stockID}", removeStock)

	mux.Get("/quote/{ticker}", quote)

	return mux
}

func provideController() {

	controller := repository.InitController(repository.NewListRepo())

	//handler
	Controller(controller)

	//restclient
	restClient.Controller(controller)
}

func provideRC() {

	//restclient initalization for handlers
	newRC(restClient.NewRestClient())
}
