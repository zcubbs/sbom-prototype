// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Scan struct {
	ID        sql.NullInt32  `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Image     string         `json:"image"`
	Status    string         `json:"status"`
	Sbom      sql.NullString `json:"sbom"`
	Report    sql.NullString `json:"report"`
}
