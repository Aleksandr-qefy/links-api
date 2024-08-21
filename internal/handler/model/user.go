package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

// https://www.sohamkamani.com/golang/json/
// JSON output

type User struct {
	Id           uuid.UUID `json:"-"`
	Name         string    `json:"name" binding:"required"`
	PasswordHash string    `json:"password" binding:"required"`
}

type UserAccount struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
