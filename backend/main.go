package main

import (
	"fmt"
	"log"
	"server/config"
	"server/internal/database"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/model"
	"server/internal/pkg/oss"
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

	// 3. Initialize OSS
	ossClient, err := oss.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize OSS client: %v", err)
	}

	// 4. Initialize Layers
	mealRepo := repository.NewMealRepository(database.DB)
	// Update MealService to include ossClient
	mealService := service.NewMealService(mealRepo, ossClient)
	mealHandler := handler.NewMealHandler(mealService)
	uploadHandler := handler.NewUploadHandler(ossClient)

	// 5. Setup Router
	r := gin.Default()

	// Apply Middleware
	r.Use(middleware.Cors())

	// Max upload size 10MB
	r.MaxMultipartMemory = 10 << 20

	// Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API Routes
	api := r.Group("/api")
	{
		api.POST("/upload", uploadHandler.Upload)

		meals := api.Group("/meals")
		{
			meals.POST("", mealHandler.Create)
			meals.GET("", mealHandler.List)
			meals.DELETE("/:id", mealHandler.Delete)
		}
	}

	// 6. Start Server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on port %d...", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
