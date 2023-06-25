-- name: CreateScan :one
INSERT INTO scan (id, uuid)
values ($1, $2)
RETURNING *;

-- name: GetScanByUUID :one
SELECT *
FROM scan
WHERE uuid = $1;

-- name: GetScanById :one
SELECT *
FROM scan
WHERE id = $1;

-- name: DeleteScanById :exec
DELETE
FROM scan
WHERE id = $1;

-- name: DeleteScanByUUID :exec
DELETE
FROM scan
WHERE uuid = $1;

-- name: GetScans :many
SELECT *
FROM scan
ORDER BY created_at DESC;


