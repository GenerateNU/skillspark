package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"skillspark/internal/config"
	"skillspark/internal/service"

	"github.com/stretchr/testify/assert"
)

func setupTestApp() (*service.App, func()) {
	// Use empty config; repository will be real, but you can swap with mocks if needed
	app := service.InitApp(config.Config{})

	// Return a cleanup function if needed in the future
	return app, func() {}
}

func TestHumaValidation_CreateLocation(t *testing.T) {
	t.Parallel()

	app, cleanup := setupTestApp()
	defer cleanup()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"body": map[string]interface{}{
					"latitude":  40.7128,
					"longitude": -74.0060,
					"address":   "123 Broadway",
					"city":      "New York",
					"state":     "NY",
					"zip_code":  "10001",
					"country":   "USA",
				},
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing required field city",
			payload: map[string]interface{}{
				"body": map[string]interface{}{
					"latitude":  40.7128,
					"longitude": -74.0060,
					"address":   "123 Broadway",
					"state":     "NY",
					"zip_code":  "10001",
					"country":   "USA",
				},
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "latitude out of range",
			payload: map[string]interface{}{
				"body": map[string]interface{}{
					"latitude":  123.456, // invalid
					"longitude": -74.0060,
					"address":   "123 Broadway",
					"city":      "New York",
					"state":     "NY",
					"zip_code":  "10001",
					"country":   "USA",
				},
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "longitude out of range",
			payload: map[string]interface{}{
				"body": map[string]interface{}{
					"latitude":  40.7128,
					"longitude": -200.0, // invalid
					"address":   "123 Broadway",
					"city":      "New York",
					"state":     "NY",
					"zip_code":  "10001",
					"country":   "USA",
				},
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "address too short",
			payload: map[string]interface{}{
				"body": map[string]interface{}{
					"latitude":  40.7128,
					"longitude": -74.0060,
					"address":   "12", // too short
					"city":      "New York",
					"state":     "NY",
					"zip_code":  "10001",
					"country":   "USA",
				},
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt // capture for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/v1/locations", bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Server.Test(req, -1)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}

func TestHumaValidation_GetLocationByID(t *testing.T) {
	t.Parallel()

	app, cleanup := setupTestApp()
	defer cleanup()

	tests := []struct {
		name       string
		locationID string
		statusCode int
	}{
		{
			name:       "valid UUID",
			locationID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			statusCode: http.StatusOK, // passes Huma validation; repo may still return 404
		},
		{
			name:       "invalid UUID",
			locationID: "not-a-uuid",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(http.MethodGet, "/api/v1/locations/"+tt.locationID, nil)
			assert.NoError(t, err)

			resp, err := app.Server.Test(req, -1)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}
