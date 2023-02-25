# from dataclasses import dataclass, asdict
from datetime import datetime, timedelta
import json

def allArgsPresent(obj: object, args: list[str]) -> bool:
    """ 
        this method should be reassigned from the __init__.py file
    """
    raise NotImplementedError("this shouldnt happen")

class TimeSlot:
    def __init__(self, starttime: str, endtime: str) -> None:
        self.starttime = datetime.strptime(starttime, "%Y-%m-%d %H:%M:%S")
        self.endtime = datetime.strptime(endtime, "%Y-%m-%d %H:%M:%S")
        self.validate()
        
    def validate(self) -> None:
        if self.starttime >= self.endtime:
            raise ValueError("endtime must be after starttime")
        
    @property
    def Duration(self) -> timedelta:
        return self.endtime - self.starttime
    
    @property
    def Minutes(self) -> int:
        return self.Duration.total_seconds() // 60
    
    def __cloneWithMachingDates(self, start:datetime, end:datetime) -> "TimeSlot":
        is_day_after = (end - start).days > 0
        
        new_start = start.replace(self.starttime.year, self.starttime.month, self.starttime.day)
        new_end = end.replace(self.endtime.year, self.endtime.month, self.endtime.day + (1 if is_day_after else 0))
        
        return TimeSlot(new_start.strftime("%Y-%m-%d %H:%M:%S"), new_end.strftime("%Y-%m-%d %H:%M:%S"))