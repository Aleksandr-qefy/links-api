package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

type User struct {
	Id           uuid.UUID `json:"-"`
	Name         string    `json:"name" binding:"required"`
	PasswordHash string    `json:"password" binding:"required"`
}

type UserAccount struct {
	Name     string `json:"name" binding:"required" example:"john"`
	Password string `json:"password" binding:"required" example:"1234"`
}
