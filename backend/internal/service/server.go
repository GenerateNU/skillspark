package service

import (
	"context"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/errs"
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

// Health check types
type HealthOutput struct {
	Body struct {
		Status  string `json:"status" doc:"Health status" example:"ok"`
		Version string `json:"version" doc:"API version" example:"1.0.0"`
	}
}

// Example request/response types for testing
type GreetingInput struct {
	Name string `path:"name" maxLength:"30" doc:"Name to greet"`
}

type GreetingOutput struct {
	Body struct {
		Message   string `json:"message" doc:"Greeting message"`
		Timestamp string `json:"timestamp" doc:"Server timestamp"`
	}
}

type CreateUserInput struct {
	Body struct {
		Email    string `json:"email" format:"email" doc:"User email address"`
		Username string `json:"username" minLength:"3" maxLength:"50" doc:"Username"`
		FullName string `json:"full_name,omitempty" doc:"Full name (optional)"`
	}
}

type UserOutput struct {
	Body struct {
		ID       string `json:"id" doc:"User ID"`
		Email    string `json:"email" doc:"User email"`
		Username string `json:"username" doc:"Username"`
		FullName string `json:"full_name,omitempty" doc:"Full name"`
	}
}

type ListUsersOutput struct {
	Body struct {
		Users []struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"users"`
		Total int `json:"total"`
	}
}

type ErrorOutput struct {
	Body struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
}

// Initialize the App union type containing a fiber app and repository.
func InitApp(config config.Config) *App {
	ctx := context.Background()
	repo := postgres.NewRepository(ctx, config.DB)
	app, humaAPI := SetupApp(config, repo)
	return &App{
		Server: app,
		Repo:   repo,
		API:    humaAPI,
	}
}

// Setup the fiber app with the specified configuration and database.
func SetupApp(config config.Config, repo *storage.Repository) (*fiber.App, huma.API) {
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

	// Documentation routes (Huma provides built-in docs at /docs and /openapi.json)
	setupDocsRoutes(app, "/app/api")

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to SkillSpark!")
	})

	// Register Huma endpoints
	setupHumaRoutes(humaAPI, repo)

	return app, humaAPI
}

// Setup example Huma routes for testing
func setupHumaRoutes(api huma.API, repo *storage.Repository) {
	// Health check endpoint
	huma.Register(api, huma.Operation{
		OperationID: "health-check",
		Method:      http.MethodGet,
		Path:        "/api/v1/health",
		Summary:     "Health check",
		Description: "Check if the API is running and healthy",
		Tags:        []string{"Health"},
	}, func(ctx context.Context, input *struct{}) (*HealthOutput, error) {
		resp := &HealthOutput{}
		resp.Body.Status = "ok"
		resp.Body.Version = "1.0.0"
		return resp, nil
	})

	// Example 1: Simple greeting endpoint with path parameter
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/api/v1/greet/{name}",
		Summary:     "Get a greeting",
		Description: "Returns a personalized greeting message",
		Tags:        []string{"Examples"},
	}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = "Hello, " + input.Name + "!"
		resp.Body.Timestamp = "2024-01-01T00:00:00Z" // Fixed for example
		return resp, nil
	})

	// Example 2: Create user endpoint with request body validation
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/v1/users",
		Summary:     "Create a new user",
		Description: "Creates a new user account with the provided information",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *CreateUserInput) (*UserOutput, error) {
		// Simulate user creation
		resp := &UserOutput{}
		resp.Body.ID = "user_123456"
		resp.Body.Email = input.Body.Email
		resp.Body.Username = input.Body.Username
		resp.Body.FullName = input.Body.FullName
		return resp, nil
	})

	// Example 3: List users endpoint
	huma.Register(api, huma.Operation{
		OperationID: "list-users",
		Method:      http.MethodGet,
		Path:        "/api/v1/users",
		Summary:     "List all users",
		Description: "Returns a list of all registered users",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct{}) (*ListUsersOutput, error) {
		resp := &ListUsersOutput{}
		resp.Body.Users = []struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			{ID: "user_1", Email: "alice@example.com", Username: "alice"},
			{ID: "user_2", Email: "bob@example.com", Username: "bob"},
		}
		resp.Body.Total = len(resp.Body.Users)
		return resp, nil
	})

	// Example 4: Get user by ID with path parameter
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        "/api/v1/users/{id}",
		Summary:     "Get user by ID",
		Description: "Returns a single user by their ID",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"User ID"`
	}) (*UserOutput, error) {
		resp := &UserOutput{}
		resp.Body.ID = input.ID
		resp.Body.Email = "user@example.com"
		resp.Body.Username = "testuser"
		resp.Body.FullName = "Test User"
		return resp, nil
	})

	// Example 5: Delete user endpoint
	huma.Register(api, huma.Operation{
		OperationID: "delete-user",
		Method:      http.MethodDelete,
		Path:        "/api/v1/users/{id}",
		Summary:     "Delete a user",
		Description: "Deletes a user account by ID",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"User ID to delete"`
	}) (*struct{}, error) {
		// Simulate deletion (return empty response with 204 status)
		return &struct{}{}, nil
	})
}
