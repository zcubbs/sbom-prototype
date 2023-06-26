-- name: InsertArtifact
INSERT INTO artifacts (artifact_name,
                       artifact_type,
                       artifact_version)
VALUES ($1,
        $2,
        $3)
RETURNING id;

-- name: GetArtifactByID
SELECT *
FROM artifacts
WHERE id = $1;

-- name: GetArtifactByUUID
SELECT *
FROM artifacts
WHERE uuid = $1;

-- name: UpdateArtifact
UPDATE artifacts
SET artifact_name    = $1,
    artifact_type    = $2,
    artifact_version = $3
WHERE id = $4;

-- name: DeleteArtifact
DELETE
FROM artifacts
WHERE id = $1;
