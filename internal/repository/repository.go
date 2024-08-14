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
	Create(link repoModel.Link) (uuid.UUID, error)
	GetAll(userId uuid.UUID) ([]repoModel.Link, error)
	GetById(userId, linkId uuid.UUID) (repoModel.Link, error)
	DeleteById(userId, linkId uuid.UUID) error
	Update(link repoModel.Link) error
}

type Repository struct {
	Authorization
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Link:          NewLinkPostgres(db),
	}
}
