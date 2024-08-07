package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Link struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	Ref         string    `json:"ref"`
	Description string    `json:"description,omitempty"`
}

type Category struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type LinkCategory struct {
	LinkId     uuid.UUID `json:"linkId"`
	CategoryId uuid.UUID `json:"categoryId"`
}
