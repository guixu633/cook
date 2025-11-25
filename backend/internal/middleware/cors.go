package middleware

import (
	"server/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(cfg *config.Config) gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	if len(cfg.Cors.AllowOrigins) > 0 {
		config.AllowOrigins = cfg.Cors.AllowOrigins
	} else {
		config.AllowAllOrigins = true
	}

	return cors.New(config)
}

