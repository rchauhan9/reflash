package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is an interface for a pgxpool.
type Pool interface {
	Close()

	Acquire(ctx context.Context) (*pgxpool.Conn, error)

	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error

	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn

	Config() *pgxpool.Config

	Stat() *pgxpool.Stat

	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults

	Begin(ctx context.Context) (pgx.Tx, error)

	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error

	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)

	Ping(ctx context.Context) error
}
