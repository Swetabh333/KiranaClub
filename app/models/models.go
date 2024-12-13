package models

import (
	"encoding/csv"
	"errors"
	"os"
	"time"
)

type JobStatus string

// all types of job statuses
const (
	JobOngoing   JobStatus = "ongoing"
	JobCompleted JobStatus = "completed"
	JobFailed    JobStatus = "failed"
)

// Model for storing job in database
type Job struct {
	ID        uint      `gorm:"primaryKey"`
	JobID     string    `gorm:"unique;not null"`
	Status    JobStatus `gorm:"not null"`  // "ongoing", "completed", "failed"
	Error     string    `gorm:"type:text"` // JSON-encoded error details
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model for storing image results
type ImageResult struct {
	ID        uint    `gorm:"primaryKey"`
	JobID     string  `gorm:"not null"`
	StoreID   string  `gorm:"not null"`
	ImageURL  string  `gorm:"not null"`
	Perimeter float64 `gorm:"not null"`
	Status    string  `gorm:"not null"` // "success", "failed" ,"pending"
	Error     string  `gorm:"type:text"`
}

// Store master for keeping track of csv file
type StoreMaster struct {
	StoreID   string
	StoreName string
	AreaCode  string
}

// map to keep data of csv rows
var storeMaster = map[string]StoreMaster{}

// function for loading the csv file
func LoadStoreMaster(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] { //skip the header
		storeMaster[record[2]] = StoreMaster{
			StoreID:   record[2],
			StoreName: record[1],
			AreaCode:  record[0],
		}
	}
	return nil
}

// function for validating if a store exits
func ValidateStoreID(storeID string) error {
	if _, exists := storeMaster[storeID]; !exists {
		return errors.New("Invalid store ID:" + storeID)
	}
	return nil
}
