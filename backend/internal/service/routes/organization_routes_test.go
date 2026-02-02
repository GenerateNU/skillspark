package routes

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
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

	writer.Close()
	return &body, writer.FormDataContentType()
}

func createOrgTestS3Client(t *testing.T) *s3_client.Client {

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

func setupOrganizationTestAPI(
	organizationRepo *repomocks.MockOrganizationRepository,
	locationRepo *repomocks.MockLocationRepository,
	s3Client *s3_client.Client,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		Organization: organizationRepo,
		Location:     locationRepo,
	}
	SetupOrganizationRoutes(api, repo, s3Client)
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

			s3Client := createOrgTestS3Client(t)
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, s3Client)

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
		name        string
		formFields  map[string]string
		includeFile bool
		mockSetup   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode  int
	}{
		{
			name: "valid payload without location",
			formFields: map[string]string{
				"name":   "Tech Innovations",
				"active": "true",
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				// Handler calls GetLocationByID with zero UUID when location_id is not provided
				l.On(
					"GetLocationByID",
					mock.Anything,
					uuid.UUID{},
				).Return(nil, nil)
				m.On(
					"CreateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.CreateOrganizationInput"),
					mock.Anything,
				).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Innovations",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)

				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationInput"),
					mock.Anything,
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
			formFields: map[string]string{
				"name":        "Tech Innovations",
				"active":      "true",
				"location_id": "10000000-0000-0000-0000-000000000001",
			},
			includeFile: true,
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
					mock.Anything,
				).Return(&models.Organization{
					ID:         uuid.New(),
					Name:       "Tech Innovations",
					Active:     true,
					LocationID: &locationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)

				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationInput"),
					mock.Anything,
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
			formFields: map[string]string{
				"name":   "",
				"active": "true",
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
		{
			name: "name above maximum length",
			formFields: map[string]string{
				"name":   string(make([]byte, 256)),
				"active": "true",
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createOrgTestS3Client(t)
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, s3Client)

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

			s3Client := createOrgTestS3Client(t)
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, s3Client)

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
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		statusCode     int
	}{
		{
			name:           "valid partial update",
			organizationID: orgID.String(),
			formFields: map[string]string{
				"name": "Updated Name",
			},
			includeFile: true,
			mockSetup: func(m *repomocks.MockOrganizationRepository, l *repomocks.MockLocationRepository) {
				// Handler calls GetLocationByID with zero UUID when location_id is not provided
				l.On(
					"GetLocationByID",
					mock.Anything,
					uuid.UUID{},
				).Return(nil, nil)
				m.On(
					"GetOrganizationByID",
					mock.Anything,
					orgID,
				).Return(&models.Organization{
					ID:        orgID,
					Name:      "Old Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				m.On(
					"UpdateOrganization",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateOrganizationInput"),
					mock.Anything,
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
			formFields:     map[string]string{},
			includeFile:    true,
			mockSetup:      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:     http.StatusUnprocessableEntity,
		},
		{
			name:           "name below minimum length",
			organizationID: orgID.String(),
			formFields: map[string]string{
				"name": "",
			},
			includeFile: true,
			mockSetup:   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository) {},
			statusCode:  http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createOrgTestS3Client(t)
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, s3Client)

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

			s3Client := createOrgTestS3Client(t)
			app, _ := setupOrganizationTestAPI(mockOrgRepo, mockLocRepo, s3Client)

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
