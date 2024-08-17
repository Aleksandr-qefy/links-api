package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Link struct {
	Id          uuid.UUID `db:"id"`
	UserId      uuid.UUID `db:"user_id"`
	Ref         string    `db:"ref"`
	Description *string   `db:"description"`
	Categories  []Category
}
