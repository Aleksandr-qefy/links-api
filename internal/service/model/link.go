package links_api

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type LinkUpdate struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Ref         string
	Description *string
	Categories  []uuid.UUID
}

type Link struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Ref         string
	Description *string
	Categories  []Category
}
