package migrations

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	EventID int
	Event   Event
	GuestID int
	Guest   Guest
}
