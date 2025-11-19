package main

import (
	"daily-meal-tracker-backend/internal/database"
	"daily-meal-tracker-backend/internal/handlers"
	"daily-meal-tracker-backend/internal/oss"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
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

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/meals", handlers.GetMeals)
		api.POST("/meals", handlers.CreateMeal)
		api.DELETE("/meals/:id", handlers.DeleteMeal)
	}

	// Handle SPA routing: redirect unknown routes to index.html
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
