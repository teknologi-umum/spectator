"""
This file contains the code for inserting the generated data
into the InfluxDB. Before running this file, we expect you
to have existing JSON files resides in the 'generated' directory.

To run this file, you will need to set .env file with the following
variables:

    INFLUX_URL
    INFLUX_ORG
    INFLUX_TOKEN

That is the easiest way to put the environment variable.
After that, you can just chill and enjoy.
"""

import os
import json
import influxdb_client
from dotenv import load_dotenv
from influxdb_client.client.write.point import Point
from influxdb_client.client.write_api import ASYNCHRONOUS
from influxdb_client.domain.write_precision import WritePrecision

load_dotenv(".env")


def main():
    print("Let's verify your environment variables.")
    influx_url = os.environ.get("INFLUX_URL")
    influx_org = os.environ.get("INFLUX_ORG")
    influx_token = os.environ.get("INFLUX_TOKEN")

    if influx_url is None:
        print("INFLUX_URL is not set")
        return

    if influx_org is None:
        print("INFLUX_ORG is not set")
        return

    if influx_token is None:
        print("INFLUX_TOKEN is not set")
        return

    print("Environment variables are set. Let's connect to InfluxDB.")

    client = influxdb_client.InfluxDBClient(url=influx_url, token=influx_token,
                                            org=influx_org)

    # Check whether the bucket exists
    buckets_api = client.buckets_api()
    session_events_bucket = buckets_api.find_bucket_by_name("session_events")
    input_events_bucket = buckets_api.find_bucket_by_name("input_events")

    if session_events_bucket is None:
        print("The bucket 'session_events' does not exist. Creating it now.")
        buckets_api.create_bucket(bucket_name="session_events", org=influx_org)

    if input_events_bucket is None:
        print("The bucket 'input_events' does not exist. Creating it now.")
        buckets_api.create_bucket(bucket_name="input_events", org=influx_org)

    with client.write_api(write_options=ASYNCHRONOUS) as write_client:

        # Write user personal first
        with open("generated/user_personal.json") as f:
            print("Reading user personal data from file.")
            users = json.load(f)
            print(f"Found {len(users)} users. Writing into InfluxDB")
            for user in users:
                point = Point(user["type"]) \
                    .tag("session_id", user["session_id"]) \
                    .field("student_number", user["student_number"]) \
                    .field("hours_of_practice", user["hours_of_practice"]) \
                    .field("years_of_experience", user["years_of_experience"]) \
                    .field("familiar_language", user["familiar_language"])
                write_client.write(
                    bucket="session_events",
                    org=influx_org,
                    record=point,
                )

        with open("generated/input_events.json") as f:
            print("Reading events data from file.")
            events = json.load(f)
            print( f"Found {len(events)} input_events. Writing into InfluxDB" )
            for event in events:
                point = Point(event["type"]) \
                    .tag("session_id", event["session_id"]) \
                    .tag("question_number", event["question_number"])\
                    .time(event["time"], write_precision=WritePrecision.S)
                fields = set(event.keys())-set(["session_id","type","time","question_number"])
                for field in fields:
                    point= point.field(field,event[field])
                write_client.write(
                    bucket="input_events",
                    org=influx_org,
                    record=point,
                )

        with open("generated/session_events.json") as f:
            print("Reading events data from file.")
            events = json.load(f)
            print( f"Found {len(events)} session_events. Writing into InfluxDB" )
            for event in events:
                point = Point(event["type"]) \
                    .tag("session_id", event["session_id"]) \
                    .time(event["time"], write_precision=WritePrecision.S)
                fields = set(event.keys())-set(["session_id","type","time"])
                for field in fields:
                    if type(event[field]) == type([]):
                        point= point.field(field," ".join([str(i) for i in event[field]]))     
                    else:
                        point= point.field(field,event[field])

                write_client.write(
                    bucket="session_events",
                    org=influx_org,
                    record=point,
                )

            print("Done. Please don't do anything until the script exits itself.")

if __name__ == "__main__":
    main()
