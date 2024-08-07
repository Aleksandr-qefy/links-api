package repository

import (
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user api.User) (api.UUID, error)
}

type Link interface {
}

type Repository struct {
	Authorization
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
	}
}
