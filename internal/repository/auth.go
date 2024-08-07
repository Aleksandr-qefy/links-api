package repository

import (
	"fmt"
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user api.User) (api.UUID, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash) values ($1, $2) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Password)
	var id api.UUID
	if err := row.Scan(&id); err != nil {
		return api.UUID(""), err
	}
	return api.UUID(""), nil
}
