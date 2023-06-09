# IOx 
### FlightSQL c++
Example gRPC Flight SQL API client using c++

#### Requirements
```
apt update \
    && apt install -y -V ca-certificates lsb-release wget g++ \
    && apt install -y -V ca-certificates lsb-release wget \
    && wget https://apache.jfrog.io/artifactory/arrow/$(lsb_release --id --short | tr 'A-Z' 'a-z')/apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb \
    && apt install -y -V ./apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb \
    && apt update \
    && apt install -y -V libarrow-dev \
    && apt install -y -V libarrow-glib-dev \
    && apt install -y -V libarrow-dataset-dev \
    && apt install -y -V libarrow-dataset-glib-dev \
    && apt install -y -V libarrow-flight-dev \
    && apt install -y -V libarrow-flight-glib-dev \
    && apt-get install -y -V  libgflags-dev \
    && apt-get install -y -V libarrow-flight-sql-dev  \
    && apt-get install -y -V libgrpc++-dev
```
#### Build
```
g++ -o flightc main.cc -larrow_flight -larrow -lgflags  -larrow_flight_sql -std=c++17
```

#### Usage

```
# Set environment variables
export INFLUX_DATABASE=my-v3-bucket && \
export INFLUX_HOST=us-east-1-1.aws.cloud2.influxdata.com && \
export INFLUX_TOKEN=WIIiwererffkdfoiwe==
```
```
./flightc
```
