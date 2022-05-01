package migrations

import (
	"time"

	"gorm.io/gorm"
)

type EventGuest struct {
	gorm.Model
	GuestID  int
	Guest    Guest
	EventID  int
	Event    Event
	Status   string
	Attend   bool
	CheckIn  time.Time
	Checkout time.Time
}
