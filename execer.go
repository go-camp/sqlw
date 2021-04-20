package sqlw

import (
	"context"
	"database/sql"
)

var _ Execer = (*sql.DB)(nil)
var _ Execer = (*sql.Conn)(nil)
var _ Execer = (*sql.Tx)(nil)

// Execer is a set of common methods from sql.DB, sql.Conn and sql.Tx.
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}
