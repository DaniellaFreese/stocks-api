package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	createListRepo()

	mux := chi.NewRouter()
	mux.Get("/", homePage)
	mux.Get("/stock/{stockID}", stockDetails)

	return mux
}
