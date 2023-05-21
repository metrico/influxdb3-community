package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	flightsql "github.com/apache/arrow/go/v12/arrow/flight/flightsql/driver"
)

const parameterized = false

func main() {
	config := flightsql.DriverConfig{
		Address: "localhost:8082",
		Token:   "",
		Timeout: 10 * time.Second,
		Params: map[string]string{
			"iox-namespace-name": "company_sensors",
		},
	}
	db, err := sql.Open("flightsql", config.DSN())
	if err != nil {
		log.Fatalf("open failed: %v", err)
	}
	defer db.Close()

	var query string
	var params []interface{}
	if !parameterized {
		query = "SELECT * FROM cpu WHERE time < '2020-06-11T16:54:00Z'::TIMESTAMP AND cpu = 'cpu2'"
	} else {
		query = "SELECT * FROM cpu WHERE cpu = '?' LIMIT 10"
		params = append(params, "cpu2")
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		log.Fatalf("query failed: %s", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("getting columns failed: %s", err)
	}
	log.Println("columns:", columns)

	// Prepare the list of data-points according to the received row
	columnData := make([]interface{}, len(columns))
	columnDataPtr := make([]interface{}, len(columns))

	for i := range columnData {
		columnDataPtr[i] = &columnData[i]
	}

	idx := 0
	for rows.Next() {
		idx++
		if err := rows.Scan(columnDataPtr...); err != nil {
			log.Fatalf("scanning row %d failed: %s", idx, err)
		}
		for j := range columns {
			fmt.Printf("%.22v    ", columnData[j])
		}
		fmt.Println()
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("iterating rows failed: %s", err)
	}
}
