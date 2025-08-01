<img width="300" alt="influx-unlocked" src="https://github.com/user-attachments/assets/cabf6239-2d95-4531-8bf0-054b6a97d5fe" />

# InfluxDB 3.x "IOx" Community Builds
Unlocked Community Builds and Containers for InfluxDB 3.x IOx _(eye-ox)_ aka _InfluxDB3 Core_

### Motivation
😄 Restore community access to low-cost storage, unlimited cardinality and flight sql<br>
🥵‍ Unlocking _"Open Core"_ limitations only designed to promote Cloud/Enterprise<br>

This builder is based on the unlocked "Core" fork: https://github.com/metrico/influxdb3-unlocked

#### Nightly Builds
###### amd64
  - [x] [docker](https://github.com/metrico/influxdb3-community/pkgs/container/influxdb3-unlocked): `docker pull ghcr.io/metrico/influxdb3-unlocked:latest`
  - [x] [binary](https://github.com/metrico/influxdb3-community/releases): `github.com/metrico/influxdb3-community/releases/latest/download/influxdb3`

<br>

## Get Started

This guide uses Docker and docker-compose. You can run locally using a [static build](https://github.com/metrico/influxdb3-community/releases).

#### Static
```bash
curl -fsSL github.com/metrico/influxdb3-community/releases/latest/download/influxdb3 -O \
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

```bash
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
```

##### Insert
```bash
wget -qO- "https://raw.githubusercontent.com/metrico/influxdb_iox/main/test_fixtures/lineproto/metrics.lp" | \
curl -v "http://127.0.0.1:8181/api/v2/write?org=company&bucket=sensors" --data-binary @-
```

#### Logs
```
syslog,appname=myapp,facility=console,host=myhost,hostname=myhost,severity=warning facility_code=14i,message="warning message here",severity_code=4i,procid="12345",timestamp=1613769568895331700,version=1
```

##### Insert
```bash
echo 'syslog,appname=myapp,facility=console,host=myhost,hostname=myhost,severity=warning facility_code=14i,message="warning message here",severity_code=4i,procid="12345",timestamp=1434055562000000000,version=1' | \
 curl -v "http://127.0.0.1:8181/api/v2/write?org=company&bucket=logs" --data-binary @-
```

#### Traces
```
spans end_time_unix_nano="2021-02-19 20:50:25.6893952 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="okey-dokey",net.peer.ip="1.2.3.4",parent_span_id="d5270e78d85f570f",peer.service="tracegen-client",service.name="tracegen",span.kind="server",span_id="4c28227be6a010e1",status_code="STATUS_CODE_OK",trace_id="7d4854815225332c9834e6dbf85b9380" 1613767825689169000
spans end_time_unix_nano="2021-02-19 20:50:25.6893952 +0000 UTC",instrumentation_library_name="tracegen",kind="SPAN_KIND_INTERNAL",name="lets-go",net.peer.ip="1.2.3.4",peer.service="tracegen-server",service.name="tracegen",span.kind="client",span_id="d5270e78d85f570f",status_code="STATUS_CODE_OK",trace_id="7d4854815225332c9834e6dbf85b9380" 1613767825689135000
```

##### Insert
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
- This project is not connected or endorsed by Influxdata or the IOx project.
- Changes released under the same License terms as the original software. 
  
</details>
