from .redis_cache import RedisCache
from .redis_utils import REDIS_CONFIG

def createRedisCache() -> RedisCache:
    return RedisCache(REDIS_CONFIG.createFromEnv())