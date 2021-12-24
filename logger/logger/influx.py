import os
import json
from influxdb_client import InfluxDBClient, Point
from influxdb_client.client.write_api import SYNCHRONOUS
from .payload import Data
from dotenv import load_dotenv

load_dotenv(dotenv_path=".env")

INFLUX_URL = os.environ.get("INFLUX_URL")
INFLUX_TOKEN = os.environ.get("INFLUX_TOKEN")
INFLUX_ORG = os.environ.get("INFLUX_ORG")

def connect() -> InfluxDBClient:
    if INFLUX_URL == "":
        raise Exception("INFLUX_URL was not set")
    if INFLUX_TOKEN == "":
        raise Exception("INFLUX_TOKEN was not set")
    if INFLUX_ORG == "":
        raise Exception("INFLUX_ORG was not set")
    return InfluxDBClient(url=INFLUX_URL,
                            token=INFLUX_TOKEN,
                            org=INFLUX_ORG)

def insert_log(client: InfluxDBClient, data: Data) -> bool:
    try:
        print("writing into influx")
        write_api = client.write_api(write_options=SYNCHRONOUS)
        p = Point("log") \
            .tag("request_id", data.request_id) \
            .tag("level", data.level) \
            .tag("application", data.application) \
            .tag("environment", data.environment) \
            .field("language", data.language) \
            .field("body", json.dumps(data.body)) \
            .field("message", data.message) \
            .time(data.timestamp)
        write_api.write(bucket="log", record=p)
        print("written into influx")
        return True
    except:
        return False

def check_bucket(client: InfluxDBClient) -> None:
    buckets_api = client.buckets_api()
    log_bucket = buckets_api.find_bucket_by_name("log")
    if log_bucket != None:
        return None
    else:
        buckets_api.create_bucket(
            bucket=None,
            bucket_name="log",
            org_id=INFLUX_ORG,
            retention_rules=None,
            description="Logs"
        )
        return None
