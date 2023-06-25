// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createScanStmt, err = db.PrepareContext(ctx, createScan); err != nil {
		return nil, fmt.Errorf("error preparing query CreateScan: %w", err)
	}
	if q.deleteScanByIdStmt, err = db.PrepareContext(ctx, deleteScanById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteScanById: %w", err)
	}
	if q.deleteScanByUUIDStmt, err = db.PrepareContext(ctx, deleteScanByUUID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteScanByUUID: %w", err)
	}
	if q.getScanByIdStmt, err = db.PrepareContext(ctx, getScanById); err != nil {
		return nil, fmt.Errorf("error preparing query GetScanById: %w", err)
	}
	if q.getScanByUUIDStmt, err = db.PrepareContext(ctx, getScanByUUID); err != nil {
		return nil, fmt.Errorf("error preparing query GetScanByUUID: %w", err)
	}
	if q.getScansStmt, err = db.PrepareContext(ctx, getScans); err != nil {
		return nil, fmt.Errorf("error preparing query GetScans: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createScanStmt != nil {
		if cerr := q.createScanStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScanStmt: %w", cerr)
		}
	}
	if q.deleteScanByIdStmt != nil {
		if cerr := q.deleteScanByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScanByIdStmt: %w", cerr)
		}
	}
	if q.deleteScanByUUIDStmt != nil {
		if cerr := q.deleteScanByUUIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScanByUUIDStmt: %w", cerr)
		}
	}
	if q.getScanByIdStmt != nil {
		if cerr := q.getScanByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScanByIdStmt: %w", cerr)
		}
	}
	if q.getScanByUUIDStmt != nil {
		if cerr := q.getScanByUUIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScanByUUIDStmt: %w", cerr)
		}
	}
	if q.getScansStmt != nil {
		if cerr := q.getScansStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScansStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createScanStmt       *sql.Stmt
	deleteScanByIdStmt   *sql.Stmt
	deleteScanByUUIDStmt *sql.Stmt
	getScanByIdStmt      *sql.Stmt
	getScanByUUIDStmt    *sql.Stmt
	getScansStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createScanStmt:       q.createScanStmt,
		deleteScanByIdStmt:   q.deleteScanByIdStmt,
		deleteScanByUUIDStmt: q.deleteScanByUUIDStmt,
		getScanByIdStmt:      q.getScanByIdStmt,
		getScanByUUIDStmt:    q.getScanByUUIDStmt,
		getScansStmt:         q.getScansStmt,
	}
}
