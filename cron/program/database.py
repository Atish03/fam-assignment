import psycopg2
from psycopg2.extras import execute_values
import datetime

import conf

class Database:
    def __init__(self) -> None:
        self.conn = psycopg2.connect(
            host=conf.DB_HOST,
            dbname=conf.DB_NAME,
            user=conf.DB_USER,
            password=conf.DB_PASSWORD
        )
        self.cur = self.conn.cursor()
    
    def insert_videos(self, videos: list) -> None:
        if len(videos) == 0:
            return

        try:
            sql = """
            INSERT INTO videos (video_id, title, description, published_at, thumbnail_url)
            VALUES %s
            ON CONFLICT (video_id) DO NOTHING;
            """

            execute_values(self.cur, sql, videos)

        except Exception as e:
            print(f"Database Error: {e}")
            
    def get_last_video_publish(self) -> str:
        try:
            sql = """
            SELECT MAX(published_at) FROM videos;
            """
            
            self.cur.execute(sql)
            last_pub = self.cur.fetchone()[0]
            
            if not last_pub:
                return (datetime.datetime.now() - datetime.timedelta(hours=1)).isoformat("T") + "Z"
            
            return last_pub.isoformat("T")
        
        except Exception as e:
            print(f"Cannot query for publish time of latest video: {e}")
            
    def commit(self):
        self.conn.commit()
        self.cur.close()
        self.conn.close()