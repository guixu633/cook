package model

import (
	"time"

	"gorm.io/gorm"
)

type Meal struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"` // User provided date
	// ImageURLs will be stored as a JSON array string in the database
	ImageURLs []string `gorm:"serializer:json" json:"image_urls"`
}
