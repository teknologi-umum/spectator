import os
from fastapi import FastAPI, Body, status, HTTPException, Response
from influx import connect, insert_log
from payload import Payload
from dotenv import load_dotenv, dotenv_values

load_dotenv(dotenv_path=".env")

ACCESS_TOKEN = os.environ.get("ACCESS_TOKEN")

app = FastAPI(
    app="Spectator Logging Service",
    version="0.0.1"
)

client = connect()

@app.post("/", status_code=status.HTTP_200_OK)
def collect(response: Response, payload: Payload = Body(None)):
    """
    This is the only endpoint for the logging service.
    The client should send a JSON body with the schema
    as defined in the Payload class.
    """
    token = payload.access_token
    if (not ACCESS_TOKEN) or token == ACCESS_TOKEN:
        data = payload.data
        succeed = insert_log(client, data)
        if succeed:
            response.status_code = status.HTTP_200_OK
            return { status: "ok" }
        else:
            response.status_code = status.HTTP_500_INTERNAL_SERVER_ERROR
            return { status: "failed" }
    else:
        raise HTTPException(status.HTTP_403_FORBIDDEN, detail="access_token is required")

