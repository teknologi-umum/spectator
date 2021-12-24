"""
This file handles anything to do with the InfluxDB
"""
import os
import json
from dotenv import load_dotenv
from influxdb_client import InfluxDBClient, Point
from influxdb_client.client.write_api import SYNCHRONOUS
from logger.payload import Data

load_dotenv(dotenv_path=".env")

INFLUX_URL = os.environ.get("INFLUX_URL")
INFLUX_TOKEN = os.environ.get("INFLUX_TOKEN")
INFLUX_ORG = os.environ.get("INFLUX_ORG")

def connect() -> InfluxDBClient:
    """
    simply returns an InfluxDB client instance
    """
    if INFLUX_URL == "":
        raise Exception("INFLUX_URL was not set")
    if INFLUX_TOKEN == "":
        raise Exception("INFLUX_TOKEN was not set")
    if INFLUX_ORG == "":
        raise Exception("INFLUX_ORG was not set")
    return InfluxDBClient(url=INFLUX_URL,
                            debug=None,
                            token=INFLUX_TOKEN,
                            org=INFLUX_ORG)

def insert_log(client: InfluxDBClient, data: Data) -> bool:
    """
    insert a log into the InfluxDB, synchronously.
    no fancy stuff.
    """
    print("writing into influx")
    write_api = client.write_api(write_options=SYNCHRONOUS)
    point = Point("log") \
        .tag("request_id", data.request_id) \
        .tag("level", data.level) \
        .tag("application", data.application) \
        .tag("environment", data.environment) \
        .field("language", data.language) \
        .field("body", json.dumps(data.body)) \
        .field("message", data.message) \
        .time(data.timestamp)
    write_api.write(bucket="log", record=point)
    print("written into influx")
    return True

def check_bucket(client: InfluxDBClient) -> None:
    """
    Checks whether the InfluxDB has a bucket named log
    """
    buckets_api = client.buckets_api()
    log_bucket = buckets_api.find_bucket_by_name("log")
    if log_bucket is not None:
        return None

    buckets_api.create_bucket(
        bucket=None,
        bucket_name="log",
        org_id=INFLUX_ORG,
        retention_rules=None,
        description="Logs"
    )
    return None

def check_connection(client: InfluxDBClient):
    """Check that the InfluxDB is running."""
    res = client.api_client.call_api("/ping", "GET")
    return res
