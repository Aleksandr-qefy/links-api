package links_api

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type Category struct {
	Id     uuid.UUID
	UserId uuid.UUID
	Name   string
}

type LinkCategory struct {
	LinkId     uuid.UUID
	CategoryId uuid.UUID
}
