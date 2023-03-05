from typing import Any
from models import ExpertPersonalData

experts_interface = [
    "getExpertPersonalData",
]

def checkMethodImplemented(cls: Any, method_name: str) -> bool:
    # hoisted function, definition should be in the __init__.py file
    raise NotImplementedError("This function wasnt defined")

class ExpertsMeta(type):
    """
        Defines behavioral contract of the Expert Repository. Dont use directly
    """
    
    def __instancecheck__(cls, __instance: Any) -> bool:
        return cls.__subclasscheck__(type(__instance))
    
    def __subclasscheck__(cls, __subclass: type) -> bool:
        return all(checkMethodImplemented(cls, method_name) for method_name in experts_interface)

class ExpertsRepository(metaclass=ExpertsMeta):
    """
        Defines behavioral contract of the Redis Cache. Dont use directly
    """
    
    def getExpertPersonalData(self, expert_id: int) -> tuple[ExpertPersonalData, Exception]:
        raise NotImplementedError("This method wasnt implemented")

implementation: ExpertsRepository = None

def setExpertsRepository(repo: ExpertsRepository) -> None:
    global implementation
    assert isinstance(repo, ExpertsRepository)
    implementation = repo
    
def getExpertPersonalData(expert_id: int) -> tuple[ExpertPersonalData, Exception]:
    return implementation.getExpertPersonalData(expert_id)