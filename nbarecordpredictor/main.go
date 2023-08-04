package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
)

type RecordsAPI struct {
	datastore datastore.RecordDataStore
}

type RecordsAPIResponse struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	StatusCode   int    `json:"statusCode"`
}

func makeHandlerFunc[T any](f func(w http.ResponseWriter, r *http.Request) (T, error)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer log.Println("Request URI: ", r.RequestURI, "Duration: ", time.Since(startTime))

		v, err := f(w, r)

		var res RecordsAPIResponse

		if err != nil {
			log.Println(err)
			res = RecordsAPIResponse{
				v,
				err.Error(),
				200,
			}

			json.NewEncoder(w).Encode(res)
			return
		}

		res = RecordsAPIResponse{
			v,
			"",
			200,
		}

		json.NewEncoder(w).Encode(res)
	})
}

func (a *RecordsAPI) GetAllRecords(w http.ResponseWriter, r *http.Request) ([]datastore.Record, error) {
	if r.Method != "GET" {
		return []datastore.Record{}, fmt.Errorf("%s only supports GET requests", r.RequestURI)
	}

	records, err := a.datastore.GetAll()
	if err != nil {
		return []datastore.Record{}, err
	}

	return records, nil
}

func (a *RecordsAPI) GetRecordsByYear(w http.ResponseWriter, r *http.Request) ([]datastore.Record, error) {
	if r.Method != "GET" {
		return []datastore.Record{}, fmt.Errorf("%s only supports GET requests", r.RequestURI)
	}

	records, err := a.datastore.Get([]string{"1997"})
	if err != nil {
		return []datastore.Record{}, err
	}

	return records, nil
}

func main() {
	fmt.Println("Initializing CSV datastore...")
	datastore, err := datastore.NewCSVStore("data/")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Done!")

	a := &RecordsAPI{datastore}
	r := mux.NewRouter()

	r.HandleFunc("/records", makeHandlerFunc(a.GetAllRecords))
	r.HandleFunc("/records/{year}", makeHandlerFunc(a.GetRecordsByYear))

	http.ListenAndServe(":3000", r)
}
