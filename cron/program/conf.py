import json
import os
import time

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
    min_usage = 100

    for i, key in enumerate(data['keys']):
        for key_name, usage in key.items():
            if usage["count"] < min_usage:
                min_usage = usage["count"]
                least_used_key = key_name
            else:
                time_diff = time.time() - usage["start_time"]
                if time_diff >= 24 * 60 * 60:
                    data["keys"][i][key_name]["count"] = 0
                    with open(KEYS_FILEPATH, 'w') as f:
                        json.dump(data, f, indent=2)

    return least_used_key

def get_api_key():
    data = read_json()

    least_used_key = __get_least_used_key(data)
    
    if not least_used_key:
        print("No keys available.")
        return None
    
    for key in data['keys']:
        if least_used_key in key:
            if key[least_used_key]["count"] == 0:
                key[least_used_key]["start_time"] = int(time.time())
            key[least_used_key]["count"] += 1
            break

    with open(KEYS_FILEPATH, 'w') as f:
        json.dump(data, f, indent=2)

    return least_used_key