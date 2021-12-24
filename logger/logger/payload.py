from datetime import datetime
from typing import Optional
from pydantic import BaseModel

class Data(BaseModel):
    request_id: str
    application: str
    environment: Optional[str] = "unset"
    language: Optional[str] = None
    level: Optional[str] = "debug"
    message: str
    body: Optional[dict] = None
    timestamp: Optional[datetime] = datetime.utcnow()

class Payload(BaseModel):
    access_token: str
    data: Data
