package hello

import "context"

type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
}

func NewService(repository Repository) Service {
	return service{
		repository: &repository,
	}
}

type service struct {
	repository *Repository
}

func (service) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name + "!", nil
}
