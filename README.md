<img width="718" height="415" alt="image" src="https://github.com/user-attachments/assets/db372f4c-c33a-47b1-a233-cab6ff86ff03" />

# RefluxDB
Unlocked Community Builds for InfluxDB 3.x IOx _(eye-ox)_ aka _InfluxDB3-Core_

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

## Configuration
Your service can be configured by following the [InfluxDB3-Core Instructions](https://docs.influxdata.com/influxdb3/core/get-started/setup/)
  

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
                                                        
<details open=true>
    <summary><h2>IOx Integrations</h2> Integrate your IOx Setup with Go, Rust, Python, etc</summary>  

  ### FlightSQL

Your service can be used with any FlightSQL client or InfluxDB3 Flight client. 
  
  ### InfluxDB3 Explorer

Your service can be used with the [InfluxDB3 Explorer](https://docs.influxdata.com/influxdb3/explorer/)

<img width="1317" height="952" alt="InfluxDB-3-Explorer-08-01-2025_09_33_PM" src="https://github.com/user-attachments/assets/c015cedf-677e-4493-9721-2e99eeb2d359" />


  ### Grafana Client

Your service can be used with the [FlightSQL datasource](https://github.com/influxdata/grafana-flightsql-datasource) in Grafana:

![image](https://user-images.githubusercontent.com/1423657/239708678-8e947ae0-6710-4ae4-85c1-903f4c06b085.png)

Once ready, we can perform queries against our data using the FlightSQL query builder:

![image](https://user-images.githubusercontent.com/1423657/239708634-30b48942-d630-4feb-887d-5b6dc37f54d3.png)
                                                        

</details>

<details>
    <summary><h2>Legal Disclaimers</h2></summary>  
  
- All rights reserved by their respective owners. IOx and InfluxDB are a trademark of Influxdata.   
- This project is not connected or endorsed by Influxdata or the IOx project.
- Changes released under the same License terms as the original software. 
  
</details>
