package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ryanbabida/nba-record-predictor-go/datastore"
	"golang.org/x/exp/slog"
)

type recordsAPI struct {
	Datastore datastore.RecordDataStore
	Router    *mux.Router
	Logger    *slog.Logger
}

func NewRecordsAPI(datastore datastore.RecordDataStore, logger *slog.Logger) *recordsAPI {
	a := &recordsAPI{Datastore: datastore, Logger: logger, Router: mux.NewRouter()}

	a.Router.HandleFunc("/records", a.GetAllRecords).Methods("GET")
	a.Router.HandleFunc("/records/{year}", a.GetRecordsByYear).Methods("GET")
	a.Router.HandleFunc("/data", a.GetRawDataSet).Methods("GET")

	return a
}

func (a *recordsAPI) Start() {
	m := http.TimeoutHandler(a.Router, time.Second*5, "request exceeded timeout")
	http.ListenAndServe(":3000", m)
}

type recordsAPIResponse struct {
	Data       any `json:"data"`
	StatusCode int `json:"statusCode"`
}

type recordsAPIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	res := recordsAPIResponse{
		Data:       data,
		StatusCode: statusCode,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func (a *recordsAPI) WriteError(w http.ResponseWriter, err error, message string, statusCode int) {
	a.Logger.Error(
		"Error occurred",
		"internalMessage", err,
		"statusCode", statusCode,
		"userFriendlyMessage", message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := recordsAPIError{
		Message:    message,
		StatusCode: statusCode,
	}
	json.NewEncoder(w).Encode(res)
}

func (a *recordsAPI) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	records, err := a.Datastore.GetAll()
	time.Sleep(time.Second * 10)
	if err != nil {
		e := fmt.Errorf("GetAllRecords: %w", err)
		a.WriteError(w, e, "unable to get records", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, records, http.StatusOK)
}

func (a *recordsAPI) GetRecordsByYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		e := fmt.Errorf("GetRecordsByYear - unable to parse year: %w", err)
		a.WriteError(w, e, "unable to parse year", http.StatusBadRequest)
		return
	}
	if year < 1996 || year > 2018 {
		e := fmt.Errorf("GetRecordsByYear - invalid year")
		a.WriteError(w, e, "invalid year, year must be betwen [1996, 2018]", http.StatusBadRequest)
		return
	}

	records, err := a.Datastore.Get([]string{vars["year"]})
	if err != nil {
		e := fmt.Errorf("GetRecordsByYear: %w", err)
		a.WriteError(w, e, "unable to get records by year", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, records, http.StatusOK)
}

func (a *recordsAPI) GetRawDataSet(w http.ResponseWriter, r *http.Request) {
	data, err := a.Datastore.GetDataSet()
	if err != nil {
		e := fmt.Errorf("GetRawDataSet: %w", err)
		a.WriteError(w, e, "unable to get raw data set", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, data, http.StatusOK)
}
