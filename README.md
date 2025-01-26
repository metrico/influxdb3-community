![image](https://github.com/metrico/iox-community/assets/1423657/52ff076a-4261-48f3-ad43-e3183e0645cd)

# InfluxDB 3.x "IOx" Community
Community Builds and Containers for InfluxDB 3.0 IOx _(eye-ox)_ aka _Edge_

### Motivation
😄 You want to try and experiment with IOx low-cost storage, unlimited cardinality and flight sql<br>
🥵‍ The IOx project is in _"Cloud Only"_ mode and InfluxDB is not producing builds yet...<br>
😄 No problem! Meet the _unofficial_ InfluxDB 3.0 _"IOx"_ musl + docker builder for early adopters

#### Nightly Builds
###### amd64/musl
  - [x] [docker](https://github.com/metrico/iox-community/pkgs/container/influxdb-edge-musl): `docker pull ghcr.io/metrico/influxdb-edge-musl:latest`
  - [x] [binary](https://github.com/metrico/iox-community/releases): `github.com/metrico/iox-community/releases/latest/download/influxdb3`

<br>

## Get Started

This guide uses Docker and docker-compose. You can run locally using a [static build](https://github.com/metrico/iox-community/releases).

#### Static
```bash
curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb3 -O \
&& chmod +x influxdb3
```

#### Docker

The default compose uses local filesystem. Use the full recipe for Object Storage and Postgres Catalog.

```
docker-compose up -f docker-compose.yml -d
```

Your local IOx endpoint should be ready on port `8086`

#### Health

Check the instance health: 
```
curl http://127.0.0.1:8086/health
```

The expected response is `OK`


<br>
  
<details>
    <summary><h2>IOx Settings</h2> Deploy IOx using different settings</summary>  
    
<br>  

```
      INFLUXDB3_MAX_HTTP_REQUEST_SIZE: "10485760"
      INFLUXDB3_GEN1_DURATION: "10m"
      INFLUXDB3_WAL_FLUSH_INTERVAL: "1s"
      INFLUXDB3_WAL_SNAPSHOT_SIZE: "600"
      INFLUXDB3_NUM_WAL_FILES_TO_KEEP: "50"
      INFLUXDB3_WAL_MAX_WRITE_BUFFER_SIZE: "100000"
      INFLUXDB3_BUFFER_MEM_LIMIT_MB: "5000"
      INFLUXDB3_PARQUET_MEM_CACHE_SIZE_MB: "1000"
      INFLUXDB3_FORCE_SNAPSHOT_MEM_THRESHOLD: "70%"
      INFLUXDB3_TELEMETRY_DISABLE_UPLOAD: true
      INFLUXDB3_NODE_IDENTIFIER_PREFIX: "bucket-id"
      INFLUXDB3_OBJECT_STORE: "file"
      INFLUXDB3_DB_DIR: "/data"
```

As write requests come in to the server, they are parsed, validated, and put into an in-memory WAL buffer. This buffer is flushed every second by default (can be changed through configuration), which will create a WAL file. Once the data is flushed to disk, it is put into a queryable in-memory buffer and then a response is sent back to the client that the write was successful. That data will now show up in queries to the server.

InfluxDB periodically snapshots the WAL to persist the oldest data in the queryable buffer, allowing the server to remove old WAL files. By default, the server will keep up to 900 WAL files buffered up (15 minutes of data) and attempt to persist the oldest 10 minutes, keeping the most recent 5 minutes around.

When the data is persisted out of the queryable buffer it is put into the configured object store as Parquet files. Those files are also put into an in-memory cache so that queries against the most recently persisted data do not have to go to object storage.

Each server needs an identifier for writing to object storage and as an identifier that is added to replicated writes, Write Buffer segments and Chunks. Must be unique in a group of connected or semi-connected IOx servers. Must be a number that can be represented by a 32-bit unsigned integer.

```
      INFLUXDB3_NODE_IDENTIFIER_PREFIX: 1
```


The demo catalog uses *sqlite* by default. To enable persistent catalog using *postgres*, use the following:

```
      - INFLUXDB_IOX_CATALOG_DSN=postgres://postgres@localhost:5432/postgres
```

The demo stores to filesystem. To enable S3/R2/Minio object storage use the following parameters:

```
      - INFLUXDB3_OBJECT_STORE=s3
      - AWS_ACCESS_KEY_ID=access_key_value
      - AWS_SECRET_ACCESS_KEY=secret_access_key_value
      - AWS_DEFAULT_REGION=us-east-2
      - INFLUXDB3_BUCKET=bucket-name
      - AWS_ENDPOINT = http://minio:9000
```

For other storage options refer to [env example](https://github.com/metrico/iox-builder/blob/main/env.example)

### API Proxy

To emulate InfluxDB3.0 Cloud works, an nginx proxy is included to serve all IOx services from a single endpoint.
```
events {}
http {
  server {
    listen 8086;
    http2 on;
    location /api {
       proxy_pass_request_headers on;
       proxy_pass http://iox:8080;
    }
    location /health {
       proxy_pass http://iox:8080;
    }
    location / {
       proxy_pass_request_headers on;
       if ($http_content_type = "application/grpc") {
            grpc_pass iox:8082;
       }
       proxy_pass http://iox:8082;
    }
  }
}
```


</details>
  

<details open=true>
    <summary><h2>IOX Insert & Query</h2> Validate your IOx Setup</summary>

<br>

Let's start testing and using your brand new IOx instance!  
  

### Line Protocol Examples

Our goal is observability formats ingestion into IOx. Here are some scope examples:

#### Metrics
```
avalanche_metric_mmmmm_0_71 cycle_id="0",gauge=29,host.name="generate-metrics-avalanche",label_key_kkkkk_0="label_val_vvvvv_0",label_key_kkkkk_1="label_val_vvvvv_1",label_key_kkkkk_2="label_val_vvvvv_2",label_key_kkkkk_3="label_val_vvvvv_3",label_key_kkkkk_4="label_val_vvvvv_4",label_key_kkkkk_5="label_val_vvvvv_5",label_key_kkkkk_6="label_val_vvvvv_6",label_key_kkkkk_7="label_val_vvvvv_7",label_key_kkkkk_8="label_val_vvvvv_8",label_key_kkkkk_9="label_val_vvvvv_9",port="9090",scheme="http",series_id="3",service.name="otel-collector" 1613772311130000000
avalanche_metric_mmmmm_0_71 cycle_id="0",gauge=16,host.name="generate-metrics-avalanche",label_key_kkkkk_0="label_val_vvvvv_0",label_key_kkkkk_1="label_val_vvvvv_1",label_key_kkkkk_2="label_val_vvvvv_2",label_key_kkkkk_3="label_val_vvvvv_3",label_key_kkkkk_4="label_val_vvvvv_4",label_key_kkkkk_5="label_val_vvvvv_5",label_key_kkkkk_6="label_val_vvvvv_6",label_key_kkkkk_7="label_val_vvvvv_7",label_key_kkkkk_8="label_val_vvvvv_8",label_key_kkkkk_9="label_val_vvvvv_9",port="9090",scheme="http",series_id="4",service.name="otel-collector" 1613772311130000000
avalanche_metric_mmmmm_0_71 cycle_id="0",gauge=22,host.name="generate-metrics-avalanche",label_key_kkkkk_0="label_val_vvvvv_0",label_key_kkkkk_1="label_val_vvvvv_1",label_key_kkkkk_2="label_val_vvvvv_2",label_key_kkkkk_3="label_val_vvvvv_3",label_key_kkkkk_4="label_val_vvvvv_4",label_key_kkkkk_5="label_val_vvvvv_5",label_key_kkkkk_6="label_val_vvvvv_6",label_key_kkkkk_7="label_val_vvvvv_7",label_key_kkkkk_8="label_val_vvvvv_8",label_key_kkkkk_9="label_val_vvvvv_9",port="9090",scheme="http",series_id="5",service.name="otel-collector" 1613772311130000000
avalanche_metric_mmmmm_0_71 cycle_id="0",gauge=90,host.name="generate-metrics-avalanche",label_key_kkkkk_0="label_val_vvvvv_0",label_key_kkkkk_1="label_val_vvvvv_1",label_key_kkkkk_2="label_val_vvvvv_2",label_key_kkkkk_3="label_val_vvvvv_3",label_key_kkkkk_4="label_val_vvvvv_4",label_key_kkkkk_5="label_val_vvvvv_5",label_key_kkkkk_6="label_val_vvvvv_6",label_key_kkkkk_7="label_val_vvvvv_7",label_key_kkkkk_8="label_val_vvvvv_8",label_key_kkkkk_9="label_val_vvvvv_9",port="9090",scheme="http",series_id="6",service.name="otel-collector" 1613772311130000000
avalanche_metric_mmmmm_0_71 cycle_id="0",gauge=51,host.name="generate-metrics-avalanche",label_key_kkkkk_0="label_val_vvvvv_0",label_key_kkkkk_1="label_val_vvvvv_1",label_key_kkkkk_2="label_val_vvvvv_2",label_key_kkkkk_3="label_val_vvvvv_3",label_key_kkkkk_4="label_val_vvvvv_4",label_key_kkkkk_5="label_val_vvvvv_5",label_key_kkkkk_6="label_val_vvvvv_6",label_key_kkkkk_7="label_val_vvvvv_7",label_key_kkkkk_8="label_val_vvvvv_8",label_key_kkkkk_9="label_val_vvvvv_9",port="9090",scheme="http",series_id="7",service.name="otel-collector" 1613772311130000000
```

#### Logs
```
logs fluent.tag="fluent.info",pid=18i,ppid=9i,worker=0i 1613769568895331700
logs fluent.tag="fluent.debug",instance=1720i,queue_size=0i,stage_size=0i 1613769568895697200
logs fluent.tag="fluent.info",worker=0i 1613769568896515100
```

#### Tracing Spans
```
spans end_time_unix_nano="2021-02-19 20:50:25.6893952 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="okey-dokey",net.peer.ip="1.2.3.4",parent_span_id="d5270e78d85f570f",peer.service="tracegen-client",service.name="tracegen",span.kind="server",span_id="4c28227be6a010e1",status_code="STATUS_CODE_OK",trace_id="7d4854815225332c9834e6dbf85b9380" 1613767825689169000
spans end_time_unix_nano="2021-02-19 20:50:25.6893952 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="lets-go",net.peer.ip="1.2.3.4",peer.service="tracegen-server",service.name="tracegen",span.kind="client",span_id="d5270e78d85f570f",status_code="STATUS_CODE_OK",trace_id="7d4854815225332c9834e6dbf85b9380" 1613767825689135000
spans end_time_unix_nano="2021-02-19 20:50:25.6895667 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="okey-dokey",net.peer.ip="1.2.3.4",parent_span_id="b57e98af78c3399b",peer.service="tracegen-client",service.name="tracegen",span.kind="server",span_id="a0643a156d7f9f7f",status_code="STATUS_CODE_OK",trace_id="fd6b8bb5965e726c94978c644962cdc8" 1613767825689388000
spans end_time_unix_nano="2021-02-19 20:50:25.6895667 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="lets-go",net.peer.ip="1.2.3.4",peer.service="tracegen-server",service.name="tracegen",span.kind="client",span_id="b57e98af78c3399b",status_code="STATUS_CODE_OK",trace_id="fd6b8bb5965e726c94978c644962cdc8" 1613767825689303300
spans end_time_unix_nano="2021-02-19 20:50:25.6896741 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="okey-dokey",net.peer.ip="1.2.3.4",parent_span_id="6a8e6a0edcc1c966",peer.service="tracegen-client",service.name="tracegen",span.kind="server",span_id="d68f7f3b41eb8075",status_code="STATUS_CODE_OK",trace_id="651dadde186b7834c52b13a28fc27bea" 1613767825689480300
```

<br>

### Insert

Insert a sample dataset using the Influx V2 API and line protocol to test the `router` API on port `8080`

#### Metrics
```bash
wget -qO- "https://raw.githubusercontent.com/metrico/influxdb_iox/main/test_fixtures/lineproto/metrics.lp" | \
curl -v "http://127.0.0.1:8181/api/v2/write?org=company&bucket=sensors" --data-binary @-
```

#### Logs
```bash
echo 'syslog,appname=myapp,facility=console,host=myhost,hostname=myhost,severity=warning facility_code=14i,message="warning message here",severity_code=4i,procid="12345",timestamp=1434055562000000000,version=1' | \
 curl -v "http://127.0.0.1:8181/api/v2/write?org=company&bucket=logs" --data-binary @-
```

#### Traces
```bash
echo 'spans end_time_unix_nano="2025-01-26 20:50:25.6893952 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="okey-dokey",net.peer.ip="1.2.3.4",parent_span_id="d5270e78d85f570f",peer.service="tracegen-client",service.name="tracegen",span.kind="server",span_id="4c28227be6a010e1",status_code="STATUS_CODE_OK",trace_id="7d4854815225332c9834e6dbf85b9380"' | \
 curl -v "http://127.0.0.1:8181/api/v2/write?org=company&bucket=traces" --data-binary @-
```

The expected response is `204`

### Query
Let's launch the `sql` client using the `querier` gRPC API on port `8082`

* Using Binary: `./influxdb3 --host http://localhost:8082 sql`
* Using Docker: `docker run -ti --rm ghcr.io/metrico/iox:latest --host http://iox:8082 sql`

#### [Flight SQL](https://github.com/influxdata/influxdb_iox/blob/main/docs/sql.md)

The first requirement is to choose a namespace _(or bucket)_ from the available ones:
```sql
influxdb3 show databases
+---------------+
| iox::database |
+---------------+
| logs          |
| mydb          |
| sensors       |
| traces        |
+---------------+

```

We can query our data using the influxdb CLI or the GRPC Flight API

#### Metric Search
```sql
influxdb3 query --database sensors "SELECT * FROM home WHERE temp > 95 LIMIT 4" 
+--------+------+-------------------------------+
| room   | temp | time                          |
+--------+------+-------------------------------+
| Garden | 99.0 | 2025-01-25T23:00:01.309084551 |
| Garden | 96.0 | 2025-01-25T16:10:11.745454525 |
| Garden | 98.0 | 2025-01-25T16:10:25.745146130 |
| Garden | 98.0 | 2025-01-25T16:10:37.742997495 |
+--------+------+-------------------------------++---------------------------------+----------------------+-------------+------------------+-------------------+--------------+-----------+------------+---------------+-------------+-------------------+-------------------+
```

#### Log Search
##### LIKE
```sql
influxdb3 query --database logs "SELECT * FROM syslog WHERE message LIKE '%here%'" 
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+
| appname | facility | facility_code | host   | hostname | message              | procid | severity | severity_code | time                          | timestamp      | version |
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+
| myapp   | console  | 14            | myhost | myhost   | warning message here | 12345  | warning  | 4             | 2025-01-25T23:57:02.459118350 | 1.434055562e18 | 1.0     |
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+

```
##### Regex
```sql
influxdb3 query --database logs "SELECT * FROM syslog WHERE message ~ '.+here'" 
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+
| appname | facility | facility_code | host   | hostname | message              | procid | severity | severity_code | time                          | timestamp      | version |
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+
| myapp   | console  | 14            | myhost | myhost   | warning message here | 12345  | warning  | 4             | 2025-01-25T23:57:02.459118350 | 1.434055562e18 | 1.0     |
+---------+----------+---------------+--------+----------+----------------------+--------+----------+---------------+-------------------------------+----------------+---------+
```

#### Trace Search
```sql
influxdb3 query --database traces "SELECT * FROM spans" 
+---------------------------------------+------------------------------+--------------------+------------+-------------+------------------+-----------------+--------------+-----------+------------------+----------------+-------------------------------+----------------------------------+
| end_time_unix_nano                    | instrumentation_library_name | kind               | name       | net.peer.ip | parent_span_id   | peer.service    | service.name | span.kind | span_id          | status_code    | time                          | trace_id                         |
+---------------------------------------+------------------------------+--------------------+------------+-------------+------------------+-----------------+--------------+-----------+------------------+----------------+-------------------------------+----------------------------------+
| 2025-01-26 20:50:25.6893952 +0000 UTC | tracegen                     | SPAN_KIND_INTERNAL | okey-dokey | 1.2.3.4     | d5270e78d85f570f | tracegen-client | tracegen     | server    | 4c28227be6a010e1 | STATUS_CODE_OK | 2025-01-26T00:01:02.450652384 | 7d4854815225332c9834e6dbf85b9380 |
| 2025-01-26 20:50:25.6893952 +0000 UTC | tracegen                     | SPAN_KIND_INTERNAL | okey-dokey | 1.2.3.4     | d5270e78d85f570f | tracegen-client | tracegen     | server    | 4c28227be6a010e1 | STATUS_CODE_OK | 2025-01-26T00:01:03.713172859 | 7d4854815225332c9834e6dbf85b9380 |
+---------------------------------------+------------------------------+--------------------+------------+-------------+------------------+-----------------+--------------+-----------+------------------+----------------+-------------------------------+----------------------------------+
```


</details>
                                                        
<details>
    <summary><h2>IOx Integrations</h2> Integrate your IOx Setup with Go, Rust, Python, etc</summary>  
  
> Official IOx FlightSQL clients:

  * [influxdb3-go](https://github.com/InfluxCommunity/influxdb3-go)
  * [influxdb3-js](https://github.com/InfluxCommunity/influxdb3-js)
  * [influxdb3-python](https://github.com/InfluxCommunity/influxdb3-python)
  * [influxdb3-java](https://github.com/InfluxCommunity/influxdb3-java)
  * [influxdb3-csharp](https://github.com/InfluxCommunity/influxdb3-csharp)

> Generic FlightSQL Drivers
  
  * [iox-community/python](https://github.com/metrico/iox-static-distro/tree/main/examples/python)
  * [iox-community/go](https://github.com/metrico/iox-static-distro/tree/main/examples/go)
  * [iox-community/rust](https://github.com/metrico/iox-static-distro/tree/main/examples/rust)
  * [influxdb-iox-client-go](https://github.com/influxdata/influxdb-iox-client-go)
  * [flightsql-dbapi-python](https://github.com/influxdata/flightsql-dbapi)
  * [influxdb_iox_client-rust](https://crates.io/crates/influxdb_iox_client)

  
<details>
    <summary><h3>Grafana</h3> Integrate your IOx Setup with Grafana</summary>  

  ### Grafana Client

Your service can be used with the [FlightSQL datasource](https://github.com/influxdata/grafana-flightsql-datasource) in Grafana:

![image](https://user-images.githubusercontent.com/1423657/239708678-8e947ae0-6710-4ae4-85c1-903f4c06b085.png)

Once ready, we can perform queries against our data using the FlightSQL query builder:

![image](https://user-images.githubusercontent.com/1423657/239708634-30b48942-d630-4feb-887d-5b6dc37f54d3.png)
                                                        
</details>

</details>

<details>
    <summary><h2>Legal Disclaimers</h2></summary>  
  
- All rights reserved by their respective owners. IOx and InfluxDB are a trademark of Influxdata.   
- This project is not connected or endorsed by Influxdata or the IOx project. Hopefully one day!
- Original, unstable, nightly. The IOx code is not modified in any way as part of the build process. 
  
</details>
