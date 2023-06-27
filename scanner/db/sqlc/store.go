package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateScanJobParamsTx struct {
	InsertScanJobParams
}

type CreateScanJobParamsTxResult struct {
	ScanJob
}

func (store *Store) CreateScanJobTx(ctx context.Context, arg CreateScanJobParamsTx) (CreateScanJobParamsTxResult, error) {
	var result CreateScanJobParamsTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.ScanJob, err = q.InsertScanJob(ctx, InsertScanJobParams{
			ArtifactUuid:    uuid.NullUUID{},
			ArtifactName:    arg.ArtifactName,
			ArtifactVersion: arg.ArtifactVersion,
			ArtifactType:    arg.ArtifactType,
			Status:          arg.Status,
			Report:          pqtype.NullRawMessage{},
			JobLog:          sql.NullString{},
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
