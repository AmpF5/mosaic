-- name: CreateFile :one
INSERT INTO files(file_path, mime_type, size)
VALUES(?, ?, ?)
RETURNING *;

-- name: GetFiles :many
SELECT * FROM files