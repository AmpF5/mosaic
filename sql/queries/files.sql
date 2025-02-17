-- name: CreateFile :one
INSERT INTO files(id, file_path, mime_type, size, modified_at)
VALUES(value1, value2, value3, value4, value5)
RETURNING *;

-- name: GetFiles :many
SELECT * FROM files