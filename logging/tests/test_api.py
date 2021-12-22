import pytest
import os
from ..src import app
from fastapi.testclient import TestClient

client = TestClient(app)

def test_not_allowed():
    response = client.get("/")
    assert response.status_code == 405

