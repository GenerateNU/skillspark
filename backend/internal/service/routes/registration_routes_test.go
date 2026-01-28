package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/models"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRegistrationTestAPI(
	registrationRepo *repomocks.MockRegistrationRepository,
	childRepo *repomocks.MockChildRepository,
	guardianRepo *repomocks.MockGuardianRepository,
	eventOccurrenceRepo *repomocks.MockEventOccurrenceRepository,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Registration:    registrationRepo,
		Child:           childRepo,
		Guardian:        guardianRepo,
		EventOccurrence: eventOccurrenceRepo,
	}

	SetupRegistrationRoutes(api, repo)

	return app, api
}

func TestHumaValidation_GetRegistrationByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		ID         string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name: "valid UUID",
			ID:   "80000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				m.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
				).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID:                  uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			ID:         "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/"+tt.ID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByChildID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		childID    string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name:    "valid UUID",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByChildIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                  uuid.New(),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}
				m.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			childID:    "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/child/"+tt.childID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByGuardianID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByGuardianIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                  uuid.New(),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}
				m.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/guardian/"+tt.guardianID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByEventOccurrenceID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		eventOccurrenceID string
		mockSetup         func(*repomocks.MockRegistrationRepository)
		statusCode        int
	}{
		{
			name:              "valid UUID",
			eventOccurrenceID: "70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                  uuid.New(),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}
				m.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:              "invalid UUID",
			eventOccurrenceID: "not-a-uuid",
			mockSetup:         func(*repomocks.MockRegistrationRepository) {},
			statusCode:        http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/event_occurrence/"+tt.eventOccurrenceID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_CreateRegistration(t *testing.T) {
	t.Parallel()

	childID := "30000000-0000-0000-0000-000000000001"
	guardianID := "11111111-1111-1111-1111-111111111111"
	eventOccurrenceID := "70000000-0000-0000-0000-000000000001"

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"status":              "registered",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse(eventOccurrenceID)).Return(&models.EventOccurrence{
					ID: uuid.MustParse(eventOccurrenceID),
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, uuid.MustParse(childID)).Return(&models.Child{
					ID: uuid.MustParse(childID),
				}, nil)
				guardianRepo.On("GetGuardianByID", mock.Anything, uuid.MustParse(guardianID)).Return(&models.Guardian{
					ID: uuid.MustParse(guardianID),
				}, nil)
				regRepo.On(
					"CreateRegistration",
					mock.Anything,
					mock.AnythingOfType("*models.CreateRegistrationInput"),
				).Return(&models.CreateRegistrationOutput{
					Body: models.Registration{
						ID:                  uuid.New(),
						ChildID:             uuid.MustParse(childID),
						GuardianID:          uuid.MustParse(guardianID),
						EventOccurrenceID:   uuid.MustParse(eventOccurrenceID),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing child_id",
			payload: map[string]interface{}{
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid child_id format",
			payload: map[string]interface{}{
				"child_id":            "not-a-uuid",
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/registrations",
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
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_UpdateRegistration(t *testing.T) {
	t.Parallel()

	registrationID := "80000000-0000-0000-0000-000000000001"

	tests := []struct {
		name       string
		id         string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			id:   registrationID,
			payload: map[string]interface{}{
				"status": "cancelled",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On(
					"UpdateRegistration",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateRegistrationInput"),
				).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                  uuid.MustParse(registrationID),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusCancelled,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "invalid id format",
			id:   "not-a-uuid",
			payload: map[string]interface{}{
				"status": "cancelled",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/registrations/"+tt.id,
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
			mockRegRepo.AssertExpectations(t)
		})
	}
}