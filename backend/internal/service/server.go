package service

import (
	"context"
	"net/http"
	"os"
	"skillspark/internal/auth"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/geocoding"
	"skillspark/internal/notification"
	"skillspark/internal/s3_client"
	"skillspark/internal/service/routes"
	"skillspark/internal/sqs_client"
	"skillspark/internal/storage"
	"skillspark/internal/storage/postgres"
	"skillspark/internal/stripeClient"
	translations "skillspark/internal/translation"
	"skillspark/jobs"

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
	Server       *fiber.App
	Repo         *storage.Repository
	StripeClient stripeClient.StripeClientInterface
	API          huma.API
	NotifService *notification.Service
}

// Initialize the App union type containing a fiber app and repository.
func InitApp(config config.Config) (*App, error) {
	ctx := context.Background()
	repo := postgres.NewRepository(ctx, config.DB)
	s3Client, err := s3_client.NewClient(config.S3)

	if err != nil {
		return nil, err
	}

	// Initialize SQS client
	sqsClient, err := sqs_client.NewClient(config.SQS)
	if err != nil {
		return nil, err
	}

	// Initialize notification service and scheduler
	var notifService *notification.Service

	if config.TestMode {
		notifService = notification.NewService(repo, sqsClient)
	}

	c := &http.Client{}
	translateClient := translations.NewClient(c)
	newStripeClient, err := stripeClient.NewStripeClient("")
	if err != nil {
		return nil, err
	}

	jobScheduler := jobs.NewJobScheduler(repo, newStripeClient, *notifService)
	jobScheduler.Start()
	defer jobScheduler.Stop()

	app, humaAPI, err := SetupApp(config, repo, s3Client, translateClient, newStripeClient, *notifService)
	if err != nil {
		return nil, err
	}
	return &App{
		Server:       app,
		Repo:         repo,
		API:          humaAPI,
		NotifService: notifService,
		StripeClient: newStripeClient,
	}, nil
}

// Setup the fiber app with the specified configuration and database.
func SetupApp(config config.Config, repo *storage.Repository, s3Client *s3_client.Client, translateClient *translations.TranslateClient, newStripeClient stripeClient.StripeClientInterface, notifService notification.Service) (*fiber.App, huma.API, error) {
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

	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000,http://localhost:8080,https://cdn.scalar.com,http://127.0.0.1:8080,http://10.0.2.2:8080,http://localhost:5173,http://localhost"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
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

	// Register public routes BEFORE auth middleware
	routes.SetupAuthRoutes(humaAPI, repo, config)
	routes.SetupOrganizationRoutes(humaAPI, repo, s3Client)
	routes.SetupManagerRoutes(humaAPI, repo, config)

	// Apply auth middleware — only affects routes registered after this point
	if !config.TestMode {
		humaAPI.UseMiddleware(auth.AuthMiddleware(humaAPI, &config.Supabase))
	}

	// Documentation routes (Huma provides built-in docs at /docs and /openapi.json)
	setupDocsRoutes(app, "/app/api")

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to SkillSpark!")
	})

	// Register protected Huma endpoints
	if err := setupProtectedHumaRoutes(humaAPI, repo, config, s3Client, translateClient, newStripeClient, notifService); err != nil {
		return nil, nil, err
	}

	routes.SetupWebhookRoutes(app, repo,
		os.Getenv("STRIPE_WEBHOOK_SECRET"),
		os.Getenv("STRIPE_ACCOUNT_WEBHOOK_SECRET"),
	)

	return app, humaAPI, nil
}

// Setup protected Huma routes (behind auth middleware)
func setupProtectedHumaRoutes(api huma.API, repo *storage.Repository, config config.Config, s3Client *s3_client.Client, translateClient *translations.TranslateClient, sc stripeClient.StripeClientInterface, notifService notification.Service) error {
	geocodingClient, err := geocoding.NewClient()
	if err != nil {
		return err
	}
	geocodingService := geocoding.NewService(geocodingClient)

	routes.SetupBaseRoutes(api)
	routes.SetupLocationsRoutes(api, repo, geocodingService)
	routes.SetupOrganizationRoutes(api, repo, s3Client)
	routes.SetupSchoolsRoutes(api, repo)
	routes.SetupEventRoutes(api, repo, s3Client, translateClient)
	routes.SetupManagerRoutes(api, repo, config)
	routes.SetupRegistrationRoutes(api, repo, sc, &notifService)
	routes.SetupGuardiansRoutes(api, repo, sc, config)
	routes.SetupChildRoutes(api, repo)
	routes.SetupEventOccurrencesRoutes(api, repo, s3Client, sc)
	routes.SetUpReviewRoutes(api, repo, translateClient)
	routes.SetupPaymentRoutes(api, repo, sc)
	routes.SetUpSavedRoutes(api, repo)
	routes.SetupGeocodingRoutes(api, geocodingService)
	routes.SetupEmergencyContactRoutes(api, repo)
	routes.SetupRecommendationRoutes(api, repo, s3Client)
	routes.SetupUserRoutes(api, repo)
	return nil
}
