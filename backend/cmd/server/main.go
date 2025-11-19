package main

import (
	"daily-meal-tracker-backend/internal/database"
	"daily-meal-tracker-backend/internal/handlers"
	"daily-meal-tracker-backend/internal/oss"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.InitDB()
	oss.InitOSS()

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins for simplicity in dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/meals", handlers.GetMeals)
	r.POST("/api/meals", handlers.CreateMeal)
	r.DELETE("/api/meals/:id", handlers.DeleteMeal)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
