package eventoccurrence

import (
	"context"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateEventOccurrence(t *testing.T) {
	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")

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
		name      string
		input     *models.CreateEventOccurrenceInput
		mockSetup func(*repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "successful creation",
			input: func() *models.CreateEventOccurrenceInput {
				input := &models.CreateEventOccurrenceInput{}
				input.Body.ManagerId = &mid
				input.Body.EventId = event.ID
				input.Body.LocationId = location.ID
				input.Body.StartTime = start
				input.Body.EndTime = end
				input.Body.MaxAttendees = 15
				input.Body.Language = "en"
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"CreateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.CreateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.New(),
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
				}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			eventOccurrence, err := handler.CreateEventOccurrence(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, eventOccurrence)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				if tt.input.Body.ManagerId != nil {
					assert.Equal(t, *tt.input.Body.ManagerId, *eventOccurrence.ManagerId)
				} else {
					assert.Nil(t, eventOccurrence.ManagerId)
				}
				assert.Equal(t, tt.input.Body.EventId, eventOccurrence.Event.ID)
				assert.Equal(t, tt.input.Body.LocationId, eventOccurrence.Location.ID)
				assert.Equal(t, tt.input.Body.StartTime, eventOccurrence.StartTime)
				assert.Equal(t, tt.input.Body.EndTime, eventOccurrence.EndTime)
				assert.Equal(t, tt.input.Body.MaxAttendees, eventOccurrence.MaxAttendees)
				assert.Equal(t, tt.input.Body.Language, eventOccurrence.Language)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetEventOccurrenceById(t *testing.T) {
	statusCodeNotFound := http.StatusNotFound
	messageSubstringNotFound := "Not found"

	start, _ := time.Parse(time.RFC3339, "2026-02-15 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-15 11:00:00+07")

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
		mockSetup        func(*repomocks.MockEventOccurrenceRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get event occurrence by id",
			id:   "70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				).Return(&models.EventOccurrence{
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
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "event occurrence not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("00000000-0000-0000-0000-000000000000"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("EventOccurrence", "id", "00000000-0000-0000-0000-000000000000").GetStatus(),
					Message: "Not found",
				})
			},
			wantErr:          true,
			statusCode:       &statusCodeNotFound,
			messageSubstring: &messageSubstringNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt 
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetEventOccurrenceByIDInput{ID: uuid.MustParse(tt.id)}
			eventOccurrence, err := handler.GetEventOccurrenceByID(ctx, input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, eventOccurrence)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				assert.Equal(t, tt.id, eventOccurrence.ID.String())
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
		mockSetup        func(*repomocks.MockEventOccurrenceRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get event occurrence by event id",
			id:   "60000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
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
	}

	for _, tt := range tests {
		tt := tt 
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetEventOccurrencesByEventIDInput{ID: uuid.MustParse(tt.id)}
			eventOccurrences, err := handler.GetEventOccurrencesByEventID(ctx, input)

			assert.Nil(t, err)
			assert.NotNil(t, eventOccurrences)
			assert.Equal(t, len(eventOccurrences), 2)

			mockRepo.AssertExpectations(t)
		})
	}
}