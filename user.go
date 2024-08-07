package links_api

import "time"

// https://www.sohamkamani.com/golang/json/
// JSON output

type UUID string

type User struct {
	Id       UUID   `json:"-"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Statistic struct {
	Id        UUID      `json:"id"`
	UserId    UUID      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Activity  string    `json:"activity"`
	Comment   string    `json:"comment,omitempty"`
}
