package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"time"
)

type Statistic struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Activity  string    `json:"activity"`
	Comment   *string   `json:"comment,omitempty"`
}

type AllStatistics struct {
	Data []Statistic `json:"data"`
}
