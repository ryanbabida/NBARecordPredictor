package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryanbabida/nba-record-predictor-go/api"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
)

func main() {
	datastore, err := datastore.NewCSVStore("data/")
	if err != nil {
		log.Fatalf("%s", err)
	}

	a := &api.RecordsAPI{Datastore: datastore}
	r := mux.NewRouter()

	r.HandleFunc("/records", api.MakeHandlerFunc(a.GetAllRecords)).Methods("GET")
	r.HandleFunc("/records/{year}", api.MakeHandlerFunc(a.GetRecordsByYear)).Methods("GET")
	r.HandleFunc("/data", api.MakeHandlerFunc(a.GetRawDataSet)).Methods("GET")

	http.ListenAndServe(":3000", r)
}
