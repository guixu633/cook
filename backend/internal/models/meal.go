package models

import (
	"gorm.io/gorm"
)

type Meal struct {
	gorm.Model
	Name     string `json:"name"`
	Date     string `json:"date"`
	ImageURL string `json:"image_url"`
}
