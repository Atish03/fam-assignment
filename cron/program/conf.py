import json
import os

SEARCH_QUERY = os.environ.get("QUERY", "Formula 1")
SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
MAX_RESULTS = 50

# Database connection settings
DB_HOST = os.environ.get("DB_HOST", "localhost")
DB_NAME = os.environ.get("DB_NAME", "postgres")
DB_USER = os.environ.get("DB_USER", "postgres")
DB_PASSWORD = os.environ.get("DB_PASS", "passwd")

KEYS_FILEPATH = os.environ.get("KEYS_FILE", "/config/api-keys.json")

def read_json():
    if not os.path.exists(KEYS_FILEPATH):
        return {"keys": []}

    with open(KEYS_FILEPATH, 'r') as f:
        data = json.load(f)
    return data

def __get_least_used_key(data):
    if not data['keys']:
        return None

    least_used_key = None
    min_usage = float('inf')

    for key in data['keys']:
        for key_name, usage_count in key.items():
            if usage_count < min_usage:
                min_usage = usage_count
                least_used_key = key_name

    return least_used_key

def get_api_key():
    data = read_json()

    least_used_key = __get_least_used_key(data)
    
    if not least_used_key:
        print("No keys available.")
        return None
    
    for key in data['keys']:
        if least_used_key in key:
            key[least_used_key] += 1
            if key[least_used_key] == 100:
                data['keys'].remove(key)
            break

    with open(KEYS_FILEPATH, 'w') as f:
        json.dump(data, f, indent=2)

    return least_used_key