package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DaniellaFreese/stocks-api/model"
	"github.com/DaniellaFreese/stocks-api/repository"
	"github.com/DaniellaFreese/stocks-api/restClient"
	"github.com/go-chi/chi/v5"
)

var control *repository.Controller

func newController(controller *repository.Controller) {
	control = controller
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome page")
}

//should make an api call to get a new stock
func stockID(w http.ResponseWriter, r *http.Request) {
	stock := &model.Stock{}

	idS := chi.URLParam(r, "stockID")
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(stock)
		return
	}
	stock, err = control.Service.GetStock(id)
	json.NewEncoder(w).Encode(stock)
}

//make a call to return the current user's list of stocks
func stockList(w http.ResponseWriter, r *http.Request) {
	stocks, err := control.Service.GetStocks()
	if err != nil {
		http.Error(w, "error returning list", 400)
		return
	}
	json.NewEncoder(w).Encode(stocks)

}

func addStock(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	stock := model.Stock{ID: 4}

	if ticker == "AXP" {
		stock.Company = "American Express"
		stock.Ticker = "AXP"
		stock.Price = 180
		stock.PERatio = 37
	}
	stocks, err := control.Service.AddStock(&stock)
	if err != nil {
		http.Error(w, "error adding stock from list", 400)
		return
	}
	json.NewEncoder(w).Encode(stocks)
}

func removeStock(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "stockID")
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(err)
		http.Error(w, "error deleting stock from list", 400)
		return
	}

	stocks, err := control.Service.DeleteStock(id)
	if err != nil {
		http.Error(w, "error deleting stock from list", 400)
		return
	}

	json.NewEncoder(w).Encode(stocks)
}

func quote(w http.ResponseWriter, r *http.Request) {

	body, err := restClient.NewRestClient().GetFinanceQuote("MSFT")
	if err != nil {
		http.Error(w, "error getting quote details", 500)
	}

	var cont model.Response
	json.Unmarshal(body, &cont)
	json.NewEncoder(w).Encode(cont.Key.Key[0])
}
