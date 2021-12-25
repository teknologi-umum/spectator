"""
This file contains the payload body schema
"""
from datetime import datetime
from typing import Optional
from pydantic import BaseModel

class Data(BaseModel):
    """
    Data contains the payload body that is sent
    to this logger worker
    """
    request_id: str
    application: str
    environment: Optional[str] = "unset"
    language: Optional[str] = None
    level: Optional[str] = "debug"
    message: str
    body: Optional[dict] = None
    timestamp: Optional[datetime] = datetime.utcnow()

class Payload(BaseModel):
    """
    Payload is the request body payload schema
    """
    access_token: str
    data: Data
