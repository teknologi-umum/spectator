import os
import json
from influxdb_client import InfluxDBClient, Point
from influxdb_client.client.write_api import SYNCHRONOUS
from payload import Data


INFLUX_URL = os.environ.get("INFLUX_URL")
INFLUX_TOKEN = os.environ.get("INFLUX_TOKEN")
INFLUX_ORG = os.environ.get("INFLUX_ORG")

def connect() -> InfluxDBClient :
    if INFLUX_URL == "":
        raise Exception("INFLUX_URL was not set")
    if INFLUX_TOKEN == "":
        raise Exception("INFLUX_TOKEN was not set")
    if INFLUX_ORG == "":
        raise Exception("INFLUX_ORG was not set")
    return InfluxDBClient(url=INFLUX_URL, token=INFLUX_TOKEN,
                          org=INFLUX_ORG)

def insert_log(client: InfluxDBClient, data: Data) -> bool:
    try:
        write_api = client.write_api(write_options=SYNCHRONOUS)
        p = Point("log").tag("platform", data.platform).tag("level",
                data.level).tag("language", data.language).tag("environment"
                , data.environment).field("body",
                json.dumps(data.body)).field("custom",
                json.dumps(data.custom)).field("notifier",
                json.dumps(data.notifier)).field("server",
                json.dumps(data.server))
        write_api.write(bucket="log", record=p)
        return True
    except:
        return False
