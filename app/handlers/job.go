package handlers

import (
	"fmt"
	"net/http"

	"github.com/Swetabh333/KiranaClub/app/models"
	"github.com/Swetabh333/KiranaClub/app/workers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// structure for job request payload
type JobRequest struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURLs []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}

// fucntion for creating a job
func SubmitJobHandler(c *gin.Context, db *gorm.DB) {
	var req JobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate request payload
	if req.Count != len(req.Visits) {
		err := fmt.Sprintf("Count does not match the number of visits count = %d and len = %d", req.Count, len(req.Visits))
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	jobID := uuid.NewString()
	job := models.Job{
		JobID:  jobID,
		Status: models.JobOngoing,
	}
	db.Create(&job)

	for _, visit := range req.Visits {
		if err := models.ValidateStoreID(visit.StoreID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		for _, url := range visit.ImageURLs {
			db.Create(&models.ImageResult{
				JobID:    jobID,
				StoreID:  visit.StoreID,
				ImageURL: url,
				Status:   "pending",
			})
		}
	}

	go workers.ProcessJob(jobID, db)
	c.JSON(http.StatusCreated, gin.H{
		"job_id": jobID,
	})
}

// function for getting id and status of job
func GetJobStatusHandler(c *gin.Context, db *gorm.DB) {
	jobID := c.Query("jobid")
	var job models.Job
	if err := db.First(&job, "job_id = ?", jobID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Job not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": job.Status,
		"job_id": job.JobID,
	})
}
