package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v12/arrow/flight/flightsql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func dbQuery(ctx context.Context) error {
	url := "iox-api.fly.dev:8082"
  // INFLUX_TOKEN is an environment variable you created for your database READ token
	token := os.Getenv("INFLUX_TOKEN")
	database := "company_sensors"

	// Create a gRPC transport
	pool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("x509: %s", err)
	}

	transport := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(pool, ""))
	opts := []grpc.DialOption{
		transport,
	}

	// Create query client
	client, err := flightsql.NewClient(url, nil, nil, opts...)
	if err != nil {
		return fmt.Errorf("flightsql: %s", err)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	ctx = metadata.AppendToOutgoingContext(ctx, "database", database)

	// Execute query
	query := `SELECT
		count(*)
	FROM
		cpu`
//	WHERE
//		time >= '2020-06-11T18:00:00Z'
//		AND time <= '2020-06-11T19:00:00Z'`

	info, err := client.Execute(ctx, query)
	if err != nil {
		return fmt.Errorf("flightsql flight info: %s", err)
	}
	reader, err := client.DoGet(ctx, info.Endpoint[0].Ticket)
	if err != nil {
		return fmt.Errorf("flightsql do get: %s", err)
	}

	// Print results as JSON
	for reader.Next() {
		record := reader.Record()
		b, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println("RECORD BATCH")
		fmt.Println(string(b))

		if err := reader.Err(); err != nil {
			return fmt.Errorf("flightsql reader: %s", err)
		}
	}

	return nil
}

func main() {
	if err := dbQuery(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
