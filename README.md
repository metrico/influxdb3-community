<img src="https://pbs.twimg.com/profile_banners/1967601206/1682514855/1500x500" width=800>

# IOx
Pronounced _(eye-ox)_ short for iron oxide. The new core of InfluxDB written in Rust on top of Apache Arrow.

### Motivation
😄 You want to try and experiment with IOx low-cost storage, unlimited cardinality and flight sql<br>
😮‍ The IOx project is in active development, which is why InfluxDB is not producing builds yet...<br>
😄 No problem! Meet the _unofficial_ InfluxDB 3.0 _"IOx"_ musl + docker builder for early adopters

##### Releases
###### amd64/musl
  - [x] [docker](https://github.com/metrico/iox-builder/pkgs/container/iox-musl): `docker pull ghcr.io/metrico/iox-musl:latest`
  - [x] [binary](https://github.com/metrico/iox-builder/releases): `github.com/metrico/iox-builder/releases/latest/download/influxdb_iox`

<br>

## Get Started

Pull the latest IOx static builds or use docker images. This demo will work with either.

#### Static
```bash
curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb_iox -O \
&& chmod +x influxdb_iox
```

#### Docker
```
docker-compose up -d
```

<br>
  
<details>
    <summary><h2>IOx Settings</h2> Deploy IOx using different settings</summary>  
  
This demo will launch IOx `router`, `querier`, `ingester` and `compactor` on the same host using local storage:
  
```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Router    │    │  Ingester   │    │   Querier   │    │   Compactor │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
       │                  │                  │                  │       
       │                  │                  │                  │       
       └──────────────────┼──────────────────┼──────────────────┘             
                          │                          
                          │                          
                       .──▼──.                       
                      (       )   Shared sqlite      
                      │`─────'│   file database      
                      │       │  /tmp/db.sqlite      
                      │.─────.│                      
                      (       )                      
                       `─────'                       
```
  
![image](https://github.com/metrico/iox-static-distro/assets/1423657/55175f98-6b0a-4097-8a34-06ab6c4fd8fe)
  

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

Each server needs an identifier for writing to object storage and as an identifier that is added to replicated writes, Write Buffer segments and Chunks. Must be unique in a group of connected or semi-connected IOx servers. Must be a number that can be represented by a 32-bit unsigned integer.

```
      - INFLUXDB_IOX_ID=1
```


To enable persisten catalog using postgres, use the following:

```
      - INFLUXDB_IOX_CATALOG_DSN=postgres://postgres@localhost:5432/postgres
```

To enable S3/R2/Minio object storage use the following parameters:

```
      - INFLUXDB_IOX_OBJECT_STORE=s3
      - AWS_ACCESS_KEY_ID=access_key_value
      - AWS_SECRET_ACCESS_KEY=secret_access_key_value
      - AWS_DEFAULT_REGION=us-east-2
      - INFLUXDB_IOX_BUCKET=bucket-name
      - AWS_ENDPOINT = http://minio:9000
```

For other storage options refer to [env example](https://github.com/metrico/iox-builder/blob/main/env.example)

</details>
  

<details>
    <summary><h2>IOX Insert & Query</h2> Validate your IOx Setup</summary>

<br>

Let's start testing and using your brand new IOx instance!  
  
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

</details>
                                                        
<details>
    <summary><h2>IOx Integrations</h2> Integrate your IOx Setup with Go, Rust, Python, etc</summary>  
  
> IOx Flight SQL [generic client examples](https://github.com/metrico/iox-static-distro/tree/main/examples)
  
* <img src="https://github.com/metrico/iox-community/assets/1423657/37838c84-6504-4a69-bff2-a30bf6f20a5e" width=50 /> [python](https://github.com/metrico/iox-static-distro/tree/main/examples/python)
* <img src="https://github.com/metrico/iox-community/assets/1423657/bd8f21e3-cea8-4423-be08-08e4eee4ff7f" width=50 /> [go](https://github.com/metrico/iox-static-distro/tree/main/examples/go)
* <img src="https://github.com/metrico/iox-community/assets/1423657/e7a4d53f-9e8a-45a1-b72b-fd427c7cf9ef" width=50 /> [rust](https://github.com/metrico/iox-static-distro/tree/main/examples/rust)
  
> IOx Flight SQL Official clients:
  
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
