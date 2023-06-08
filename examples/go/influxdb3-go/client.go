package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/InfluxCommunity/influxdb3-go/influx"
	"os"
)

func main() {

	// Get Query settings
	query := flag.String("query", "SELECT 'flight' as hello", "FlightSQL Query")
	database := flag.String("database", "company_sensors", "FlightSQL Bucket")
	url := flag.String("url", "http://localhost:8082", "FlightSQL API URL")
	token := flag.String("token", "", "FlightSQL Auth Token")
	if os.Getenv("INFLUXDB_URL") != "" {
		flag.Set("url", os.Getenv("INFLUXDB_URL"))
	}
	if os.Getenv("INFLUXDB_TOKEN") != "" {
		flag.Set("token", os.Getenv("INFLUXDB_TOKEN"))
	}
	if os.Getenv("INFLUXDB_DATABASE") != "" {
		flag.Set("database", os.Getenv("INFLUXDB_DATABASE"))
	}
	flag.Parse()

	// Initialize IOx Client
	client, err := influx.New(influx.Configs{
		HostURL:   *url,
		AuthToken: *token,
	})

	if err != nil {
		panic(err)
	}
	defer func(client *influx.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)


	// Execute Query
	iterator, err := client.Query(context.Background(), *database, *query)

	if err != nil {
		panic(err)
	}

	// Return JSON
	results := []map[string]interface{}{}

	for iterator.Next() {
		value := iterator.Value()
		results = append(results, value)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
