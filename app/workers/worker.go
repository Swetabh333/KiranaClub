package workers

import (
	"image"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Swetabh333/KiranaClub/app/models"
	"gorm.io/gorm"
)

// Function for calculating perimeter and simulating gpu processing
func ProcessJob(jobID string, db *gorm.DB) {
	var results []models.ImageResult
	db.Where("job_id = ?", jobID).Find(&results)
	for _, result := range results {
		resp, err := http.Get(result.ImageURL)
		if err != nil {
			db.Model(&result).Updates(models.ImageResult{Status: "failed", Error: "image download failed"})
			continue
		}
		img, _, err := image.Decode(resp.Body)
		if err != nil {
			resp.Body.Close()
			db.Model(&result).Updates(models.ImageResult{Status: "failed", Error: "image decoding failed"})
			continue
		}
		resp.Body.Close()

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		//Simulating GPU processing
		time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

		perimeter := 2 * float64(width+height)
		db.Model(&result).Updates(models.ImageResult{Perimeter: perimeter, Status: "success"})
	}
	db.Model(&models.Job{}).Where("job_id = ?", jobID).Update("status", "completed")
	log.Printf("Job %s completed\n", jobID)
}
