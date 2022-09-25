import redis
import config
import os
from config import get_string

def getCacheHostname(default):
  value = os.getenv("CACHE_HOSTNAME")
  if value is None:
      value = default
  return value

def connectToCache():
  host = getCacheHostname("localhost")
  port = config.get_string(["cache", "port"], 6379)
  db = config.get_string(["cache", "db"], 0)
  return redis.Redis(host=host, port=port, db=db)

cache = connectToCache()

def write():
  try:
    key = config.get_string(["cache", "key"], "count")
    cache.incrby(key, 1)
  except BaseException as err:
    print(f"Unexpected {err=}, {type(err)=}")
    raise

def read():
  try:
    key = config.get_string(["cache", "key"], "count")
    value = cache.get(key)
    if value is None:
      return 0
    return value
  except:
    raise
