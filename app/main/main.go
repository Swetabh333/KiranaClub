package main

import (
	"log"
	"os"

	"github.com/Swetabh333/KiranaClub/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)

	}
	db.AutoMigrate(&models.Job{}, &models.Visit{}, &models.Image{})
}
