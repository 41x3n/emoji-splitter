package config

import (
	"os"

	"41x3n.yama/emoji-splitter/utils"
	"github.com/joho/godotenv"
)

var errorLogger = utils.ErrorLogger

type Config struct {
	// Add any configuration values you need here
	// For example, the port the server should run on
	Port     string
	GIN_MODE string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		errorLogger.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	// Parse configuration values
	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")

	return &Config{
		Port:     port,
		GIN_MODE: ginMode,
	}, nil
}
