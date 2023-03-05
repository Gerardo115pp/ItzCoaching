from typing import Any
from models import Payment

payments_interface = [
    "insertPayment",
    "updatePayment",
    "getAllPayments",
    "getPaymentBySession",
    "getPaymentByAppointment"
]

def checkMethodImplemented(cls: Any, method_name: str) -> bool:
    # hoisted function, definition should be in the __init__.py file
    raise NotImplementedError("This function wasnt defined")

class PaymentMeta(type):
    """
        Defines behavioral contract of the Payment repo. Dont use directly
    """
    
    def __instancecheck__(cls, __instance: Any) -> bool:
        return cls.__subclasscheck__(type(__instance))
    
    def __subclasscheck__(cls, __subclass: type) -> bool:
        return all(checkMethodImplemented(cls, method_name) for method_name in payments_interface)

class PaymentsRepository(metaclass=PaymentMeta):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    
    def insertPayment(self, payment: Payment) -> Exception:
        raise NotImplementedError("This method wasnt implemented")
    
    def updatePayment(self, payment: Payment) -> Exception:
        raise NotImplementedError("This method wasnt implemented")
    
    def getAllPayments(self) -> tuple[list[Payment], Exception]:
        raise NotImplementedError("This method wasnt implemented")
    
    def getPaymentBySession(self, session_id: str) -> tuple[Payment, Exception]:
        raise NotImplementedError("This method wasnt implemented")
    
    def getPaymentByAppointment(self, appointment_id: int) -> tuple[Payment, Exception]:
        raise NotImplementedError("This method wasnt implemented")
    
implementation: PaymentsRepository = None

def setPaymentsRepository(repo: PaymentsRepository) -> None:
    global implementation
    assert isinstance(repo, PaymentsRepository)
    implementation = repo
    
def insertPayment(payment: Payment) -> Exception:
    return implementation.insertPayment(payment)

def updatePayment(payment: Payment) -> Exception:
    return implementation.updatePayment(payment)

def getAllPayments() -> tuple[list[Payment], Exception]:
    return implementation.getAllPayments()

def getPaymentBySession(session_id: str) -> tuple[Payment, Exception]:
    return implementation.getPaymentBySession(session_id)

def getPaymentByAppointment(appointment_id: int) -> tuple[Payment, Exception]:
    return implementation.getPaymentByAppointment(appointment_id)