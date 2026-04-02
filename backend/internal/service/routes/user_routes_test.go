package routes_test

import (
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupUserTestAPI(
	userRepo *repomocks.MockUserRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		User: userRepo,
	}

	routes.SetupUserRoutes(api, repo)

	return app, api
}

func TestHumaValidation_GetUserByUsername(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		username   string
		mockSetup  func(*repomocks.MockUserRepository)
		statusCode int
	}{
		{
			name:     "valid username",
			username: "jamesw",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On(
					"GetUserByUsername",
					mock.Anything,
					"jamesw",
				).Return(&models.User{
					ID:                  uuid.MustParse("b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e"),
					Name:                "James Wilson",
					Email:               "james.wilson@email.com",
					Username:            "jamesw",
					ProfilePictureS3Key: nil,
					LanguagePreference:  "en",
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
					AuthID:              uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:     "user not found",
			username: "randomusername",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On(
					"GetUserByUsername",
					mock.Anything,
					"randomusername",
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("User", "username", "randomusername").GetStatus(),
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

			mockRepo := new(repomocks.MockUserRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupUserTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/user/"+tt.username,
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
