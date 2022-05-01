package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	StartedAt   time.Time
}
