package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetRegistrationByID(t *testing.T) {
	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "successful get registration",
			id:   uuid.MustParse("80000000-0000-0000-0000-000000000001"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID:                  uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			id:   uuid.MustParse("80000000-0000-0000-0000-000000000099"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Registration", "id", uuid.MustParse("80000000-0000-0000-0000-000000000099")).Code,
					Message: "Registration not found",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.GetRegistrationByID(context.TODO(), &models.GetRegistrationByIDInput{
				ID: tt.id,
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.id, output.Body.ID)
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetRegistrationsByChildID(t *testing.T) {
	tests := []struct {
		name        string
		childID     uuid.UUID
		mockSetup   func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr     bool
		expectedLen int
	}{
		{
			name:    "successful get registrations by child",
			childID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(&models.GetRegistrationsByChildIDOutput{
					Body: struct {
						Registrations []models.Registration `json:"registrations" doc:"List of registrations for the child"`
					}{
						Registrations: []models.Registration{
							{ID: uuid.New(), ChildID: uuid.MustParse("30000000-0000-0000-0000-000000000001")},
							{ID: uuid.New(), ChildID: uuid.MustParse("30000000-0000-0000-0000-000000000001")},
						},
					},
				}, nil)
			},
			wantErr:     false,
			expectedLen: 2,
		},
		{
			name:    "no registrations found",
			childID: uuid.New(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(&models.GetRegistrationsByChildIDOutput{
					Body: struct {
						Registrations []models.Registration `json:"registrations" doc:"List of registrations for the child"`
					}{
						Registrations: []models.Registration{},
					},
				}, nil)
			},
			wantErr:     false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.GetRegistrationsByChildID(context.TODO(), &models.GetRegistrationsByChildIDInput{
				ChildID: tt.childID,
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.expectedLen, len(output.Body.Registrations))
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetRegistrationsByGuardianID(t *testing.T) {
	tests := []struct {
		name        string
		guardianID  uuid.UUID
		mockSetup   func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr     bool
		expectedLen int
	}{
		{
			name:       "successful get registrations by guardian",
			guardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(&models.GetRegistrationsByGuardianIDOutput{
					Body: struct {
						Registrations []models.Registration `json:"registrations" doc:"List of registrations for the guardian"`
					}{
						Registrations: []models.Registration{
							{ID: uuid.New(), GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111")},
							{ID: uuid.New(), GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111")},
							{ID: uuid.New(), GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111")},
						},
					},
				}, nil)
			},
			wantErr:     false,
			expectedLen: 3,
		},
		{
			name:       "no registrations found",
			guardianID: uuid.New(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(&models.GetRegistrationsByGuardianIDOutput{
					Body: struct {
						Registrations []models.Registration `json:"registrations" doc:"List of registrations for the guardian"`
					}{
						Registrations: []models.Registration{},
					},
				}, nil)
			},
			wantErr:     false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.GetRegistrationsByGuardianID(context.TODO(), &models.GetRegistrationsByGuardianIDInput{
				GuardianID: tt.guardianID,
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.expectedLen, len(output.Body.Registrations))
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateRegistration(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateRegistrationInput
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "successful create",
			input: &models.CreateRegistrationInput{
				Body: struct {
					ChildID           uuid.UUID                  `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
					GuardianID        uuid.UUID                  `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
					EventOccurrenceID uuid.UUID                  `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
					Status            models.RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
				}{
					ChildID:           uuid.MustParse("30000000-0000-0000-0000-000000000001"),
					GuardianID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					EventOccurrenceID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
					Status:            models.RegistrationStatusRegistered,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse("70000000-0000-0000-0000-000000000001")).Return(&models.EventOccurrence{
					ID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).Return(&models.Child{
					ID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
				}, nil)
				guardianRepo.On("GetGuardianByID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).Return(&models.Guardian{
					ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				}, nil)
				regRepo.On("CreateRegistration", mock.Anything, mock.AnythingOfType("*models.CreateRegistrationInput")).Return(&models.CreateRegistrationOutput{
					Body: models.Registration{
						ID:                  uuid.New(),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid event_occurrence_id",
			input: &models.CreateRegistrationInput{
				Body: struct {
					ChildID           uuid.UUID                  `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
					GuardianID        uuid.UUID                  `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
					EventOccurrenceID uuid.UUID                  `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
					Status            models.RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
				}{
					ChildID:           uuid.MustParse("30000000-0000-0000-0000-000000000001"),
					GuardianID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					EventOccurrenceID: uuid.New(),
					Status:            models.RegistrationStatusRegistered,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("EventOccurrence", "id", "").Code,
					Message: "Event occurrence not found",
				})
			},
			wantErr: true,
		},
		{
			name: "invalid child_id",
			input: &models.CreateRegistrationInput{
				Body: struct {
					ChildID           uuid.UUID                  `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
					GuardianID        uuid.UUID                  `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
					EventOccurrenceID uuid.UUID                  `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
					Status            models.RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
				}{
					ChildID:           uuid.New(),
					GuardianID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					EventOccurrenceID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
					Status:            models.RegistrationStatusRegistered,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse("70000000-0000-0000-0000-000000000001")).Return(&models.EventOccurrence{
					ID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Child", "id", "").Code,
					Message: "Child not found",
				})
			},
			wantErr: true,
		},
		{
			name: "invalid guardian_id",
			input: &models.CreateRegistrationInput{
				Body: struct {
					ChildID           uuid.UUID                  `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
					GuardianID        uuid.UUID                  `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
					EventOccurrenceID uuid.UUID                  `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
					Status            models.RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
				}{
					ChildID:           uuid.MustParse("30000000-0000-0000-0000-000000000001"),
					GuardianID:        uuid.New(),
					EventOccurrenceID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
					Status:            models.RegistrationStatusRegistered,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse("70000000-0000-0000-0000-000000000001")).Return(&models.EventOccurrence{
					ID: uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).Return(&models.Child{
					ID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
				}, nil)
				guardianRepo.On("GetGuardianByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Guardian", "id", "").Code,
					Message: "Guardian not found",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.CreateRegistration(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.input.Body.ChildID, output.Body.ChildID)
				assert.Equal(t, tt.input.Body.GuardianID, output.Body.GuardianID)
				assert.Equal(t, tt.input.Body.EventOccurrenceID, output.Body.EventOccurrenceID)
			}

			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateRegistration(t *testing.T) {
	existingID := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	newStatus := models.RegistrationStatusCancelled

	tests := []struct {
		name      string
		input     *models.UpdateRegistrationInput
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "successful update status only",
			input: &models.UpdateRegistrationInput{
				ID: existingID,
				Body: struct {
					ChildID           *uuid.UUID                  `json:"child_id,omitempty" doc:"Updated child ID (optional)" format:"uuid"`
					GuardianID        *uuid.UUID                  `json:"guardian_id,omitempty" doc:"Updated guardian ID (optional)" format:"uuid"`
					EventOccurrenceID *uuid.UUID                  `json:"event_occurrence_id,omitempty" doc:"Updated event occurrence ID (optional)" format:"uuid"`
					Status            *models.RegistrationStatus `json:"status,omitempty" doc:"Updated registration status (optional)" enum:"registered,cancelled"`
				}{
					Status: &newStatus,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                  existingID,
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusCancelled,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			input: &models.UpdateRegistrationInput{
				ID: existingID,
				Body: struct {
					ChildID           *uuid.UUID                  `json:"child_id,omitempty" doc:"Updated child ID (optional)" format:"uuid"`
					GuardianID        *uuid.UUID                  `json:"guardian_id,omitempty" doc:"Updated guardian ID (optional)" format:"uuid"`
					EventOccurrenceID *uuid.UUID                  `json:"event_occurrence_id,omitempty" doc:"Updated event occurrence ID (optional)" format:"uuid"`
					Status            *models.RegistrationStatus `json:"status,omitempty" doc:"Updated registration status (optional)" enum:"registered,cancelled"`
				}{
					Status: &newStatus,
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Registration", "id", existingID).Code,
					Message: "Registration not found",
				})
			},
			wantErr: true,
		},
		{
			name: "invalid child_id on update",
			input: &models.UpdateRegistrationInput{
				ID: existingID,
				Body: struct {
					ChildID           *uuid.UUID                  `json:"child_id,omitempty" doc:"Updated child ID (optional)" format:"uuid"`
					GuardianID        *uuid.UUID                  `json:"guardian_id,omitempty" doc:"Updated guardian ID (optional)" format:"uuid"`
					EventOccurrenceID *uuid.UUID                  `json:"event_occurrence_id,omitempty" doc:"Updated event occurrence ID (optional)" format:"uuid"`
					Status            *models.RegistrationStatus `json:"status,omitempty" doc:"Updated registration status (optional)" enum:"registered,cancelled"`
				}{
					ChildID: func() *uuid.UUID { id := uuid.New(); return &id }(),
				},
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				childRepo.On("GetChildByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Child", "id", "").Code,
					Message: "Child not found",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.UpdateRegistration(context.TODO(), tt.input)

			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, output)
			} else {
				require.NoError(t, err)
				require.NotNil(t, output)
				if tt.input.Body.Status != nil {
					assert.Equal(t, *tt.input.Body.Status, output.Body.Status)
				}
			}

			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteRegistration(t *testing.T) {
	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   uuid.MustParse("80000000-0000-0000-0000-000000000001"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("DeleteRegistration", mock.Anything, mock.AnythingOfType("*models.DeleteRegistrationInput")).Return(&models.DeleteRegistrationOutput{
					Body: models.Registration{
						ID:                  uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:             uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:   uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:              models.RegistrationStatusRegistered,
						EventName:           "STEM Club",
						OccurrenceStartTime: time.Now(),
						CreatedAt:           time.Now(),
						UpdatedAt:           time.Now(),
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			id:   uuid.MustParse("80000000-0000-0000-0000-000000000099"),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository) {
				regRepo.On("DeleteRegistration", mock.Anything, mock.AnythingOfType("*models.DeleteRegistrationInput")).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Registration", "id", "80000000-0000-0000-0000-000000000099").Code,
					Message: "Registration not found",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo)
			output, err := handler.DeleteRegistration(context.TODO(), &models.DeleteRegistrationInput{
				ID: tt.id,
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.id, output.Body.ID)
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}