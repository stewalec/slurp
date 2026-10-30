CREATE TABLE IF NOT EXISTS saved_item_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    item_url TEXT NOT NULL,
    item_title TEXT NOT NULL,
    archive_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
);

INSERT INTO
    saved_item_new
SELECT
    id,
    user_id,
    item_url,
    item_title,
    archive_url,
    created_at
FROM
    saved_item;

DROP TABLE saved_item;

ALTER TABLE saved_item_new
RENAME TO saved_item;

CREATE INDEX IF NOT EXISTS idx_saved_item_user ON saved_item (user_id);

CREATE INDEX IF NOT EXISTS idx_saved_item_url ON saved_item (item_url);

CREATE INDEX IF NOT EXISTS idx_saved_item_created ON saved_item (created_at);