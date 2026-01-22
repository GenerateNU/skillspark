package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
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
			name: "missing title",
			payload: map[string]interface{}{
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
			},
			mockSetup:  func(*repomocks.MockEventRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "title too short",
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

func TestHumaValidation_UpdateEvent(t *testing.T) {
	t.Parallel()

	validID := uuid.New().String()
	notFoundID := "00000000-0000-0000-0000-000000000000"

	tests := []struct {
		name       string
		eventID    string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockEventRepository)
		statusCode int
	}{
		{
			name:    "valid update",
			eventID: validID,
			payload: map[string]interface{}{
				"title": "Advanced Robotics",
			},
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"UpdateEvent",
					mock.Anything,
					mock.MatchedBy(func(input *models.UpdateEventInput) bool {
						return input.ID.String() == validID && input.Body.Title != nil && *input.Body.Title == "Advanced Robotics"
					}),
				).Return(&models.Event{
					ID:    uuid.MustParse(validID),
					Title: "Advanced Robotics",
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "event not found",
			eventID: notFoundID,
			payload: map[string]interface{}{
				"title": "Advanced Robotics",
			},
			mockSetup: func(m *repomocks.MockEventRepository) {
				httpErr := errs.NotFound("Event", "id", uuid.MustParse(notFoundID))
				m.On(
					"UpdateEvent",
					mock.Anything,
					mock.MatchedBy(func(input *models.UpdateEventInput) bool {
						return input.ID.String() == notFoundID
					}),
				).Return(nil, &httpErr)
			},
			statusCode: http.StatusNotFound,
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
			name:    "invalid validation in body",
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
	notFoundID := "00000000-0000-0000-0000-000000000000"

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
				).Return(nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "event not found",
			eventID: notFoundID,
			mockSetup: func(m *repomocks.MockEventRepository) {
				httpErr := errs.NotFound("Event", "id", uuid.MustParse(notFoundID))
				m.On(
					"DeleteEvent",
					mock.Anything,
					uuid.MustParse(notFoundID),
				).Return(&httpErr)
			},
			statusCode: http.StatusNotFound,
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

func TestHumaValidation_GetEventOccurrencesByEventId(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")
	start2, _ := time.Parse(time.RFC3339, "2026-02-22 09:00:00+07")
	end2, _ := time.Parse(time.RFC3339, "2026-02-22 11:00:00+07")

	category_arr := []string{"science","technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	event := models.Event{
		ID: 				uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title: 				"Junior Robotics Workshop",
		Description: 		"Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID: 	uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin: 		&eight,
		AgeRangeMax: 		&twelve,
		Category: 			category_arr,
		HeaderImageS3Key: 	&jpg,
		CreatedAt: 			time.Now(),
		UpdatedAt: 			time.Now(),
	}

	location := models.Location{
		ID:           uuid.MustParse("10000000-0000-0000-0000-000000000004"),
		Latitude:     13.7650000,
		Longitude:    100.5380000,
		AddressLine1: "321 Phetchaburi Road",
		AddressLine2: &addr,
		Subdistrict:  "Ratchathewi",
		District:     "Ratchathewi",
		Province:     "Bangkok",
		PostalCode:   "10400",
		Country:      "Thailand",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name       	string
		eventID 	string
		mockSetup  	func(*repomocks.MockEventRepository)
		statusCode 	int
	}{
		{
			name:      	"valid UUID",
			eventID: 	"60000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
					mock.Anything,
					uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				).Return([]models.EventOccurrence{
					{
						ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						ManagerId: 		&mid,
						Event: 			event,
						Location: 		location,
						StartTime: 		start,
						EndTime: 		end,
						MaxAttendees: 	15,
						Language: 		"en",
						CurrEnrolled: 	8,
						CreatedAt:    	time.Now(),
						UpdatedAt:    	time.Now(),
					},
					{
						ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
						ManagerId: 		&mid,
						Event: 			event,
						Location: 		location,
						StartTime: 		start2,
						EndTime: 		end2,
						MaxAttendees: 	15,
						Language: 		"en",
						CurrEnrolled: 	5,
						CreatedAt:    	time.Now(),
						UpdatedAt:    	time.Now(),
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			eventID: 	"not-a-uuid",
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
				http.MethodGet,
				"/api/v1/events/"+tt.eventID+"/event-occurrences/",
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