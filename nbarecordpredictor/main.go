package main

import (
	"os"

	"github.com/ryanbabida/nba-record-predictor-go/api"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
	"golang.org/x/exp/slog"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	datastore, err := datastore.NewCSVStore("data/")
	if err != nil {
		logger.Error(err.Error())
	}

	a := api.NewRecordsAPI(datastore, logger)
	a.Start()
}
