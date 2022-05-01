package db

import (
	"log"

	"github.com/SalMashuri/invitation-online/config"
	"github.com/SalMashuri/invitation-online/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMYSQLConnectionString()
	log.Println(conString)
	DB, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(
		&migrations.User{},
		&migrations.Event{},
		&migrations.Comments{},
		&migrations.EventGuest{},
		&migrations.EventPic{},
		&migrations.Guest{},
	)
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
