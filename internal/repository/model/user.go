package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type User struct {
	Id           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	PasswordHash string    `db:"password_hash"`
}
