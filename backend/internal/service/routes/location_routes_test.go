package routes_test

import (
	"bytes"
	"encoding/json"
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

func setupTestAPI(
	locationRepo *repomocks.MockLocationRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Location: locationRepo,
	}

	routes.SetupLocationsRoutes(api, repo)

	return app, api
}

func TestHumaValidation_CreateLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockLocationRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"latitude":  40.7128,
				"longitude": -74.0060,
				"address":   "123 Broadway",
				"city":      "New York",
				"state":     "NY",
				"zip_code":  "10001",
				"country":   "USA",
			},
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"CreateLocation",
					mock.Anything,
					mock.AnythingOfType("*models.CreateLocationInput"),
				).Return(&models.Location{
					ID:        uuid.New(),
					Latitude:  40.7128,
					Longitude: -74.0060,
					Address:   "123 Broadway",
					City:      "New York",
					State:     "NY",
					ZipCode:   "10001",
					Country:   "USA",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing city",
			payload: map[string]interface{}{
				"latitude":  40.7128,
				"longitude": -74.0060,
				"address":   "123 Broadway",
				"state":     "NY",
				"zip_code":  "10001",
				"country":   "USA",
			},
			mockSetup:  func(*repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity, // Huma returns 422 for validation errors
		},
		{
			name: "latitude out of range",
			payload: map[string]interface{}{
				"latitude":  123.456,
				"longitude": -74.0060,
				"address":   "123 Broadway",
				"city":      "New York",
				"state":     "NY",
				"zip_code":  "10001",
				"country":   "USA",
			},
			mockSetup:  func(*repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/locations",
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

func TestHumaValidation_GetLocationByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		locationID string
		mockSetup  func(*repomocks.MockLocationRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			locationID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"GetLocationByID",
					mock.Anything,
					uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				).Return(&models.Location{
					ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					City:      "New York",
					Latitude:  40.7128,
					Longitude: -74.0060,
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			locationID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity, // Huma returns 422
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/locations/"+tt.locationID,
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
