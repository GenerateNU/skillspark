package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupmanagerTestAPI(
	managerRepo *repomocks.MockManagerRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Manager: managerRepo,
	}

	routes.SetupManagerRoutes(api, repo)

	return app, api
}

func TestHumaValidation_GetManagerByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		ID         string
		mockSetup  func(*repomocks.MockManagerRepository)
		statusCode int
	}{
		{
			name: "valid UUID",
			ID:   "50000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"GetManagerByID",
					mock.Anything,
					uuid.MustParse("50000000-0000-0000-0000-000000000001"),
				).Return(&models.Manager{
					ID:             uuid.MustParse("50000000-0000-0000-0000-000000000001"),
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			ID:         "not-a-uuid",
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity, // Huma returns 422
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupmanagerTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/manager/"+tt.ID,
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

func TestHumaValidation_GetManagerByOrgID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		OrganizationID string
		mockSetup      func(*repomocks.MockManagerRepository)
		statusCode     int
	}{
		{
			name:           "valid UUID",
			OrganizationID: "40000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"GetManagerByOrgID",
					mock.Anything,
					uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				).Return(&models.Manager{
					ID:             uuid.MustParse("50000000-0000-0000-0000-000000000001"),
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			OrganizationID: "not-a-uuid",
			mockSetup:      func(*repomocks.MockManagerRepository) {},
			statusCode:     http.StatusUnprocessableEntity, // Huma returns 422
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupmanagerTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/manager/org/"+tt.OrganizationID,
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

func TestHumaValidation_CreateManager(t *testing.T) {
	t.Parallel()

	orgID := "40000000-0000-0000-0000-000000000006"

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockManagerRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"name":                "Alice Smith",
				"email":               "alice@org.com",
				"username":            "alices",
				"language_preference": "en",
				"organization_id":     orgID,
				"role":                "Assistant Director",
			},
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"CreateManager",
					mock.Anything,
					mock.MatchedBy(func(input *models.CreateManagerInput) bool {
						return input.Body.Name == "Alice Smith" && *input.Body.OrganizationID == uuid.MustParse(orgID)
					}),
				).Return(&models.Manager{
					ID:             uuid.New(),
					UserID:         uuid.New(),
					Name:           "Alice Smith",
					Email:          "alice@org.com",
					Username:       "alices",
					OrganizationID: uuid.MustParse(orgID),
					Role:           "Assistant Director",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing required fields (name)",
			payload: map[string]interface{}{
				"organization_id": orgID,
				"role":            "Assistant Director",
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing role",
			payload: map[string]interface{}{
				"name":            "Alice",
				"organization_id": orgID,
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupmanagerTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/manager",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_PatchManager(t *testing.T) {
	t.Parallel()

	orgID := "40000000-0000-0000-0000-000000000006"
	managerID := "50000000-0000-0000-0000-000000000001"

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockManagerRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"id":                  managerID,
				"name":                "Alice Updated",
				"email":               "alice@org.com",
				"username":            "aliceu",
				"language_preference": "en",
				"organization_id":     orgID,
				"role":                "Senior Director",
			},
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"PatchManager",
					mock.Anything,
					mock.MatchedBy(func(input *models.PatchManagerInput) bool {
						return input.Body.ID == uuid.MustParse(managerID) && input.Body.Name == "Alice Updated"
					}),
				).Return(&models.Manager{
					ID:             uuid.MustParse(managerID),
					UserID:         uuid.New(),
					Name:           "Alice Updated",
					OrganizationID: uuid.MustParse(orgID),
					Role:           "Senior Director",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "valid payload without organization",
			payload: map[string]interface{}{
				"id":                  managerID,
				"name":                "Alice",
				"role":                "Manager",
				"email":               "a@a.com",
				"username":            "aa",
				"language_preference": "en",
			},
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"PatchManager",
					mock.Anything,
					mock.AnythingOfType("*models.PatchManagerInput"),
				).Return(&models.Manager{
					ID:             uuid.MustParse(managerID),
					UserID:         uuid.New(),
					OrganizationID: uuid.Nil,
					Role:           "Manager",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing id",
			payload: map[string]interface{}{
				"name":            "Alice",
				"organization_id": orgID,
				"role":            "Director",
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing required field (name)",
			payload: map[string]interface{}{
				"id":              managerID,
				"organization_id": orgID,
				"role":            "Director",
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing role",
			payload: map[string]interface{}{
				"id":              managerID,
				"name":            "Alice",
				"organization_id": orgID,
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid id format",
			payload: map[string]interface{}{
				"id":   "not-a-valid-uuid",
				"name": "Alice",
				"role": "Director",
			},
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupmanagerTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/manager",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_DeleteManager(t *testing.T) {
	t.Parallel()

	managerID := "50000000-0000-0000-0000-000000000001"
	orgID := "40000000-0000-0000-0000-000000000006"

	tests := []struct {
		name       string
		id         string
		mockSetup  func(*repomocks.MockManagerRepository)
		statusCode int
	}{
		{
			name: "valid id",
			id:   managerID,
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On(
					"DeleteManager",
					mock.Anything,
					uuid.MustParse(managerID),
				).Return(&models.Manager{
					ID:             uuid.MustParse(managerID),
					UserID:         uuid.New(),
					OrganizationID: uuid.MustParse(orgID),
					Role:           "Director",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid id format",
			id:         "not-a-valid-uuid",
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "empty id",
			id:         "",
			mockSetup:  func(*repomocks.MockManagerRepository) {},
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupmanagerTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/manager/"+tt.id,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
