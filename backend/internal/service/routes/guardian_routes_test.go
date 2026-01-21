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

func setupGuardianTestAPI(
	guardianRepo *repomocks.MockGuardianRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Guardian: guardianRepo,
	}

	routes.SetupGuardiansRoutes(api, repo)

	return app, api
}

func TestHumaValidation_CreateGuardian(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"user_id": "d1c8d767-c3cf-42e9-848f-15756491e02e",
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByUserID", mock.Anything, uuid.MustParse("d1c8d767-c3cf-42e9-848f-15756491e02e")).Return(nil, nil)
				m.On(
					"CreateGuardian",
					mock.Anything,
					mock.AnythingOfType("*models.CreateGuardianInput"),
				).Return(&models.Guardian{
					ID:          uuid.New(),
					UserID:      uuid.MustParse("d1c8d767-c3cf-42e9-848f-15756491e02e"),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		// {
		// 	name: "missing required fields",
		// 	payload: map[string]interface{}{
		// 		"first_name": "John",
		// 	},
		// 	mockSetup:  func(*repomocks.MockGuardianRepository) {},
		// 	statusCode: http.StatusUnprocessableEntity,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/guardians",
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

func TestHumaValidation_GetGuardianByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"GetGuardianByID",
					mock.Anything,
					uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				).Return(&models.Guardian{
					ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					UserID:    uuid.MustParse("d1c8d767-c3cf-42e9-848f-15756491e02e"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/guardians/"+tt.guardianID,
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

func TestHumaValidation_UpdateGuardian(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid payload",
			guardianID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			payload: map[string]interface{}{
				"user_id": "d1c8d767-c3cf-42e9-848f-15756491e02e",
			},
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"UpdateGuardian",
					mock.Anything,
					mock.MatchedBy(func(input *models.UpdateGuardianInput) bool {
						return input.ID == uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
					}),
				).Return(&models.Guardian{
					ID:          uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					UserID:      uuid.MustParse("d1c8d767-c3cf-42e9-848f-15756491e02e"),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			payload: map[string]interface{}{
				"first_name": "Jane",
			},
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPut,
				"/api/v1/guardians/"+tt.guardianID,
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

func TestHumaValidation_DeleteGuardian(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"DeleteGuardian",
					mock.Anything,
					uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				).Return(&models.Guardian{
					ID: uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/guardians/"+tt.guardianID,
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

func TestHumaValidation_GetGuardianByChildID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		childID    string
		mockSetup  func(*repomocks.MockGuardianRepository)
		statusCode int
	}{
		{
			name:    "valid UUID",
			childID: "b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On(
					"GetGuardianByChildID",
					mock.Anything,
					uuid.MustParse("b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22"),
				).Return(&models.Guardian{
					ID: uuid.New(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			childID:    "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupGuardianTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/guardians/child/"+tt.childID,
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
