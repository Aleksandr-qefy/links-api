package repository

import (
	"errors"
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CategPostgres struct {
	db *sqlx.DB
}

func NewCategPostgres(db *sqlx.DB) *CategPostgres {
	return &CategPostgres{db: db}
}

func (r *CategPostgres) Create(category repoModel.Category) (uuid.UUID, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, name) VALUES ($1, $2) RETURNING id",
		categoriesTable,
	)
	row := r.db.QueryRowx(query, category.UserId, category.Name)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *CategPostgres) GetAll(userId uuid.UUID) ([]repoModel.Category, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id=$1",
		categoriesTable,
	)

	var categories []repoModel.Category
	err := r.db.Select(&categories, query, userId)

	return categories, err
}

func (r *CategPostgres) GetById(userId, categoryId uuid.UUID) (repoModel.Category, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE id=$1 AND user_id=$2",
		categoriesTable,
	)

	var category repoModel.Category
	err := r.db.Get(&category, query, categoryId, userId)

	return category, err
}

func (r *CategPostgres) DeleteById(userId, categoryId uuid.UUID) error {
	query := fmt.Sprintf(
		"SELECT COALESCE(SUM(1), 0) FROM %s WHERE id=$1 AND user_id=$2",
		categoriesTable,
	)

	var categoriesCount int
	err := r.db.Get(&categoriesCount, query, categoryId, userId)

	if categoriesCount == 0 {
		return errors.New("no category found with such id for this user")
	}

	query = fmt.Sprintf(
		"DELETE FROM %s WHERE id=$1 AND user_id=$2",
		categoriesTable,
	)
	_, err = r.db.Exec(query, categoryId, userId)
	return err
}

func (r *CategPostgres) Update(categoryUpdate repoModel.Category) error {
	query := fmt.Sprintf(
		"SELECT COALESCE(SUM(1), 0) FROM %s WHERE id=$1 AND user_id=$2",
		categoriesTable,
	)

	var categoriesCount int
	err := r.db.Get(&categoriesCount, query, categoryUpdate.Id, categoryUpdate.UserId)

	if categoriesCount == 0 {
		return errors.New("no category found with such id for this user")
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if categoryUpdate.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		argId++
		args = append(args, categoryUpdate.Name)
	}

	setValuesQueryPart := strings.Join(setValues, ", ")

	query = fmt.Sprintf(
		"UPDATE %s SET %s WHERE id=$%d AND user_id=$%d",
		categoriesTable,
		setValuesQueryPart,
		argId,
		argId+1,
	)

	args = append(args, string(categoryUpdate.Id))
	args = append(args, string(categoryUpdate.UserId))

	_, err = r.db.Exec(query, args...)
	return err
}
