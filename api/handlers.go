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
var rc *restClient.RestClient

func Controller(controller *repository.Controller) {
	control = controller
}

func newRC(restClient *restClient.RestClient) {
	rc = restClient
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome page")
}

//Return a specific stock from your watch list
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

//return the entire watch list of stocks
func stockList(w http.ResponseWriter, r *http.Request) {
	stocks, err := control.Service.GetStocks()
	if err != nil {
		http.Error(w, "error returning list", 400)
		return
	}
	json.NewEncoder(w).Encode(stocks)

}

//add a stock to the watch list
//need to call yahoo finance to get ticker data, and create stock model and add it to the watchlist
func addStock(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	body, err := rc.GetFinanceQuote(ticker)
	if err != nil {
		http.Error(w, "error getting quote details", 500)
	}

	var cont model.Response
	json.Unmarshal(body, &cont)

	stock := model.Stock{
		Company: cont.Key.Key[0].Name,
		Ticker:  cont.Key.Key[0].Symbol,
		Price:   0.00,
		PERatio: cont.Key.Key[0].PERatio,
	}

	stocks, err := control.Service.AddStock(&stock)
	if err != nil {
		http.Error(w, "error adding stock from list", 400)
		return
	}
	json.NewEncoder(w).Encode(stocks)
}

//remove a stock from the watch list
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

//get latest data on a particular stock ticker
func quote(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	body, err := rc.GetFinanceQuote(ticker)
	if err != nil {
		http.Error(w, "error getting quote details", 500)
	}

	var cont model.Response
	json.Unmarshal(body, &cont)
	json.NewEncoder(w).Encode(cont.Key.Key[0])
}
