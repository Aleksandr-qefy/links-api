package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Category struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AllCategories struct {
	Data []Category `json:"data"`
}

type CategoryUpdate struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}

type LinkCategory struct {
	LinkId     uuid.UUID `json:"linkId"`
	CategoryId uuid.UUID `json:"categoryId"`
}
