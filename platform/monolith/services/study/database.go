package study

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	db "github.com/rchauhan9/reflash/monolith/common/database"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDatabasePool(ctx context.Context, connectionString string) (db.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to database %s", connectionString)
	}
	return pool, nil
}
