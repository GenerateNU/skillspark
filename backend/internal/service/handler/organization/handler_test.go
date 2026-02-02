package organization

import (
	"context"
	"net/http/httptest"
	"net/url"
	"os"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// createTestS3Client creates an S3 client for testing using credentials from .env
func createTestS3Client(t *testing.T) *s3_client.Client {
	// Load .env file from backend root
	_ = godotenv.Load("../../../../.env")

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

func TestHandler_GetOrganizationById(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockSetup      func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		expectedStatus int
		wantErr        bool
	}{
		{
			name: "successful get organization",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				pfpKey := "orgs/profile.jpg"
				locationID := uuid.MustParse("b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
				orgRepo.On("GetOrganizationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&models.Organization{
					ID:         uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Name:       "Babel Street",
					Active:     true,
					PfpS3Key:   &pfpKey,
					LocationID: &locationID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},
			expectedStatus: 200,
			wantErr:        false,
		},
		{
			name: "organization not found",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Internal server error").Code,
					Message: "Internal server error",
				})
			},
			expectedStatus: 500,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New(fiber.Config{
				ErrorHandler: errs.ErrorHandler,
			})
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createTestS3Client(t)
			handler := NewHandler(mockOrgRepo, mockLocRepo, s3Client)
			app.Get("/organizations/:id", func(c *fiber.Ctx) error {
				output, err := handler.GetOrganizationById(c.Context(), &models.GetOrganizationByIDInput{
					ID: uuid.MustParse(c.Params("id")),
				}, s3Client)
				if err != nil {
					return err
				}
				return c.Status(fiber.StatusOK).JSON(output)
			})

			req := httptest.NewRequest("GET", "/organizations/"+tt.id, nil)
			res, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateOrganization(t *testing.T) {
	// Dummy image data (PNG header bytes)
	dummyImageData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
	}
	pfpKey := "orgs/a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11/pfp.jpg"

	tests := []struct {
		name       string
		input      *models.CreateOrganizationInput
		updateBody *models.UpdateOrganizationBody
		imageData  *[]byte
		mockSetup  func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr    bool
		wantURL    bool
	}{
		{
			name: "successful create without location",
			input: &models.CreateOrganizationInput{
				Body: models.CreateOrganizationBody{
					Name: "Tech Corp",
				},
			},
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Corp",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful create with location",
			input: &models.CreateOrganizationInput{
				Body: models.CreateOrganizationBody{
					Name:       "Tech Corp",
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&models.Location{
					ID: uuid.New(),
				}, nil)
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Corp",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful create with image - returns valid presigned URL",
			input: &models.CreateOrganizationInput{
				Body: models.CreateOrganizationBody{
					Name: "Tech Corp",
				},
			},
			updateBody: &models.UpdateOrganizationBody{},
			imageData:  &dummyImageData,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        orgID,
					Name:      "Tech Corp",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        orgID,
					Name:      "Tech Corp",
					Active:    true,
					PfpS3Key:  &pfpKey,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: true,
		},
		{
			name: "invalid location_id",
			input: &models.CreateOrganizationInput{
				Body: models.CreateOrganizationBody{
					Name:       "Tech Corp",
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Location not found").Code,
					Message: "Location not found",
				})
			},
			wantErr: true,
			wantURL: false,
		},
		{
			name: "database error on create",
			input: &models.CreateOrganizationInput{
				Body: models.CreateOrganizationBody{
					Name: "Tech Corp",
				},
			},
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput"), mock.Anything).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Database error").Code,
					Message: "Database error",
				})
			},
			wantErr: true,
			wantURL: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createTestS3Client(t)
			handler := NewHandler(mockOrgRepo, mockLocRepo, s3Client)
			output, err := handler.CreateOrganization(context.TODO(), tt.input, tt.updateBody, tt.imageData, s3Client)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.input.Body.Name, output.Name)
			}

			if tt.wantURL {
				require.NotNil(t, output.PresignedURL, "expected presigned URL to be returned")
				// Validate URL structure
				parsedURL, parseErr := url.Parse(*output.PresignedURL)
				require.NoError(t, parseErr, "presigned URL should be valid")
				assert.True(t, strings.HasPrefix(parsedURL.Scheme, "http"), "URL should have http/https scheme")
				assert.Contains(t, parsedURL.Host, "amazonaws.com", "URL should be an AWS S3 URL")
				assert.NotEmpty(t, parsedURL.RawQuery, "presigned URL should have query parameters")
			} else if !tt.wantErr {
				assert.Nil(t, output.PresignedURL, "expected no presigned URL when no image data provided")
			}

			mockOrgRepo.AssertExpectations(t)
			mockLocRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateOrganization(t *testing.T) {
	existingID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	newName := "Updated Name"
	activeFalse := false
	pfpKey := "orgs/a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11/pfp.jpg"

	// Dummy image data (PNG header)
	dummyImageData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
	}

	tests := []struct {
		name      string
		input     *models.UpdateOrganizationInput
		imageData *[]byte
		mockSetup func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr   bool
		wantURL   bool
	}{
		{
			name: "successful update name only",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					Name: &newName,
				},
			},
			imageData: nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:       existingID,
					Name:     "Old Name",
					Active:   true,
					PfpS3Key: nil,
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful update multiple fields",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					Name:   &newName,
					Active: &activeFalse,
				},
			},
			imageData: nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:       existingID,
					Name:     "Old Name",
					Active:   true,
					PfpS3Key: nil,
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful update with image - returns valid presigned URL",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					Name: &newName,
				},
			},
			imageData: &dummyImageData,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:       existingID,
					Name:     "Old Name",
					Active:   true,
					PfpS3Key: nil,
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    true,
					PfpS3Key:  &pfpKey,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: true,
		},
		{
			name: "successful update with image - existing key reused",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					Name: &newName,
				},
			},
			imageData: &dummyImageData,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:       existingID,
					Name:     "Old Name",
					Active:   true,
					PfpS3Key: &pfpKey, // existing key
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput"), mock.Anything).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    true,
					PfpS3Key:  &pfpKey,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
			wantURL: true,
		},
		{
			name: "organization not found on get",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					Name: &newName,
				},
			},
			imageData: nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Organization", "id", existingID.String()).Code,
					Message: "Organization not found",
				})
			},
			wantErr: true,
			wantURL: false,
		},
		{
			name: "invalid location_id on update",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: models.UpdateOrganizationBody{
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			imageData: nil,
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Location not found").Code,
					Message: "Location not found",
				})
			},
			wantErr: true,
			wantURL: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createTestS3Client(t)
			handler := NewHandler(mockOrgRepo, mockLocRepo, s3Client)
			output, err := handler.UpdateOrganization(context.TODO(), tt.input, tt.imageData, s3Client)

			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, output)
			} else {
				require.NoError(t, err)
				require.NotNil(t, output)
				if tt.input.Body.Name != nil {
					assert.Equal(t, *tt.input.Body.Name, output.Name)
				}
			}

			if tt.wantURL {
				require.NotNil(t, output.PresignedURL, "expected presigned URL to be returned")
				parsedURL, parseErr := url.Parse(*output.PresignedURL)
				require.NoError(t, parseErr, "presigned URL should be valid")
				assert.True(t, strings.HasPrefix(parsedURL.Scheme, "http"), "URL should have http/https scheme")
				assert.Contains(t, parsedURL.Host, "amazonaws.com", "URL should be an AWS S3 URL")
				assert.NotEmpty(t, parsedURL.RawQuery, "presigned URL should have query parameters")
			} else if !tt.wantErr {
				assert.Nil(t, output.PresignedURL, "expected no presigned URL when no image data provided")
			}

			mockOrgRepo.AssertExpectations(t)
			mockLocRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteOrganization(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("DeleteOrganization", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&models.Organization{
					ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Name:      "Deleted Org",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "organization not found",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("DeleteOrganization", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Organization", "id", "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19").Code,
					Message: "Organization not found",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createTestS3Client(t)
			handler := NewHandler(mockOrgRepo, mockLocRepo, s3Client)
			output, err := handler.DeleteOrganization(context.TODO(), &models.DeleteOrganizationInput{
				ID: uuid.MustParse(tt.id),
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.id, output.Body.ID.String())
			}

			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAllOrganizations(t *testing.T) {
	tests := []struct {
		name        string
		pagination  utils.Pagination
		mockSetup   func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr     bool
		expectedLen int
	}{
		{
			name:       "successful get all with defaults",
			pagination: utils.Pagination{Page: 1, Limit: 20},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgs := []models.Organization{
					{ID: uuid.New(), Name: "Org 1", Active: true},
					{ID: uuid.New(), Name: "Org 2", Active: true},
				}
				orgRepo.On("GetAllOrganizations", mock.Anything, mock.AnythingOfType("utils.Pagination")).Return(orgs, nil)
			},
			wantErr:     false,
			expectedLen: 2,
		},
		{
			name:       "successful get all with pagination",
			pagination: utils.Pagination{Page: 2, Limit: 10},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgs := []models.Organization{
					{ID: uuid.New(), Name: "Org 3", Active: true},
				}
				orgRepo.On("GetAllOrganizations", mock.Anything, mock.AnythingOfType("utils.Pagination")).Return(orgs, nil)
			},
			wantErr:     false,
			expectedLen: 1,
		},
		{
			name:       "database error",
			pagination: utils.Pagination{Page: 1, Limit: 20},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetAllOrganizations", mock.Anything, mock.AnythingOfType("utils.Pagination")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Database error").Code,
					Message: "Database error",
				})
			},
			wantErr:     true,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			s3Client := createTestS3Client(t)
			handler := NewHandler(mockOrgRepo, mockLocRepo, s3Client)
			output, err := handler.GetAllOrganizations(context.TODO(), tt.pagination, s3Client)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.expectedLen, len(output))
			}

			mockOrgRepo.AssertExpectations(t)
		})
	}
}
