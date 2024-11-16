package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database ", err)

	}
	log.Println("Connected")
}
