package organization

import (
	"context"
	"net/http/httptest"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			app.Get("/organizations/:id", func(c *fiber.Ctx) error {
				output, err := handler.GetOrganizationById(c.Context(), &models.GetOrganizationByIDInput{
					ID: uuid.MustParse(c.Params("id")),
				})
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
	tests := []struct {
		name      string
		input     *models.CreateOrganizationInput
		mockSetup func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr   bool
	}{
		{
			name: "successful create without location",
			input: &models.CreateOrganizationInput{
				Body: struct {
					Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name: "Tech Corp",
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.Organization")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "successful create with location",
			input: &models.CreateOrganizationInput{
				Body: struct {
					Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name:       "Tech Corp",
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&models.Location{
					ID: uuid.New(),
				}, nil)
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.Organization")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid location_id",
			input: &models.CreateOrganizationInput{
				Body: struct {
					Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name:       "Tech Corp",
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Location not found").Code,
					Message: "Location not found",
				})
			},
			wantErr: true,
		},
		{
			name: "database error on create",
			input: &models.CreateOrganizationInput{
				Body: struct {
					Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name: "Tech Corp",
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.Organization")).Return(&errs.HTTPError{
					Code:    errs.InternalServerError("Database error").Code,
					Message: "Database error",
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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			output, err := handler.CreateOrganization(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.input.Body.Name, output.Body.Name)
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

	tests := []struct {
		name      string
		input     *models.UpdateOrganizationInput
		mockSetup func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr   bool
	}{
		{
			name: "successful update name only",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: struct {
					Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name: &newName,
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:        existingID,
					Name:      "Old Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.Organization")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "successful update multiple fields",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: struct {
					Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name:   &newName,
					Active: &activeFalse,
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:        existingID,
					Name:      "Old Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.Organization")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "organization not found",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: struct {
					Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					Name: &newName,
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Organization not found").Code,
					Message: "Organization not found",
				})
			},
			wantErr: true,
		},
		{
			name: "invalid location_id on update",
			input: &models.UpdateOrganizationInput{
				ID: existingID,
				Body: struct {
					Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
					Active     *bool      `json:"active,omitempty" doc:"Active status"`
					PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
					LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
				}{
					LocationID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetOrganizationByID", mock.Anything, existingID).Return(&models.Organization{
					ID:        existingID,
					Name:      "Old Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				locRepo.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Location not found").Code,
					Message: "Location not found",
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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			output, err := handler.UpdateOrganization(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				if tt.input.Body.Name != nil {
					assert.Equal(t, *tt.input.Body.Name, output.Body.Name)
				}
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
				orgRepo.On("DeleteOrganization", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "organization not found",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("DeleteOrganization", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&errs.HTTPError{
					Code:    errs.InternalServerError("Organization not found").Code,
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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			output, err := handler.DeleteOrganization(context.TODO(), &models.DeleteOrganizationInput{
				ID: uuid.MustParse(tt.id),
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.id, output.Body.ID)
			}

			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAllOrganizations(t *testing.T) {
	tests := []struct {
		name          string
		input         *models.GetAllOrganizationsInput
		mockSetup     func(*repomocks.MockOrganizationRepository, *repomocks.MockLocationRepository)
		wantErr       bool
		expectedTotal int
	}{
		{
			name: "successful get all with defaults",
			input: &models.GetAllOrganizationsInput{
				Page:     1,
				PageSize: 20,
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgs := []models.Organization{
					{ID: uuid.New(), Name: "Org 1", Active: true},
					{ID: uuid.New(), Name: "Org 2", Active: true},
				}
				orgRepo.On("GetAllOrganizations", mock.Anything, 0, 20).Return(orgs, 2, nil)
			},
			wantErr:       false,
			expectedTotal: 2,
		},
		{
			name: "successful get all with pagination",
			input: &models.GetAllOrganizationsInput{
				Page:     2,
				PageSize: 10,
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgs := []models.Organization{
					{ID: uuid.New(), Name: "Org 3", Active: true},
				}
				orgRepo.On("GetAllOrganizations", mock.Anything, 10, 10).Return(orgs, 11, nil)
			},
			wantErr:       false,
			expectedTotal: 11,
		},
		{
			name: "database error",
			input: &models.GetAllOrganizationsInput{
				Page:     1,
				PageSize: 20,
			},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, locRepo *repomocks.MockLocationRepository) {
				orgRepo.On("GetAllOrganizations", mock.Anything, 0, 20).Return(nil, 0, &errs.HTTPError{
					Code:    errs.InternalServerError("Database error").Code,
					Message: "Database error",
				})
			},
			wantErr:       true,
			expectedTotal: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockLocRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockOrgRepo, mockLocRepo)

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			output, err := handler.GetAllOrganizations(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.expectedTotal, output.Body.TotalCount)
			}

			mockOrgRepo.AssertExpectations(t)
		})
	}
}