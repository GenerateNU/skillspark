package review

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	translatemocks "skillspark/internal/translation/mocks"
	"skillspark/internal/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateReview(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateReviewInput
		mockSetup func(*repomocks.MockReviewRepository, *repomocks.MockRegistrationRepository, *repomocks.MockGuardianRepository, *translatemocks.TranslateMock)
		wantErr   bool
	}{
		{
			name: "successful create review",
			input: func() *models.CreateReviewInput {
				in := &models.CreateReviewInput{}
				in.AcceptLanguage = "en-US"
				in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				in.Body.Description = "Great event!"
				in.Body.Categories = []string{"fun", "engaging"}
				return in
			}(),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository, regRepo *repomocks.MockRegistrationRepository, guardianRepo *repomocks.MockGuardianRepository, translateMock *translatemocks.TranslateMock) {
				// translation succeeds
				translated := "งานยอดเยี่ยม!"
				result := map[string]*string{
					"Great event!": &translated,
				}
				translateMock.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(result, nil)

				// registration exists
				regRepo.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
					mock.Anything,
				).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID: uuid.MustParse("99999999-0000-0000-0000-000000000001"),
					},
				}, nil)

				// guardian exists
				guardianRepo.On("GetGuardianByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
					Return(&models.Guardian{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				// create review
				reviewRepo.On("CreateReview", mock.Anything, mock.AnythingOfType("*models.CreateReviewDBInput")).
					Return(&models.Review{
						ID:             uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						RegistrationID: uuid.MustParse("10000000-0000-0000-0000-000000000001"),
						GuardianID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						Description:    "Great event!",
						Categories:     []string{"fun", "engaging"},
						CreatedAt:      time.Now(),
						UpdatedAt:      time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "translation fails",
			input: func() *models.CreateReviewInput {
				in := &models.CreateReviewInput{}
				in.AcceptLanguage = "en-US"
				in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				in.Body.Description = "Great event!"
				in.Body.Categories = []string{"fun", "engaging"}
				return in
			}(),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository, regRepo *repomocks.MockRegistrationRepository, guardianRepo *repomocks.MockGuardianRepository, translateMock *translatemocks.TranslateMock) {
				// translation fails
				translateMock.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("translation service unavailable"))
			},
			wantErr: true,
		},
		{
			name: "invalid registration_id",
			input: func() *models.CreateReviewInput {
				in := &models.CreateReviewInput{}
				in.AcceptLanguage = "en-US"
				in.Body.RegistrationID = uuid.MustParse("99999999-0000-0000-0000-000000000000")
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				in.Body.Description = "Great event!"
				in.Body.Categories = []string{"fun", "engaging"}
				return in
			}(),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository, regRepo *repomocks.MockRegistrationRepository, guardianRepo *repomocks.MockGuardianRepository, translateMock *translatemocks.TranslateMock) {
				// translation succeeds
				translated := "งานยอดเยี่ยม!"
				result := map[string]*string{
					"Great event!": &translated,
				}
				translateMock.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(result, nil)

				regRepo.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
					mock.Anything,
				).Return(nil, errors.New("not found"))

			},
			wantErr: true,
		},
		{
			name: "invalid guardian_id",
			input: func() *models.CreateReviewInput {
				in := &models.CreateReviewInput{}
				in.AcceptLanguage = "en-US"
				in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
				in.Body.GuardianID = uuid.MustParse("22222222-2222-2222-2222-222222222222")
				in.Body.Description = "Great event!"
				in.Body.Categories = []string{"fun", "engaging"}
				return in
			}(),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository, regRepo *repomocks.MockRegistrationRepository, guardianRepo *repomocks.MockGuardianRepository, translateMock *translatemocks.TranslateMock) {
				// translation succeeds
				translated := "งานยอดเยี่ยม!"
				result := map[string]*string{
					"Great event!": &translated,
				}
				translateMock.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(result, nil)

				// registration exists
				regRepo.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
					mock.Anything,
				).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID: uuid.MustParse("99999999-0000-0000-0000-000000000001"),
					},
				}, nil)

				// guardian does NOT exist
				guardianRepo.On("GetGuardianByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
					Return(nil, &errs.HTTPError{
						Code:    errs.BadRequest("guardian not found").Code,
						Message: "guardian not found",
					})
			},
			wantErr: true,
		},
		{
			name: "create review fails in repository",
			input: func() *models.CreateReviewInput {
				in := &models.CreateReviewInput{}
				in.AcceptLanguage = "en-US"
				in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				in.Body.Description = "Great event!"
				in.Body.Categories = []string{"fun", "engaging"}
				return in
			}(),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository, regRepo *repomocks.MockRegistrationRepository, guardianRepo *repomocks.MockGuardianRepository, translateMock *translatemocks.TranslateMock) {
				// translation succeeds
				translated := "งานยอดเยี่ยม!"
				result := map[string]*string{
					"Great event!": &translated,
				}
				translateMock.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(result, nil)

				// registration exists
				regRepo.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
					mock.Anything,
				).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID: uuid.MustParse("99999999-0000-0000-0000-000000000001"),
					},
				}, nil)

				// guardian exists
				guardianRepo.On("GetGuardianByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
					Return(&models.Guardian{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				// repository returns error
				reviewRepo.On("CreateReview", mock.Anything, mock.AnythingOfType("*models.CreateReviewDBInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.BadRequest("Invalid review").Code,
						Message: "Invalid review",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReviewRepo := new(repomocks.MockReviewRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockTranslate := new(translatemocks.TranslateMock)
			tt.mockSetup(mockReviewRepo, mockRegRepo, mockGuardianRepo, mockTranslate)

			handler := &Handler{
				ReviewRepository:       mockReviewRepo,
				RegistrationRepository: mockRegRepo,
				GuardianRepository:     mockGuardianRepo,
				TranslateClient:        mockTranslate,
			}
			ctx := context.Background()

			out, err := handler.CreateReview(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, out)
			} else {
				assert.NotNil(t, out)
				assert.Equal(t, tt.input.Body.Description, out.Body.Description)
				assert.Equal(t, tt.input.Body.RegistrationID, out.Body.RegistrationID)
				assert.Equal(t, tt.input.Body.GuardianID, out.Body.GuardianID)
				assert.Equal(t, tt.input.Body.Categories, out.Body.Categories)
			}

			mockReviewRepo.AssertExpectations(t)
			mockRegRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockTranslate.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateReview_AcceptLanguageInvariant(t *testing.T) {
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

			in := &models.CreateReviewInput{}
			in.AcceptLanguage = tt.lang
			in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
			in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
			in.Body.Description = "Great event!"
			in.Body.Categories = []string{"fun", "engaging"}

			handler := &Handler{
				ReviewRepository:       new(repomocks.MockReviewRepository),
				RegistrationRepository: new(repomocks.MockRegistrationRepository),
				GuardianRepository:     new(repomocks.MockGuardianRepository),
				TranslateClient:        new(translatemocks.TranslateMock),
			}

			out, err := handler.CreateReview(context.Background(), in)
			assert.Nil(t, out, "expected nil output for invalid AcceptLanguage")
			assert.NotNil(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}

	// Verify both valid languages are accepted (no early rejection)
	validLanguages := []string{"en-US", "th-TH"}
	for _, lang := range validLanguages {
		lang := lang
		t.Run("accepted_"+lang, func(t *testing.T) {
			t.Parallel()

			in := &models.CreateReviewInput{}
			in.AcceptLanguage = lang
			in.Body.RegistrationID = uuid.MustParse("10000000-0000-0000-0000-000000000001")
			in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
			in.Body.Description = "Great event!"
			in.Body.Categories = []string{"fun", "engaging"}

			mockTranslate := new(translatemocks.TranslateMock)
			translated := "translated"
			result := map[string]*string{"Great event!": &translated}
			mockTranslate.On("CallTranslateAPI", mock.Anything, mock.Anything, mock.Anything).Return(result, nil)

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockRegRepo.On("GetRegistrationByID", mock.Anything, mock.Anything, mock.Anything).Return(&models.GetRegistrationByIDOutput{
				Body: models.Registration{ID: uuid.MustParse("99999999-0000-0000-0000-000000000001")},
			}, nil)

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockGuardianRepo.On("GetGuardianByID", mock.Anything, mock.Anything).Return(&models.Guardian{
				ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			}, nil)

			mockReviewRepo := new(repomocks.MockReviewRepository)
			mockReviewRepo.On("CreateReview", mock.Anything, mock.Anything).Return(&models.Review{
				ID:          uuid.New(),
				Description: "Great event!",
				Categories:  []string{"fun", "engaging"},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, nil)

			handler := &Handler{
				ReviewRepository:       mockReviewRepo,
				RegistrationRepository: mockRegRepo,
				GuardianRepository:     mockGuardianRepo,
				TranslateClient:        mockTranslate,
			}

			out, err := handler.CreateReview(context.Background(), in)
			assert.Nil(t, err, "expected no error for valid AcceptLanguage %s", lang)
			assert.NotNil(t, out, "expected non-nil output for valid AcceptLanguage %s", lang)
		})
	}
}

func TestHandler_GetReviewsByEventID_AcceptLanguageInvariant(t *testing.T) {
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

			handler := &Handler{
				EventRepository:  new(repomocks.MockEventRepository),
				ReviewRepository: new(repomocks.MockReviewRepository),
			}

			pagination := utils.Pagination{Page: 1, Limit: 10}
			reviews, err := handler.GetReviewsByEventID(context.Background(), uuid.New(), tt.lang, pagination)
			assert.Nil(t, reviews, "expected nil reviews for invalid AcceptLanguage")
			assert.Error(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}
}

func TestHandler_GetReviewsByGuardianID_AcceptLanguageInvariant(t *testing.T) {
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

			handler := &Handler{
				GuardianRepository: new(repomocks.MockGuardianRepository),
				ReviewRepository:   new(repomocks.MockReviewRepository),
			}

			pagination := utils.Pagination{Page: 1, Limit: 10}
			reviews, err := handler.GetReviewsByGuardianID(context.Background(), uuid.New(), tt.lang, pagination)
			assert.Nil(t, reviews, "expected nil reviews for invalid AcceptLanguage")
			assert.Error(t, err, "expected error for invalid AcceptLanguage")
			assert.Contains(t, err.Error(), "Invalid AcceptLanguage")
		})
	}
}

func TestHandler_DeleteReview(t *testing.T) {
	tests := []struct {
		name      string
		reviewID  uuid.UUID
		mockSetup func(*repomocks.MockReviewRepository)
		wantMsg   string
		wantErr   bool
	}{
		{
			name:     "successful deletion",
			reviewID: uuid.MustParse("20000000-0000-0000-0000-000000000001"),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository) {
				reviewRepo.On("DeleteReview", mock.Anything, mock.AnythingOfType("uuid.UUID")).
					Return(nil)
			},
			wantMsg: "Review successfully deleted.",
			wantErr: false,
		},
		{
			name:     "repository returns error",
			reviewID: uuid.MustParse("20000000-0000-0000-0000-000000000002"),
			mockSetup: func(reviewRepo *repomocks.MockReviewRepository) {
				reviewRepo.On("DeleteReview", mock.Anything, mock.AnythingOfType("uuid.UUID")).
					Return(&errs.HTTPError{
						Code:    errs.BadRequest("cannot delete review").Code,
						Message: "cannot delete review",
					})
			},
			wantMsg: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockReviewRepo)

			handler := &Handler{
				ReviewRepository: mockReviewRepo,
			}

			msg, err := handler.DeleteReview(context.Background(), tt.reviewID)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, "", msg)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantMsg, msg)
			}

			mockReviewRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetReviewsByEventID(t *testing.T) {
	tests := []struct {
		name           string
		eventID        uuid.UUID
		acceptLanguage string
		mockSetup      func(*repomocks.MockEventRepository, *repomocks.MockReviewRepository)
		wantReviews    []models.Review
		wantErr        bool
	}{
		{
			name:           "successful fetch reviews",
			eventID:        uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			acceptLanguage: "en-US",
			mockSetup: func(eventRepo *repomocks.MockEventRepository, reviewRepo *repomocks.MockReviewRepository) {
				// Event exists
				eventRepo.On("GetEventByID", mock.Anything, uuid.MustParse("10000000-0000-0000-0000-000000000001"), "en-US").
					Return(&models.Event{ID: uuid.MustParse("10000000-0000-0000-0000-000000000001")}, nil)

				// Reviews returned
				reviewRepo.On("GetReviewsByEventID", mock.Anything, uuid.MustParse("10000000-0000-0000-0000-000000000001"), "en-US", mock.AnythingOfType("utils.Pagination")).
					Return([]models.Review{
						{
							ID:             uuid.MustParse("20000000-0000-0000-0000-000000000001"),
							RegistrationID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
							GuardianID:     uuid.MustParse("40000000-0000-0000-0000-000000000001"),
							Description:    "Great event!",
							Categories:     []string{"fun", "engaging"},
							CreatedAt:      time.Now(),
							UpdatedAt:      time.Now(),
						},
					}, nil)
			},
			wantReviews: []models.Review{
				{
					ID:             uuid.MustParse("20000000-0000-0000-0000-000000000001"),
					RegistrationID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
					GuardianID:     uuid.MustParse("40000000-0000-0000-0000-000000000001"),
					Description:    "Great event!",
					Categories:     []string{"fun", "engaging"},
				},
			},
			wantErr: false,
		},
		{
			name:           "event does not exist",
			eventID:        uuid.MustParse("99999999-0000-0000-0000-000000000000"),
			acceptLanguage: "en-US",
			mockSetup: func(eventRepo *repomocks.MockEventRepository, reviewRepo *repomocks.MockReviewRepository) {
				eventRepo.On("GetEventByID", mock.Anything, uuid.MustParse("99999999-0000-0000-0000-000000000000"), "en-US").
					Return(nil, errs.BadRequest("event does not exist"))
			},
			wantReviews: nil,
			wantErr:     true,
		},
		{
			name:           "review repository error",
			eventID:        uuid.MustParse("10000000-0000-0000-0000-000000000002"),
			acceptLanguage: "en-US",
			mockSetup: func(eventRepo *repomocks.MockEventRepository, reviewRepo *repomocks.MockReviewRepository) {
				eventRepo.On("GetEventByID", mock.Anything, uuid.MustParse("10000000-0000-0000-0000-000000000002"), "en-US").
					Return(&models.Event{ID: uuid.MustParse("10000000-0000-0000-0000-000000000002")}, nil)

				reviewRepo.On("GetReviewsByEventID", mock.Anything, uuid.MustParse("10000000-0000-0000-0000-000000000002"), "en-US", mock.AnythingOfType("utils.Pagination")).
					Return(nil, errs.BadRequest("cannot fetch reviews"))
			},
			wantReviews: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockEventRepo := new(repomocks.MockEventRepository)
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockEventRepo, mockReviewRepo)

			handler := &Handler{
				EventRepository:  mockEventRepo,
				ReviewRepository: mockReviewRepo,
			}

			pagination := utils.Pagination{Page: 1, Limit: 10}
			reviews, err := handler.GetReviewsByEventID(context.Background(), tt.eventID, tt.acceptLanguage, pagination)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, reviews)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, reviews)
				assert.Equal(t, len(tt.wantReviews), len(reviews))
				for i := range reviews {
					assert.Equal(t, tt.wantReviews[i].ID, reviews[i].ID)
					assert.Equal(t, tt.wantReviews[i].Description, reviews[i].Description)
					assert.Equal(t, tt.wantReviews[i].Categories, reviews[i].Categories)
				}
			}

			mockEventRepo.AssertExpectations(t)
			mockReviewRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetReviewsByGuardianID(t *testing.T) {
	tests := []struct {
		name           string
		guardianID     uuid.UUID
		acceptLanguage string
		mockSetup      func(*repomocks.MockGuardianRepository, *repomocks.MockReviewRepository)
		wantReviews    []models.Review
		wantErr        bool
	}{
		{
			name:           "successful fetch reviews",
			guardianID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			acceptLanguage: "en-US",
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, reviewRepo *repomocks.MockReviewRepository) {
				// Guardian exists
				guardianRepo.On("GetGuardianByID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&models.Guardian{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				// Reviews returned
				reviewRepo.On("GetReviewsByGuardianID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111"), "en-US", mock.AnythingOfType("utils.Pagination")).
					Return([]models.Review{
						{
							ID:             uuid.MustParse("20000000-0000-0000-0000-000000000001"),
							RegistrationID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
							GuardianID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							Description:    "Amazing!",
							Categories:     []string{"fun", "informative"},
							CreatedAt:      time.Now(),
							UpdatedAt:      time.Now(),
						},
					}, nil)
			},
			wantReviews: []models.Review{
				{
					ID:             uuid.MustParse("20000000-0000-0000-0000-000000000001"),
					RegistrationID: uuid.MustParse("30000000-0000-0000-0000-000000000001"),
					GuardianID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Description:    "Amazing!",
					Categories:     []string{"fun", "informative"},
				},
			},
			wantErr: false,
		},
		{
			name:           "review repository error",
			guardianID:     uuid.MustParse("11111111-1111-1111-1111-111111111112"),
			acceptLanguage: "en-US",
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, reviewRepo *repomocks.MockReviewRepository) {
				// Guardian exists
				guardianRepo.On("GetGuardianByID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111112")).
					Return(&models.Guardian{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					}, nil)

				// Repository error
				reviewRepo.On("GetReviewsByGuardianID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111112"), "en-US", mock.AnythingOfType("utils.Pagination")).
					Return(nil, errs.BadRequest("cannot fetch reviews"))
			},
			wantReviews: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockReviewRepo := new(repomocks.MockReviewRepository)
			tt.mockSetup(mockGuardianRepo, mockReviewRepo)

			handler := &Handler{
				GuardianRepository: mockGuardianRepo,
				ReviewRepository:   mockReviewRepo,
			}

			pagination := utils.Pagination{Page: 1, Limit: 10}
			reviews, err := handler.GetReviewsByGuardianID(context.Background(), tt.guardianID, tt.acceptLanguage, pagination)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, reviews)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, reviews)
				assert.Equal(t, len(tt.wantReviews), len(reviews))
				for i := range reviews {
					assert.Equal(t, tt.wantReviews[i].ID, reviews[i].ID)
					assert.Equal(t, tt.wantReviews[i].Description, reviews[i].Description)
					assert.Equal(t, tt.wantReviews[i].Categories, reviews[i].Categories)
				}
			}

			mockGuardianRepo.AssertExpectations(t)
			mockReviewRepo.AssertExpectations(t)
		})
	}
}
