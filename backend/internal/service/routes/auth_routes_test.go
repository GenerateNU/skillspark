package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"skillspark/internal/config"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupAuthTestAPI(
	userRepo *repomocks.MockUserRepository,
	guardianRepo *repomocks.MockGuardianRepository,
	managerRepo *repomocks.MockManagerRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		User:     userRepo,
		Guardian: guardianRepo,
		Manager:  managerRepo,
	}

	cfg := config.Config{
		Supabase: config.Supabase{
			URL:            "https://example.supabase.co",
			AnonKey:        "dummy-anon-key",
			ServiceRoleKey: "dummy-service-role-key",
		},
	}

	routes.SetupAuthRoutes(api, repo, cfg)

	return app, api
}

func TestHumaValidation_GuardianSignUp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockUserRepository, *repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"email":      "guardian@example.com",
				"password":   "password123",
				"username":   "guardianuser",
				"name":       "Guardian User",
				"language_preference": "eng", 
				"profile_picture_s3_key": "/guardian123.jpg",
			},
			mockSetup: func(u *repomocks.MockUserRepository, g *repomocks.MockGuardianRepository) {},
			statusCode: http.StatusOK,
		},
		{
			name: "missing required fields (email)",
			payload: map[string]interface{}{
				"password": "password123",
			},
			mockSetup:  func(u *repomocks.MockUserRepository, g *repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid email format",
			payload: map[string]interface{}{
				"email":    "not-an-email",
				"password": "password123",
			},
			mockSetup:  func(u *repomocks.MockUserRepository, g *repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()


			if tt.name == "valid payload" {
				t.Skip("Skipping valid payload test because Supabase client cannot be mocked easily in route tests without refactoring")
			}

			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)

			tt.mockSetup(mockUserRepo, mockGuardianRepo)

			app, _ := setupAuthTestAPI(mockUserRepo, mockGuardianRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/auth/signup/guardian",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}

func TestHumaValidation_ManagerSignUp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		statusCode int
	}{
		{
			name: "missing required fields",
			payload: map[string]interface{}{
				"organization_name": "Test Org",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid email format",
			payload: map[string]interface{}{
				"email": "not-an-email",
				"password": "password123",
				"organization_name": "Test Org",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)

			app, _ := setupAuthTestAPI(mockUserRepo, mockGuardianRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/auth/signup/manager",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}

func TestHumaValidation_GuardianLogin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		statusCode int
	}{
		{
			name: "missing email",
			payload: map[string]interface{}{
				"password": "password123",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing password",
			payload: map[string]interface{}{
				"email": "test@example.com",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)

			app, _ := setupAuthTestAPI(mockUserRepo, mockGuardianRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/auth/login/guardian",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}

func TestHumaValidation_ManagerLogin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		statusCode int
	}{
		{
			name: "missing email",
			payload: map[string]interface{}{
				"password": "password123",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing password",
			payload: map[string]interface{}{
				"email": "test@example.com",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockUserRepo := new(repomocks.MockUserRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)

			app, _ := setupAuthTestAPI(mockUserRepo, mockGuardianRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/auth/login/manager",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}
