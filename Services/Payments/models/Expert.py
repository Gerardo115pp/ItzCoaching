from dataclasses import dataclass, asdict
from datetime import datetime, timedelta
from typing import Any
import json

def allArgsPresent(obj: object, args: list[str]) -> bool:
    """ 
        this method should be reassigned from the __init__.py file
    """
    raise NotImplementedError("this shouldnt happen")

@dataclass(frozen=True)
class ExpertPersonalData:
    name: str
    email: str
    username: str
    minute_price: float
    
    @property
    def MinutePrice(self) -> float:
        return float(self.minute_price)