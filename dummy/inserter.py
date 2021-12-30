import os
import json
import influxdb_client
from dotenv import load_dotenv
from influxdb_client.client.write_api import ASYNCHRONOUS

load_dotenv(".env")


def main():
    print("Let's verify your environment variables.")
    influx_url = os.environ.get("INFLUX_URL")
    influx_org = os.environ.get("INFLUX_ORG")
    influx_token = os.environ.get("INFLUX_TOKEN")
    influx_bucket = os.environ.get("INFLUX_BUCKET")

    if influx_url is None:
        print("INFLUX_URL is not set")
        return

    if influx_org is None:
        print("INFLUX_ORG is not set")
        return

    if influx_token is None:
        print("INFLUX_TOKEN is not set")
        return

    if influx_bucket is None:
        print("INFLUX_BUCKET is not set")
        return

    print("Environment variables are set. Let's connect to InfluxDB.")

    client = influxdb_client.InfluxDBClient(url=influx_url, token=influx_token, org=influx_org)
    write_api = client.write_api(write_options=ASYNCHRONOUS)

    # Write user personal first
    with open("generated/user_personal.json") as f:
        users = json.load(f)
        # TODO: do a for loop to write the personal user

    with open("generated/events.json") as f:
        events = json.load(f)
        # TODO: do a for loop to write the events
        # might as well use batch write
        # see this: https://github.com/influxdata/influxdb-client-python#batching
        # the idiomatic way maybe is by calling "with client.write_api() as write_api"


if __name__ == "__main__":
    main()
