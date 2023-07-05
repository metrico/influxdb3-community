# IOx 
### FlightSQL NodeJS
Example gRPC Flight SQL API client using NodeJS

#### Usage
Install and import the [Influxdb3-js library](https://github.com/bonitoo-io/influxdb3-js)

```javascript
const {InfluxDBClient, Point} = require('@influxdata/influxdb3-client');

const token = "abc123"
const org = "company"
const bucket = "sensors"
const database = org + "_" + bucket

async function main() {
    await insert()
    await query()
}

main()

async function insert(){
    // Connect to HTTP API for inserting line protocol
    const host = "http://localhost:8080"
    const client = new InfluxDBClient({host, token})
    const line = `stat,unit=temperature avg=20.5,max=45.0`
    await client.write(line, bucket, org)
    client.close()
}

async function query(){
    // Connect to gRPc API for querying using flightsql
    const host = "http://localhost:8082"
    const client = new InfluxDBClient({host, token })
    // Execute query
    const query = `
        SELECT * FROM "stat"
        WHERE
        time >= now() - interval '5 minute'
        AND
        "unit" IN ('temperature')
    `
    const queryResult = await client.query(query, database, queryType)
    for await (const row of queryResult) {
        console.log(`avg is ${row.avg}`)
        console.log(`max is ${row.max}`)
    }
    client.close()

}
```

### TBD
