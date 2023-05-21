#!/bin/bash

curl "http://127.0.0.1:8080/health"

wget -qO- "https://github.com/influxdata/influxdb_iox/raw/main/test_fixtures/lineproto/metrics.lp" | \
curl -v "http://127.0.0.1:8080/api/v2/write?org=company&bucket=sensors" --data-binary @-


