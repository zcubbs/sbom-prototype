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
                  updated_at,
                  log
                  )
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetScanByUUID :one
SELECT *
FROM scan
WHERE uuid = $1;

-- name: GetScanById :one
SELECT *
FROM scan
WHERE id = $1;

-- name: UpdateScan :one
UPDATE scan
SET
    uuid = $2,
    image = $3,
    image_tag = $4,
    sbom_id = $5,
    status = $6,
    artifact_id = $7,
    artifact_name = $8,
    artifact_version = $9,
    created_at = $10,
    updated_at = $11,
    log = $12
WHERE uuid = $1
RETURNING *;

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
