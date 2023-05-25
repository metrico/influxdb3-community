<img src="https://pbs.twimg.com/profile_banners/1967601206/1682514855/1500x500" width=800>

# iox
Pronounced _(eye-ox)_ short for iron oxide. The new core of InfluxDB written in Rust on top of Apache Arrow.

### Motivation
😄 You want to try and experiment with iox low-cost storage, unlimited cardinality and flight sql<br>
😮‍ The iox project is in active development, which is why InfluxDB is not producing builds yet...<br>
😄 No problem! Meet the _unofficial_ InfluxDB 3.0 _"iox"_ musl + docker builder for early adopters

##### Releases
###### amd64/musl
  - [x] docker: `docker pull ghcr.io/metrico/iox-musl:latest`
  - [x] binary: `github.com/metrico/iox-builder/releases/latest/download/influxdb_iox`

<br>

## Get Started

Pull the latest iox static build or docker image. This demo will use both.

#### Static
```
curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb_iox -O && chmod +x influxdb_iox
```
#### Docker
```
docker pull ghcr.io/metrico/iox-musl:latest
```

## Docker Compose
Our compose will start an `all-in-one` iox instance using the provided [compose file](https://gist.github.com/lmangani/c48cf7ef997ed5273ec05a15937c7ad5/raw/a87a13ecad33512ea902705f19ef5866f9a95245/docker-compose.yml)
```
docker-compose up -d
```

This will start iox `router`, `querier`, `ingester` and `compactor` on the same host using the following settings:

```
      - INFLUXDB_IOX_OBJECT_STORE=file
      - INFLUXDB_IOX_DB_DIR=/data
      - INFLUXDB_IOX_BUCKET=iox
      - INFLUXDB_IOX_CATALOG_DSN=sqlite:///data/catalog.sqlite
      - INFLUXDB_IOX_ROUTER_HTTP_BIND_ADDR=iox:8080
      - INFLUXDB_IOX_ROUTER_GRPC_BIND_ADDR=iox:8081
      - INFLUXDB_IOX_QUERIER_GRPC_BIND_ADDR=iox:8082
      - INFLUXDB_IOX_INGESTER_GRPC_BIND_ADDR=iox:8083
      - INFLUXDB_IOX_COMPACTOR_GRPC_BIND_ADDR=iox:8084
```

If you want to use S3/R2/Minio for storage refer to [this example](https://github.com/metrico/iox-builder/blob/main/env.example)

## Testing

It's time to test our fresh instance with inserts and queries

### Health

Check the instance health: `curl http://127.0.0.1:8080/health`

The expected response is `OK`

### Insert

Insert a sample dataset using the Influx V2 API and line protocol to test the `router` API on port 8080
```
wget -qO- "https://github.com/influxdata/influxdb_iox/raw/main/test_fixtures/lineproto/metrics.lp" | \
curl -v "http://127.0.0.1:8080/api/v2/write?org=company&bucket=sensors" --data-binary @-
```

The expected response is `204`

### Query
Let's launch the `sql` client using the `querier` gRPC API on port 8082

* Using Binary: `./influxdb_iox --host http://localhost:8082 sql`
* Using Docker: `docker run -ti --rm ghcr.io/metrico/iox:latest --host http://iox:8082 sql`

#### [Datafusion SQL](https://github.com/influxdata/influxdb_iox/blob/main/docs/sql.md)

The first requirement is to choose a namespace _(or bucket)_ from the available ones:
```
> show namespaces;
+--------------+-----------------+
| namespace_id | name            |
+--------------+-----------------+
| 1            | company_sensors |
+--------------+-----------------+

> use company_sensors;
You are now in remote mode, querying namespace company_sensors
```

Once a namespace is selected, we can display any contained tables:
```
company_sensors> show tables;
+---------------+--------------------+-------------+------------+
| table_catalog | table_schema       | table_name  | table_type |
+---------------+--------------------+-------------+------------+
| public        | iox                | cpu         | BASE TABLE |
| public        | iox                | disk        | BASE TABLE |
| public        | iox                | diskio      | BASE TABLE |
| public        | iox                | mem         | BASE TABLE |
| public        | iox                | net         | BASE TABLE |
| public        | iox                | processes   | BASE TABLE |
| public        | iox                | swap        | BASE TABLE |
| public        | iox                | system      | BASE TABLE |
| public        | system             | queries     | BASE TABLE |
| public        | information_schema | tables      | VIEW       |
| public        | information_schema | views       | VIEW       |
| public        | information_schema | columns     | VIEW       |
| public        | information_schema | df_settings | VIEW       |
+---------------+--------------------+-------------+------------+
```

From any of the available tables, we can select data:

```
company_sensors> select count(*) from cpu;
+-----------------+
| COUNT(UInt8(1)) |
+-----------------+
| 248             |
+-----------------+
```


```
company_sensors> select * from cpu WHERE usage_idle <= 96 limit 1;
+------+---------------------------------+----------------------+-------------+------------------+-------------------+--------------+-----------+------------+---------------+-------------+-------------------+-------------------+
| cpu  | host                            | time                 | usage_guest | usage_guest_nice | usage_idle        | usage_iowait | usage_irq | usage_nice | usage_softirq | usage_steal | usage_system      | usage_user        |
+------+---------------------------------+----------------------+-------------+------------------+-------------------+--------------+-----------+------------+---------------+-------------+-------------------+-------------------+
| cpu0 | Andrews-MBP.hsd1.ma.comcast.net | 2020-06-11T16:52:00Z | 0.0         | 0.0              | 89.56262425447316 | 0.0          | 0.0       | 0.0        | 0.0           | 0.0         | 5.964214711729622 | 4.473161033797217 |
+------+---------------------------------+----------------------+-------------+------------------+-------------------+--------------+-----------+------------+---------------+-------------+-------------------+-------------------+
```

### Golang Client

Query iox from your code using the included [client examples](https://github.com/metrico/iox-builder/tree/main/client)

### Grafana Client

Your service can be used with the [FlightSQL datasource](https://github.com/influxdata/grafana-flightsql-datasource) in Grafana:

![image](https://user-images.githubusercontent.com/1423657/239708678-8e947ae0-6710-4ae4-85c1-903f4c06b085.png)

Once ready, we can perform queries against our data using the FlightSQL query builder:

![image](https://user-images.githubusercontent.com/1423657/239708634-30b48942-d630-4feb-887d-5b6dc37f54d3.png)
