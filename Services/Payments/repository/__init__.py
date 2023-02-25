from . import redis_cache
from typing import Any

def _checkMethodImplemented(cls: Any, method_name: str) -> bool:
    return hasattr(cls, method_name) and (callable(getattr(cls, method_name)) or isinstance(getattr(cls, method_name), property))

redis_cache.checkMethodImplemented = _checkMethodImplemented