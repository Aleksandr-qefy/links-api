package repository

import (
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
)

type StatPostgres struct {
	db *sqlx.DB
}

func NewStatPostgres(db *sqlx.DB) *StatPostgres {
	return &StatPostgres{db: db}
}

func (r StatPostgres) Create(category repoModel.Statistic) (uuid.UUID, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, created_at, activity, comment) VALUES ($1, $2, $3, $4) RETURNING id",
		statisticsTable,
	)

	var id uuid.UUID
	if err := r.db.Get(
		&id,
		query,
		category.UserId,
		category.CreatedAt,
		category.Activity,
		category.Comment,
	); err != nil {
		return "", err
	}

	return id, nil
}

func (r StatPostgres) GetAll(userId uuid.UUID) ([]repoModel.Statistic, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id = $1 ORDER BY created_at",
		statisticsTable,
	)

	var statistics []repoModel.Statistic
	if err := r.db.Select(&statistics, query, userId); err != nil {
		return nil, err
	}

	return statistics, nil
}
