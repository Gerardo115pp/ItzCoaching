from . import redis_cache
from . import appointments
from . import payments
from . import experts
from typing import Any

def _checkMethodImplemented(cls: Any, method_name: str) -> bool:
    method_exists = hasattr(cls, method_name) and (callable(getattr(cls, method_name)) or isinstance(getattr(cls, method_name), property))
    if not method_exists:
        print(f"Method {method_name} not implemented in {cls}")
    return method_exists

redis_cache.checkMethodImplemented = _checkMethodImplemented
appointments.checkMethodImplemented = _checkMethodImplemented
experts.checkMethodImplemented = _checkMethodImplemented
payments.checkMethodImplemented = _checkMethodImplemented