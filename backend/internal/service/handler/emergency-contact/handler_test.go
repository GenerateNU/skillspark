package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetEmergencyContactByGuardianID(t *testing.T) {
	tests := []struct {
		name       string
		guardianID uuid.UUID
		mockSetup  func(*repomocks.MockEmergencyContactRepository)
		wantLen    int
		wantErr    bool
	}{
		{
			name:       "successful fetch",
			guardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				repo.On("GetEmergencyContactByGuardianID",
					mock.Anything,
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				).Return([]*models.EmergencyContact{
					{
						ID:         uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						Name:       "Jane Doe",
					},
				}, nil)
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name:       "repository error",
			guardianID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				err := errs.InternalServerError("repo failure", "")
				repo.On("GetEmergencyContactByGuardianID",
					mock.Anything,
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				).Return(nil, &err)
			},
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEmergencyContactRepository)
			tt.mockSetup(mockRepo)

			handler := &Handler{EmergencyContactRepository: mockRepo}

			result, err := handler.GetEmergencyContactByGuardianID(context.Background(), tt.guardianID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, tt.wantLen)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateEmergencyContact(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateEmergencyContactInput
		mockSetup func(*repomocks.MockEmergencyContactRepository)
		wantErr   bool
	}{
		{
			name:  "successful creation",
			input: &models.CreateEmergencyContactInput{},
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				repo.On("CreateEmergencyContact",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEmergencyContactInput"),
				).Return(&models.CreateEmergencyContactOutput{
					Body: &models.EmergencyContact{
						ID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						Name: "Jane Doe",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name:  "repository error",
			input: &models.CreateEmergencyContactInput{},
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				err := errs.InternalServerError("repo failure", "")
				repo.On("CreateEmergencyContact",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEmergencyContactInput"),
				).Return(nil, &err)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEmergencyContactRepository)
			tt.mockSetup(mockRepo)

			handler := &Handler{EmergencyContactRepository: mockRepo}

			output, err := handler.CreateEmergencyContact(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateEmergencyContact(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.UpdateEmergencyContactInput
		mockSetup func(*repomocks.MockEmergencyContactRepository)
		wantErr   bool
	}{
		{
			name:  "successful update",
			input: &models.UpdateEmergencyContactInput{ID: uuid.MustParse("20000000-0000-0000-0000-000000000001")},
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				repo.On("UpdateEmergencyContact",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEmergencyContactInput"),
				).Return(&models.UpdateEmergencyContactOutput{
					Body: &models.EmergencyContact{
						ID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						Name: "Updated Name",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name:  "repository error",
			input: &models.UpdateEmergencyContactInput{ID: uuid.MustParse("20000000-0000-0000-0000-000000000002")},
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				err := errs.InternalServerError("repo failure", "")
				repo.On("UpdateEmergencyContact",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEmergencyContactInput"),
				).Return(nil, &err)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEmergencyContactRepository)
			tt.mockSetup(mockRepo)

			handler := &Handler{EmergencyContactRepository: mockRepo}

			output, err := handler.UpdateEmergencyContact(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteEmergencyContact(t *testing.T) {
	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(*repomocks.MockEmergencyContactRepository)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				repo.On("DeleteEmergencyContact",
					mock.Anything,
					uuid.MustParse("20000000-0000-0000-0000-000000000001"),
				).Return(&models.DeleteEmergencyContactOutput{
					SuccessMessage: "success",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			id:   uuid.MustParse("20000000-0000-0000-0000-000000000002"),
			mockSetup: func(repo *repomocks.MockEmergencyContactRepository) {
				err := errs.NotFound("EmergencyContact", "id", uuid.MustParse("20000000-0000-0000-0000-000000000002"))
				repo.On("DeleteEmergencyContact",
					mock.Anything,
					uuid.MustParse("20000000-0000-0000-0000-000000000002"),
				).Return(nil, &err)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEmergencyContactRepository)
			tt.mockSetup(mockRepo)

			handler := &Handler{EmergencyContactRepository: mockRepo}

			output, err := handler.DeleteEmergencyContact(context.Background(), tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
