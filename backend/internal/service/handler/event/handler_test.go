package event

import (
	"context"
	"net/url"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	s3mocks "skillspark/internal/s3_client/mocks"
	repomocks "skillspark/internal/storage/repo-mocks"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// createMockS3Client creates a mock S3 client for testing
func createMockS3Client() *s3mocks.S3ClientMock {
	return new(s3mocks.S3ClientMock)
}

// createDummyImageData
func createDummyImageData() *[]byte {
	data := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
	}
	return &data
}

func TestHandler_CreateEvent(t *testing.T) {
	headerKey := "events/a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11/header.jpg"

	tests := []struct {
		name         string
		input        *models.CreateEventInput
		updateBody   *models.UpdateEventBody
		imageData    *[]byte
		mockSetup    func(*repomocks.MockEventRepository)
		mockS3Setup  func(*s3mocks.S3ClientMock)
		wantErr      bool
		wantURL      bool
	}{
		{
			name: "successful create event without image",
			input: func() *models.CreateEventInput {
				input := &models.CreateEventInput{}
				ageMin := 8
				ageMax := 12

				input.Body.Title = "Junior Robotics"
				input.Body.Description = "Intro to robotics"
				input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
				input.Body.AgeRangeMin = &ageMin
				input.Body.AgeRangeMax = &ageMax
				input.Body.Category = []string{"stem", "robotics"}
				return input
			}(),
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*models.CreateEventInput"), mock.Anything).Return(&models.Event{
					ID:             uuid.New(),
					Title:          "Junior Robotics",
					Description:    "Intro to robotics",
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					AgeRangeMin:    func() *int { i := 8; return &i }(),
					AgeRangeMax:    func() *int { i := 12; return &i }(),
					Category:       []string{"stem", "robotics"},
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				// No S3 calls expected when no image
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful create event with image - returns valid presigned URL",
			input: func() *models.CreateEventInput {
				input := &models.CreateEventInput{}
				ageMin := 8
				ageMax := 12

				input.Body.Title = "Junior Robotics"
				input.Body.Description = "Intro to robotics"
				input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
				input.Body.AgeRangeMin = &ageMin
				input.Body.AgeRangeMax = &ageMax
				input.Body.Category = []string{"stem", "robotics"}
				return input
			}(),
			updateBody: &models.UpdateEventBody{},
			imageData:  createDummyImageData(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				eventID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*models.CreateEventInput"), mock.Anything).Return(&models.Event{
					ID:             eventID,
					Title:          "Junior Robotics",
					Description:    "Intro to robotics",
					OrganizationID: uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					AgeRangeMin:    func() *int { i := 8; return &i }(),
					AgeRangeMax:    func() *int { i := 12; return &i }(),
					Category:       []string{"stem", "robotics"},
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput"), mock.Anything).Return(&models.Event{
					ID:               eventID,
					Title:            "Junior Robotics",
					Description:      "Intro to robotics",
					OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					HeaderImageS3Key: &headerKey,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				mockURL := "https://test-bucket.s3.amazonaws.com/events/test/header.jpg?X-Amz-Signature=abc123"
				m.On("UploadImage", mock.Anything, mock.Anything, mock.Anything).Return(&mockURL, nil)
			},
			wantErr: false,
			wantURL: true,
		},
		{
			name: "internal server error on create",
			input: func() *models.CreateEventInput {
				input := &models.CreateEventInput{}
				input.Body.Title = "Error Event"
				return input
			}(),
			updateBody: nil,
			imageData:  nil,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("CreateEvent", mock.Anything, mock.AnythingOfType("*models.CreateEventInput"), mock.Anything).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				// No S3 calls expected on error
			},
			wantErr: true,
			wantURL: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			mockS3 := createMockS3Client()
			tt.mockS3Setup(mockS3)

			handler := NewHandler(mockRepo, mockS3)
			ctx := context.Background()

			event, err := handler.CreateEvent(ctx, tt.input, tt.updateBody, tt.imageData, mockS3)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, event)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, event)
				assert.Equal(t, tt.input.Body.Title, event.Title)
				assert.Equal(t, tt.input.Body.OrganizationID, event.OrganizationID)
			}

			if tt.wantURL {
				require.NotNil(t, event.PresignedURL, "expected presigned URL to be returned")
				parsedURL, parseErr := url.Parse(*event.PresignedURL)
				require.NoError(t, parseErr, "presigned URL should be valid")
				assert.True(t, strings.HasPrefix(parsedURL.Scheme, "http"), "URL should have http/https scheme")
			} else if !tt.wantErr {
				assert.Nil(t, event.PresignedURL, "expected no presigned URL when no image data provided")
			}

			mockRepo.AssertExpectations(t)
			mockS3.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateEvent(t *testing.T) {
	eventID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	headerKey := "events/a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11/header.jpg"

	tests := []struct {
		name        string
		input       *models.UpdateEventInput
		imageData   *[]byte
		mockSetup   func(*repomocks.MockEventRepository)
		mockS3Setup func(*s3mocks.S3ClientMock)
		wantErr     bool
		wantURL     bool
	}{
		{
			name: "successful update event without image",
			input: func() *models.UpdateEventInput {
				input := &models.UpdateEventInput{}
				input.ID = eventID
				title := "Updated Robotics"
				input.Body.Title = &title
				return input
			}(),
			imageData: nil,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput"), (*string)(nil)).Return(&models.Event{
					ID:          eventID,
					Title:       "Updated Robotics",
					Description: "Intro to robotics",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				// No S3 calls expected when no image
			},
			wantErr: false,
			wantURL: false,
		},
		{
			name: "successful update event with image - returns valid presigned URL",
			input: func() *models.UpdateEventInput {
				input := &models.UpdateEventInput{}
				input.ID = eventID
				title := "Updated Robotics with Image"
				input.Body.Title = &title
				return input
			}(),
			imageData: createDummyImageData(),
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput"), mock.Anything).Return(&models.Event{
					ID:               eventID,
					Title:            "Updated Robotics with Image",
					Description:      "Intro to robotics",
					HeaderImageS3Key: &headerKey,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				mockURL := "https://test-bucket.s3.amazonaws.com/events/test/header.jpg?X-Amz-Signature=abc123"
				m.On("UploadImage", mock.Anything, mock.Anything, mock.Anything).Return(&mockURL, nil)
			},
			wantErr: false,
			wantURL: true,
		},
		{
			name: "no occurrences found - returns message string",
			input: func() *models.UpdateEventInput {
				input := &models.UpdateEventInput{}
				input.ID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				return input
			}(),
			imageData: nil,
			mockSetup: func(m *repomocks.MockEventRepository) {
				m.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*models.UpdateEventInput"), (*string)(nil)).Return(&models.Event{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				// No S3 calls expected
			},
			wantErr: false,
			wantURL: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repomocks.MockEventRepository)
			tt.mockSetup(mockRepo)

			mockS3 := createMockS3Client()
			tt.mockS3Setup(mockS3)

			handler := NewHandler(mockRepo, mockS3)
			ctx := context.Background()

			event, err := handler.UpdateEvent(ctx, tt.input, tt.imageData, mockS3)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, event)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, event)
				if tt.input.Body.Title != nil {
					assert.Equal(t, *tt.input.Body.Title, event.Title)
				}
				assert.Equal(t, tt.input.ID, event.ID)
			}

			if tt.wantURL {
				require.NotNil(t, event.PresignedURL, "expected presigned URL to be returned")
				parsedURL, parseErr := url.Parse(*event.PresignedURL)
				require.NoError(t, parseErr, "presigned URL should be valid")
				assert.True(t, strings.HasPrefix(parsedURL.Scheme, "http"), "URL should have http/https scheme")
			}

			mockRepo.AssertExpectations(t)
			mockS3.AssertExpectations(t)
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

			mockS3 := createMockS3Client()
			handler := NewHandler(mockRepo, mockS3)
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

	category_arr := []string{"science", "technology"}
	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	addr := "Suite 15"
	mid := uuid.MustParse("50000000-0000-0000-0000-000000000001")
	event := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         category_arr,
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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
		mockS3Setup      func(*s3mocks.S3ClientMock)
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
						ID:           uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						ManagerId:    &mid,
						Event:        event,
						Location:     location,
						StartTime:    start,
						EndTime:      end,
						MaxAttendees: 15,
						Language:     "en",
						CurrEnrolled: 8,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
					{
						ID:           uuid.MustParse("70000000-0000-0000-0000-000000000002"),
						ManagerId:    &mid,
						Event:        event,
						Location:     location,
						StartTime:    start2,
						EndTime:      end2,
						MaxAttendees: 15,
						Language:     "en",
						CurrEnrolled: 5,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				}, nil)
			},
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				m.On("GeneratePresignedURL", mock.Anything, jpg, time.Hour).Return("https://test-bucket.s3.amazonaws.com/events/robotics_workshop.jpg", nil)
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
			mockS3Setup: func(m *s3mocks.S3ClientMock) {
				// No S3 calls expected when no occurrences
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

			mockS3 := createMockS3Client()
			tt.mockS3Setup(mockS3)

			handler := NewHandler(mockRepo, mockS3)
			ctx := context.Background()

			input := &models.GetEventOccurrencesByEventIDInput{ID: uuid.MustParse(tt.id)}
			eventOccurrences, err := handler.GetEventOccurrencesByEventID(ctx, input, mockS3)

			assert.Nil(t, err)
			assert.NotNil(t, eventOccurrences)
			if len(eventOccurrences) != 0 {
				assert.Equal(t, 2, len(eventOccurrences))
			} else {
				assert.Equal(t, 0, len(eventOccurrences))
			}
			mockRepo.AssertExpectations(t)
			mockS3.AssertExpectations(t)
		})
	}
}
