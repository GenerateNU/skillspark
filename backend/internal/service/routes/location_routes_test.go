package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"

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
				"latitude":      40.7128,
				"longitude":     -74.0060,
				"address_line1": "123 Broadway",
				"address_line2": nil,
				"district":      "New York",
				"subdistrict":   "Manhattan",
				"province":      "NY",
				"postal_code":   "10001",
				"country":       "USA",
			},
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"CreateLocation",
					mock.Anything,
					mock.AnythingOfType("*models.CreateLocationInput"),
				).Return(&models.Location{
					ID:           uuid.New(),
					Latitude:     40.7128,
					Longitude:    -74.0060,
					AddressLine1: "123 Broadway",
					AddressLine2: nil,
					Subdistrict:  "Manhattan",
					District:     "New York",
					Province:     "NY",
					PostalCode:   "10001",
					Country:      "USA",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing city",
			payload: map[string]interface{}{
				"latitude":      40.7128,
				"longitude":     -74.0060,
				"address_line1": "123 Broadway",
				"district":      "New York",
				"subdistrict":   "Manhattan",
				"postal_code":   "10001",
				"country":       "USA",
			},
			mockSetup:  func(*repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity, // Huma returns 422 for validation errors
		},
		{
			name: "latitude out of range",
			payload: map[string]interface{}{
				"latitude":      123.456,
				"longitude":     -74.0060,
				"address_line1": "123 Broadway",
				"district":      "New York",
				"subdistrict":   "Manhattan",
				"province":      "NY",
				"postal_code":   "10001",
				"country":       "USA",
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

			// Add this debugging code
			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

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
					District:  "New York",
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
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "location not found",
			locationID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"GetLocationByID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Location", "id", "00000000-0000-0000-0000-000000000000").GetStatus(),
					Message: "Not found",
				})
			},
			statusCode: http.StatusNotFound,
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

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetAllLocations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		queryParams string
		mockSetup   func(*repomocks.MockLocationRepository)
		statusCode  int
	}{
		{
			name:        "valid request default pagination",
			queryParams: "",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"GetAllLocations",
					mock.Anything,
					utils.Pagination{Page: 1, Limit: 100},
				).Return([]models.Location{
					{
						ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
						District:  "New York",
						Latitude:  40.7128,
						Longitude: -74.0060,
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:        "valid request custom pagination",
			queryParams: "?page=2&limit=5",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"GetAllLocations",
					mock.Anything,
					utils.Pagination{Page: 2, Limit: 5},
				).Return([]models.Location{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:        "internal server error",
			queryParams: "",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On(
					"GetAllLocations",
					mock.Anything,
					mock.AnythingOfType("utils.Pagination"),
				).Return(nil, &errs.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				})
			},
			statusCode: http.StatusInternalServerError,
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
				"/api/v1/locations"+tt.queryParams,
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
