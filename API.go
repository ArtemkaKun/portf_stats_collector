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
	CVViews   string = "cv_views"
	Watchers  string = "watchers"
	Starts    string = "stars"
	Forks     string = "forks"
)

func init() {
	router.HandleFunc("/newSiteView", AddSiteView).Methods("POST")
	router.HandleFunc("/viewsStats", GetViewsStats).Methods("GET")

	router.HandleFunc("/newCVView", AddCVView).Methods("POST")
	router.HandleFunc("/CVStats", GetCVViewsStats).Methods("GET")

	router.HandleFunc("/watchers", GetWatchersStats).Methods("GET")
	router.HandleFunc("/stars", GetStarsStats).Methods("GET")
	router.HandleFunc("/forks", GetForksStats).Methods("GET")
}

func GetForksStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(Forks))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetStarsStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(Starts))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetWatchersStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(Watchers))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetCVViewsStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(CVViews))
	if err != nil {
		EncodingJSONError(err)
	}
}

func AddCVView(_ http.ResponseWriter, _ *http.Request) {
	AddNewDate(CVViews)
}

func AddSiteView(_ http.ResponseWriter, _ *http.Request) {
	AddNewDate(SiteViews)
}

func GetViewsStats(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(GetDailyData(SiteViews))
	if err != nil {
		EncodingJSONError(err)
	}
}

func AddWatchersCount(username string) {
	stats := GetGITStat(fmt.Sprintf("http://localhost:8001/watchers/%v", username))
	if stats == 0 {
		return
	}

	AddNewDateWithStats(Watchers, stats)
}

func AddStarsCount(username string) {
	stats := GetGITStat(fmt.Sprintf("http://localhost:8001/stars/%v", username))
	if stats == 0 {
		return
	}

	AddNewDateWithStats(Starts, stats)
}

func AddForksCount(username string) {
	stats := GetGITStat(fmt.Sprintf("http://localhost:8001/forks/%v", username))
	if stats == 0 {
		return
	}

	AddNewDateWithStats(Forks, stats)
}

func GetGITStat(request string) (statCount uint16) {
	resp, err := http.Get(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&statCount)
	if err != nil {
		DecodingJSONError(err)
		return
	}

	return
}

func DecodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while encoding JSON: %v\n", err))
}
