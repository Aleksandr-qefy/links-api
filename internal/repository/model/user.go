package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type User struct {
	Id           uuid.UUID `db:"id"`
	Name         string
	PasswordHash string
}
