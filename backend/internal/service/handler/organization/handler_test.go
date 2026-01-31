package organization

import (
	"context"
	"net/http/httptest"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput")).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Corp",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
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
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput")).Return(&models.Organization{
					ID:        uuid.New(),
					Name:      "Tech Corp",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
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
				orgRepo.On("CreateOrganization", mock.Anything, mock.AnythingOfType("*models.CreateOrganizationInput")).Return(nil, &errs.HTTPError{
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
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput")).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
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
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput")).Return(&models.Organization{
					ID:        existingID,
					Name:      "Updated Name",
					Active:    false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
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
				orgRepo.On("UpdateOrganization", mock.Anything, mock.AnythingOfType("*models.UpdateOrganizationInput")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Organization", "id", existingID.String()).Code,
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
				require.Error(t, err)
				assert.Nil(t, output)
			} else {
				require.NoError(t, err)
				require.NotNil(t, output)
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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
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

			handler := NewHandler(mockOrgRepo, mockLocRepo)
			output, err := handler.GetAllOrganizations(context.TODO(), tt.pagination)

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

func TestHandler_GetEventOccurrencesByOrganizationId(t *testing.T) {
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
		name             string
		id               string
		mockSetup        func(*repomocks.MockOrganizationRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get event occurrence by organization id",
			id:   "40000000-0000-0000-0000-000000000001",
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
			wantErr: false,
		},
		{
			name: "no event occurrences with the organization id",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockOrganizationRepository) {
				m.On(
					"GetEventOccurrencesByOrganizationID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(make([]models.EventOccurrence, 0), nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockOrganizationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockLocationRepo)
			ctx := context.Background()

			input := &models.GetEventOccurrencesByOrganizationIDInput{ID: uuid.MustParse(tt.id)}
			eventOccurrences, err := handler.GetEventOccurrencesByOrganizationID(ctx, input)

			assert.Nil(t, err)
			assert.NotNil(t, eventOccurrences)
			if len(eventOccurrences) != 0 {
				assert.Equal(t, 2, len(eventOccurrences))
			} else {
				assert.Equal(t, 0, len(eventOccurrences))
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
