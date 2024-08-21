package links_api

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"time"
)

type Statistic struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	CreatedAt time.Time
	Activity  string
	Comment   *string
}
