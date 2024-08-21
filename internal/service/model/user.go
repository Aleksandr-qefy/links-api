package links_api

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

type User struct {
	Id           uuid.UUID
	Name         string
	PasswordHash string
}

type UserAccount struct {
	Name     string
	Password string
}
