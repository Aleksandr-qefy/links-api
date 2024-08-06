package repository

type Authorization interface {
}

type Link interface {
}

type Repository struct {
	Authorization
	Link
}

func NewRepository() *Repository {
	return &Repository{}
}
