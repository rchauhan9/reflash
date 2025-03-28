package study

import (
	"context"
	db "github.com/rchauhan9/reflash/monolith/common/database"
)

type Repository interface {
	CreateStudy(ctx context.Context) (StudyProject, error)
	CreateCards(ctx context.Context, studyProjectID string, cards []StudyProjectCard) error
	DeleteCards(ctx context.Context, studyProjectID string) error
}

func NewRepository(dbPool db.Pool) Repository {
	return &repository{
		pool: dbPool,
	}
}

type repository struct {
	pool db.Pool
}

func (r *repository) CreateStudy(ctx context.Context) (StudyProject, error) {

	return StudyProject{}, nil
}

func (r *repository) CreateCards(ctx context.Context, studyProjectID string, cards []StudyProjectCard) error {
	return nil
}

func (r *repository) DeleteCards(ctx context.Context, studyProjectID string) error {
	return nil
}
