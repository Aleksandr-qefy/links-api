package model

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"time"
)

/*
CREATE TABLE statistics(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    activity VARCHAR(255) NOT NULL,
    comment TEXT
);
*/

type Statistic struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	Activity  string    `db:"activity"`
	Comment   *string   `db:"comment"`
}
