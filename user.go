package links_api

import "time"

// https://www.sohamkamani.com/golang/json/
// JSON output

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Statistic struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Action    string    `json:"action"`
	Comment   string    `json:"comment,omitempty"`
}
