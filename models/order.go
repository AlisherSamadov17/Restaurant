package models

import (
	"time"
	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID
	Amount string
	User_id uuid.UUID
	Created_at time.Time
}
