import os
from fastapi import FastAPI, Body, status, HTTPException, Response
from influx import connect, insert_log
from payload import Payload

ACCESS_TOKEN = os.environ.get("ACCESS_TOKEN")

app = FastAPI(
    app="Spectator Logging Service",
    version="0.0.1"
)

client = connect()

@app.post("/{path}", status_code=status.HTTP_200_OK)
def collect(path: str, response: Response, payload: Payload = Body(None)):
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

