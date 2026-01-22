package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupChildTestAPI(
	childRepo *repomocks.MockChildRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Child: childRepo,
	}

	routes.SetupChildRoutes(api, repo)

	return app, api
}

func TestHumaValidation_CreateChild(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockChildRepository)
		statusCode int
	}{
		{
			name: "invalid birth month (too large)",
			payload: map[string]interface{}{
				"name":        "Test Child",
				"school_id":   "20000000-0000-0000-0000-000000000001",
				"birth_month": 13,
				"birth_year":  2019,
				"guardian_id": "11111111-1111-1111-1111-111111111111",
				"interests":   []string{"math"},
			},
			mockSetup:  func(*repomocks.MockChildRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid birth year (future year)",
			payload: map[string]interface{}{
				"name":        "Test Child",
				"school_id":   "20000000-0000-0000-0000-000000000001",
				"birth_month": 5,
				"birth_year":  2040,
				"guardian_id": "11111111-1111-1111-1111-111111111111",
				"interests":   []string{"math"},
			},
			mockSetup:  func(*repomocks.MockChildRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupChildTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/child",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				body, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(body))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
