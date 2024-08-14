package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Category struct {
	Id     uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`
	Name   string    `db:"name"`
}
