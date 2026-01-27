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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo)
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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo)
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

func TestHandler_UpdateEventOccurrence(t *testing.T) {
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	mid_new := uuid.MustParse("50000000-0000-0000-0000-000000000005")
	eid := uuid.MustParse("60000000-0000-0000-0000-00000000000e")
	lid := uuid.MustParse("10000000-0000-0000-0000-000000000008")
	start, _ := time.Parse(time.RFC3339, "2026-02-22 09:00:00+07")
	end, _ := time.Parse(time.RFC3339, "2026-02-22 11:00:00+07")
	start_new, _ := time.Parse(time.RFC3339, "2026-02-15 10:00:00+07")
	end_new, _ := time.Parse(time.RFC3339, "2026-02-15 12:00:00+07")
	max := 10
	lang := "th"
	curr := 8
	curr_bad := 20

	category_arr := []string{"science","technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	event := models.Event{
		ID: 				eid,
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

	category_arr_new := []string{"technology","math"}
	ten := 10
	fifteen := 15
	event_new := models.Event{
		ID: 				uuid.MustParse("60000000-0000-0000-0000-00000000000e"),
		Title: 				"Python for Kids",
		Description: 		"Introduction to Python programming. Build simple programs and games while learning core concepts.",
		OrganizationID: 	uuid.MustParse("40000000-0000-0000-0000-000000000005"),
		AgeRangeMin: 		&ten,
		AgeRangeMax: 		&fifteen,
		Category: 			category_arr_new,
		HeaderImageS3Key: 	nil,
		CreatedAt: 			time.Now(),
		UpdatedAt: 			time.Now(),
	}

	addr := "Suite 15"
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

	location_new := models.Location{
		ID:           lid,
		Latitude:     13.7400000,
		Longitude:    100.5450000,
		AddressLine1: "369 Wireless Road",
		AddressLine2: nil,
		Subdistrict:  "Lumphini",
		District:     "Pathum Wan",
		Province:     "Bangkok",
		PostalCode:   "10330",
		Country:      "Thailand",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	

	tests := []struct {
		name      string
		input     *models.UpdateEventOccurrenceInput
		mockSetup func(*repomocks.MockEventOccurrenceRepository)
		wantErr   bool
	}{
		{
			name: "new current enrolled exceeds original max attendees",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
				input.Body.ManagerId = nil
				input.Body.EventId = nil
				input.Body.LocationId = nil
				input.Body.StartTime = nil
				input.Body.EndTime = nil
				input.Body.MaxAttendees = nil
				input.Body.Language = nil
				input.Body.CurrEnrolled = &curr_bad
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000002"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId: 		&mid,
					Event: 			event,
					Location: 		location,
					StartTime: 		start,
					EndTime: 		end,
					MaxAttendees: 	15,
					Language: 		"en",
					CurrEnrolled: 	5,
					CreatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "new current enrolled exceeds new max attendees",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
				input.Body.ManagerId = nil
				input.Body.EventId = nil
				input.Body.LocationId = nil
				input.Body.StartTime = nil
				input.Body.EndTime = nil
				input.Body.MaxAttendees = &max
				input.Body.Language = nil
				input.Body.CurrEnrolled = &curr_bad
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000002"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId: 		&mid,
					Event: 			event,
					Location: 		location,
					StartTime: 		start,
					EndTime: 		end,
					MaxAttendees: 	15,
					Language: 		"en",
					CurrEnrolled: 	5,
					CreatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "successfully updated current enrolled",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
				input.Body.ManagerId = nil
				input.Body.EventId = nil
				input.Body.LocationId = nil
				input.Body.StartTime = nil
				input.Body.EndTime = nil
				input.Body.MaxAttendees = nil
				input.Body.Language = nil
				input.Body.CurrEnrolled = &curr
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"UpdateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId: 		&mid,
					Event: 			event,
					Location: 		location,
					StartTime: 		start,
					EndTime: 		end,
					MaxAttendees: 	15,
					Language: 		"en",
					CurrEnrolled: 	8,
					CreatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    	time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "successfully updated all fields except current enrolled",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")
				input.Body.ManagerId = &mid_new
				input.Body.EventId = &eid
				input.Body.LocationId = &lid
				input.Body.StartTime = &start_new
				input.Body.EndTime = &end_new
				input.Body.MaxAttendees = &max
				input.Body.Language = &lang
				input.Body.CurrEnrolled = nil
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On(
					"UpdateEventOccurrence",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateEventOccurrenceInput"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId: 		&mid_new,
					Event: 			event_new,
					Location: 		location_new,
					StartTime: 		start_new,
					EndTime: 		end_new,
					MaxAttendees: 	10,
					Language: 		"th",
					CurrEnrolled: 	5,
					CreatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
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
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo)
			ctx := context.Background()

			if !tt.wantErr {
				mockRepo.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000002"),
				).Return(&models.EventOccurrence{
					ID:        		uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId: 		&mid,
					Event: 			event,
					Location: 		location,
					StartTime: 		start,
					EndTime: 		end,
					MaxAttendees: 	15,
					Language: 		"en",
					CurrEnrolled: 	5,
					CreatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    	time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
				}, nil)
			}
			eventOccurrence, err := handler.UpdateEventOccurrence(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, eventOccurrence)
			} else if tt.name == "successfully updated current enrolled" {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				assert.Equal(t, mid, *eventOccurrence.ManagerId)
				assert.Equal(t, event.ID, eventOccurrence.Event.ID)
				assert.Equal(t, location.ID, eventOccurrence.Location.ID)
				assert.Equal(t, start, eventOccurrence.StartTime)
				assert.Equal(t, end, eventOccurrence.EndTime)
				assert.Equal(t, 15, eventOccurrence.MaxAttendees)
				assert.Equal(t, "en", eventOccurrence.Language)
				assert.Equal(t, *tt.input.Body.CurrEnrolled, eventOccurrence.CurrEnrolled)
			} else if tt.name == "successfully updated all fields except current enrolled" {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				assert.Equal(t, mid_new, *eventOccurrence.ManagerId)
				assert.Equal(t, eid, eventOccurrence.Event.ID)
				assert.Equal(t, lid, eventOccurrence.Location.ID)
				assert.Equal(t, start_new, eventOccurrence.StartTime)
				assert.Equal(t, end_new, eventOccurrence.EndTime)
				assert.Equal(t, max, eventOccurrence.MaxAttendees)
				assert.Equal(t, lang, eventOccurrence.Language)
				assert.Equal(t, 5, eventOccurrence.CurrEnrolled)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}