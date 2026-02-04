package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	// "skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/config"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupGuardianTestAPI(
	guardianRepo *repomocks.MockGuardianRepository,
	managerRepo *repomocks.MockManagerRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Guardian: guardianRepo,
		Manager:  managerRepo,
	}

	cfg := config.Config {
		Supabase: config.Supabase{
			URL:            "https://example.supabase.co",
			AnonKey:        "dummy-anon-key",
			ServiceRoleKey: "dummy-service-role-key",
		},
	}

	routes.SetupGuardiansRoutes(api, repo, cfg)

	return app, api
}

func TestHumaValidation_CreateGuardian(t *testing.T) {
	t.Parallel()

	userID := uuid.New()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockGuardianRepository, *repomocks.MockManagerRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"name":                "John Doe",
				"email":               "john@example.com",
				"username":            "johndoe",
				"language_preference": "en",
			},
			mockSetup: func(m *repomocks.MockGuardianRepository, mm *repomocks.MockManagerRepository) {

				m.On("GetGuardianByUserID", mock.Anything, mock.Anything).Return(nil, nil).Maybe()
				m.On(
					"CreateGuardian",
					mock.Anything,
					mock.MatchedBy(func(input *models.CreateGuardianInput) bool {
						return input.Body.Name == "John Doe" && input.Body.Email == "john@example.com"
					}),
				).Return(&models.Guardian{
					ID:                 uuid.New(),
					UserID:             userID,
					Name:               "John Doe",
					Email:              "john@example.com",
					Username:           "johndoe",
					LanguagePreference: "en",
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing required fields",
			payload: map[string]interface{}{
				"username": "johndoe",
			},
			mockSetup:  func(*repomocks.MockGuardianRepository, *repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo, mockManagerRepo)

			app, _ := setupGuardianTestAPI(mockRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/guardians",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
			mockManagerRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetGuardianByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"GetGuardianByID",
					mock.Anything,
					uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				).Return(&models.Guardian{
					ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo, mockManagerRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/guardians/"+tt.guardianID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_UpdateGuardian(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid payload",
			guardianID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			payload: map[string]interface{}{
				"name":                "Jane Doe",
				"email":               "jane@example.com",
				"username":            "janedoe",
				"language_preference": "es",
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"UpdateGuardian",
					mock.Anything,
					mock.MatchedBy(func(input *models.UpdateGuardianInput) bool {
						return input.ID == uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11") &&
							input.Body.Name == "Jane Doe"
					}),
				).Return(&models.Guardian{
					ID:                 uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					UserID:             uuid.New(),
					Name:               "Jane Doe",
					Email:              "jane@example.com",
					Username:           "janedoe",
					LanguagePreference: "es",
					CreatedAt:          time.Now(),
					UpdatedAt:          time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			payload: map[string]interface{}{
				"name": "Jane",
			},
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo, mockManagerRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPut,
				"/api/v1/guardians/"+tt.guardianID,
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_DeleteGuardian(t *testing.T) {
	t.Parallel()

	guardianID := "88888888-8888-8888-8888-888888888888"
	userID := "b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e"

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: guardianID,
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"DeleteGuardian",
					mock.Anything,
					uuid.MustParse(guardianID),
				).Return(&models.Guardian{
					ID: uuid.MustParse(guardianID),
					UserID: uuid.MustParse(userID),
					Name: "James Wilson",
					Email: "james.wilson@email.com",
					Username: "jamesw",
					LanguagePreference: "en",
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.name == "valid UUID" {
				t.Skip("Skipping valid test (cannot mock Supabase client easily for route tests)")
			}

			mockRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo, mockManagerRepo)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/guardians/"+tt.guardianID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetGuardianByChildID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		childID    string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:    "valid UUID",
			childID: "b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"GetGuardianByChildID",
					mock.Anything,
					uuid.MustParse("b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22"),
				).Return(&models.Guardian{
					ID: uuid.New(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			childID:    "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo, mockManagerRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/guardians/child/"+tt.childID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
