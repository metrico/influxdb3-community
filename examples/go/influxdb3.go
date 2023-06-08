package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/InfluxCommunity/influxdb3-go/influx"
)

func main() {

	query := flag.String("query", "SELECT 1", "FlightSQL Query")
	database := flag.String("database", "company_sensors", "FlightSQL Bucket")
	url := flag.String("url", "http://localhost:8082", "FlightSQL API URL")
	token := flag.String("token", "", "FlightSQL Auth Token")
	flag.Parse()

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

	iterator, err := client.Query(context.Background(), *database, *query)

	if err != nil {
		panic(err)
	}

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
