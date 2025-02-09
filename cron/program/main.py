import requests
import datetime

import conf
import database

def get_videos(api_key, after: str) -> list:
    params = {
        "part": "snippet",
        "q": conf.SEARCH_QUERY,
        "maxResults": conf.MAX_RESULTS,
        "order": "date",
        "type": "video",
        "publishedAfter": after,
        "key": api_key
    }
    
    return_data = []

    # Requesting youtube api for 50 videos matching the parameters
    response = requests.get(conf.SEARCH_URL, params=params)
    if response.status_code == 200:
        data = response.json()
        if "items" in data and data["items"]:
            videos = data["items"]
            
            for video in videos:
                video_id = video["id"]["videoId"]
                title = video["snippet"]["title"]
                description = video["snippet"]["description"]
                published_at = video["snippet"]["publishedAt"]
                thumbnail_url = video["snippet"]["thumbnails"]["medium"]["url"]
                return_data.append((
                    video_id,
                    title,
                    description,
                    published_at,
                    thumbnail_url
                ))
    else:
        print(f"API Error: {response.status_code}, {response.text}")
        
    return return_data

if __name__ == "__main__":
    db = database.Database()
    
    api_key = conf.get_api_key()
    published_after = db.get_last_video_publish()
    
    if api_key:
        videos = get_videos(api_key, published_after)
        db.insert_videos(videos)
        
        print(f"Inserted {len(videos) - 1} videos at {datetime.datetime.now()}")
        
        db.commit()
