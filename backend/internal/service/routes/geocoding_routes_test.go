package routes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"skillspark/internal/errs"
	geocodingHandler "skillspark/internal/service/handler/geocoding"
	"skillspark/internal/service/routes"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockGeocoder is a testify mock implementing geocodingHandler.Geocoder.
type mockGeocoder struct {
	mock.Mock
}

func (m *mockGeocoder) Geocode(ctx context.Context, address string) (float64, float64, *errs.HTTPError) {
	args := m.Called(ctx, address)
	lat := args.Get(0).(float64)
	lng := args.Get(1).(float64)
	if args.Get(2) == nil {
		return lat, lng, nil
	}
	return lat, lng, args.Get(2).(*errs.HTTPError)
}

func setupGeocodingTestAPI(svc geocodingHandler.Geocoder) *fiber.App {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	routes.SetupGeocodingRoutesWithGeocoder(api, svc)
	return app
}

func TestGeocodingRoute_PostGeocode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*mockGeocoder)
		statusCode int
	}{
		{
			name:    "valid address returns 200 with coordinates",
			payload: map[string]interface{}{"address": "1 Sukhumvit Rd, Bangkok"},
			mockSetup: func(m *mockGeocoder) {
				m.On("Geocode", mock.Anything, "1 Sukhumvit Rd, Bangkok").
					Return(13.7563, 100.5018, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "invalid address returns 400",
			payload: map[string]interface{}{"address": "zzzzz not a real place"},
			mockSetup: func(m *mockGeocoder) {
				e := errs.BadRequest("address is invalid or could not be geocoded with sufficient confidence")
				m.On("Geocode", mock.Anything, "zzzzz not a real place").
					Return(0.0, 0.0, &e)
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name:    "geocoder failure returns 500",
			payload: map[string]interface{}{"address": "Some Address"},
			mockSetup: func(m *mockGeocoder) {
				e := errs.InternalServerError("geocoding failed: connection refused")
				m.On("Geocode", mock.Anything, "Some Address").
					Return(0.0, 0.0, &e)
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "missing address field returns 422",
			payload:    map[string]interface{}{},
			mockSetup:  func(*mockGeocoder) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "empty address string returns 422",
			payload:    map[string]interface{}{"address": ""},
			mockSetup:  func(*mockGeocoder) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := new(mockGeocoder)
			tt.mockSetup(svc)

			app := setupGeocodingTestAPI(svc)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/v1/geocode", bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			svc.AssertExpectations(t)
		})
	}
}

func TestGeocodingRoute_PostGeocode_ResponseBody(t *testing.T) {
	t.Parallel()

	svc := new(mockGeocoder)
	svc.On("Geocode", mock.Anything, "Nimman Rd, Chiang Mai").
		Return(18.7883, 98.9853, nil)

	app := setupGeocodingTestAPI(svc)

	body, _ := json.Marshal(map[string]interface{}{"address": "Nimman Rd, Chiang Mai"})
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/geocode", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
	assert.Equal(t, 18.7883, result.Latitude)
	assert.Equal(t, 98.9853, result.Longitude)

	svc.AssertExpectations(t)
}
