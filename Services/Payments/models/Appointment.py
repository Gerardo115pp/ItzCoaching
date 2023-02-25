from dataclasses import dataclass, asdict
from datetime import datetime, timedelta
from enum import Enum
from typing import Any
import json

def allArgsPresent(obj: object, args: list[str]) -> bool:
    """ 
        this method should be reassigned from the __init__.py file
    """
    raise NotImplementedError("this shouldnt happen")


class AppointmentStatus(str, Enum):
    UNPAID = "unpaid"
    PAID = "paid"
    FINALIZED = "finalized"
    CANCELLED = "cancelled"

@dataclass
class Appointment:
    id: int
    expert: int
    customer_email: str
    utc_start: datetime
    duration: timedelta
    status: AppointmentStatus = AppointmentStatus.UNPAID
    created_at: datetime = datetime.utcnow()
    
    # def __post_init__(self) -> None:
    
    def toJson(self) -> dict:
        return json.dumps(asdict(self), cls=AppointmentENCODER)

@dataclass
class AppointmentConfirmation:
    appointment: int
    expert_confirmed: bool
    consumer_confirmed: bool
    
class AppointmentENCODER(json.JSONEncoder):
    
    def default(self, o: Any) -> Any:
        if isinstance(o, timedelta):
            return o.total_seconds()
        elif isinstance(o, datetime):
            return o.strftime("%Y-%m-%d %H:%M:%S")
        return super().default(o)