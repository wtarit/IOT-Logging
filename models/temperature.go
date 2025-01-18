package models

import (
	"time"

	"github.com/google/uuid"
)

type Temperature struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	Temperature float32
}
