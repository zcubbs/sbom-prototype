-- name: InsertScanJob
INSERT INTO sboms (created_at,
                   manifest,
                   job_log,
                   status,
                   artifact_uuid,
                   artifact_name,
                   artifact_version,
                   artifact_type)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8)
RETURNING id;

-- name: GetScanJobByID
SELECT *
FROM sboms
WHERE id = $1;

-- name: GetScanJobByUUID
SELECT *
FROM sboms
WHERE uuid = $1;

-- name: UpdateScanJob
UPDATE sboms
SET created_at       = $1,
    manifest         = $2,
    job_log          = $3,
    status           = $4,
    artifact_uuid    = $5,
    artifact_name    = $6,
    artifact_version = $7,
    artifact_type    = $8,
    updated_at       = current_timestamp
WHERE id = $9;

-- name: DeleteScanJob
DELETE
FROM sboms
WHERE id = $1;

-- name: GetScanJobsList
SELECT *
FROM sboms
ORDER BY id
LIMIT $1 OFFSET $2;
