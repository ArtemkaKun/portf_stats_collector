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

func AddNewDate(table string) {
	_, err := Connection.Exec(context.Background(), fmt.Sprintf("INSERT INTO %v VALUES($1)", table), time.Now())
	if err != nil {
		InsertErrorHandler(err)
	}
}

func AddNewDateWithStats(table string, statCount uint16) {
	_, err := Connection.Exec(context.Background(), fmt.Sprintf("INSERT INTO %v VALUES($1, $2)", table), time.Now(), statCount)
	if err != nil {
		InsertErrorHandler(err)
	}
}

func GetDailyData(tableName string) (dailyStats []OneDayStats) {
	var dataTable pgx.Rows
	var err error

	if tableName == "watchers" || tableName == "stars" || tableName == "forks" {
		dataTable, err = Connection.Query(context.Background(), fmt.Sprintf("SELECT * FROM %v", tableName))
		if err != nil {
			QueryErrorHandler(err)
		}

		goto getData
	}

	dataTable, err = Connection.Query(context.Background(), fmt.Sprintf("SELECT date, COUNT(date) FROM %v GROUP BY date", tableName))
	if err != nil {
		QueryErrorHandler(err)
	}

getData:
	defer dataTable.Close()

	dailyStats = collectData(dataTable)

	return
}

func collectData(dataTable pgx.Rows) (dailyStats []OneDayStats) {
	for dataTable.Next() {
		dayStats := new(OneDayStats)

		err := dataTable.Scan(&dayStats.Day, &dayStats.NumberOfStats)
		if err != nil {
			QueryErrorHandler(err)
			continue
		}
		dailyStats = append(dailyStats, *dayStats)
	}

	if dataTable.Err() != nil {
		QueryErrorHandler(dataTable.Err())
	}
	return dailyStats
}

func QueryErrorHandler(err error) {
	fmt.Println(fmt.Errorf("Error when try using pgx.Query/pgx.QueryRow: %v", err))
}

func InsertErrorHandler(err error) {
	fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
}
