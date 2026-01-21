package routes_test

import (
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
			ID:         schoolID,
			Name:       "Test School",
			LocationID: locationID,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	mockRepo.On("GetAllSchools", mock.Anything, mock.Anything).
		Return(expectedSchools, nil)

	app, _ := setupSchoolsTestAPI(mockRepo)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/schools", nil)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode as direct array (not wrapped)
	var decoded []models.School
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded, 1)
	assert.Equal(t, expectedSchools[0].ID, decoded[0].ID)
	assert.Equal(t, expectedSchools[0].Name, decoded[0].Name)
	assert.Equal(t, expectedSchools[0].LocationID, decoded[0].LocationID)

	mockRepo.AssertExpectations(t)
}

func TestGetAllSchools_WithPagination(t *testing.T) {
	t.Parallel()
	mockRepo := new(repomocks.MockSchoolRepository)
	now := time.Now()

	expectedSchools := []models.School{
		{
			ID:         uuid.New(),
			Name:       "School 1",
			LocationID: uuid.New(),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         uuid.New(),
			Name:       "School 2",
			LocationID: uuid.New(),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	mockRepo.On("GetAllSchools", mock.Anything, mock.Anything).
		Return(expectedSchools, nil)

	app, _ := setupSchoolsTestAPI(mockRepo)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/schools?page=2&limit=5", nil)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode as direct array (not wrapped)
	var decoded []models.School
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded, 2)

	mockRepo.AssertExpectations(t)
}
