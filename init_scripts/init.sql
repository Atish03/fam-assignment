-- Create the table if it doesn't exist
CREATE TABLE IF NOT EXISTS videos (
    video_id TEXT UNIQUE,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    thumbnail_url TEXT
);

CREATE INDEX IF NOT EXISTS idx_published_at ON videos (published_at DESC);
