package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Link interface {
}

type Repository struct {
	Authorization
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
