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
	resp, err := http.Get(fmt.Sprintf("http://localhost:8001/watchers/%v", username))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var watchersCount uint16
	err = json.NewDecoder(resp.Body).Decode(&watchersCount)
	if err != nil {
		DecodingJSONError(err)
	}

	AddNewDateWithStats(Watchers, watchersCount)
}

func AddStarsCount(username string) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8001/stars/%v", username))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var starsCount uint16
	err = json.NewDecoder(resp.Body).Decode(&starsCount)
	if err != nil {
		DecodingJSONError(err)
	}

	AddNewDateWithStats(Starts, starsCount)
}

func AddForksCount(username string) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8001/forks/%v", username))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var forksCount uint16
	err = json.NewDecoder(resp.Body).Decode(&forksCount)
	if err != nil {
		DecodingJSONError(err)
	}

	AddNewDateWithStats(Forks, forksCount)
}

func DecodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while encoding JSON: %v\n", err))
}
