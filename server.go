package main

import (
	"log"

	"github.com/SalMashuri/invitation-online/db"
	"github.com/SalMashuri/invitation-online/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some err occured, Err %s", err)
	}
	db.NewDB()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
