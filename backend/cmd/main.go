package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"skillspark/internal/config"
	"skillspark/internal/service"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize application with config
	app, err := service.InitApp(*cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Close database connection when main exits
	defer func() {
		slog.Info("Closing database connection")
		if err := app.Repo.Close(); err != nil {
			slog.Error("failed to close database", "error", err)
		}
	}()

	port := cfg.Application.Port

	// Listen for connections with a goroutine
	go func() {
		if err := app.Server.Listen(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for termination signal (SIGINT or SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")

	// Shutdown server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Server.ShutdownWithContext(ctx); err != nil {
		slog.Error("failed to shutdown server gracefully", "error", err)
	}

	slog.Info("Server shutdown complete")
}

func LoadConfig() (*config.Config, error) {
	environment := os.Getenv("ENVIRONMENT")
	testMode := os.Getenv("TEST_MODE")

	var cfg config.Config
	switch environment {
	case "production":
		// Load configuration from environment variables for production
		err := envconfig.Process(context.Background(), &cfg)
		if err != nil {
			log.Fatalln("Error processing environment variables: ", err)
		}
	case "development":
		log.Println("Loading configuration from environment variables for development")
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

	cfg.TestMode = testMode == "true"

	return &cfg, nil
}
