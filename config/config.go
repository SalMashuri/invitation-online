package config

import (
	"fmt"
	"os"
)

func GetDBType() string {
	dbtype := os.Getenv("DB_TYPE")
	return dbtype
}

func GetMYSQLConnectionString() string {
	db_user := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_pass,
		db_host,
		db_port,
		db_name,
	)

	return dataBase
}
