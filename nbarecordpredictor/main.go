package main

import (
	"log"

	"github.com/ryanbabida/nba-record-predictor-go/api"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
)

func main() {
	datastore, err := datastore.NewCSVStore("data/")
	if err != nil {
		log.Fatalf("%s", err)
	}

	a := api.NewRecordsAPI(datastore)
	a.Start()
}
