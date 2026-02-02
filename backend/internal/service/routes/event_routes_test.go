package routes_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"testing"
	"time"

	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// dummyImageData for testing purposes
func dummyImageData() []byte {
	return []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
	}
}

// createEventMultipartForm creates a multipart form with fields and optionally a file
func createEventMultipartForm(fields map[string]string, includeFile bool) (*bytes.Buffer, string) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	if includeFile {

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", `form-data; name="header_image"; filename="test.png"`)
		h.Set("Content-Type", "image/png")
		part, _ := writer.CreatePart(h)
		_, _ = part.Write(dummyImageData())
	}

	writer.Close()
	return &body, writer.FormDataContentType()
}

func createEventTestS3Client(t *testing.T) *s3_client.Client {

	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Logf("Warning: Could not load .env file: %v", err)
	}

	s3Config := config.S3{
		Bucket:    os.Getenv("AWS_S3_BUCKET"),
		Region:    os.Getenv("AWS_REGION"),
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	client, err := s3_client.NewClient(s3Config)
	require.NoError(t, err)
	return client
}

func setupEventTestAPI(eventRepo *repomocks.MockEventRepository, s3Client *s3_client.Client) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test Event API", "1.0.0"))
	repo := &storage.Repository{
		Event: eventRepo,
	}
	routes.SetupEventRoutes(api, repo, s3Client)
	return app, api
}

func TestHumaValidation_CreateEvent(t *testing.T) {
	t.Parallel()

	orgID := uuid.New()

	tests := []struct {
		name        string
		formFields  map[string]string
		includeFile bool
		mockSetup   func(*repomocks.MockEventRepository)
		statusCode  int
	}{
		{
			name: "valid payload",
			formFields: map[string]string{
				"title":           "Junior Robotics",
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
				"age_range_min":   "10",
				"age_range_max":   "14",
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockEventRepository) {
				eventID := uuid.New()
				m.On(
					"CreateEvent",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEventInput"),
					mock.Anything,
				).Return(&models.Event{
					ID:             eventID,
					Title:          "Junior Robotics",
					Description:    "Introduction to robotics",
					OrganizationID: orgID,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)

				m.On(
					"UpdateEvent",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEventInput"),
					mock.Anything,
				).Return(&models.Event{
					ID:             eventID,
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
			formFields: map[string]string{
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockEventRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
		{
			name: "title too short",
			formFields: map[string]string{
				"title":           "A",
				"description":     "Introduction to robotics",
				"organization_id": orgID.String(),
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockEventRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			s3Client := createEventTestS3Client(t)
			app, _ := setupEventTestAPI(mockRepo, s3Client)

			body, contentType := createEventMultipartForm(tt.formFields, tt.includeFile)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/events",
				body,
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", contentType)

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
		name        string
		eventID     string
		formFields  map[string]string
		includeFile bool
		mockSetup   func(*repomocks.MockEventRepository)
		statusCode  int
	}{
		{
			name:    "valid update",
			eventID: validID,
			formFields: map[string]string{
				"title": "Advanced Robotics",
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
					mock.Anything,
					uuid.MustParse(validID),
				).Return([]models.EventOccurrence{
					{ID: uuid.New(), Event: models.Event{ID: uuid.MustParse(validID)}},
				}, nil)
				m.On(
					"UpdateEvent",
					mock.Anything,
					mock.MatchedBy(func(input *models.UpdateEventInput) bool {
						return input.ID.String() == validID && input.Body.Title != nil && *input.Body.Title == "Advanced Robotics"
					}),
					mock.Anything,
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
			formFields: map[string]string{
				"title": "Advanced Robotics",
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
					mock.Anything,
					uuid.MustParse(notFoundID),
				).Return([]models.EventOccurrence{}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "invalid UUID",
			eventID: "not-a-uuid",
			formFields: map[string]string{
				"title": "New Title",
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockEventRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
		{
			name:    "invalid validation in body",
			eventID: validID,
			formFields: map[string]string{
				"title": "A",
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockEventRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			s3Client := createEventTestS3Client(t)
			app, _ := setupEventTestAPI(mockRepo, s3Client)

			body, contentType := createEventMultipartForm(tt.formFields, tt.includeFile)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/events/"+tt.eventID,
				body,
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", contentType)

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

			s3Client := createEventTestS3Client(t)
			app, _ := setupEventTestAPI(mockRepo, s3Client)

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
		name       string
		eventID    string
		mockSetup  func(*repomocks.MockEventRepository)
		statusCode int
	}{
		{
			name:    "valid UUID",
			eventID: "60000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
					mock.Anything,
					uuid.MustParse("60000000-0000-0000-0000-000000000001"),
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

			s3Client := createEventTestS3Client(t)
			app, _ := setupEventTestAPI(mockRepo, s3Client)

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
