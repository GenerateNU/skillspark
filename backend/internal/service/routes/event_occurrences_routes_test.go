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
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupEventOccurrencesTestAPI(
	eventOccurrenceRepo *repomocks.MockEventOccurrenceRepository,
	managerRepo *repomocks.MockManagerRepository,
	eventRepo *repomocks.MockEventRepository,
	locationRepo *repomocks.MockLocationRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		EventOccurrence: eventOccurrenceRepo,
		Manager:         managerRepo,
		Event:           eventRepo,
		Location:        locationRepo,
	}
	routes.SetupEventOccurrencesRoutes(api, repo)
	return app, api
}

func TestHumaValidation_GetEventOccurrenceById(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")

	category_arr := []string{"science", "technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	event := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         category_arr,
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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
		name              string
		eventOccurrenceID string
		mockSetup         func(*repomocks.MockEventOccurrenceRepository)
		statusCode        int
	}{
		{
			name:              "valid UUID",
			eventOccurrenceID: "70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				).Return(&models.EventOccurrence{
					ID:           uuid.MustParse("70000000-0000-0000-0000-000000000001"),
					ManagerId:    &mid,
					Event:        event,
					Location:     location,
					StartTime:    start,
					EndTime:      end,
					MaxAttendees: 15,
					Language:     "en",
					CurrEnrolled: 8,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:              "invalid UUID",
			eventOccurrenceID: "not-a-uuid",
			mockSetup:         func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode:        http.StatusUnprocessableEntity,
		},
		{
			name:              "event occurrence not found",
			eventOccurrenceID: "00000000-0000-0000-0000-000000000000",
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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(
				mockRepo,
				mockManagerRepo,
				mockEventRepo,
				mockLocationRepo,
			)

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

func TestHumaValidation_CreateEventOccurrence(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-01T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2026-02-01T01:00:00Z")

	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")

	category_arr := []string{"science", "technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	event := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         category_arr,
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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
				"manager_id":    nil,
				"event_id":      uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				"location_id":   uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				"start_time":    start,
				"end_time":      end,
				"max_attendees": 10,
				"language":      "en",
			},
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"CreateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID:           uuid.New(),
					ManagerId:    nil,
					Event:        event,
					Location:     location,
					StartTime:    start,
					EndTime:      end,
					MaxAttendees: 10,
					Language:     "en",
					CurrEnrolled: 0,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "max attendees below minimum",
			payload: map[string]interface{}{
				"manager_id":    nil,
				"event_id":      uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				"location_id":   uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				"start_time":    start,
				"end_time":      end,
				"max_attendees": 0,
				"language":      "en",
			},
			mockSetup:  func(*repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "language below minimum length",
			payload: map[string]interface{}{
				"manager_id":    nil,
				"event_id":      uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				"location_id":   uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
				"start_time":    start,
				"end_time":      end,
				"max_attendees": 10,
				"language":      "e",
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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(
				mockRepo,
				mockManagerRepo,
				mockEventRepo,
				mockLocationRepo,
			)
			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/event-occurrences",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			if tt.statusCode == http.StatusOK {
				mockManagerRepo.On("GetManagerByID", mock.Anything, mock.Anything).Return(&models.Manager{
					ID:             mid,
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
				}, nil)
				mockEventRepo.On("GetEventByID", mock.Anything, mock.Anything).Return(&event, nil)
				mockLocationRepo.On("GetLocationByID", mock.Anything, mock.Anything).Return(&location, nil)
			}

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

func TestHumaValidation_UpdateEventOccurrence(t *testing.T) {
	t.Parallel()

	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	mid_new := uuid.MustParse("50000000-0000-0000-0000-000000000005")
	eid := uuid.MustParse("60000000-0000-0000-0000-00000000000e")
	lid := uuid.MustParse("10000000-0000-0000-0000-000000000008")
	start, _ := time.Parse(time.RFC3339, "2026-02-22 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-22 11:00:00+07")
	start_new, _ := time.Parse(time.RFC3339, "2026-02-15 10:00:00+07")
	end_new, _ := time.Parse(time.RFC3339, "2026-02-15 12:00:00+07")
	max := 10
	lang := "th"
	curr := 8

	category_arr := []string{"science", "technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	event := models.Event{
		ID:               eid,
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         category_arr,
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	category_arr_new := []string{"technology", "math"}
	ten := 10
	fifteen := 15
	event_new := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-00000000000e"),
		Title:            "Python for Kids",
		Description:      "Introduction to Python programming. Build simple programs and games while learning core concepts.",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000005"),
		AgeRangeMin:      &ten,
		AgeRangeMax:      &fifteen,
		Category:         category_arr_new,
		HeaderImageS3Key: nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	addr := "Suite 15"
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

	location_new := models.Location{
		ID:           lid,
		Latitude:     13.7400000,
		Longitude:    100.5450000,
		AddressLine1: "369 Wireless Road",
		AddressLine2: nil,
		Subdistrict:  "Lumphini",
		District:     "Pathum Wan",
		Province:     "Bangkok",
		PostalCode:   "10330",
		Country:      "Thailand",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name       string
		id         uuid.UUID
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockEventOccurrenceRepository)
		statusCode int
	}{
		{
			name: "valid payload with all fields changed except current enrolled",
			id:   uuid.MustParse("70000000-0000-0000-0000-000000000002"),
			payload: map[string]interface{}{
				"manager_id":    mid_new,
				"event_id":      eid,
				"location_id":   lid,
				"start_time":    start_new,
				"end_time":      end_new,
				"max_attendees": max,
				"language":      lang,
				"curr_enrolled": curr,
			},
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"UpdateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID:           uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId:    &mid_new,
					Event:        event_new,
					Location:     location_new,
					StartTime:    start_new,
					EndTime:      end_new,
					MaxAttendees: 10,
					Language:     "th",
					CurrEnrolled: 8,
					CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "current enrolled below minimum",
			id:   uuid.MustParse("70000000-0000-0000-0000-000000000002"),
			payload: map[string]interface{}{
				"manager_id":    nil,
				"event_id":      nil,
				"location_id":   nil,
				"start_time":    nil,
				"end_time":      nil,
				"max_attendees": nil,
				"language":      nil,
				"curr_enrolled": -1,
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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(
				mockRepo,
				mockManagerRepo,
				mockEventRepo,
				mockLocationRepo,
			)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/event-occurrences/"+tt.id.String(),
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			if tt.statusCode == http.StatusOK {
				mockRepo.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000002"),
				).Return(&models.EventOccurrence{
					ID:           uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId:    &mid,
					Event:        event,
					Location:     location,
					StartTime:    start,
					EndTime:      end,
					MaxAttendees: 15,
					Language:     "en",
					CurrEnrolled: 5,
					CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
				}, nil)
				mockManagerRepo.On("GetManagerByID", mock.Anything, mock.Anything).Return(&models.Manager{
					ID:             mid,
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
				}, nil)
				mockEventRepo.On("GetEventByID", mock.Anything, mock.Anything).Return(&event, nil)
				mockLocationRepo.On("GetLocationByID", mock.Anything, mock.Anything).Return(&location, nil)
			}

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

func TestHumaValidation_GetAllEventOccurrences(t *testing.T) {
	t.Parallel()

	search := "robotics"
	lat := 13.75
	lng := 100.55
	radius := 5.0
	minDuration := 30
	maxDuration := 120
	priceTier := "$$"

	tests := []struct {
		name       string
		query      string
		mockSetup  func(*repomocks.MockEventOccurrenceRepository)
		statusCode int
	}{
		{
			name:  "happy path with all filters",
			query: "?page=1&limit=10&search=robotics&lat=13.75&lng=100.55&radius_km=5&price=$$&min_duration=30&max_duration=120",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetAllEventOccurrences",
					mock.Anything,
					utils.Pagination{Page: 1, Limit: 10},
					models.GetAllEventOccurrencesFilter{
						Search:             &search,
						Latitude:           &lat,
						Longitude:          &lng,
						RadiusKm:           &radius,
						PriceTier:          &priceTier,
						MinDurationMinutes: &minDuration,
						MaxDurationMinutes: &maxDuration,
					},
				).Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "lat/lng/radius incomplete should return 400",
			query:      "?lat=13.75&lng=100.55",
			mockSetup:  func(m *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "min_duration > max_duration should return 400",
			query:      "?min_duration=120&max_duration=30",
			mockSetup:  func(m *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid price tier should return 400",
			query:      "?price=invalid",
			mockSetup:  func(m *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "min_age greater than max_age should return 400",
			query:      "?min_age=18&max_age=10",
			mockSetup:  func(m *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "min_date after max_date should return 400",
			query:      "?min_date=2026-02-10T00:00:00Z&max_date=2026-02-01T00:00:00Z",
			mockSetup:  func(m *repomocks.MockEventOccurrenceRepository) {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:  "only min_age provided should succeed",
			query: "?min_age=18",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetAllEventOccurrences", mock.Anything, mock.Anything, mock.Anything).
					Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "valid date range should succeed",
			query: "?min_date=2026-02-01T00:00:00Z&max_date=2026-02-10T00:00:00Z",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetAllEventOccurrences", mock.Anything, mock.Anything, mock.Anything).
					Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "only max_date provided should succeed",
			query: "?max_date=2026-02-10T00:00:00Z",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetAllEventOccurrences", mock.Anything, mock.Anything, mock.Anything).
					Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "no filters, just default pagination",
			query: "?page=1&limit=5",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetAllEventOccurrences",
					mock.Anything,
					utils.Pagination{Page: 1, Limit: 5},
					models.GetAllEventOccurrencesFilter{},
				).Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)

			tt.mockSetup(mockRepo)

			app, _ := setupEventOccurrencesTestAPI(
				mockRepo,
				mockManagerRepo,
				mockEventRepo,
				mockLocationRepo,
			)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/event-occurrences"+tt.query,
				nil,
			)
			assert.NoError(t, err)

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
