from typing import Any
from models import Appointment

redis_cache_interface = [
    "getPendingAppointment",
    "setPendingAppointment",
]

def checkMethodImplemented(cls: Any, method_name: str) -> bool:
    # hoisted function, definition should be in the __init__.py file
    raise NotImplementedError("This function wasnt defined")

class RedisCacheMeta(type):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    
    def __instancecheck__(cls, __instance: Any) -> bool:
        return cls.__subclasscheck__(type(__instance))
    
    def __subclasscheck__(cls, __subclass: type) -> bool:
        return all(checkMethodImplemented(cls, method_name) for method_name in redis_cache_interface)

class RedisCache(metaclass=RedisCacheMeta):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    def getPendingAppointment(self, appointment_hash: str) -> tuple[Appointment, Exception]:
        raise NotImplementedError("This method wasnt implemented")

    def setPendingAppointment(self, appointment: Appointment) -> Exception:
        raise NotImplementedError("This method wasnt implemented")
    
implementation: RedisCache = None

def setCache(cache: RedisCache) -> None:
    global implementation
    assert isinstance(cache, RedisCache)
    implementation = cache

def getPendindAppointment(appointment_hash: str) -> tuple[Appointment, Exception]:
    return implementation.getPendingAppointment(appointment_hash)

def setPendingAppointment(appointment: Appointment) -> Exception:
    implementation.setPendingAppointment(appointment)