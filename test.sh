#!/bin/bash

curl "http://127.0.0.1:8080/health"

wget -qO- "https://github.com/influxdata/influxdb_iox/raw/main/test_fixtures/lineproto/metrics.lp" | \
curl -v "http://127.0.0.1:8080/api/v2/write?org=company&bucket=sensors" --data-binary @-

if [ ! -f "influxdb_iox" ]; then
    curl -fsSL github.com/metrico/iox-builder/releases/latest/download/influxdb_iox -O && chmod +x influxdb_iox
fi

./influxdb_iox --host http://127.0.0.1:8082 query company_sensors 'select count(*), cpu as cpu_num from cpu group by cpu'
# ./influxdb_iox --host http://127.0.0.1:8082 debug schema get company_sensors
