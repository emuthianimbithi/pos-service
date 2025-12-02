package main

import (
	"log"

	"github.com/emuthianimbithi/pos-service/internal/api"
	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/internal/container"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize container
	c := container.NewContainer(cfg)

	// Initialize router
	router := api.SetupRouter(c)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
