package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

const (
	SiteViews string = "site_views"
)

func init() {
	router.HandleFunc("/newView", AddView).Methods("POST")
	router.HandleFunc("/viewsStats", GetViewsStats).Methods("GET")
}

func AddView(_ http.ResponseWriter, _ *http.Request) {
	AddNewSiteView()
}

func GetViewsStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(SiteViews))
	if err != nil {
		EncodingJSONError(err)
	}
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}
