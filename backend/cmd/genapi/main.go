package main

import (
	"fmt"
	"os"
	"path/filepath"
	"skillspark/internal/config"
	notifications "skillspark/internal/notification"
	"skillspark/internal/s3_client"
	"skillspark/internal/service"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
	translations "skillspark/internal/translation"

	"gopkg.in/yaml.v3"
)

func main() {
	// Load config (or use defaults)
	cfg := config.Config{
		// Add minimal config needed for service setup
	}

	// Create an empty repository for API generation
	// The handlers won't be called, so nil fields are fine
	repo := &storage.Repository{}

	s3Client, err := s3_client.NewClient(cfg.S3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create S3 Client: %v\n", err)
	}

	notificationsService := notifications.NewService(nil, nil)
	translateClient := translations.NewClient(nil)
	newStripeClient, err := stripeClient.NewStripeClient("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Stripe Client: %v\n", err)
	}

	// genapi only registers routes to produce the OpenAPI spec — no real API calls
	// are ever made, so a placeholder key is sufficient.
	if os.Getenv("OPENCAGE_API_KEY") == "" {
		os.Setenv("OPENCAGE_API_KEY", "genapi-placeholder")
	}

	// Initialize app to get Huma API
	_, humaAPI, err := service.SetupApp(cfg, repo, s3Client, translateClient, newStripeClient, *notificationsService)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup app: %v\n", err)
		os.Exit(1)
	}

	// Get OpenAPI spec
	openAPI := humaAPI.OpenAPI()

	// Create api directory if it doesn't exist
	apiDir := "api"
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create api directory: %v\n", err)
		os.Exit(1)
	}

	// Write YAML file
	yamlPath := filepath.Join(apiDir, "openapi.yaml")
	yamlFile, err := os.Create(yamlPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create YAML file: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := yamlFile.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close YAML file: %v\n", err)
		}
	}()

	encoder := yaml.NewEncoder(yamlFile)
	encoder.SetIndent(2)
	if err := encoder.Encode(openAPI); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode OpenAPI spec: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ OpenAPI spec generated: %s\n", yamlPath)
}
