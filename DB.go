package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

var Connection *pgx.Conn

func init() {
	var err error
	Connection, err = pgx.Connect(context.Background(), "postgres://postgres:1337@localhost:5432/portf-stats")
	if err != nil {
		log.Panic(fmt.Errorf("Unable to Connection to database: %v\n", err))
	} else {
		fmt.Println("Connected to PSQL!")
	}
}

func AddNewSiteView() {
	_, err := Connection.Exec(context.Background(), "INSERT INTO public.site_views VALUES($1)", time.Now())
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}

func GetViewsData() (viewsStats []OneDayStats) {
	dataTable, err := Connection.Query(context.Background(), "SELECT date, COUNT(date) FROM site_views GROUP BY date")
	if err != nil {
		QueryErrorHandler(err)
	}

	defer dataTable.Close()

	for dataTable.Next() {
		dayStats := new(OneDayStats)

		err = dataTable.Scan(&dayStats.Day, &dayStats.NumberOfViews)
		if err != nil {
			QueryErrorHandler(err)
			continue
		}
		viewsStats = append(viewsStats, *dayStats)
	}

	if dataTable.Err() != nil {
		QueryErrorHandler(err)
	}

	return
}

func QueryErrorHandler(err error) {
	fmt.Println(fmt.Errorf("Error when try using pgx.Query/pgx.QueryRow: %v", err))
}
