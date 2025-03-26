package study

import (
	"context"
	db "github.com/rchauhan9/reflash/monolith/common/database"
)

type Repository interface {
	CreateStudy(ctx context.Context) (Study, error)
}

func NewRepository(dbPool db.Pool) Repository {
	return &repository{
		pool: dbPool,
	}
}

type repository struct {
	pool db.Pool
}

func (r *repository) CreateStudy(ctx context.Context) (Study, error) {

	return Study{}, nil
}
