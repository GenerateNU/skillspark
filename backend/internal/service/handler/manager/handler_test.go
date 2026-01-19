package manager

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

func TestHandler_GetManagerByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockManagerRepository)
		wantErr   bool
	}{
		{
			name: "successful get manager by id - Director",
			id:   "50000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByID", mock.Anything, uuid.MustParse("50000000-0000-0000-0000-000000000001")).Return(&models.Manager{
					ID:             uuid.MustParse("50000000-0000-0000-0000-000000000001"),
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "successful get manager by id - Assistant Coach",
			id:   "50000000-0000-0000-0000-000000000006",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByID", mock.Anything, uuid.MustParse("50000000-0000-0000-0000-000000000006")).
					Return(&models.Manager{
						ID:             uuid.MustParse("50000000-0000-0000-0000-000000000006"),
						UserID:         uuid.MustParse("d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a"),
						OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000002"),
						Role:           "Assistant Coach",
						CreatedAt:      time.Now(),
						UpdatedAt:      time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "manager not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Manager", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name: "internal server error",
			id:   "ffffffff-ffff-ffff-ffff-ffffffffffff",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByID", mock.Anything, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")).
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

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetManagerByIDInput{ID: uuid.MustParse(tt.id)}
			location, err := handler.GetManagerByID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, location)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.id, location.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetManagerByOrgID(t *testing.T) {
	tests := []struct {
		name            string
		organization_id string
		mockSetup       func(*repomocks.MockManagerRepository)
		wantErr         bool
	}{
		{
			name:            "successful get manager by org_id - Director",
			organization_id: "40000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByOrgID", mock.Anything, uuid.MustParse("40000000-0000-0000-0000-000000000001")).Return(&models.Manager{
					ID:             uuid.MustParse("50000000-0000-0000-0000-000000000001"),
					UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Role:           "Director",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name:            "successful get manager by org_id - Head Coach",
			organization_id: "40000000-0000-0000-0000-000000000002",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByOrgID", mock.Anything, uuid.MustParse("40000000-0000-0000-0000-000000000002")).
					Return(&models.Manager{
						ID:             uuid.MustParse("50000000-0000-0000-0000-000000000002"),
						UserID:         uuid.MustParse("d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a"),
						OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000002"),
						Role:           "Head Coach",
						CreatedAt:      time.Now(),
						UpdatedAt:      time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name:            "manager not found",
			organization_id: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByOrgID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Location", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name:            "internal server error",
			organization_id: "ffffffff-ffff-ffff-ffff-ffffffffffff",
			mockSetup: func(m *repomocks.MockManagerRepository) {
				m.On("GetManagerByOrgID", mock.Anything, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")).
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

			mockRepo := new(repomocks.MockManagerRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetManagerByOrgIDInput{OrganizationID: uuid.MustParse(tt.organization_id)}
			manager, err := handler.GetManagerByOrgID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, manager)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, manager)
				assert.Equal(t, tt.organization_id, manager.OrganizationID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
