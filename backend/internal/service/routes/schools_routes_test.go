package routes_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/errs"
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

func setupSchoolsTestAPI(
	schoolRepo *repomocks.MockSchoolRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test Schools API", "1.0.0"))

	repo := &storage.Repository{
		School: schoolRepo,
	}

	routes.SetupSchoolsRoutes(api, repo)

	return app, api
}

func TestGetAllSchools_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockSchoolRepository)

	now := time.Now()
	schoolID := uuid.New()
	locationID := uuid.New()

	expectedSchools := []models.School{
		{
			ID:   schoolID,
			Name: "Test School",
			Location: models.Location{
				ID:        locationID,
				Latitude:  40.7128,
				Longitude: -74.0060,
				Address:   "123 Main St",
				City:      "New York",
				State:     "NY",
				ZipCode:   "10001",
				Country:   "USA",
				CreatedAt: now,
				UpdatedAt: now,
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	mockRepo.On("GetAllSchools", mock.Anything).Return(expectedSchools, (*errs.HTTPError)(nil))

	app, _ := setupSchoolsTestAPI(mockRepo)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/schools", nil)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded models.GetAllSchoolsOutput
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded.Body, 1)
	assert.Equal(t, expectedSchools[0].ID, decoded.Body[0].ID)
	assert.Equal(t, expectedSchools[0].Name, decoded.Body[0].Name)
	assert.Equal(t, expectedSchools[0].Location.City, decoded.Body[0].Location.City)

	mockRepo.AssertExpectations(t)
}

