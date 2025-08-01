<img width="300" alt="image" src="https://github.com/user-attachments/assets/44d2c1b0-4ec8-4d39-a75c-384e9e0a3f89" />

Unlocked Community Builds for InfluxDB 3.x IOx _(eye-ox)_ aka _InfluxDB3-Core_

### Motivation
😄 Provide community access to low-cost storage, unlimited cardinality and flightsql<br>
🥵‍ Unlocking _"Open Core"_ limitations only designed to promote Cloud/Enterprise<br>

🔥 This builder is based on the unlocked "InfluxDB3-Core" fork: https://github.com/refluxdb/refluxdb<br>

#### Nightly Builds
- [x] [docker](https://github.com/refluxdb/influxdb3-community/pkgs/container/influxdb3-unlocked): `docker pull ghcr.io/refluxdb/refluxdb:latest`
- [x] [binary](https://github.com/refluxdb/influxdb3-community/releases): `github.com/refluxdb/influxdb3-community/releases/latest/download/influxdb3`

<br>

## Get Started

👉 For a full liste of unlocked features refer to the [RefluxDB](https://github.com/refluxdb/refluxdb) README

This guide uses Docker and docker-compose. You can run locally using a [static build](https://github.com/refluxdb/influxdb3-community/releases).

#### Static
```bash
curl -fsSL github.com/refluxdb/influxdb3-community/releases/latest/download/influxdb3 -O \
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


<br>

## Service Configuration

Your service can be configured by following the [InfluxDB3-Core Instructions](https://docs.influxdata.com/influxdb3/core/get-started/setup/)

                                                        
## IOx Integrations

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
                                                        

## Legal Disclaimers
  
- All rights reserved by their respective owners. IOx and InfluxDB are a trademark of Influxdata.   
- This project is not connected or endorsed by Influxdata or the InfluxDB project.
- Changes in the fork are released under the same License terms as the original software. 
  
