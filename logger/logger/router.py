import os
from fastapi import APIRouter, Response, status, HTTPException, Body
from fastapi.param_functions import Depends
from influxdb_client.client.influxdb_client import InfluxDBClient
from logger.payload import Payload
from logger.influx import check_connection, insert_log

router = APIRouter()

ACCESS_TOKEN = os.environ.get("ACCESS_TOKEN")


@router.get("/")
def ping(response: Response, client: InfluxDBClient = Depends(InfluxDBClient)):
    """
    For healthcheck purposes
    """
    if check_connection(client):
        response.status_code = status.HTTP_200_OK
        return { "db": "ok" }

    response.status_code = status.HTTP_500_INTERNAL_SERVER_ERROR
    return { "db": "failed" }

@router.post("/")
def collect(response: Response, client: InfluxDBClient = Depends(InfluxDBClient), payload: Payload = Body(None)):
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
            return { "status": "ok" }

        response.status_code = status.HTTP_500_INTERNAL_SERVER_ERROR
        return { "status": "failed" }

    raise HTTPException(status.HTTP_403_FORBIDDEN, detail="access_token is required")
