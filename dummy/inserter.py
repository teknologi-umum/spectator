import os
import json
import influxdb_client
from dotenv import load_dotenv
from influxdb_client.client.write.point import Point
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

    # Check whether the bucket exists
    buckets_api = client.buckets_api()
    session_events_bucket = buckets_api.find_bucket_by_name("session_events")
    input_events_bucket = buckets_api.find_bucket_by_name("input_events")

    if session_events_bucket is None:
        buckets_api.create_bucket(bucket_name="session_events", org=influx_org)

    if input_events_bucket is None:
        buckets_api.create_bucket(bucket_name="input_events", org=influx_org)

    with client.write_api(write_options=ASYNCHRONOUS) as write_client:

        # Write user personal first
        with open("generated/user_personal.json") as f:
            users = json.load(f)
            for i in range(len(users)):
                user = users[i]

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

        with open("generated/events.json") as f:
            events = json.load(f)
            for i in range(len(events)):
                event = events[i]

                point = Point(event["type"]) \
                    .tag("session_id", event["session_id"]) \
                    .tag("question_number", event["question_number"])

                if event["type"] == "coding_event_keystroke":
                    point = point.field("key_char", event["key_char"]) \
                        .field("key_code", event["key_code"]) \
                        .field("shift", event["shift"]) \
                        .field("alt", event["alt"]) \
                        .field("control", event["control"]) \
                        .field("meta", event["meta"]) \
                        .field("unrelated_key", event["unrelated_key"]) \
                        .time(event["time"])
                elif event["type"] == "coding_event_mousemove":
                    point = point.field("direction", event["direction"]) \
                        .field("x_position", event["x_position"]) \
                        .field("y_position", event["y_position"]) \
                        .field("window_width", event["window_width"]) \
                        .field("window_height", event["window_height"]) \
                        .time(event["time"])
                elif event["type"] == "coding_event_mouseclick":
                    point = point.field("right_click", event["right_click"]) \
                        .field("left_click", event["left_click"]) \
                        .field("middle_click", event["middle_click"]) \
                        .time(event["time"])

                write_client.write(
                    bucket="input_events",
                    org=influx_org,
                    record=point,
                )

if __name__ == "__main__":
    main()
