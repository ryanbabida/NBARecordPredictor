package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
)

type RecordsAPI struct {
	Datastore datastore.RecordDataStore
}

type ApiError struct {
	internalMessage     error
	userFriendlyMessage string
	statusCode          int
}

type recordsAPIResponse struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	StatusCode   int    `json:"statusCode"`
}

func WriteJSON(w http.ResponseWriter, data any, errorMessage string, statusCode int) {
	res := recordsAPIResponse{
		Data:         data,
		ErrorMessage: errorMessage,
		StatusCode:   statusCode,
	}

	json.NewEncoder(w).Encode(res)
}

func MakeHandlerFunc[T any](f func(w http.ResponseWriter, r *http.Request) (T, *ApiError)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer log.Println("Request URI: ", r.RequestURI, "Duration: ", time.Since(startTime).Nanoseconds())

		v, err := f(w, r)
		if err != nil {
			log.Println(err)
			WriteJSON(w, v, err.userFriendlyMessage, err.statusCode)
			return
		}

		WriteJSON(w, v, "", 200)
	})
}

func (a *RecordsAPI) GetAllRecords(w http.ResponseWriter, r *http.Request) ([]datastore.Record, *ApiError) {
	records, err := a.Datastore.GetAll()
	if err != nil {
		e := fmt.Errorf("GetAllRecords: %w", err)
		return []datastore.Record{}, &ApiError{internalMessage: e, userFriendlyMessage: "Unable to get all records", statusCode: 500}
	}

	return records, nil
}

func (a *RecordsAPI) GetRecordsByYear(w http.ResponseWriter, r *http.Request) ([]datastore.Record, *ApiError) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		return []datastore.Record{}, &ApiError{userFriendlyMessage: "unable to parse year", statusCode: 500}
	}
	if year < 1996 || year > 2018 {
		return []datastore.Record{}, &ApiError{userFriendlyMessage: "invalid year", statusCode: 500}
	}

	records, err := a.Datastore.Get([]string{vars["year"]})
	if err != nil {
		e := fmt.Errorf("GetRecordsByYear: %w", err)
		return []datastore.Record{}, &ApiError{internalMessage: e, userFriendlyMessage: "Unable to get records by year", statusCode: 500}
	}

	return records, nil
}

func (a *RecordsAPI) GetRawDataSet(w http.ResponseWriter, r *http.Request) (datastore.RecordData, *ApiError) {
	data, err := a.Datastore.GetDataSet()
	if err != nil {
		e := fmt.Errorf("GetRawDataSet: %w", err)
		return datastore.RecordData{}, &ApiError{internalMessage: e, userFriendlyMessage: "Unable to get data set", statusCode: 500}
	}

	return data, nil
}
