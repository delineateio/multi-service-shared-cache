import yaml
import os

def get_full_path():
  path = os.path.abspath(__file__)
  dir_path = os.path.dirname(path)
  return dir_path + "/config.yaml"

def load_config():

  with open(get_full_path(), "r") as file:
    return yaml.safe_load(file)

config = load_config()

def get_string(keys, default):
  current = config
  for key in keys:
    if key == keys[len(keys) -1]:
      return current[key]
    else:
      current = current[key]
  return default
