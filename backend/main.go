package main

import (
	"fmt"
	"log"
	"server/config"
	"server/internal/database"
	"server/internal/handler"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Initialize Database
	database.InitDB(cfg)
	// Auto Migrate
	if err := database.DB.AutoMigrate(&model.Meal{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 3. Initialize Layers
	mealRepo := repository.NewMealRepository(database.DB)
	mealService := service.NewMealService(mealRepo)
	mealHandler := handler.NewMealHandler(mealService)

	// 4. Setup Router
	r := gin.Default()

	// Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API Routes
	api := r.Group("/api")
	{
		meals := api.Group("/meals")
		{
			meals.POST("", mealHandler.Create)
			meals.GET("", mealHandler.List)
			meals.DELETE("/:id", mealHandler.Delete)
		}
	}

	// 5. Start Server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on port %d...", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
