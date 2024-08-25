package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"time"
)

type Statistic struct {
	Id        uuid.UUID `json:"id" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
	CreatedAt time.Time `json:"createdAt" example:"2024-08-24T15:18:09.055118+03:00"`
	Activity  string    `json:"activity" example:"get_category"`
	Comment   *string   `json:"comment,omitempty" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
}

type AllStatistics struct {
	Data []Statistic `json:"data"`
}
