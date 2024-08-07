package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type User struct {
	Id           uuid.UUID
	Name         string
	PasswordHash string
}
