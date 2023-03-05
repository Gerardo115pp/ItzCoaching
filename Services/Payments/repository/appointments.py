from typing import Any
from models import Appointment

appointments_interface = [
    "insertAppointment",
    "getAppointmentByCustomerExpert",
    "verifyAvailability",
    
]

def checkMethodImplemented(cls: Any, method_name: str) -> bool:
    # hoisted function, definition should be in the __init__.py file
    raise NotImplementedError("This function wasnt defined")

class AppointmentMeta(type):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    
    def __instancecheck__(cls, __instance: Any) -> bool:
        return cls.__subclasscheck__(type(__instance))
    
    def __subclasscheck__(cls, __subclass: type) -> bool:
        return all(checkMethodImplemented(cls, method_name) for method_name in appointments_interface)

class AppointmentRepository(metaclass=AppointmentMeta):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    
    def insertAppointment(self, appointment: Appointment) -> tuple[int,Exception]:
        raise NotImplementedError("This method wasnt implemented")
    
    def getAppointmentByCustomerExpert(self, customer_email: str, expert: int) -> tuple[Appointment, Exception]:
        raise NotImplementedError("This method wasnt implemented")
    
    def verifyAvailability(self, appointment: Appointment) -> Exception:
        raise NotImplementedError("This method wasnt implemented")
    
implementation: AppointmentRepository = None

def setAppointmentRepository(repo: AppointmentRepository) -> None:
    global implementation
    assert isinstance(repo, AppointmentRepository)
    implementation = repo
    
def insertAppointment(appointment: Appointment) -> tuple[int, Exception]:
    return implementation.insertAppointment(appointment)

def getAppointmentByCustomerExpert(customer_email: str, expert: int) -> tuple[Appointment, Exception]:
    return implementation.getAppointmentByCustomerExpert(customer_email, expert)

def verifyAvailability(appointment: Appointment) -> Exception:
    return implementation.verifyAvailability(appointment)