package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetRegistrationByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockRegistrationRepository)
		wantErr   bool
	}{
		{
			name: "successful get registration",
			id:   "80000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				m.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID:                     uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:                uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "usd",
						PaymentIntentStatus:    "requires_capture",
						CancelledAt:            nil,
						PaidAt:                 nil,
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			id:   "80000000-0000-0000-0000-000000000099",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				m.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Registration", "id", "80000000-0000-0000-0000-000000000099").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetRegistrationByIDInput{ID: uuid.MustParse(tt.id)}
			registration, err := handler.GetRegistrationByID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
				assert.Equal(t, tt.id, registration.Body.ID.String())
				assert.NotEmpty(t, registration.Body.StripePaymentIntentID)
				assert.NotEmpty(t, registration.Body.Currency)
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetRegistrationsByChildID(t *testing.T) {
	tests := []struct {
		name        string
		childID     string
		mockSetup   func(*repomocks.MockRegistrationRepository)
		wantErr     bool
		expectedLen int
	}{
		{
			name:    "successful get registrations by child",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByChildIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                    uuid.New(),
						ChildID:               uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_1",
						StripeCustomerID:      "cus_test_123",
						TotalAmount:           10000,
						Currency:              "usd",
					},
					{
						ID:                    uuid.New(),
						ChildID:               uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_2",
						StripeCustomerID:      "cus_test_123",
						TotalAmount:           10000,
						Currency:              "usd",
					},
				}
				m.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 2,
		},
		{
			name:    "no registrations found",
			childID: "30000000-0000-0000-0000-000000000099",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByChildIDOutput{}
				output.Body.Registrations = []models.Registration{}
				m.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetRegistrationsByChildIDInput{ChildID: uuid.MustParse(tt.childID)}
			registrations, err := handler.GetRegistrationsByChildID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registrations)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registrations)
				assert.Equal(t, tt.expectedLen, len(registrations.Body.Registrations))
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetRegistrationsByGuardianID(t *testing.T) {
	tests := []struct {
		name        string
		guardianID  string
		mockSetup   func(*repomocks.MockRegistrationRepository)
		wantErr     bool
		expectedLen int
	}{
		{
			name:       "successful get registrations by guardian",
			guardianID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByGuardianIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentIntentID: "pi_test_1",
						Currency:              "usd",
					},
					{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentIntentID: "pi_test_2",
						Currency:              "usd",
					},
					{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentIntentID: "pi_test_3",
						Currency:              "usd",
					},
				}
				m.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 3,
		},
		{
			name:       "no registrations found",
			guardianID: "11111111-1111-1111-1111-111111111199",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByGuardianIDOutput{}
				output.Body.Registrations = []models.Registration{}
				m.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetRegistrationsByGuardianIDInput{GuardianID: uuid.MustParse(tt.guardianID)}
			registrations, err := handler.GetRegistrationsByGuardianID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registrations)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registrations)
				assert.Equal(t, tt.expectedLen, len(registrations.Body.Registrations))
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetRegistrationsByEventOccurrenceID(t *testing.T) {
	tests := []struct {
		name              string
		eventOccurrenceID string
		mockSetup         func(*repomocks.MockRegistrationRepository)
		wantErr           bool
		expectedLen       int
	}{
		{
			name:              "successful get registrations by event occurrence",
			eventOccurrenceID: "70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                    uuid.New(),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_1",
						Currency:              "usd",
					},
					{
						ID:                    uuid.New(),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_2",
						Currency:              "usd",
					},
					{
						ID:                    uuid.New(),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_3",
						Currency:              "usd",
					},
				}
				m.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 3,
		},
		{
			name:              "no registrations found",
			eventOccurrenceID: "70000000-0000-0000-0000-000000000099",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				output.Body.Registrations = []models.Registration{}
				m.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).Return(output, nil)
			},
			wantErr:     false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetRegistrationsByEventOccurrenceIDInput{EventOccurrenceID: uuid.MustParse(tt.eventOccurrenceID)}
			registrations, err := handler.GetRegistrationsByEventOccurrenceID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registrations)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registrations)
				assert.Equal(t, tt.expectedLen, len(registrations.Body.Registrations))
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateRegistration(t *testing.T) {
	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	eventOccurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")
	organizationID := uuid.MustParse("10000000-0000-0000-0000-000000000001")
	invalidChildID := uuid.New()
	invalidGuardianID := uuid.New()
	invalidEventOccurrenceID := uuid.New()

	tests := []struct {
		name      string
		input     *models.CreateRegistrationInput
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name: "successful create",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				i.Body.Currency = "usd"
				i.Body.PaymentMethodID = "pm_test_123"
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				stripeAccountID := "acct_test_123"
				stripeCustomerID := "cus_test_123"
				
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID).Return(&models.EventOccurrence{
					ID:        eventOccurrenceID,
					Price:     10000,
					StartTime: time.Now().Add(25 * time.Hour),
					Event: models.Event{
						ID:             uuid.New(),
						OrganizationID: organizationID,
						Title:          "STEM Club",
					},
				}, nil)
				
				childRepo.On("GetChildByID", mock.Anything, childID).Return(&models.Child{
					ID:         childID,
					GuardianID: guardianID,
				}, nil)
				
				guardianRepo.On("GetGuardianByID", mock.Anything, guardianID).Return(&models.Guardian{
					ID:               guardianID,
					StripeCustomerID: &stripeCustomerID,
				}, nil)
				
				orgRepo.On("GetOrganizationByID", mock.Anything, organizationID).Return(&models.Organization{
					ID:              organizationID,
					StripeAccountID: &stripeAccountID,
				}, nil)
				
				sc.On("CreatePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CreatePaymentIntentInput")).Return(&models.CreatePaymentIntentOutput{
					Body: struct {
						PaymentIntentID   string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
						ClientSecret      string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
						Status            string `json:"status" doc:"Payment intent status"`
						TotalAmount       int    `json:"total_amount" doc:"Total amount in cents"`
						ProviderAmount    int    `json:"provider_amount" doc:"Amount provider receives in cents"`
						PlatformFeeAmount int    `json:"platform_fee_amount" doc:"Platform fee in cents"`
						Currency          string `json:"currency" doc:"Currency code"`
					}{
						PaymentIntentID:   "pi_test_123",
						TotalAmount:       10000,
						ProviderAmount:    8500,
						PlatformFeeAmount: 1500,
						Currency:          "usd",
						Status:            "requires_capture",
					},
				}, nil)
				
				regRepo.On("CreateRegistration", mock.Anything, mock.AnythingOfType("*models.CreateRegistrationWithPaymentData")).Return(&models.CreateRegistrationOutput{
					Body: models.Registration{
						ID:                     uuid.New(),
						ChildID:                childID,
						GuardianID:             guardianID,
						EventOccurrenceID:      eventOccurrenceID,
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "usd",
						PaymentIntentStatus:    "requires_capture",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid event_occurrence_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = invalidEventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, invalidEventOccurrenceID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("EventOccurrence", "id", invalidEventOccurrenceID.String()).Code,
						Message: "Event occurrence not found",
					})
			},
			wantErr: true,
		},
		{
			name: "invalid child_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.Body.ChildID = invalidChildID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID).Return(&models.EventOccurrence{
					ID: eventOccurrenceID,
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, invalidChildID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "id", invalidChildID.String()).Code,
						Message: "Child not found",
					})
			},
			wantErr: true,
		},
		{
			name: "invalid guardian_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.Body.ChildID = childID
				i.Body.GuardianID = invalidGuardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID).Return(&models.EventOccurrence{
					ID: eventOccurrenceID,
				}, nil)
				childRepo.On("GetChildByID", mock.Anything, childID).Return(&models.Child{
					ID: childID,
				}, nil)
				guardianRepo.On("GetGuardianByID", mock.Anything, invalidGuardianID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", invalidGuardianID.String()).Code,
						Message: "Guardian not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			registration, err := handler.CreateRegistration(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
				assert.Equal(t, tt.input.Body.ChildID, registration.Body.ChildID)
			}

			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateRegistration(t *testing.T) {
	existingID := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	newChildID := uuid.MustParse("30000000-0000-0000-0000-000000000002")
	invalidChildID := uuid.New()

	tests := []struct {
		name      string
		input     *models.UpdateRegistrationInput
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name: "successful update child",
			input: func() *models.UpdateRegistrationInput {
				i := &models.UpdateRegistrationInput{ID: existingID}
				i.Body.ChildID = newChildID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				childRepo.On("GetChildByID", mock.Anything, newChildID).Return(&models.Child{
					ID: newChildID,
				}, nil)
				
				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                     existingID,
						ChildID:                newChildID,
						GuardianID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "usd",
						PaymentIntentStatus:    "requires_capture",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			input: func() *models.UpdateRegistrationInput {
				i := &models.UpdateRegistrationInput{ID: existingID}
				i.Body.ChildID = newChildID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				childRepo.On("GetChildByID", mock.Anything, newChildID).Return(&models.Child{
					ID: newChildID,
				}, nil)
				
				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Registration", "id", existingID.String()).Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name: "invalid child_id on update",
			input: func() *models.UpdateRegistrationInput {
				i := &models.UpdateRegistrationInput{ID: existingID}
				i.Body.ChildID = invalidChildID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				childRepo.On("GetChildByID", mock.Anything, invalidChildID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "id", invalidChildID.String()).Code,
						Message: "Child not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo, mockChildRepo)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)
			ctx := context.Background()

			registration, err := handler.UpdateRegistration(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
				assert.Equal(t, tt.input.Body.ChildID, registration.Body.ChildID)
			}
			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
		})
	}
}