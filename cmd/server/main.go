package main

import (
	"41x3n.yama/emoji-splitter/config"
	"41x3n.yama/emoji-splitter/routes"
	"41x3n.yama/emoji-splitter/utils"
	"github.com/gin-gonic/gin"
)

var infoLogger = utils.InfoLogger
var errorLogger = utils.ErrorLogger

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		errorLogger.Fatalf("Error loading configuration: %v", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.GIN_MODE)

	// Create custom router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start server
	infoLogger.Fatal(router.Run(":" + cfg.Port))
}
