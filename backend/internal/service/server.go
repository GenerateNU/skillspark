package service

import (
	"context"
	"skillspark/internal/auth"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/s3_client"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	"skillspark/internal/storage/postgres"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	Server *fiber.App
	Repo   *storage.Repository
	API    huma.API
}

// Initialize the App union type containing a fiber app and repository.
func InitApp(config config.Config) (*App, error) {
	ctx := context.Background()
	repo := postgres.NewRepository(ctx, config.DB)
	s3Client, err := s3_client.NewClient(config.S3)
	if err != nil {
		return nil, err
	}
	app, humaAPI := SetupApp(config, repo, s3Client)
	return &App{
		Server: app,
		Repo:   repo,
		API:    humaAPI,
	}, nil
}

// Setup the fiber app with the specified configuration and database.
func SetupApp(config config.Config, repo *storage.Repository, s3Client *s3_client.Client) (*fiber.App, huma.API) {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: errs.ErrorHandler,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:8080,https://cdn.scalar.com,http://127.0.0.1:8080",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, X-Request-ID",
	}))

	// Create Huma API with OpenAPI configuration
	humaConfig := huma.DefaultConfig("SkillSpark API", "1.0.0")
	humaConfig.Info.Description = "API for the SkillSpark application"
	humaConfig.Info.Contact = &huma.Contact{
		Name: "SkillSpark Team",
	}
	humaConfig.Servers = []*huma.Server{
		{URL: "http://localhost:8080", Description: "Local development server"},
	}

	humaAPI := humafiber.New(app, humaConfig)

	if !config.TestMode {
		humaAPI.UseMiddleware(auth.AuthMiddleware(humaAPI, &config.Supabase))
	}

	// Documentation routes (Huma provides built-in docs at /docs and /openapi.json)
	setupDocsRoutes(app, "/app/api")

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to SkillSpark!")
	})

	// Register Huma endpoints
	setupHumaRoutes(humaAPI, repo, config, s3Client)

	return app, humaAPI
}

// Setup Huma routes
func setupHumaRoutes(api huma.API, repo *storage.Repository, config config.Config, s3Client *s3_client.Client) {
	routes.SetupBaseRoutes(api)
	routes.SetupLocationsRoutes(api, repo)
	routes.SetupExamplesRoutes(api, repo)
	routes.SetupOrganizationRoutes(api, repo, s3Client)
	routes.SetupSchoolsRoutes(api, repo)
	routes.SetupEventRoutes(api, repo, s3Client)
	routes.SetupManagerRoutes(api, repo, config)
	routes.SetupRegistrationRoutes(api, repo)
	routes.SetupGuardiansRoutes(api, repo, config)
	routes.SetupChildRoutes(api, repo)
	routes.SetupEventOccurrencesRoutes(api, repo)
	routes.SetupAuthRoutes(api, repo, config)
}
