CREATE TABLE files
(
    id UUID PRIMARY KEY,
    file_path TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size Integer,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);