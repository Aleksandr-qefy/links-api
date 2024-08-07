package repository

import (
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user repoModel.User) (uuid.UUID, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name, password_hash) values ($1, $2) RETURNING id",
		usersTable,
	)
	row := r.db.QueryRow(query, user.Name, user.PasswordHash)
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}
