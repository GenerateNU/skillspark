package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"skillspark/internal/config"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func LoadConfig() (*config.Config, error) {
	environment := os.Getenv("ENVIRONMENT")

	var cfg config.Config
	switch environment {
	case "production":
		// Load configuration from environment variables for production
		err := envconfig.Process(context.Background(), &cfg)
		if err != nil {
			log.Fatalln("Error processing environment variables: ", err)
		}
	case "development":
		// Load configuration from environment variables for development
		err := godotenv.Overload("../.local.env")
		if err != nil {
			log.Fatalln("Error loading .local.env file: ", err)
		}
		err = envconfig.Process(context.Background(), &cfg)
		if err != nil {
			log.Fatalln("Error processing environment variables: ", err)
		}
	default:
		log.Fatalln("Invalid environment name: ", environment, "The environment name must be one of either production or development")
		return nil, fmt.Errorf("invalid environment name: %s", environment)
	}

	return &cfg, nil
}
