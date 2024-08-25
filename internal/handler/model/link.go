package model

import "github.com/Aleksandr-qefy/links-api/internal/uuid"

type LinkCreate struct {
	Id          uuid.UUID   `json:"id" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	UserId      uuid.UUID   `json:"userId" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	Ref         string      `json:"ref" binding:"required" example:"\n \n https://github.com/Aleksandr-qefy/ \n\t"`
	Description *string     `json:"description,omitempty" example:"\n \n My Github \n\t"`
	Categories  []uuid.UUID `json:"categories" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
}

type Link struct {
	Id          uuid.UUID  `json:"id" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	Ref         string     `json:"ref" binding:"required" example:"https://mpei.ru/Pages/default.aspx \n "`
	Description *string    `json:"description,omitempty" example:""`
	Categories  []Category `json:"categories"`
}

type AllLinks struct {
	Data []Link `json:"data"`
}

type LinkUpdate struct {
	Id          uuid.UUID    `json:"id" binding:"required" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	Ref         string       `json:"ref,omitempty" example:"https://mpei.ru/Pages/default.aspx \n "`
	Description *string      `json:"description,omitempty" example:""`
	Categories  *[]uuid.UUID `json:"categories,omitempty"`
}
