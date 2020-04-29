package main

import "time"

type OneDayStats struct {
	Day           time.Time `json:"day"`
	NumberOfStats uint16    `json:"numberOfViews"`
}
