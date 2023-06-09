# IOx 
### FlightSQL GO
This client example is based on the Influxdb3-go official library from InfluxCommunity.

## Build
```
go mod init flight
go mod tidy
go build -ldflags='-extldflags=-static -s -w' -o flight client.go
```

## Usage
```
Usage of ./flight:
  -database string
    	FlightSQL Bucket (default "company_sensors")
  -query string
    	FlightSQL Query (default "SELECT 'flight' as hello")
  -token string
    	FlightSQL Auth Token
  -url string
    	FlightSQL API URL (default "http://localhost:8082")
```

### ENV
The following ENV parameters will override command line settings:

- `INFLUXDB_URL`, `INFLUXDB_TOKEN`, `INFLUXDB_DATABASE`
