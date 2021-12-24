from logger import __version__, app
from fastapi.testclient import TestClient

client = TestClient(app)

def test_not_allowed():
    response = client.get("/")
    assert response.status_code == 405

def test_version():
    assert __version__ == '0.1.0'
