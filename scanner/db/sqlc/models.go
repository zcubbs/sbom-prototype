// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type ScanJob struct {
	ID              int64                 `json:"id"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	SbomUuid        uuid.NullUUID         `json:"sbom_uuid"`
	ArtifactUuid    uuid.NullUUID         `json:"artifact_uuid"`
	ArtifactName    string                `json:"artifact_name"`
	ArtifactVersion string                `json:"artifact_version"`
	ArtifactType    string                `json:"artifact_type"`
	Status          string                `json:"status"`
	Report          pqtype.NullRawMessage `json:"report"`
	JobLog          sql.NullString        `json:"job_log"`
}
