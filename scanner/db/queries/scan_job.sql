-- name: InsertScanJob
INSERT INTO scan_jobs (
  artifact_uuid,
  artifact_name,
  artifact_version,
  artifact_type,
  status,
  report,
  job_log
) VALUES (
           $1,
           $2,
           $3,
           $4,
           $5,
           $6,
           $7
         ) RETURNING id;

-- name: GetScanJobByID
SELECT *
FROM scan_jobs
WHERE id = $1;

-- name: UpdateScanJob
UPDATE scan_jobs
SET
  artifact_uuid = $1,
  artifact_name = $2,
  artifact_version = $3,
  artifact_type = $4,
  status = $5,
  report = $6,
  job_log = $7,
  updated_at = current_timestamp
WHERE id = $8;

-- name: DeleteScanJob
DELETE FROM scan_jobs
WHERE id = $1;

-- name: GetScanJobsList
SELECT *
FROM scan_jobs
ORDER BY id
LIMIT $1
OFFSET $2;
