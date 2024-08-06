package links_api

import "time"

type Link struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Ref         string    `json:"ref"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
