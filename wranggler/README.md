## Wranggler

This job is to aggregate data to response backend request and Storing file csv and JSON file from influxDB.

## How to test

```
INFLUX_TOKEN="" \
INFLUX_HOST="http://localhost:8086" \
INFLUX_ORG="teknum" MINIO_HOST="" \ 
MINIO_ACCESS_ID="" \ 
MINIO_SECRET_KEY="" \ 
LOGGER_SERVER_ADDRESS="" \
LOGGER_TOKEN="" \
go rum main.go file.go funfact.go ping.go
```


