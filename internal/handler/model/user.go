package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"time"
)

// https://www.sohamkamani.com/golang/json/
// JSON output

type User struct {
	Id       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserSign struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Statistic struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Activity  string    `json:"activity"`
	Comment   string    `json:"comment,omitempty"`
}
