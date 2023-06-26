-- name: InsertArtifactTag
INSERT INTO artifact_tags (
  artifact_id,
  tag_label,
  tag_value
) VALUES (
           $1,
           $2,
           $3
         ) RETURNING id;

-- name: GetArtifactTagByID
SELECT *
FROM artifact_tags
WHERE id = $1;

-- name: DeleteArtifactTag
DELETE FROM artifact_tags
WHERE id = $1;
