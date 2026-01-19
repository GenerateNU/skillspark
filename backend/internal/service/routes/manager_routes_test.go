package routes_test

import (
	"net/http"
	"testing"

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
