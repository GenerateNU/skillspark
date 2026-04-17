package routes

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"testing"
	"time"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	s3mocks "skillspark/internal/s3_client/mocks"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	translateMocks "skillspark/internal/translation/mocks"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// dummyImageData for testing organization routes
func dummyImageData() []byte {
	return []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
	}
}

// createMultipartForm creates a multipart form with fields and optionally a file
func createMultipartForm(fields map[string]string, includeFile bool, fileFieldName string) (*bytes.Buffer, string) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	if includeFile {
		// Create form file
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", `form-data; name="`+fileFieldName+`"; filename="test.png"`)
		h.Set("Content-Type", "image/png")
		part, _ := writer.CreatePart(h)
		_, _ = part.Write(dummyImageData())
	}

	err := writer.Close()
	if err != nil {
		return nil, ""
	}
	return &body, writer.FormDataContentType()
}

// createMockS3Client creates a mock S3 client for testing
func createMockS3Client() *s3mocks.S3ClientMock {
	return new(s3mocks.S3ClientMock)
}

func setupOrganizationTestAPI(
	organizationRepo *repomocks.MockOrganizationRepository,
	locationRepo *repomocks.MockLocationRepository,
	reviewRepo *repomocks.MockReviewRepository,
	s3Client s3_client.S3Interface,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		Organization: organizationRepo,
		Location:     locationRepo,
		Review:       reviewRepo,
	}
	translateClient := new(translateMocks.TranslateMock)
	translateClient.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).
		Return(map[string]*string{}, nil).Maybe()
	SetupOrganizationRoutes(api, repo, s3Client, translateClient)
	return app, api
}

var testLocationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")

