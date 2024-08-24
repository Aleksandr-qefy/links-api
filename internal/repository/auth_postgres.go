package repository

import (
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user repoModel.User) (uuid.UUID, error) {
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

func (r *AuthPostgres) GetUser(user repoModel.User) (repoModel.User, error) {
	var outputUser repoModel.User
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE name=$1 AND password_hash=$2",
		usersTable,
	)
	err := r.db.Get(&outputUser, query, user.Name, user.PasswordHash)
	return outputUser, err
}

func (r *AuthPostgres) DeleteAccount(userId uuid.UUID) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id=$1",
		usersTable,
	)
	_, err := r.db.Exec(query, userId)
	return err
}
