package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Category struct {
	Id   uuid.UUID `json:"id" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	Name string    `json:"name" binding:"required" example:"IT Category"`
}

type AllCategories struct {
	Data []Category `json:"data"`
}

type CategoryUpdate struct {
	Id   uuid.UUID `json:"id" binding:"required" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	Name string    `json:"name" binding:"required" example:"Golang API"`
}