func TestHumaValidation_GetOrganizationById(t *testing.T) {
	t.Parallel()

	pfpKey := "orgs/babel_street.jpg"

	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository)
		statusCode     int
	}{
		{
			name:           "valid UUID",
			organizationID: "40000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"GetOrganizationByID",
					mock.Anything,
					uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					mock.Anything,
				).Return(&models.Organization{
					ID:         uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Name:       "Science Academy Bangkok",
					Active:     true,
					PfpS3Key:   &pfpKey,
					LocationID: &testLocationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
				r.On(
					"GetAggregateReviewsForOrganization",
					mock.Anything,
					uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				).Return(&models.ReviewAggregate{
					TotalReviews:  5,
					AverageRating: 4.2,
					Breakdown:     []models.ReviewRatingCount{},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "organization not found",
			organizationID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"GetOrganizationByID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					mock.Anything,
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
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo, mockReviewRepo)

			mockS3 := createMockS3Client()
			if tt.statusCode == http.StatusOK {
				mockURL := "https://test-bucket.s3.amazonaws.com/orgs/babel_street.jpg"
				mockS3.On("GeneratePresignedURL", mock.Anything, mock.Anything, mock.Anything).Return(mockURL, nil)
			}
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, mockReviewRepo, mockS3)

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
			mockReviewRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_CreateOrganization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		formFields  map[string]string
		includeFile bool
		mockSetup   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository)
		statusCode  int
	}{
		{
			name: "valid payload with location",
			formFields: map[string]string{
				"name":        "Tech Innovations",
				"active":      "true",
				"location_id": testLocationID.String(),
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				l.On(
					"GetLocationByID",
					mock.Anything,
					testLocationID,
				).Return(&models.Location{
					ID: testLocationID,
				}, nil)
				m.On(
					"CreateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.CreateOrganizationDBInput"),
					mock.Anything,
				).Return(&models.Organization{
					ID:         uuid.New(),
					Name:       "Tech Innovations",
					Active:     true,
					LocationID: &testLocationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)

				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationDBInput"),
					mock.Anything,
				).Return(&models.Organization{
					ID:         uuid.New(),
					Name:       "Tech Innovations",
					Active:     true,
					LocationID: &testLocationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "name below minimum length",
			formFields: map[string]string{
				"name":        "",
				"active":      "true",
				"location_id": testLocationID.String(),
			},
			includeFile: true,
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "name above maximum length",
			formFields: map[string]string{
				"name":        string(make([]byte, 256)),
				"active":      "true",
				"location_id": testLocationID.String(),
			},
			includeFile: true,
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo, mockReviewRepo)

			mockS3 := createMockS3Client()
			if tt.statusCode == http.StatusOK {
				mockURL := "https://test-bucket.s3.amazonaws.com/orgs/test/pfp.jpg"
				mockS3.On("UploadImage", mock.Anything, mock.Anything, mock.Anything).Return(&mockURL, nil)
			}
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, mockReviewRepo, mockS3)

			body, contentType := createMultipartForm(tt.formFields, tt.includeFile, "profile_image")

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/organizations",
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
		mockSetup  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository)
		statusCode int
	}{
		{
			name:  "valid pagination defaults",
			query: "",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"GetAllOrganizations",
					mock.Anything,
					mock.AnythingOfType("utils.Pagination"),
					mock.Anything,
				).Return([]models.Organization{
					{ID: uuid.New(), Name: "Org 1", Active: true, LocationID: &testLocationID},
					{ID: uuid.New(), Name: "Org 2", Active: true, LocationID: &testLocationID},
				}, nil)
				r.On(
					"GetAggregateReviewsForOrganization",
					mock.Anything,
					mock.AnythingOfType("uuid.UUID"),
				).Return(&models.ReviewAggregate{
					TotalReviews:  0,
					AverageRating: 0,
					Breakdown:     []models.ReviewRatingCount{},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "valid pagination with page and page_size",
			query: "?page=2&page_size=5",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"GetAllOrganizations",
					mock.Anything,
					utils.Pagination{Page: 2, Limit: 5},
					mock.Anything,
				).Return([]models.Organization{
					{ID: uuid.New(), Name: "Org 3", Active: true, LocationID: &testLocationID},
				}, nil)
				r.On(
					"GetAggregateReviewsForOrganization",
					mock.Anything,
					mock.AnythingOfType("uuid.UUID"),
				).Return(&models.ReviewAggregate{
					TotalReviews:  0,
					AverageRating: 0,
					Breakdown:     []models.ReviewRatingCount{},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "page below minimum",
			query: "?page=0",
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:  "page_size above maximum",
			query: "?page_size=101",
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo, mockReviewRepo)

			mockS3 := createMockS3Client()
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, mockReviewRepo, mockS3)

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
		formFields     map[string]string
		includeFile    bool
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository)
		statusCode     int
	}{
		{
			name:           "valid partial update",
			organizationID: orgID.String(),
			formFields: map[string]string{
				"name":        "Updated Name",
				"location_id": testLocationID.String(),
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				l.On(
					"GetLocationByID",
					mock.Anything,
					testLocationID,
				).Return(&models.Location{
					ID: testLocationID,
				}, nil)
				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationDBInput"),
					mock.Anything,
				).Return(&models.Organization{
					ID:         orgID,
					Name:       "Updated Name",
					Active:     true,
					LocationID: &testLocationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			formFields:     map[string]string{},
			includeFile:    true,
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "name below minimum length",
			organizationID: orgID.String(),
			formFields: map[string]string{
				"name": "",
			},
			includeFile: true,
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo, mockReviewRepo)

			mockS3 := createMockS3Client()
			if tt.statusCode == http.StatusOK {
				mockURL := "https://test-bucket.s3.amazonaws.com/orgs/test/pfp.jpg"
				mockS3.On("UploadImage", mock.Anything, mock.Anything, mock.Anything).Return(&mockURL, nil)
			}
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, mockReviewRepo, mockS3)

			body, contentType := createMultipartForm(tt.formFields, tt.includeFile, "profile_image")

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/organizations/"+tt.organizationID,
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
			mockOrgRepo.AssertExpectations(t)
			mockLocRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_DeleteOrganization(t *testing.T) {
	t.Parallel()

	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000001")

	tests := []struct {
		name           string
		organizationID string
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository)
		statusCode     int
	}{
		{
			name:           "valid delete",
			organizationID: orgID.String(),
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"DeleteOrganization",
					mock.Anything,
					orgID,
					mock.Anything,
				).Return(&models.Organization{
					ID:         orgID,
					Name:       "Deleted Org",
					Active:     true,
					LocationID: &testLocationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:           "invalid UUID",
			organizationID: "not-a-uuid",
			mockSetup: func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository, *repomocks.MockReviewRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "organization not found",
			organizationID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository, r *repomocks.MockReviewRepository) {
				m.On(
					"DeleteOrganization",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					mock.Anything,
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
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo, mockReviewRepo)

			mockS3 := createMockS3Client()
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, mockReviewRepo, mockS3)

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
		ID:           testLocationID,
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
					mock.Anything,
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
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockRepo)

			mockS3 := createMockS3Client()
			if tt.statusCode == http.StatusOK {
				mockURL := "https://test-bucket.s3.amazonaws.com/events/robotics_workshop.jpg"
				mockS3.On("GeneratePresignedURL", mock.Anything, mock.Anything, mock.Anything).Return(mockURL, nil)
			}
			app, _ := setupOrganizationTestAPI(mockRepo, mockLocationRepo, mockReviewRepo, mockS3)

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

func TestHumaValidation_Organization_InvalidAcceptLanguage(t *testing.T) {
	t.Parallel()

	invalidLangs := []struct {
		name string
		lang string
	}{
		{name: "unsupported locale fr-FR", lang: "fr-FR"},
		{name: "lowercase en-us", lang: "en-us"},
		{name: "random string", lang: "invalid"},
	}

	type endpoint struct {
		label string
		build func() (*http.Request, error)
	}

	orgID := "40000000-0000-0000-0000-000000000001"
	endpoints := []endpoint{
		{
			label: "GetOrganizationById",
			build: func() (*http.Request, error) {
				return http.NewRequest(http.MethodGet, "/api/v1/organizations/"+orgID, nil)
			},
		},
		{
			label: "GetAllOrganizations",
			build: func() (*http.Request, error) {
				return http.NewRequest(http.MethodGet, "/api/v1/organizations", nil)
			},
		},
		{
			label: "DeleteOrganization",
			build: func() (*http.Request, error) {
				return http.NewRequest(http.MethodDelete, "/api/v1/organizations/"+orgID, nil)
			},
		},
		{
			label: "GetEventOccurrencesByOrganizationId",
			build: func() (*http.Request, error) {
				return http.NewRequest(http.MethodGet,
					"/api/v1/organizations/"+orgID+"/event-occurrences/", nil)
			},
		},
		{
			label: "CreateOrganization",
			build: func() (*http.Request, error) {
				body, contentType := createMultipartForm(map[string]string{
					"name":        "Tech Innovations",
					"active":      "true",
					"location_id": testLocationID.String(),
				}, true, "profile_image")
				req, err := http.NewRequest(http.MethodPost, "/api/v1/organizations", body)
				if err != nil {
					return nil, err
				}
				req.Header.Set("Content-Type", contentType)
				return req, nil
			},
		},
		{
			label: "UpdateOrganization",
			build: func() (*http.Request, error) {
				body, contentType := createMultipartForm(map[string]string{
					"name": "Tech Innovations",
				}, true, "profile_image")
				req, err := http.NewRequest(http.MethodPatch, "/api/v1/organizations/"+orgID, body)
				if err != nil {
					return nil, err
				}
				req.Header.Set("Content-Type", contentType)
				return req, nil
			},
		},
	}

	for _, ep := range endpoints {
		ep := ep
		for _, tt := range invalidLangs {
			tt := tt
			t.Run(ep.label+"/"+tt.name, func(t *testing.T) {
				t.Parallel()

				mockRepo := new(repomocks.MockOrganizationRepository)
				mockLocationRepo := new(repomocks.MockLocationRepository)
				mockReviewRepo := new(repomocks.MockReviewRepository)
				mockS3 := new(s3mocks.S3ClientMock)
				app, _ := setupOrganizationTestAPI(mockRepo, mockLocationRepo, mockReviewRepo, mockS3)

				req, err := ep.build()
				assert.NoError(t, err)
				req.Header.Set("Accept-Language", tt.lang)

				resp, err := app.Test(req)
				assert.NoError(t, err)
				defer func() { _ = resp.Body.Close() }()

				assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode,
					"expected 422 for invalid Accept-Language %q on %s", tt.lang, ep.label)
			})
		}
	}
}
