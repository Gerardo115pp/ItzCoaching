from datetime import datetime
from dataclasses import dataclass, asdict
from enum import Enum

def allArgsPresent(obj: object, args: list[str]) -> bool:
    """ 
        this method should be reassigned from the __init__.py file
    """
    raise NotImplementedError("this shouldnt happen")

class PaymentStatus(str, Enum):
    PENDING = "pending"
    SUCCEDED = "succeded"
    FAILED = "failed"
    REFUNDED = "refunded"


@dataclass
class Payment:
    id: int
    appointment: int
    amount: float
    currency: str
    stripe_charge_id: str
    stripe_customer_id: str
    stripe_payment_intent_id: str
    stripe_session_id: str
    stripe_refund_id: str
    description: str
    status: PaymentStatus
    created_at: datetime
    
    
def createEmptyPayment(appointment_id: int, amount: int, currency: str='mxn') -> Payment:
    return Payment(
        id=None,
        appointment=appointment_id,
        amount=amount,
        currency=currency,
        stripe_charge_id="",
        stripe_customer_id="",
        stripe_payment_intent_id="",
        stripe_session_id="",
        stripe_refund_id="",
        description="",
        status=PaymentStatus.PENDING,
        created_at=datetime.now()
    )