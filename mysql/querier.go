package mysql

import (
	"context"
	"database/sql"
)

// RowQuerier requests a row in specific context
type RowQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// RowsQuerier requests rows in specific context
type RowsQuerier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// Executer executes request in a specific context
type Executer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// Querier is a global request sender
type Querier interface {
	RowQuerier
	RowsQuerier
	Executer
}
