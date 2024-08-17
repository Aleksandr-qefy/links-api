package repository

import (
	"database/sql"
	"errors"
	"fmt"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/Aleksandr-qefy/links-api/pkg/set"
	"github.com/jmoiron/sqlx"
	"strings"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func addCategories(transaction *sql.Tx, linkId uuid.UUID, categoryIds []uuid.UUID) error {
	if categoryIds == nil || len(categoryIds) == 0 {
		return nil
	}

	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	for _, categoryId := range categoryIds {
		values = append(values, fmt.Sprintf("($%d, $%d)", argId, argId+1))
		argId += 2
		args = append(args, linkId)
		args = append(args, categoryId)
	}

	valuesQueryPart := strings.Join(values, ", ")

	query := fmt.Sprintf(
		"INSERT INTO %s (link_id, category_id) VALUES %s",
		linksCategoriesTable,
		valuesQueryPart,
	)

	_, err := transaction.Exec(query, args...)

	return err
}

func (r LinkPostgres) Create(link repoModel.Link, categories []uuid.UUID) (uuid.UUID, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	createLinkQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, ref, description) VALUES ($1, $2, $3) RETURNING id",
		linksTable,
	)

	row := tx.QueryRow(createLinkQuery, link.UserId, link.Ref, link.Description)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return "", err
	}

	if err := addCategories(tx, id, categories); err != nil {
		tx.Rollback()
		return "", err
	}

	return id, tx.Commit()
}

func getCategories(db *sqlx.DB, linkId uuid.UUID) ([]repoModel.Category, error) {
	query := fmt.Sprintf(
		"SELECT c.* FROM %s c JOIN %s lc ON c.id = lc.category_id WHERE lc.link_id = $1",
		categoriesTable,
		linksCategoriesTable,
	)

	var categories []repoModel.Category
	err := db.Select(&categories, query, linkId)

	return categories, err
}

func (r LinkPostgres) GetAll(userId uuid.UUID) ([]repoModel.Link, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE user_id = $1",
		linksTable,
	)

	var links []repoModel.Link
	if err := r.db.Select(&links, query, userId); err != nil {
		return nil, err
	}

	for i, link := range links {
		if categories, err := getCategories(r.db, link.Id); err == nil {
			links[i].Categories = categories
		} else {
			return nil, err
		}
	}

	return links, nil
}

func (r LinkPostgres) GetById(userId, linkId uuid.UUID) (repoModel.Link, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE id = $1 AND user_id = $2",
		linksTable,
	)

	var link repoModel.Link
	if err := r.db.Get(&link, query, linkId, userId); err != nil {
		return repoModel.Link{}, err
	}

	if categories, err := getCategories(r.db, linkId); err == nil {
		link.Categories = categories
	} else {
		return repoModel.Link{}, err
	}

	return link, nil
}

func (r LinkPostgres) DeleteById(userId, linkId uuid.UUID) error {
	query := fmt.Sprintf(
		"SELECT COALESCE(SUM(1), 0) FROM %s WHERE id = $1 AND user_id = $2",
		linksTable,
	)

	var linksCount int
	if err := r.db.Get(&linksCount, query, linkId, userId); err != nil {
		return err
	}

	if linksCount == 0 {
		return errors.New("no link found with such id for this user")
	}

	query = fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1 AND user_id = $2",
		linksTable,
	)
	_, err := r.db.Exec(query, linkId, userId)
	return err
}

func getCategoryIds(db *sqlx.DB, linkId uuid.UUID) ([]uuid.UUID, error) {
	query := fmt.Sprintf(
		"SELECT category_id FROM %s WHERE link_id = $1",
		linksCategoriesTable,
	)

	var categories []uuid.UUID
	err := db.Select(&categories, query, linkId)

	return categories, err
}

func removeCategories(transaction *sql.Tx, linkId uuid.UUID, categoryIds []uuid.UUID) error {
	if categoryIds == nil || len(categoryIds) == 0 {
		return nil
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	for _, categoryId := range categoryIds {
		setValues = append(setValues, fmt.Sprintf("$%d", argId))
		argId++
		args = append(args, categoryId)
	}

	valuesQueryPart := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE category_id IN (%s) AND link_id = $%d",
		linksCategoriesTable,
		valuesQueryPart,
		argId,
	)

	args = append(args, linkId)

	_, err := transaction.Exec(query, args...)
	return err
}

func (r LinkPostgres) Update(linkUpdate repoModel.Link, categories []uuid.UUID) error {
	query := fmt.Sprintf(
		"SELECT COALESCE(SUM(1), 0) FROM %s WHERE id = $1 AND user_id = $2",
		linksTable,
	)

	var linksCount int
	if err := r.db.Get(&linksCount, query, linkUpdate.Id, linkUpdate.UserId); err != nil {
		return err
	}

	if linksCount == 0 {
		return errors.New("no link found with such id for this user")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if linkUpdate.Ref != "" {
		setValues = append(setValues, fmt.Sprintf("ref=$%d", argId))
		argId++
		args = append(args, linkUpdate.Ref)
	}

	if linkUpdate.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		argId++
		if *linkUpdate.Description == "" {
			args = append(args, nil)
		} else {
			args = append(args, *linkUpdate.Description)
		}
	}

	if categories != nil {
		newCategories := set.NewSet(categories...)
		oldCategoriesSlc, err := getCategoryIds(r.db, linkUpdate.Id)
		if err != nil {
			return err
		}
		oldCategories := set.NewSet(oldCategoriesSlc...)

		categoriesToRemove := oldCategories.Minus(newCategories)
		categoriesToAdd := newCategories.Minus(oldCategories)

		if err := removeCategories(tx, linkUpdate.Id, categoriesToRemove.Slice()); err != nil {
			return err
		}

		if err := addCategories(tx, linkUpdate.Id, categoriesToAdd.Slice()); err != nil {
			return err
		}
	}

	setValuesQueryPart := strings.Join(setValues, ", ")

	query = fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d AND user_id = $%d",
		linksTable,
		setValuesQueryPart,
		argId,
		argId+1,
	)

	args = append(args, string(linkUpdate.Id))
	args = append(args, string(linkUpdate.UserId))

	if _, err := tx.Exec(query, args...); err != nil {
		return err
	}

	return tx.Commit()
}
