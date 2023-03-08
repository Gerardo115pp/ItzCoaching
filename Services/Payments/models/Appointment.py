from dataclasses import dataclass, asdict
from datetime import datetime, timedelta, timezone
from .TimeSlot import TimeSlot
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
    
    @classmethod
    def fromJson(cls, json_string: str) -> 'Appointment':
        return cls(**json.loads(json_string, cls=AppointmentDECODER))
    
    @classmethod
    def fromDict(cls, appointment_dict: dict) -> 'Appointment':
        appointment_dict["utc_start"] = datetime.strptime(appointment_dict["utc_start"], "%Y-%m-%d %H:%M:%S")
        appointment_dict["duration"] = timedelta(milliseconds=float(appointment_dict["duration"]))
        appointment_dict["status"] = AppointmentStatus(appointment_dict["status"])
        appointment_dict["expert"] = int(appointment_dict["expert"])
        appointment_dict["created_at"] = datetime.strptime(appointment_dict["created_at"], "%Y-%m-%d %H:%M:%S")
        appointment_dict["id"] = appointment_dict.get("id", None)
        
        return cls(**appointment_dict)
    
    # def __post_init__(self) -> None:
    @property
    def TimeSlot(self) -> TimeSlot:
        end= self.utc_start + self.duration
        return TimeSlot(self.utc_start.strftime("%Y-%m-%d %H:%M:%S"), end.strftime("%Y-%m-%d %H:%M:%S"))
    
    def overlaps(self, other: 'Appointment') -> bool:
        # print(f" self {type(self.TimeSlot).__name__} other {type(other.TimeSlot).__name__}")
        return self.TimeSlot.overlaps(other.TimeSlot)
    
    def toJson(self) -> dict:
        return json.dumps(asdict(self), cls=AppointmentENCODER)
    
    def toDict(self) -> dict:
        self_dict = {
            "id": self.id,
            "expert": self.expert,
            "customer_email": self.customer_email,
            "utc_start": self.utc_start.strftime("%Y-%m-%d %H:%M:%S"),
            "duration": self.duration.total_seconds() * 1000,
            "status": self.status.value,
            "created_at": self.created_at.strftime("%Y-%m-%d %H:%M:%S")
        }
        
        return self_dict
    @property
    def StartString(self) -> str:
        return self.utc_start.strftime("%Y-%m-%d %H:%M:%S")
    
    def durationMinutes(self) -> int:
        return self.duration.seconds / 60
    
    def durationMilliseconds(self) -> int:
        return self.duration.seconds * 1000
    

@dataclass
class AppointmentConfirmation:
    appointment: int
    expert_confirmed: bool
    consumer_confirmed: bool
    
class AppointmentENCODER(json.JSONEncoder):
    
    def default(self, o: Any) -> Any:
        if isinstance(o, timedelta):
            return o.total_seconds() * 1000
        elif isinstance(o, datetime):
            return o.strftime("%Y-%m-%d %H:%M:%S")
        return super().default(o)

class AppointmentDECODER(json.JSONDecoder):

    def __init__(self, *args, **kwargs) -> None:
        super().__init__(object_hook=self.object_hook, *args, **kwargs)
    
    def object_hook(self, obj: dict) -> Any:
        if "duration" in obj:
            obj["duration"] = timedelta(milliseconds=obj["duration"])
        if "utc_start" in obj:
            obj["utc_start"] = datetime.strptime(obj["utc_start"], "%Y-%m-%d %H:%M:%S")
        if "created_at" in obj:
            obj["created_at"] = datetime.strptime(obj["created_at"], "%Y-%m-%d %H:%M:%S")
        if "status" in obj:
            obj["status"] = AppointmentStatus(obj["status"])
        return obj