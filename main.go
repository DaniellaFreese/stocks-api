package main

import (
	"log"
	"net/http"

	"github.com/DaniellaFreese/stocks-api/api"
)

const portNumber = ":8090"

func main() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: api.Routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
