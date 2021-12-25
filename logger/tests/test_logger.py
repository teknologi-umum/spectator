import os
from fastapi.testclient import TestClient
from logger.app import app

client = TestClient(app)

def test_health_check():
    os.environ["ACCESS_TOKEN"] = "testing"
    response = client.get("/")

    assert response.status_code == 200
    assert response.json() == {"db": "ok"}

def test_insert_log():
    os.environ["ACCESS_TOKEN"] = "testing"
    response = client.post("/", json={
        "access_token": "testing",
        "data": {
            "request_id": "1",
            "application": "worker",
            "message": "hello world",
            "body": {}
        }
    })

    assert response.status_code == 200

def test_race(start_race):
    start_race(threads_num=2, target=test_insert_log)
    start_race(threads_num=2, target=test_health_check)
