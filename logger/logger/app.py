"""
Main entrypoint for the logger worker for Spectator project
"""
import os
from fastapi import FastAPI
from dotenv import load_dotenv
from logger.influx import check_bucket, connect
from logger.router import router

def get_application() -> FastAPI:
    load_dotenv(dotenv_path=".env")

    ACCESS_TOKEN = os.environ.get("ACCESS_TOKEN")
    # Validate the existance of the ACCESS_TOKEN
    if ACCESS_TOKEN in ("", None):
        raise EnvironmentError("ACCESS_TOKEN was not set")

    # Initialize the database connection instance
    client = connect()
    check_bucket(client)

    app = FastAPI(
        app="Spectator Logging Service",
        version="0.0.1"
    )

    app.include_router(router, dependencies=[client])

    return app

app = get_application()
