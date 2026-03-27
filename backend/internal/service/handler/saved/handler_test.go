package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetByGuardianID(t *testing.T) {
	tests := []struct {
		name       string
		guardianID uuid.UUID
		mockSetup  func(*repomocks.MockGuardianRepository, *repomocks.MockSavedRepository)
		wantSaved  []models.Saved
		wantErr    bool
	}{
		{
			name:       "successful fetch saved",
			guardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, savedRepo *repomocks.MockSavedRepository) {

				guardianRepo.On("GetGuardianByID",
					mock.Anything,
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				).Return(&models.Guardian{
					ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				}, nil)

				savedRepo.On("GetByGuardianID",
					mock.Anything,
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					mock.AnythingOfType("utils.Pagination"),
					mock.AnythingOfType("string"),
				).Return([]models.Saved{
					{
						ID:         uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					},
				}, nil)
			},
			wantSaved: []models.Saved{
				{
					ID:         uuid.MustParse("20000000-0000-0000-0000-000000000001"),
					GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			wantErr: false,
		},
		{
			name:       "guardian does not exist",
			guardianID: uuid.MustParse("99999999-0000-0000-0000-000000000000"),
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, savedRepo *repomocks.MockSavedRepository) {

				guardianRepo.On("GetGuardianByID",
					mock.Anything,
					uuid.MustParse("99999999-0000-0000-0000-000000000000"),
				).Return(nil, errs.BadRequest("guardian does not exist"))
			},
			wantSaved: nil,
			wantErr:   true,
		},
		{
			name:       "repository error",
			guardianID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, savedRepo *repomocks.MockSavedRepository) {

				guardianRepo.On("GetGuardianByID",
					mock.Anything,
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				).Return(&models.Guardian{
					ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				}, nil)

				savedRepo.On("GetByGuardianID",
					mock.Anything,
					uuid.MustParse("22222222-2222-2222-2222-222222222222"),
					mock.AnythingOfType("utils.Pagination"),
					mock.AnythingOfType("string"),
				).Return(nil, errs.BadRequest("cannot fetch saved"))
			},
			wantSaved: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockSavedRepo := new(repomocks.MockSavedRepository)

			tt.mockSetup(mockGuardianRepo, mockSavedRepo)

			handler := &Handler{
				GuardianRepository: mockGuardianRepo,
				SavedRepository:    mockSavedRepo,
			}

			pagination := utils.Pagination{Page: 1, Limit: 10}

			saved, err := handler.GetByGuardianID(context.Background(), tt.guardianID, pagination, "en-US")

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, saved)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, saved)
				assert.Equal(t, len(tt.wantSaved), len(saved))
				for i := range saved {
					assert.Equal(t, tt.wantSaved[i].ID, saved[i].ID)
					assert.Equal(t, tt.wantSaved[i].GuardianID, saved[i].GuardianID)
				}
			}

			mockGuardianRepo.AssertExpectations(t)
			mockSavedRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteSaved(t *testing.T) {

	tests := []struct {
		name      string
		input     *models.DeleteSavedInput
		mockSetup func(*repomocks.MockSavedRepository)
		wantMsg   string
		wantErr   bool
	}{
		{
			name: "repository error",
			input: &models.DeleteSavedInput{
				ID: uuid.MustParse("20000000-0000-0000-0000-000000000002"),
			},
			mockSetup: func(repo *repomocks.MockSavedRepository) {
				err := errs.BadRequest("cannot delete saved")
				repo.On("DeleteSaved",
					mock.Anything,
					uuid.MustParse("20000000-0000-0000-0000-000000000002"),
				).Return(&err)
			},
			wantMsg: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			mockSavedRepo := new(repomocks.MockSavedRepository)

			tt.mockSetup(mockSavedRepo)

			handler := &Handler{
				SavedRepository: mockSavedRepo,
			}

			msg, err := handler.DeleteSaved(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, "", msg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantMsg, msg)
			}

			mockSavedRepo.AssertExpectations(t)

		})
	}
}

func TestHandler_CreateSaved(t *testing.T) {

	tests := []struct {
		name      string
		input     *models.CreateSavedInput
		mockSetup func(*repomocks.MockSavedRepository)
		wantErr   bool
	}{
		{
			name:  "repository error",
			input: &models.CreateSavedInput{},
			mockSetup: func(repo *repomocks.MockSavedRepository) {

				err := errs.BadRequest("cannot delete saved")

				repo.On("CreateSaved",
					mock.Anything,
					mock.AnythingOfType("*models.CreateSavedInput"),
				).Return(nil, &err)

			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			mockSavedRepo := new(repomocks.MockSavedRepository)

			tt.mockSetup(mockSavedRepo)

			handler := &Handler{
				SavedRepository: mockSavedRepo,
			}

			output, err := handler.CreateSaved(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tt.input.Body.GuardianID, output.Body.GuardianID)
			}

			mockSavedRepo.AssertExpectations(t)

		})
	}
}
