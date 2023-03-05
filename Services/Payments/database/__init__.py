from .redis_cache import RedisCache
from .redis_utils import REDIS_CONFIG
from .mysql_utils import MYSQL_CONFIG
from .appointments import AppointmentsRepository
from .experts import ExpertsRepository
from .payments import PaymentsRepository

def createRedisCache() -> RedisCache:
    return RedisCache(REDIS_CONFIG.createFromEnv())

def createAppointmentsRepository() -> AppointmentsRepository:
    return AppointmentsRepository(MYSQL_CONFIG.createFromEnv())

def createExpertsRepository() -> ExpertsRepository:
    return ExpertsRepository(MYSQL_CONFIG.createFromEnv())

def createPaymentsRepository() -> PaymentsRepository:
    return PaymentsRepository(MYSQL_CONFIG.createFromEnv())