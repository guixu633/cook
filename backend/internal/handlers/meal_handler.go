package handlers

import (
	"daily-meal-tracker-backend/internal/database"
	"daily-meal-tracker-backend/internal/models"
	"daily-meal-tracker-backend/internal/oss"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMeals(c *gin.Context) {
	var meals []models.Meal
	result := database.DB.Order("date desc").Find(&meals)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": meals})
}

func CreateMeal(c *gin.Context) {
	name := c.PostForm("name")
	date := c.PostForm("date")

	file, err := c.FormFile("image")
	var imageURL string
	if err == nil {
		// Upload to OSS
		imageURL, err = oss.UploadToOSS(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image: " + err.Error()})
			return
		}
	}

	meal := models.Meal{
		Name:     name,
		Date:     date,
		ImageURL: imageURL,
	}

	result := database.DB.Create(&meal)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": meal})
}

func DeleteMeal(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Received DELETE request for ID: %s", id)
	var meal models.Meal
	if err := database.DB.First(&meal, id).Error; err != nil {
		log.Printf("Meal not found: %s", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	if err := database.DB.Delete(&meal).Error; err != nil {
		log.Printf("Failed to delete meal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete meal"})
		return
	}

	log.Printf("Meal deleted successfully: %s", id)
	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}
