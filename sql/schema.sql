CREATE TABLE files
(
    id INTEGER PRIMARY KEY,
    file_path TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size INTEGER,
    modified_at TIMESTAMP NOT NULL
);