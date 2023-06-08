# IOx 
### FlightSQL NodeJS
Example gRPC Flight SQL API client using NodeJS

#### Usage
Install and import the [Influxdb3-js library](https://github.com/bonitoo-io/influxdb3-js)

```
import {InfluxDB} from '@influxdata/influxdb3-client/src'

let url = 'https://eu-central-1-1.aws.cloud2.influxdata.com/';
let database = 'my-database';
let token = 'my-token';

let client = new InfluxDB({url: url, database: database, token: token});
```

### TBD
