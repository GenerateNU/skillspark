package event

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

func TestHandler_CreateEvent(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateEventInput
		mockSetup func(*repomocks.MockEventRepository)
		wantErr   bool
	}{
		{
			name: "successful create event",
			input: func() *models.CreateEventInput {
				input := &models.CreateEventInput{}
				ageMin := 8
				ageMax := 12
				headerImage := "events/robotics.jpg"

				input.Body.Title = "Junior Robotics"
				input.Body.Description = "Intro to robotics"
				input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
				input.Body.AgeRangeMin = &ageMin
				input.Body.AgeRangeMax = &ageMax
				input.Body.Category = []string{"stem", "robotics"}
				input.Body.HeaderImageS3Key = &headerImage
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*models.CreateEventInput")).Return(&models.Event{
					ID:               uuid.New(),
					Title:            "Junior Robotics",
					Description:      "Intro to robotics",
					OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					AgeRangeMin:      func() *int { i := 8; return &i }(),
					AgeRangeMax:      func() *int { i := 12; return &i }(),
					Category:         []string{"stem", "robotics"},
					HeaderImageS3Key: func() *string { s := "events/robotics.jpg"; return &s }(),
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			input: func() *models.CreateEventInput {
				input := &models.CreateEventInput{}
				input.Body.Title = "Error Event"
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*models.CreateEventInput")).
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

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			event, err := handler.CreateEvent(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, event)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, event)
				assert.Equal(t, tt.input.Body.Title, event.Title)
				assert.Equal(t, tt.input.Body.OrganizationID, event.OrganizationID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_PatchEvent(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.PatchEventInput
		mockSetup func(*repomocks.MockEventRepository)
		wantErr   bool
	}{
		{
			name: "successful patch event",
			input: func() *models.PatchEventInput {
				input := &models.PatchEventInput{}
				input.ID = uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
				input.Body.Title = "Updated Robotics"
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("PatchEvent", mock.Anything, mock.AnythingOfType("*models.PatchEventInput")).Return(&models.Event{
					ID:          uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Title:       "Updated Robotics",
					Description: "Intro to robotics", // Unchanged field example
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "event not found",
			input: func() *models.PatchEventInput {
				input := &models.PatchEventInput{}
				input.ID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("PatchEvent", mock.Anything, mock.AnythingOfType("*models.PatchEventInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Event", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			event, err := handler.PatchEvent(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, event)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, event)
				assert.Equal(t, tt.input.Body.Title, event.Title)
				assert.Equal(t, tt.input.ID, event.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteEvent(t *testing.T) {
	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(*repomocks.MockEventRepository)
		wantErr   bool
	}{
		{
			name: "successful delete event",
			id:   uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
			mockSetup: func(m *repomocks.MockEventRepository) {
				// Note: Ensure the mock signature matches the repo interface.
				// Based on your repo implementation, DeleteEvent usually returns (*struct{}, error)
				m.On("DeleteEvent", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")).
					Return(&struct{}{}, nil)
			},
			wantErr: false,
		},
		{
			name: "event not found",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("DeleteEvent", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, errs.NotFound("Event", "id", "00000000-0000-0000-0000-000000000000"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			err := handler.DeleteEvent(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
