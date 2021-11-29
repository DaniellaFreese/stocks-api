package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DaniellaFreese/stocks-api/model"
	"github.com/DaniellaFreese/stocks-api/repository"
	"github.com/go-chi/chi/v5"
)

var listRepository *repository.ListRepository

func createListRepo() {
	listRepository = repository.NewDatabase()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome page")
}

func stockDetails(w http.ResponseWriter, r *http.Request) {
	stock := &model.Stock{}

	idS := chi.URLParam(r, "stockID")
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(stock)
		return
	}

	stock, err = listRepository.GetStock(id)
	json.NewEncoder(w).Encode(stock)
}

func stocksList(w http.ResponseWriter, r *http.Request) {}

func dddStock(w http.ResponseWriter, r *http.Request) {}

func removeStock(w http.ResponseWriter, r *http.Request) {}

func analyzeStock(w http.ResponseWriter, r *http.Request) {}
