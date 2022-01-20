## Worker

This job is to aggregate data to response backend request and Storing file csv and JSON file from influxDB.

### What is this directory

| directory | description |
|--|--|
| file | contain function to procced aggregation data inside database into file by session_id|
| funfact | contain function to aggregate and return fun fact like word each minute, delete rate and submission |
| logger | contain code to trace error and send it to logger server|
| logger_porto | logger protocol definition to make request to logger server |
| worker_proto | worker protocol definiiton for services |

## Tool used

- [docker-compose, this for make all setup easier](https://docs.docker.com/compose/)
- [GRPC, this protocol used to service interact with other services](https://grpc.io/docs/languages/go/quickstart/)
- [influxdb, time series database](https://github.com/influxdata/influxdb-client-go)

## How to run

### The server

```
INFLUX_TOKEN="" \
INFLUX_HOST="http://localhost:8086" \
INFLUX_ORG="teknum" MINIO_HOST="" \ 
MINIO_ACCESS_ID="" \ 
MINIO_SECRET_KEY="" \ 
LOGGER_SERVER_ADDRESS="" \
LOGGER_TOKEN="" \
ENVIRONMENT="" \
go run .
```

### The test

```
INFLUX_TOKEN="" \
INFLUX_HOST="http://localhost:8086" \
INFLUX_ORG="teknum" MINIO_HOST="" \ 
MINIO_ACCESS_ID="" \ 
MINIO_SECRET_KEY="" \ 
LOGGER_SERVER_ADDRESS="" \
LOGGER_TOKEN="" \
ENVIRONMENT="" \
go test ./... -v
```

