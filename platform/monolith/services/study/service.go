package study

import "context"

type Service interface {
	CreateStudy(ctx context.Context) (Study, error)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

type service struct {
	repository Repository
}

func (s *service) CreateStudy(ctx context.Context) (Study, error) {
	return s.repository.CreateStudy(ctx)
}

type Study struct {
	ID string
}
