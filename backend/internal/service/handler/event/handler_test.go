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
				assert.NotNil(t, err)
				assert.Nil(t, event)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, event)
				assert.Equal(t, tt.input.Body.Title, event.Title)
				assert.Equal(t, tt.input.Body.OrganizationID, event.OrganizationID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateEvent(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.UpdateEventInput
		mockSetup func(*repomocks.MockEventRepository)
		wantErr   bool
	}{
		{
			name: "successful update event",
			input: func() *models.UpdateEventInput {
				input := &models.UpdateEventInput{}
				input.ID = uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
				title := "Updated Robotics"
				input.Body.Title = &title
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput")).Return(&models.Event{
					ID:          uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Title:       "Updated Robotics",
					Description: "Intro to robotics",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "event not found",
			input: func() *models.UpdateEventInput {
				input := &models.UpdateEventInput{}
				input.ID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput")).
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

			event, err := handler.UpdateEvent(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, event)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, event)
				assert.Equal(t, *tt.input.Body.Title, event.Title)
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
		wantMsg   string
		wantErr   bool
	}{
		{
			name: "successful delete event",
			id:   uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("DeleteEvent", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")).
					Return(nil)
			},
			wantMsg: "Event successfully deleted.",
			wantErr: false,
		},
		{
			name: "event not found",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			mockSetup: func(m *repomocks.MockEventRepository) {
				httpErr := errs.NotFound("Event", "id", "00000000-0000-0000-0000-000000000000")
				m.On("DeleteEvent", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(&httpErr)
			},
			wantMsg: "",
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

			msg, err := handler.DeleteEvent(ctx, tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Empty(t, msg)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantMsg, msg)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetEventOccurrencesByEventId(t *testing.T) {
	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")
	start2, _ := time.Parse(time.RFC3339, "2026-02-22 09:00:00+07")
	end2, _ := time.Parse(time.RFC3339, "2026-02-22 11:00:00+07")

	category_arr := []string{"science","technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	event := models.Event{
		ID: 				uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title: 				"Junior Robotics Workshop",
		Description: 		"Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID: 	uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin: 		&eight,
		AgeRangeMax: 		&twelve,
		Category: 			category_arr,
		HeaderImageS3Key: 	&jpg,
		CreatedAt: 			time.Now(),
		UpdatedAt: 			time.Now(),
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
		mockSetup        func(*repomocks.MockEventRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get event occurrence by event id",
			id:   "60000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
					mock.Anything,
					uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				).Return([]models.EventOccurrence{
					{
						ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						ManagerId: 		&mid,
						Event: 			event,
						Location: 		location,
						StartTime: 		start,
						EndTime: 		end,
						MaxAttendees: 	15,
						Language: 		"en",
						CurrEnrolled: 	8,
						CreatedAt:    	time.Now(),
						UpdatedAt:    	time.Now(),
					},
					{
						ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
						ManagerId: 		&mid,
						Event: 			event,
						Location: 		location,
						StartTime: 		start2,
						EndTime: 		end2,
						MaxAttendees: 	15,
						Language: 		"en",
						CurrEnrolled: 	5,
						CreatedAt:    	time.Now(),
						UpdatedAt:    	time.Now(),
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "no event occurrences with the event id",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On(
					"GetEventOccurrencesByEventID",
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

			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetEventOccurrencesByEventIDInput{ID: uuid.MustParse(tt.id)}
			eventOccurrences, err := handler.GetEventOccurrencesByEventID(ctx, input)

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