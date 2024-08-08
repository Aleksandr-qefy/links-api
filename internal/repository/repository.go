package repository

import (
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user repoModel.User) (uuid.UUID, error)
	GetUser(user repoModel.User) (repoModel.User, error)
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
