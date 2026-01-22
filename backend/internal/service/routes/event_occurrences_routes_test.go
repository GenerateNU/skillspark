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

func setupEventOccurrencesTestAPI(
	eventOccurrenceRepo *repomocks.MockEventOccurrenceRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		EventOccurrence: eventOccurrenceRepo,
	}
	routes.SetupEventOccurrencesRoutes(api, repo)
	return app, api
}

func TestHumaValidation_GetEventOccurrenceById(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")

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
		name       			string
		eventOccurrenceID 	string
		mockSetup  			func(*repomocks.MockEventOccurrenceRepository)
		statusCode 			int
	}{
		{
			name:       		"valid UUID",
			eventOccurrenceID: 	"70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				).Return(&models.EventOccurrence{
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
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       		"invalid UUID",
			eventOccurrenceID: 	"not-a-uuid",
			mockSetup:  func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       		"event occurrence not found",
			eventOccurrenceID: 	"00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("EventOccurrence", "id", "00000000-0000-0000-0000-000000000000").GetStatus(),
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

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(mockRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/event-occurrences/"+tt.eventOccurrenceID,
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
		mockSetup  	func(*repomocks.MockEventOccurrenceRepository)
		statusCode 	int
	}{
		{
			name:      	"valid UUID",
			eventID: 	"60000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
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
			mockSetup:  func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(mockRepo)

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

func TestHumaValidation_CreateEventOccurrence(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-01T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2026-02-01T01:00:00Z")

	category_arr := []string{"science","technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
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
		ID:           uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
		Latitude:     40.7128,
		Longitude:    -74.0060,
		AddressLine1: "123 Broadway",
		AddressLine2: nil,
		Subdistrict:  "Manhattan",
		District:     "New York",
		Province:     "NY",
		PostalCode:   "10001",
		Country:      "USA",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockEventOccurrenceRepository)
		statusCode int
	}{
		{
			name: "valid payload with null manager id",
			payload: map[string]interface{}{
				"manager_id": nil,
				"event_id": event.ID, 
				"location_id": location.ID, 
				"start_time": start, 
				"end_time": end, 
				"max_attendees": 10, 
				"language": "en",
			},
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"CreateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID: 			uuid.New(),
					ManagerId: 		nil,
					Event: 			event,
					Location: 		location,
					StartTime: 		start,
					EndTime: 		end,
					MaxAttendees: 	10,
					Language: 		"en",
					CurrEnrolled: 	0,
					CreatedAt:    	time.Now(),
					UpdatedAt:    	time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "max attendees below minimum",
			payload: map[string]interface{}{
				"manager_id": nil,
				"event_id": uuid.MustParse("60000000-0000-0000-0000-000000000001"), 
				"location_id": uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), 
				"start_time": start, 
				"end_time": end, 
				"max_attendees": 0, 
				"language": "en",
			},
			mockSetup:  func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "language below minimum length",
			payload: map[string]interface{}{
				"manager_id": nil,
				"event_id": uuid.MustParse("60000000-0000-0000-0000-000000000001"), 
				"location_id": uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"), 
				"start_time": start, 
				"end_time": end, 
				"max_attendees": 10, 
				"language": "e",
			},
			mockSetup:  func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(mockRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/event-occurrences",
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