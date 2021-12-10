package restClient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DaniellaFreese/stocks-api/repository"
)

var control *repository.Controller

func Controller(controller *repository.Controller) {
	control = controller
}

type RestClient struct {
	client *http.Client
}

func NewRestClient() *RestClient {
	fc := RestClient{
		client: &http.Client{},
	}

	return &fc
}

func (c RestClient) GetFinanceQuote(ticker string) ([]byte, error) {
	apiEndpoint := "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=" + ticker
	req, err := http.NewRequest("GET", apiEndpoint, nil)

	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	addHeaders(req)
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	return body, nil

}

func addHeaders(req *http.Request) {
	APIKey, ok := control.EnvVars.Get("RESTCLIENT.APIKEY_YF").(string)
	if !ok {
		log.Fatalf("Invalid type assertion on environment variables")
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("X-API-KEY", APIKey)
}
