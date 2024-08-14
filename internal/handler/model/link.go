package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Link struct {
	Id          uuid.UUID   `json:"id"`
	UserId      uuid.UUID   `json:"userId"`
	Ref         string      `json:"ref" binding:"required"`
	Description *string     `json:"description,omitempty"`
	Categories  []uuid.UUID `json:"categories"`
}

type LinkUpdate struct {
	Id          uuid.UUID    `json:"id" binding:"required"`
	Ref         string       `json:"ref,omitempty"`
	Description *string      `json:"description,omitempty"`
	Categories  *[]uuid.UUID `json:"categories,omitempty"`
}
