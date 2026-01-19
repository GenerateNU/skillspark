package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
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

func setupEventTestAPI(eventRepo *repomocks.MockEventRepository) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test Event API", "1.0.0"))
	repo := &storage.Repository{
		Event: eventRepo,
	}
	routes.SetupEventRoutes(api, repo)
	return app, api
}

func TestHumaValidation_CreateEvent(t *testing.T) {
	t.Parallel()

	orgID := uuid.New()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockEventRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"title":               "Junior Robotics",
				"description":         "Introduction to robotics",
				"organization_id":     orgID.String(),
				"age_range_min":       10,
				"age_range_max":       14,
				"category":            []string{"stem", "robotics"},
				"header_image_s3_key": "events/robotics.jpg",
			},
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"CreateEvent",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEventInput"),
				).Return(&models.Event{
					ID:             uuid.New(),
					Title:          "Junior Robotics",
					Description:    "Introduction to robotics",
					OrganizationID: orgID,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing title", // Required field
			payload: map[string]interface{}{
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
			},
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "title too short", // MinLength: 2
			payload: map[string]interface{}{
				"title":           "A",
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
			},
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/events",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_PatchEvent(t *testing.T) {
	t.Parallel()

	validID := uuid.New().String()

	tests := []struct {
		name       string
		eventID    string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockEventRepository)
		statusCode int
	}{
		{
			name:    "valid patch",
			eventID: validID,
			payload: map[string]interface{}{
				"title": "Advanced Robotics",
			},
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"PatchEvent",
					mock.Anything,
					mock.MatchedBy(func(input *models.PatchEventInput) bool {
						return input.ID.String() == validID && input.Body.Title == "Advanced Robotics"
					}),
				).Return(&models.Event{
					ID:    uuid.MustParse(validID),
					Title: "Advanced Robotics",
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "invalid UUID",
			eventID: "not-a-uuid",
			payload: map[string]interface{}{
				"title": "New Title",
			},
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:    "invalid validation in body", // Title too short
			eventID: validID,
			payload: map[string]interface{}{
				"title": "A",
			},
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/events/"+tt.eventID,
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

func TestHumaValidation_DeleteEvent(t *testing.T) {
	t.Parallel()

	validID := uuid.New().String()

	tests := []struct {
		name       string
		eventID    string
		mockSetup  func(*repomocks.MockEventRepository)
		statusCode int
	}{
		{
			name:    "successful delete",
			eventID: validID,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"DeleteEvent",
					mock.Anything,
					uuid.MustParse(validID),
				).Return(&struct{}{}, nil)
			},
			statusCode: http.StatusNoContent, // Assuming 204 for successful delete with no content
		},
		{
			name:       "invalid UUID",
			eventID:    "not-a-uuid",
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/events/"+tt.eventID,
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
