package api

import (
	"net/http"

	"github.com/DaniellaFreese/stocks-api/repository"
	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {

	controller := repository.InitController(repository.NewListRepo())
	newController(controller)

	mux := chi.NewRouter()
	mux.Get("/", homePage)
	mux.Get("/stock/{stockID}", stockDetails)
	mux.Get("/stock", stockList)
	mux.Put("/stock/{ticker}", addStock)

	return mux
}
