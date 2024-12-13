package main

import (
	"log"
	"os"

	"github.com/Swetabh333/KiranaClub/app/handlers"
	"github.com/Swetabh333/KiranaClub/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connecting to database
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)

	}
	//Migrations
	db.AutoMigrate(&models.Job{}, &models.ImageResult{})

	//Loading the csv in memory
	err = models.LoadStoreMaster("StoreMasterAssignment.csv")

	if err != nil {
		log.Fatalf("Failed to load Store Master : %v", err)
	}
	log.Println("Loaded")

	//Initializing Gin Router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//Implementing the submit endpoint
	router.POST("/api/submit", func(c *gin.Context) {
		handlers.SubmitJobHandler(c, db)
	})
	//Implementing the status endpoint
	router.GET("/api/status", func(c *gin.Context) {
		handlers.GetJobStatusHandler(c, db)
	})
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
