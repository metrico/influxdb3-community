# IOx 
### FlightSQL GO
Example gRPC Flight SQL API client using golang

#### Usage
```
go mod tidy
go run flightsql.go
```

#### Static 
```
CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o flightsql flightsql.go 
```
