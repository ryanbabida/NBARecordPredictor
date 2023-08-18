package main

import (
	"fmt"
	"os"

	"github.com/ryanbabida/nba-record-predictor/api"
	"github.com/ryanbabida/nba-record-predictor/config"
	"github.com/ryanbabida/nba-record-predictor/datastore"
	"golang.org/x/exp/slog"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	config, err := config.GetConfig("config.yaml")
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	datastore, err := datastore.NewCSVStore(config.DataStore.Filepath, config.DataStore.Files)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	fmt.Println("Starting server...")

	a := api.NewRecordsAPI(datastore, config, logger)

	fmt.Printf("live at http://localhost%s\n", config.Server.Port)

	a.Start()

}
