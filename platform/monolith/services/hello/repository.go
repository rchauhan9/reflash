package hello

type Repository interface {
}

func NewRepository() Repository {
	return repository{}
}

type repository struct{}
