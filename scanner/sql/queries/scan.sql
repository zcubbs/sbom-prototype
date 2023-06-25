-- name: CreateScan :one
INSERT INTO scan (
                  uuid,
                  image,
                  image_tag,
                  sbom_id,
                  status,
                  artifact_id,
                  artifact_name,
                  artifact_version,
                  created_at,
                  updated_at
                  )
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
ORDER BY created_at desc
LIMIT $1 OFFSET $2;

-- name: CountScans :one
SELECT count(*) FROM scan;
