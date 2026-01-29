package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/errs"
	"skillspark/internal/models"
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

func setupOrganizationTestAPI(
	organizationRepo *repomocks.MockOrganizationRepository,
	locationRepo *repomocks.MockLocationRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		Organization: organizationRepo,
		Location:     locationRepo,
	}
	SetupOrganizationRoutes(api, repo)
	return app, api
}

func TestHumaValidation_GetOrganizationById(t *testing.T) {
	t.Parallel()

	pfpKey := "orgs/babel_street.jpg"
	locationID := uuid.MustParse("10000000-0000-0000-0000-000000000001")

	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode     int
	}{
		{
			name:           "valid UUID",
			organizationID: "40000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"GetOrganizationByID",
					mock.Anything,
					uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				).Return(&models.Organization{
					ID:         uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Name:       "Science Academy Bangkok",
					Active:     true,
					PfpS3Key:   &pfpKey,
					LocationID: &locationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			mockSetup:      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:     http.StatusUnprocessableEntity,
		},
		{
			name:           "organization not found",
			organizationID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"GetOrganizationByID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Organization", "id", "00000000-0000-0000-0000-000000000000").GetStatus(),
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

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/organizations/"+tt.organizationID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_CreateOrganization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode int
	}{
		{
			name: "valid payload without location",
			payload: map[string]interface{}{
				"name":   "Tech Innovations",
				"active": true,
			},
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"CreateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.CreateOrganizationInput"),
				).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Innovations",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "valid payload with location",
			payload: map[string]interface{}{
				"name":        "Tech Innovations",
				"active":      true,
				"location_id": uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			},
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				locationID := uuid.MustParse("10000000-0000-0000-0000-000000000001")
				l.On(
					"GetLocationByID",
					mock.Anything,
					locationID,
				).Return(&models.Location{
					ID: locationID,
				}, nil)
				m.On(
					"CreateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.CreateOrganizationInput"),
				).Return(&models.Organization{
					ID:         uuid.New(),
					Name:       "Tech Innovations",
					Active:     true,
					LocationID: &locationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "name below minimum length",
			payload: map[string]interface{}{
				"name":   "",
				"active": true,
			},
			mockSetup:  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "name above maximum length",
			payload: map[string]interface{}{
				"name":   string(make([]byte, 256)),
				"active": true,
			},
			mockSetup:  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/organizations",
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
			mockOrgRepo.AssertExpectations(t)
			mockLocRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetAllOrganizations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		query      string
		mockSetup  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode int
	}{
		{
			name:  "valid pagination defaults",
			query: "",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"GetAllOrganizations",
					mock.Anything,
					mock.AnythingOfType("utils.Pagination"),
				).Return([]models.Organization{
					{ID: uuid.New(), Name: "Org 1", Active: true},
					{ID: uuid.New(), Name: "Org 2", Active: true},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "valid pagination with page and page_size",
			query: "?page=2&page_size=5",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"GetAllOrganizations",
					mock.Anything,
					utils.Pagination{Page: 2, Limit: 5},
				).Return([]models.Organization{
					{ID: uuid.New(), Name: "Org 3", Active: true},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "page below minimum",
			query:      "?page=0",
			mockSetup:  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "page_size above maximum",
			query:      "?page_size=101",
			mockSetup:  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/organizations"+tt.query,
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
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_UpdateOrganization(t *testing.T) {
	t.Parallel()

	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000001")

	tests := []struct {
		name           string
		organizationID string
		payload        map[string]interface{}
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode     int
	}{
		{
			name:           "valid partial update",
			organizationID: orgID.String(),
			payload: map[string]interface{}{
				"name": "Updated Name",
			},
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationInput"),
				).Return(&models.Organization{
					ID:        orgID,
					Name:      "Updated Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			payload:        map[string]interface{}{},
			mockSetup:      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:     http.StatusUnprocessableEntity,
		},
		{
			name:           "name below minimum length",
			organizationID: orgID.String(),
			payload: map[string]interface{}{
				"name": "",
			},
			mockSetup:  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/organizations/"+tt.organizationID,
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
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_DeleteOrganization(t *testing.T) {
	t.Parallel()

	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000001")

	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode     int
	}{
		{
			name:           "valid delete",
			organizationID: orgID.String(),
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"DeleteOrganization",
					mock.Anything,
					orgID,
				).Return(&models.Organization{
					ID:        orgID,
					Name:      "Deleted Org",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			mockSetup:      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:     http.StatusUnprocessableEntity,
		},
		{
			name:           "organization not found",
			organizationID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				m.On(
					"DeleteOrganization",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Organization", "id", "00000000-0000-0000-0000-000000000000").GetStatus(),
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

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/organizations/"+tt.organizationID,
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
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetEventOccurrencesByOrganizationId(t *testing.T) {
	t.Parallel()

	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")
	start2, _ := time.Parse(time.RFC3339, "2026-02-22 09:00:00+07")
	end2, _ := time.Parse(time.RFC3339, "2026-02-22 11:00:00+07")

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
		name           string
		organizationID string
		mockSetup      func(*repomocks.MockOrganizationRepository)
		statusCode     int
	}{
		{
			name:           "valid UUID",
			organizationID: "40000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockOrganizationRepository) {
				m.On(
					"GetEventOccurrencesByOrganizationID",
					mock.Anything,
					uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				).Return([]models.EventOccurrence{
					{
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
					},
					{
						ID:           uuid.MustParse("70000000-0000-0000-0000-000000000002"),
						ManagerId:    &mid,
						Event:        event,
						Location:     location,
						StartTime:    start2,
						EndTime:      end2,
						MaxAttendees: 15,
						Language:     "en",
						CurrEnrolled: 5,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			mockSetup:      func(*repomocks.MockOrganizationRepository) {},
			statusCode:     http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockOrganizationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			app, _ := setupOrganizationTestAPI(mockRepo, mockLocationRepo)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/organizations/"+tt.organizationID+"/event-occurrences/",
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
