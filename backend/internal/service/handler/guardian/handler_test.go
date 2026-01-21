package guardian

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
)

func TestHandler_GetGuardianById(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful get guardian by id - Sarah Johnson",
			id:   "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name: "internal server error",
			id:   "ffffffff-ffff-ffff-ffff-ffffffffffff",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByID", mock.Anything, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetGuardianByIDInput{ID: uuid.MustParse(tt.id)}
			guardian, err := handler.GetGuardianById(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.id, guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateGuardian(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateGuardianInput
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful create guardian - Michael Chen",
			input: func() *models.CreateGuardianInput {
				input := &models.CreateGuardianInput{}
				input.Body.UserID = uuid.MustParse("b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e")
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				// Expect check for existing guardian to return NotFound (allowing creation)
				m.On("GetGuardianByUserID", mock.Anything, uuid.MustParse("b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e")).
					Return(nil, &errs.HTTPError{Code: 404, Message: "Not Found"})

				m.On("CreateGuardian", mock.Anything, mock.AnythingOfType("*models.CreateGuardianInput")).Return(&models.Guardian{
					ID:        uuid.New(),
					UserID:    uuid.MustParse("b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian already exists",
			input: func() *models.CreateGuardianInput {
				input := &models.CreateGuardianInput{}
				input.Body.UserID = uuid.MustParse("c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f")
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				// Expect check to return existing guardian
				m.On("GetGuardianByUserID", mock.Anything, uuid.MustParse("c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f")).
					Return(&models.Guardian{
						ID:     uuid.New(),
						UserID: uuid.MustParse("c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f"),
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "internal server error during creation",
			input: func() *models.CreateGuardianInput {
				input := &models.CreateGuardianInput{}
				input.Body.UserID = uuid.MustParse("d4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a")
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				// Expect check to return NotFound
				m.On("GetGuardianByUserID", mock.Anything, uuid.MustParse("d4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a")).
					Return(nil, &errs.HTTPError{Code: 404, Message: "Not Found"})

				m.On("CreateGuardian", mock.Anything, mock.AnythingOfType("*models.CreateGuardianInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			guardian, err := handler.CreateGuardian(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.input.Body.UserID, guardian.UserID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateGuardian(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		input     *models.UpdateGuardianInput
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful update guardian - Sarah Johnson",
			id:   "11111111-1111-1111-1111-111111111111",
			input: func() *models.UpdateGuardianInput {
				input := &models.UpdateGuardianInput{}
				input.ID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				input.Body.UserID = uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d") // Keep same user ID for update test simplicity
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("UpdateGuardian", mock.Anything, mock.MatchedBy(func(input *models.UpdateGuardianInput) bool {
					return input.ID == uuid.MustParse("11111111-1111-1111-1111-111111111111")
				})).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			input: func() *models.UpdateGuardianInput {
				input := &models.UpdateGuardianInput{}
				input.ID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				return input
			}(),
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("UpdateGuardian", mock.Anything, mock.AnythingOfType("*models.UpdateGuardianInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			guardian, err := handler.UpdateGuardian(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.input.Body.UserID, guardian.UserID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetGuardianByChildId(t *testing.T) {
	tests := []struct {
		name      string
		childID   string
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name:    "successful get guardian by child id - Emily Johnson",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByChildID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name:    "guardian not found for child",
			childID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("GetGuardianByChildID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "child_id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetGuardianByChildIDInput{ChildID: uuid.MustParse(tt.childID)}
			guardian, err := handler.GetGuardianByChildId(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, "11111111-1111-1111-1111-111111111111", guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteGuardian(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockGuardianRepository)
		wantErr   bool
	}{
		{
			name: "successful delete guardian - Sarah Johnson",
			id:   "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("DeleteGuardian", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).Return(&models.Guardian{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:    uuid.MustParse("a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "guardian not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockGuardianRepository) {
				m.On("DeleteGuardian", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockGuardianRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.DeleteGuardianInput{ID: uuid.MustParse(tt.id)}
			guardian, err := handler.DeleteGuardian(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, guardian)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, guardian)
				assert.Equal(t, tt.id, guardian.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
