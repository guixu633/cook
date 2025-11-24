package handler

import (
	"log"
	"net/http"
	"server/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MealHandler struct {
	service service.MealService
}

func NewMealHandler(service service.MealService) *MealHandler {
	return &MealHandler{service: service}
}

type CreateMealRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	ImageURLs   []string  `json:"image_urls"`
}

func (h *MealHandler) Create(c *gin.Context) {
	var req CreateMealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use current time if date is not provided (zero value)
	date := req.Date
	if date.IsZero() {
		date = time.Now()
	}

	log.Printf("Creating meal: %s, Date: %v", req.Name, date)

	meal, err := h.service.CreateMeal(req.Name, req.Description, date, req.ImageURLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meal"})
		return
	}

	c.JSON(http.StatusCreated, meal)
}

func (h *MealHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteMeal(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete meal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}

func (h *MealHandler) List(c *gin.Context) {
	meals, err := h.service.ListMeals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list meals"})
		return
	}

	c.JSON(http.StatusOK, meals)
}
