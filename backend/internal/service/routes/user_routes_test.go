package routes_test

import (
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
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
			name:     "username exists",
			username: "jamesw",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "jamesw").Return(true, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:     "username does not exist",
			username: "randomusername",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "randomusername").Return(false, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:     "repository error",
			username: "erroruser",
			mockSetup: func(m *repomocks.MockUserRepository) {
				m.On("GetUserByUsername", mock.Anything, "erroruser").Return(false, &errs.HTTPError{
					Code:    500,
					Message: "internal server error",
				})
			},
			statusCode: http.StatusInternalServerError,
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

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
