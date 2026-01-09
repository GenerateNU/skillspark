package main

import (
	"fmt"
	"os"
	"path/filepath"
	"skillspark/internal/config"
	"skillspark/internal/service"

	"gopkg.in/yaml.v3"
)

func main() {
	// Load config (or use defaults)
	cfg := config.Config{
		// Add minimal config needed for service setup
	}

	// Initialize app to get Huma API
	_, humaAPI := service.SetupApp(cfg, nil)

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
