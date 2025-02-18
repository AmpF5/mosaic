-- +goose Up
CREATE TABLE files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size INTEGER NOT NULL,
    modified_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- +goose Down
DROP TABLE files;