package migrations

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Photo    string
	Name     string
	Gender   string
	Alamat   string
	Username string
	Password string
}
