package models

import (
	"time"
)

type JobStatus string

const (
	JobOngoing   JobStatus = "ongoing"
	JobPending   JobStatus = "pending"
	JobCompleted JobStatus = "completed"
	JobFailed    JobStatus = "failed"
)

type Job struct {
	ID        uint      `gorm:"primaryKey"`
	JobID     string    `gorm:"unique;not null"`
	Status    JobStatus `gorm:"not null"`  // "ongoing", "completed", "failed", "pending"
	Error     string    `gorm:"type:text"` // JSON-encoded error details
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ImageResult struct {
	ID        uint    `gorm:"primaryKey"`
	JobID     string  `gorm:"not null"`
	StoreID   string  `gorm:"not null"`
	ImageURL  string  `gorm:"not null"`
	Perimeter float64 `gorm:"not null"`
	Status    string  `gorm:"not null"` // "success", "failed"
	Error     string  `gorm:"type:text"`
}
