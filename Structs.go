package main

import "time"

type OneDayStats struct {
	Day           time.Time `json:"day"`
	NumberOfViews uint16    `json:"numberOfViews"`
}
