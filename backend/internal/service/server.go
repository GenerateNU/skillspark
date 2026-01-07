package service

import (
	"context"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/service/handler/session"
	"skillspark/internal/storage"
	"skillspark/internal/storage/postgres"

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
}

// Initialize the App union type containing a fiber app and repository.
func InitApp(config config.Config) *App {
	ctx := context.Background()
	repo := postgres.NewRepository(ctx, config.DB)
	app := SetupApp(config, repo)
	return &App{
		Server: app,
		Repo:   repo,
	}
}

// Setup the fiber app with the specified configuration and database.
func SetupApp(config config.Config, repo *storage.Repository) *fiber.App {
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
		AllowOrigins:     "http://localhost:3000,http://localhost:8080",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, X-Request-ID",
	}))

	// Documentation routes
	setupDocsRoutes(app, "/app/api")

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to SkillSpark!")
	})

	// API v1 routes
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Setup API routes
	setupAPIRoutes(apiV1, repo)

	return app
}

// setupAPIRoutes configures all API v1 routes
// this will get moved outside of this file later by @josh-torre
func setupAPIRoutes(router fiber.Router, repo *storage.Repository) {
	sessionHandler := session.NewHandler(repo.Session)
	router.Route("/sessions", func(r fiber.Router) {
		r.Get("/", sessionHandler.GetSessions)
	})
}
