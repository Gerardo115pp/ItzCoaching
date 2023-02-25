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


@dataclass(frozen=True)
class Payment:
    id: int
    appointment: int
    amount: float
    currency: str
    stripe_charge_id: str
    stripe_customer_id: str
    stripe_payment_intent_id: str
    stripe_payment_id: str
    description: str
    status: PaymentStatus
    created_at: datetime
    