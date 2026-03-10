package eventoccurrence

import (
	"context"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	s3mocks "skillspark/internal/s3_client/mocks"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// shared fixtures
var (
	testMid     = uuid.MustParse("50000000-0000-0000-0000-000000000001")
	testEOID    = uuid.MustParse("70000000-0000-0000-0000-000000000002")
	testStart, _ = time.Parse(time.RFC3339, "2026-02-22T09:00:00+07:00")
	testEnd, _   = time.Parse(time.RFC3339, "2026-02-22T11:00:00+07:00")
)

func makeTestEvent() models.Event {
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	return models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         []string{"science", "technology"},
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func makeTestLocation() models.Location {
	addr := "Suite 15"
	return models.Location{
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
}

func makeTestEventOccurrence(startTime time.Time) *models.EventOccurrence {
	return &models.EventOccurrence{
		ID:           testEOID,
		ManagerId:    &testMid,
		Event:        makeTestEvent(),
		Location:     makeTestLocation(),
		StartTime:    startTime,
		EndTime:      testEnd,
		MaxAttendees: 15,
		Language:     "en",
		CurrEnrolled: 5,
		CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
		UpdatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
	}
}

func newHandler(
	eoRepo *repomocks.MockEventOccurrenceRepository,
	managerRepo *repomocks.MockManagerRepository,
	eventRepo *repomocks.MockEventRepository,
	locationRepo *repomocks.MockLocationRepository,
	regRepo *repomocks.MockRegistrationRepository,
	sc *stripemocks.MockStripeClient,
) *Handler {
	return NewHandler(eoRepo, managerRepo, eventRepo, locationRepo, regRepo, sc)
}

func TestHandler_CreateEventOccurrence(t *testing.T) {
	event := makeTestEvent()
	location := makeTestLocation()
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"

	start, _ := time.Parse(time.RFC3339, "2026-02-15T09:00:00+07:00")
	end, _ := time.Parse(time.RFC3339, "2026-02-15T11:00:00+07:00")

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
				input.AcceptLanguage = "en-US"
				input.Body.ManagerId = &testMid
				input.Body.EventId = event.ID
				input.Body.LocationId = location.ID
				input.Body.StartTime = start
				input.Body.EndTime = end
				input.Body.MaxAttendees = 15
				input.Body.Language = "en"
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("CreateEventOccurrence", mock.Anything, mock.AnythingOfType("*models.CreateEventOccurrenceInput")).
					Return(&models.EventOccurrence{
						ID:           uuid.New(),
						ManagerId:    &testMid,
						Event:        event,
						Location:     location,
						StartTime:    start,
						EndTime:      end,
						MaxAttendees: 15,
						Language:     "en",
						CurrEnrolled: 8,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
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
			mockS3 := new(s3mocks.S3ClientMock)
      mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3, mockRegRepo, mockStripeClient)
			ctx := context.Background()

			mockManagerRepo.On("GetManagerByID", mock.Anything, mock.Anything).Return(&models.Manager{
				ID:             testMid,
				UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
				OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				Role:           "Director",
			}, nil)
			mockEventRepo.On("GetEventByID", mock.Anything, mock.Anything).Return(&models.Event{
				ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				Title:            "Junior Robotics Workshop",
				Description:      "Learn the basics of robotics",
				OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				AgeRangeMin:      &eight,
				AgeRangeMax:      &twelve,
				Category:         []string{"science", "technology"},
				HeaderImageS3Key: &jpg,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}, nil)
			mockLocationRepo.On("GetLocationByID", mock.Anything, mock.Anything).Return(&models.Location{
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
			}, nil)

			eventOccurrence, err := handler.CreateEventOccurrence(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, eventOccurrence)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				if tt.input.Body.ManagerId != nil {
					assert.Equal(t, *tt.input.Body.ManagerId, *eventOccurrence.ManagerId)
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

func TestHandler_CreateEventOccurrence_AcceptLanguageInvariant(t *testing.T) {
	invalidLanguages := []struct {
		name string
		lang string
	}{
		{name: "empty AcceptLanguage", lang: ""},
		{name: "unsupported locale fr-FR", lang: "fr-FR"},
		{name: "lowercase en-us", lang: "en-us"},
		{name: "random string", lang: "invalid"},
	}

	for _, tt := range invalidLanguages {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := &models.CreateEventOccurrenceInput{}
			input.AcceptLanguage = tt.lang
			input.Body.EventId = uuid.MustParse("60000000-0000-0000-0000-000000000001")
			input.Body.LocationId = uuid.MustParse("10000000-0000-0000-0000-000000000004")

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockS3 := new(s3mocks.S3ClientMock)
			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3)

			occurrence, err := handler.CreateEventOccurrence(context.Background(), input)
			assert.Nil(t, occurrence, "expected nil occurrence for invalid AcceptLanguage")
			assert.NotNil(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}
}

func TestHandler_GetEventOccurrenceById_AcceptLanguageInvariant(t *testing.T) {
	invalidLanguages := []struct {
		name string
		lang string
	}{
		{name: "empty AcceptLanguage", lang: ""},
		{name: "unsupported locale fr-FR", lang: "fr-FR"},
		{name: "lowercase en-us", lang: "en-us"},
		{name: "random string", lang: "invalid"},
	}

	for _, tt := range invalidLanguages {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := &models.GetEventOccurrenceByIDInput{
				ID:             uuid.MustParse("70000000-0000-0000-0000-000000000001"),
				AcceptLanguage: tt.lang,
			}

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockS3 := new(s3mocks.S3ClientMock)
			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3)

			occurrence, err := handler.GetEventOccurrenceByID(context.Background(), input)
			assert.Nil(t, occurrence, "expected nil occurrence for invalid AcceptLanguage")
			assert.NotNil(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}
}

func TestHandler_UpdateEventOccurrence_AcceptLanguageInvariant(t *testing.T) {
	invalidLanguages := []struct {
		name string
		lang string
	}{
		{name: "empty AcceptLanguage", lang: ""},
		{name: "unsupported locale fr-FR", lang: "fr-FR"},
		{name: "lowercase en-us", lang: "en-us"},
		{name: "random string", lang: "invalid"},
	}

	for _, tt := range invalidLanguages {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := &models.UpdateEventOccurrenceInput{}
			input.AcceptLanguage = tt.lang
			input.ID = uuid.MustParse("70000000-0000-0000-0000-000000000002")

			mockRepo := new(repomocks.MockEventOccurrenceRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockS3 := new(s3mocks.S3ClientMock)
			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3)

			occurrence, err := handler.UpdateEventOccurrence(context.Background(), input)
			assert.Nil(t, occurrence, "expected nil occurrence for invalid AcceptLanguage")
			assert.NotNil(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}
}

func TestHandler_GetEventOccurrenceById(t *testing.T) {
	statusCodeNotFound := http.StatusNotFound
	messageSubstringNotFound := "Not found"

	event := makeTestEvent()
	location := makeTestLocation()
	start, _ := time.Parse(time.RFC3339, "2026-02-15T09:00:00+07:00")
	end, _ := time.Parse(time.RFC3339, "2026-02-15T11:00:00+07:00")

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
				m.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse("70000000-0000-0000-0000-000000000001")).
					Return(&models.EventOccurrence{
						ID:           uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						ManagerId:    &testMid,
						Event:        event,
						Location:     location,
						StartTime:    start,
						EndTime:      end,
						MaxAttendees: 15,
						Language:     "en",
						CurrEnrolled: 8,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "event occurrence not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetEventOccurrenceByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
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
			mockS3 := new(s3mocks.S3ClientMock)
			mockS3.On("GeneratePresignedURL", mock.Anything, mock.Anything, mock.Anything).Return("https://test-bucket.s3.amazonaws.com/presigned", nil)
      mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3, mockRegRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetEventOccurrenceByIDInput{ID: uuid.MustParse(tt.id), AcceptLanguage: "en-US"}
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
	midNew := uuid.MustParse("50000000-0000-0000-0000-000000000005")
	eid := uuid.MustParse("60000000-0000-0000-0000-00000000000e")
	lid := uuid.MustParse("10000000-0000-0000-0000-000000000008")
	startNew, _ := time.Parse(time.RFC3339, "2026-02-15T10:00:00+07:00")
	endNew, _ := time.Parse(time.RFC3339, "2026-02-15T12:00:00+07:00")
	max := 10
	lang := "th"
	curr := 8
	currBad := 20

	event := makeTestEvent()
	location := makeTestLocation()

	ten := 10
	fifteen := 15
	eventNew := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-00000000000e"),
		Title:            "Python for Kids",
		Description:      "Introduction to Python programming.",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000005"),
		AgeRangeMin:      &ten,
		AgeRangeMax:      &fifteen,
		Category:         []string{"technology", "math"},
		HeaderImageS3Key: nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	locationNew := models.Location{
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

	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"

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
				input.ID = testEOID
				input.Body.CurrEnrolled = &currBad
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetEventOccurrenceByID", mock.Anything, testEOID).
					Return(makeTestEventOccurrence(testStart), nil)
			},
			wantErr: true,
		},
		{
			name: "new current enrolled exceeds new max attendees",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.AcceptLanguage = "en-US"
				input.ID = testEOID
				input.Body.ManagerId = nil
				input.Body.EventId = nil
				input.Body.LocationId = nil
				input.Body.StartTime = nil
				input.Body.EndTime = nil
				input.Body.MaxAttendees = &max
				input.Body.CurrEnrolled = &currBad
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("GetEventOccurrenceByID", mock.Anything, testEOID).
					Return(makeTestEventOccurrence(testStart), nil)
			},
			wantErr: true,
		},
		{
			name: "successfully updated current enrolled",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.AcceptLanguage = "en-US"
				input.ID = testEOID
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
					mock.Anything,
				).Return(&models.EventOccurrence{
					ID:           testEOID,
					ManagerId:    &testMid,
					Event:        event,
					Location:     location,
					StartTime:    start,
					EndTime:      end,
					MaxAttendees: 15,
					Language:     "en",
					CurrEnrolled: 8,
					CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "successfully updated all fields except current enrolled",
			input: func() *models.UpdateEventOccurrenceInput {
				input := &models.UpdateEventOccurrenceInput{}
				input.AcceptLanguage = "en-US"
				input.ID = testEOID
				input.Body.ManagerId = &midNew
				input.Body.EventId = &eid
				input.Body.LocationId = &lid
				input.Body.StartTime = &startNew
				input.Body.EndTime = &endNew
				input.Body.MaxAttendees = &max
				input.Body.Language = &lang
				return input
			}(),
			mockSetup: func(m *repomocks.MockEventOccurrenceRepository) {
				m.On("UpdateEventOccurrence", mock.Anything, mock.AnythingOfType("*models.UpdateEventOccurrenceInput")).
					Return(&models.EventOccurrence{
						ID:           testEOID,
						ManagerId:    &midNew,
						Event:        eventNew,
						Location:     locationNew,
						StartTime:    startNew,
						EndTime:      endNew,
						MaxAttendees: 10,
						Language:     "th",
						CurrEnrolled: 5,
						CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
						UpdatedAt:    time.Now(),
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
			mockS3 := new(s3mocks.S3ClientMock)
      mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockS3, mockRegRepo, mockStripeClient)
			ctx := context.Background()

			if !tt.wantErr {
				mockRepo.On(
					"GetEventOccurrenceByID",
					mock.Anything,
					uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					mock.Anything,
				).Return(&models.EventOccurrence{
					ID:           uuid.MustParse("70000000-0000-0000-0000-000000000002"),
					ManagerId:    &mid,
					Event:        event,
					Location:     location,
					StartTime:    start,
					EndTime:      end,
					MaxAttendees: 15,
					Language:     "en",
					CurrEnrolled: 5,
					CreatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
					UpdatedAt:    time.Date(2026, time.January, 20, 21, 41, 2, 0, time.Local),
				}, nil)
			}

			mockManagerRepo.On("GetManagerByID", mock.Anything, mock.Anything).Return(&models.Manager{
				ID:             testMid,
				UserID:         uuid.MustParse("c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f"),
				OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				Role:           "Director",
			}, nil)
			mockEventRepo.On("GetEventByID", mock.Anything, mock.Anything).Return(&models.Event{
				ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
				Title:            "Junior Robotics Workshop",
				OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
				AgeRangeMin:      &eight,
				AgeRangeMax:      &twelve,
				Category:         []string{"science", "technology"},
				HeaderImageS3Key: &jpg,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}, nil)
			mockLocationRepo.On("GetLocationByID", mock.Anything, mock.Anything).Return(&models.Location{
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
			}, nil)

			eventOccurrence, err := handler.UpdateEventOccurrence(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, eventOccurrence)
			} else if tt.name == "successfully updated current enrolled" {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				assert.Equal(t, testMid, *eventOccurrence.ManagerId)
				assert.Equal(t, event.ID, eventOccurrence.Event.ID)
				assert.Equal(t, location.ID, eventOccurrence.Location.ID)
				assert.Equal(t, testStart, eventOccurrence.StartTime)
				assert.Equal(t, testEnd, eventOccurrence.EndTime)
				assert.Equal(t, 15, eventOccurrence.MaxAttendees)
				assert.Equal(t, "en", eventOccurrence.Language)
				assert.Equal(t, *tt.input.Body.CurrEnrolled, eventOccurrence.CurrEnrolled)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, eventOccurrence)
				assert.Equal(t, midNew, *eventOccurrence.ManagerId)
				assert.Equal(t, eid, eventOccurrence.Event.ID)
				assert.Equal(t, lid, eventOccurrence.Location.ID)
				assert.Equal(t, startNew, eventOccurrence.StartTime)
				assert.Equal(t, endNew, eventOccurrence.EndTime)
				assert.Equal(t, max, eventOccurrence.MaxAttendees)
				assert.Equal(t, lang, eventOccurrence.Language)
				assert.Equal(t, 5, eventOccurrence.CurrEnrolled)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CancelEventOccurrence(t *testing.T) {
	eoID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

	cancelledPaymentIntentOutput := &models.CancelPaymentIntentOutput{}
	cancelledPaymentIntentOutput.Body.PaymentIntentID = "pi_test_123"
	cancelledPaymentIntentOutput.Body.Status = "canceled"
	cancelledPaymentIntentOutput.Body.Amount = 10000
	cancelledPaymentIntentOutput.Body.Currency = "thb"

	refundOutput := &models.RefundPaymentOutput{}
	refundOutput.Body.RefundID = "re_test_123"
	refundOutput.Body.Status = "succeeded"
	refundOutput.Body.Amount = 10000
	refundOutput.Body.Currency = "thb"

	makeRegistrations := func(paymentStatus string) *models.GetRegistrationsByEventOccurrenceIDOutput {
		out := &models.GetRegistrationsByEventOccurrenceIDOutput{}
		out.Body.Registrations = []models.Registration{
			{
				ID:                    uuid.New(),
				EventOccurrenceID:     eoID,
				StripePaymentIntentID: "pi_test_123",
				OrgStripeAccountID:    "acct_test_123",
				PaymentIntentStatus:   paymentStatus,
				Status:                models.RegistrationStatusRegistered,
			},
		}
		return out
	}

	tests := []struct {
		name      string
		eoID      uuid.UUID
		mockSetup func(*repomocks.MockEventOccurrenceRepository, *repomocks.MockRegistrationRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name: "cancel with requires_capture registrations — cancels payment intents",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(makeRegistrations("requires_capture"), nil)

				sc.On("CancelPaymentIntent", mock.Anything, mock.AnythingOfType("*models.CancelPaymentIntentInput")).
					Return(cancelledPaymentIntentOutput, nil)

				eoRepo.On("CancelEventOccurrence", mock.Anything, eoID).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "cancel with succeeded registrations — issues refunds",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(makeRegistrations("succeeded"), nil)

				sc.On("RefundPayment", mock.Anything, mock.AnythingOfType("*models.RefundPaymentInput")).
					Return(refundOutput, nil)

				eoRepo.On("CancelEventOccurrence", mock.Anything, eoID).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "cancel with no registrations",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				out := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				out.Body.Registrations = []models.Registration{}
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(out, nil)

				eoRepo.On("CancelEventOccurrence", mock.Anything, eoID).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "get registrations fails",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "db error"})
			},
			wantErr: true,
		},
		{
			name: "refund payment fails — returns error",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(makeRegistrations("succeeded"), nil)

				sc.On("RefundPayment", mock.Anything, mock.AnythingOfType("*models.RefundPaymentInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "stripe error"})
			},
			wantErr: true,
		},
		{
			name: "cancel payment intent fails — returns error",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(makeRegistrations("requires_capture"), nil)

				sc.On("CancelPaymentIntent", mock.Anything, mock.AnythingOfType("*models.CancelPaymentIntentInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "stripe error"})
			},
			wantErr: true,
		},
		{
			name: "unknown payment intent status — returns error",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(makeRegistrations("requires_payment_method"), nil)
			},
			wantErr: true,
		},
		{
			name: "cancel event occurrence repo fails",
			eoID: eoID,
			mockSetup: func(eoRepo *repomocks.MockEventOccurrenceRepository, regRepo *repomocks.MockRegistrationRepository, sc *stripemocks.MockStripeClient) {
				out := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				out.Body.Registrations = []models.Registration{}
				regRepo.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).
					Return(out, nil)

				eoRepo.On("CancelEventOccurrence", mock.Anything, eoID).
					Return(&errs.HTTPError{Code: 500, Message: "db error"})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockEventRepo := new(repomocks.MockEventRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockEORepo, mockRegRepo, mockStripeClient)

			handler := newHandler(mockEORepo, mockManagerRepo, mockEventRepo, mockLocationRepo, mockRegRepo, mockStripeClient)
			ctx := context.Background()

			msg, err := handler.CancelEventOccurrence(ctx, tt.eoID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, msg)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, msg)
			}

			mockEORepo.AssertExpectations(t)
			mockRegRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}