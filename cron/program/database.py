import psycopg2
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
    
    def insert_video(self, video: dict) -> None:
        if not video:
            return

        try:
            sql = """
            INSERT INTO videos (video_id, title, description, published_at, thumbnail_url)
            VALUES (%s, %s, %s, %s, %s)
            ON CONFLICT (video_id) DO NOTHING;
            """

            self.cur.execute(sql, (
                video["video_id"], video["title"], video["description"], video["published_at"], video["thumbnail_url"]
            ))

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