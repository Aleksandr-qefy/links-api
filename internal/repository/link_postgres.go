package repository

import (
	"errors"
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

/*
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   user_id UUID REFERENCES users ON DELETE CASCADE,
   ref TEXT NOT NULL,
   description TEXT,
   UNIQUE(user_id, ref)
*/

func (r LinkPostgres) Create(link repoModel.Link) (uuid.UUID, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var id uuid.UUID
	if link.Description != nil {
		createLinkQuery := fmt.Sprintf(
			"INSERT INTO %s (user_id, ref, description) VALUES ($1, $2, $3) RETURNING id",
			linksTable,
		)
		row := tx.QueryRow(createLinkQuery, link.UserId, link.Ref, link.Description)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return "", err
		}
	} else {
		createLinkQuery := fmt.Sprintf(
			"INSERT INTO %s (user_id, ref) VALUES ($1, $2) RETURNING id",
			linksTable,
		)
		row := tx.QueryRow(createLinkQuery, link.UserId, link.Ref)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return "", err
		}
	}

	if len(link.Categories) > 0 {
		valuesList := make([]string, len(link.Categories))
		for i := 0; i < len(link.Categories); i++ {
			valuesList[i] = fmt.Sprintf("($1, $%d)", i+2)
		}

		createLinksCategoriesQuery := fmt.Sprintf(
			"INSERT INTO %s (link_id, category_id) VALUES %s",
			linksCategoriesTable,
			strings.Join(valuesList, ", "),
		)

		values := make([]uuid.UUID, 0, len(link.Categories)+1)
		values = append(values, link.UserId)
		values = append(values, link.Categories...)
		_, err = tx.Exec(createLinksCategoriesQuery, values)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	return id, tx.Commit()
}

func (r LinkPostgres) GetAll(userId uuid.UUID) ([]repoModel.Link, error) {
	var links []repoModel.Link
	query := fmt.Sprintf(
		"SELECT * FROM %s l WHERE l.user_id = $1",
		linksTable,
	)

	err := r.db.Select(&links, query, userId)

	return links, err
}

func (r LinkPostgres) GetById(userId, linkId uuid.UUID) (repoModel.Link, error) {
	var link repoModel.Link
	query := fmt.Sprintf(
		"SELECT * FROM %s l WHERE id = $1 AND l.user_id = $2",
		linksTable,
	)

	err := r.db.Get(&link, query, linkId, userId)

	return link, err
}

func (r LinkPostgres) DeleteById(userId, linkId uuid.UUID) error {
	var linksCount int
	query := fmt.Sprintf(
		"SELECT SUM(1) FROM %s l WHERE id = $1 AND l.user_id = $2",
		linksTable,
	)
	err := r.db.Select(&linksCount, query, linkId, userId)

	if linksCount == 0 {
		return errors.New("no link with such id for this user found")
	}

	query = fmt.Sprintf(
		"DELETE FROM %s l WHERE id = $1 AND l.user_id = $2",
		linksTable,
	)
	_, err = r.db.Exec(query, linkId, userId)
	return err
}
