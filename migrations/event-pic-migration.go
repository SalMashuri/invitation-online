package migrations

import (
	"gorm.io/gorm"
)

type EventPic struct {
	gorm.Model
	ImageUrl string
	EventID  int
	Event    Event
}
