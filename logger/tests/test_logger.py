import logging
import os
from fastapi.testclient import TestClient
from logger.app import app

os.environ["ACCESS_TOKEN"] = "testing"

if os.environ.get("INFLUX_URL") == "":
    logging.error("INFLUX_URL must be set outside")

if os.environ.get("INFLUX_TOKEN") == "":
    logging.error("INFLUX_TOKEN must be set outside")

if os.environ.get("INFLUX_ORG") == "":
    logging.error("INFLUX_ORG must be set outside")

client = TestClient(app)

def test_health_check():
    response = client.get("/")

    assert response.status_code == 200
    assert response.json() == {"db": "ok"}

def test_insert_log():
    response = client.post("/", json={
	"access_token": "testing",
	"data": {
		"request_id": "1",
		"application": "worker",
		"message": "hello world",
		"body": {}
	}})

    print(response.json())
    assert response.status_code == 200

