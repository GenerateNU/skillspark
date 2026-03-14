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
						ID:                    uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:               uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                models.RegistrationStatusRegistered,
						EventName:             "STEM Club",
						OccurrenceStartTime:   time.Now(),
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
						StripePaymentIntentID: "pi_test_123",
						StripeCustomerID:      "cus_test_123",
						OrgStripeAccountID:    "acct_test_123",
						StripePaymentMethodID: "pm_test_123",
						TotalAmount:           10000,
						ProviderAmount:        8500,
						PlatformFeeAmount:     1500,
						Currency:              "thb",
						PaymentIntentStatus:   "requires_capture",
						CancelledAt:           nil,
						PaidAt:                nil,
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			input := &models.GetRegistrationByIDInput{
				AcceptLanguage: "en-US",
				ID:             uuid.MustParse(tt.id),
			}
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
						Currency:              "thb",
					},
					{
						ID:                    uuid.New(),
						ChildID:               uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_2",
						StripeCustomerID:      "cus_test_123",
						TotalAmount:           10000,
						Currency:              "thb",
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			input := &models.GetRegistrationsByChildIDInput{AcceptLanguage: "en-US", ChildID: uuid.MustParse(tt.childID)}
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
						Currency:              "thb",
					},
					{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentIntentID: "pi_test_2",
						Currency:              "thb",
					},
					{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentIntentID: "pi_test_3",
						Currency:              "thb",
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			input := &models.GetRegistrationsByGuardianIDInput{AcceptLanguage: "en-US", GuardianID: uuid.MustParse(tt.guardianID)}
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
						Currency:              "thb",
					},
					{
						ID:                    uuid.New(),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_2",
						Currency:              "thb",
					},
					{
						ID:                    uuid.New(),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						StripePaymentIntentID: "pi_test_3",
						Currency:              "thb",
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			input := &models.GetRegistrationsByEventOccurrenceIDInput{AcceptLanguage: "en-US", EventOccurrenceID: uuid.MustParse(tt.eventOccurrenceID)}
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
	stripeAccountID := "acct_test_123"
	stripeCustomerID := "cus_test_123"
	paymentMethodID := "pm_test_123"

	invalidChildID := uuid.New()
	invalidGuardianID := uuid.New()
	invalidEventOccurrenceID := uuid.New()

	validEventOccurrence := &models.EventOccurrence{
		ID:           eventOccurrenceID,
		Price:        10000,
		Currency:     "thb",
		StartTime:    time.Now().Add(25 * time.Hour),
		CurrEnrolled: 5,
		MaxAttendees: 15,
		Event: models.Event{
			ID:             uuid.New(),
			OrganizationID: organizationID,
			Title:          "STEM Club",
		},
	}

	validGuardian := &models.Guardian{
		ID:               guardianID,
		StripeCustomerID: &stripeCustomerID,
	}

	validChild := &models.Child{
		ID:         childID,
		GuardianID: guardianID,
	}

	validOrg := &models.Organization{
		ID:              organizationID,
		StripeAccountID: &stripeAccountID,
	}

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
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				i.Body.PaymentMethodID = &paymentMethodID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
					Return(validGuardian, nil)

				childRepo.On("GetChildByID", mock.Anything, childID).
					Return(validChild, nil)

				orgRepo.On("GetOrganizationByID", mock.Anything, organizationID).
					Return(validOrg, nil)

				sc.On("CreatePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CreatePaymentIntentInput")).
					Return(&models.CreatePaymentIntentOutput{
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
							Currency:          "thb",
							Status:            "requires_capture",
						},
					}, nil)

				regRepo.On("CreateRegistration", mock.Anything, mock.AnythingOfType("*models.CreateRegistrationWithPaymentData")).
					Return(&models.CreateRegistrationOutput{
						Body: models.Registration{
							ID:                    uuid.New(),
							ChildID:               childID,
							GuardianID:            guardianID,
							EventOccurrenceID:     eventOccurrenceID,
							Status:                models.RegistrationStatusRegistered,
							EventName:             "STEM Club",
							OccurrenceStartTime:   time.Now(),
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
							StripePaymentIntentID: "pi_test_123",
							StripeCustomerID:      stripeCustomerID,
							OrgStripeAccountID:    stripeAccountID,
							StripePaymentMethodID: paymentMethodID,
							TotalAmount:           10000,
							ProviderAmount:        8500,
							PlatformFeeAmount:     1500,
							Currency:              "thb",
							PaymentIntentStatus:   "requires_capture",
						},
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "invalid event_occurrence_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = invalidEventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, invalidEventOccurrenceID, mock.Anything).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("EventOccurrence", "id", invalidEventOccurrenceID.String()).Code,
						Message: "Event occurrence not found",
					})
			},
			wantErr: true,
		},
		{
			name: "invalid guardian_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = invalidGuardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, invalidGuardianID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", invalidGuardianID.String()).Code,
						Message: "Guardian not found",
					})
			},
			wantErr: true,
		},
		{
			name: "invalid child_id",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = invalidChildID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
					Return(validGuardian, nil)

				childRepo.On("GetChildByID", mock.Anything, invalidChildID).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "id", invalidChildID.String()).Code,
						Message: "Child not found",
					})
			},
			wantErr: true,
		},
		{
			name: "guardian missing stripe customer ID",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
					Return(&models.Guardian{ID: guardianID, StripeCustomerID: nil}, nil)
			},
			wantErr: true,
		},
		{
			name: "child does not belong to guardian",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				i.Body.PaymentMethodID = &paymentMethodID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
					Return(validGuardian, nil)

				childRepo.On("GetChildByID", mock.Anything, childID).
					Return(&models.Child{ID: childID, GuardianID: uuid.New()}, nil)
			},
			wantErr: true,
		},
		{
			name: "event occurrence already started",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				i.Body.PaymentMethodID = &paymentMethodID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(&models.EventOccurrence{
						ID:           eventOccurrenceID,
						Price:        10000,
						Currency:     "thb",
						StartTime:    time.Now().Add(-1 * time.Hour),
						CurrEnrolled: 5,
						MaxAttendees: 15,
						Event: models.Event{
							ID:             uuid.New(),
							OrganizationID: organizationID,
							Title:          "STEM Club",
						},
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "event occurrence at max capacity",
			input: func() *models.CreateRegistrationInput {
				i := &models.CreateRegistrationInput{}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = childID
				i.Body.GuardianID = guardianID
				i.Body.EventOccurrenceID = eventOccurrenceID
				i.Body.Status = models.RegistrationStatusRegistered
				i.Body.PaymentMethodID = &paymentMethodID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(&models.EventOccurrence{
						ID:           eventOccurrenceID,
						Price:        10000,
						Currency:     "thb",
						StartTime:    time.Now().Add(25 * time.Hour),
						CurrEnrolled: 15,
						MaxAttendees: 15,
						Event: models.Event{
							ID:             uuid.New(),
							OrganizationID: organizationID,
							Title:          "STEM Club",
						},
					}, nil)
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			registration, err := handler.CreateRegistration(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
				assert.Equal(t, tt.input.Body.ChildID, registration.Body.ChildID)
				assert.Equal(t, tt.input.Body.GuardianID, registration.Body.GuardianID)
				assert.Equal(t, tt.input.Body.EventOccurrenceID, registration.Body.EventOccurrenceID)
			}

			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
			mockOrgRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
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
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = &newChildID
				return i
			}(),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				childRepo.On("GetChildByID", mock.Anything, newChildID).Return(&models.Child{
					ID: newChildID,
				}, nil)

				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                    existingID,
						ChildID:               newChildID,
						GuardianID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                models.RegistrationStatusRegistered,
						EventName:             "STEM Club",
						OccurrenceStartTime:   time.Now(),
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
						StripePaymentIntentID: "pi_test_123",
						StripeCustomerID:      "cus_test_123",
						OrgStripeAccountID:    "acct_test_123",
						StripePaymentMethodID: "pm_test_123",
						TotalAmount:           10000,
						ProviderAmount:        8500,
						PlatformFeeAmount:     1500,
						Currency:              "thb",
						PaymentIntentStatus:   "requires_capture",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "registration not found",
			input: func() *models.UpdateRegistrationInput {
				i := &models.UpdateRegistrationInput{ID: existingID}
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = &newChildID
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
				i.AcceptLanguage = "en-US"
				i.Body.ChildID = &invalidChildID
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

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			registration, err := handler.UpdateRegistration(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
				if tt.input.Body.ChildID != nil {
					assert.Equal(t, *tt.input.Body.ChildID, registration.Body.ChildID)
				}
			}
			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CancelRegistration(t *testing.T) {
	registrationID := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	eventOccurrenceID := uuid.MustParse("70000000-0000-0000-0000-000000000001")

	validRegistration := func(paymentStatus string, startTime time.Time) *models.GetRegistrationByIDOutput {
		return &models.GetRegistrationByIDOutput{
			Body: models.Registration{
				ID:                    registrationID,
				EventOccurrenceID:     eventOccurrenceID,
				Status:                models.RegistrationStatusRegistered,
				StripePaymentIntentID: "pi_test_123",
				OrgStripeAccountID:    "acct_test_123",
				Currency:              "thb",
				TotalAmount:           10000,
				PaymentIntentStatus:   paymentStatus,
				OccurrenceStartTime:   startTime,
			},
		}
	}

	validEventOccurrence := func(startTime time.Time) *models.EventOccurrence {
		return &models.EventOccurrence{
			ID:        eventOccurrenceID,
			StartTime: startTime,
		}
	}

	cancelledOutput := &models.CancelRegistrationOutput{}
	cancelledOutput.Body.Message = "Registration cancelled successfully"
	cancelledOutput.Body.Registration = models.Registration{
		ID:     registrationID,
		Status: models.RegistrationStatusCancelled,
	}

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

	tests := []struct {
		name      string
		input     *models.CancelRegistrationInput
		mockSetup func(*repomocks.MockRegistrationRepository, *repomocks.MockEventOccurrenceRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "cancel requires_capture — cancels payment intent",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(validRegistration("requires_capture", time.Now().Add(48*time.Hour)), nil)

				sc.On("CancelPaymentIntent", mock.Anything, mock.AnythingOfType("*models.CancelPaymentIntentInput")).
					Return(cancelledPaymentIntentOutput, nil)

				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(cancelledOutput, nil)
			},
			wantErr: false,
		},
		{
			name:  "cancel succeeded — event more than 24hrs away — issues refund",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(validRegistration("succeeded", time.Now().Add(48*time.Hour)), nil)

				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence(time.Now().Add(48*time.Hour)), nil)

				sc.On("RefundPayment", mock.Anything, mock.AnythingOfType("*models.RefundPaymentInput")).
					Return(refundOutput, nil)

				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(cancelledOutput, nil)
			},
			wantErr: false,
		},
		{
			name:  "cancel succeeded — event within 24hrs — no refund",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(validRegistration("succeeded", time.Now().Add(12*time.Hour)), nil)

				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence(time.Now().Add(12*time.Hour)), nil)

				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(cancelledOutput, nil)
			},
			wantErr: false,
		},
		{
			name:  "registration already cancelled",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(&models.GetRegistrationByIDOutput{
						Body: models.Registration{
							ID:     registrationID,
							Status: models.RegistrationStatusCancelled,
						},
					}, nil)
			},
			wantErr: true,
		},
		{
			name:  "registration not found",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: uuid.New()},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Registration", "id", uuid.New().String()).Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
		{
			name:  "stripe cancel fails",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(validRegistration("requires_capture", time.Now().Add(48*time.Hour)), nil)

				sc.On("CancelPaymentIntent", mock.Anything, mock.AnythingOfType("*models.CancelPaymentIntentInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "stripe error"})
			},
			wantErr: true,
		},
		{
			name:  "stripe refund fails",
			input: &models.CancelRegistrationInput{AcceptLanguage: "en-US", ID: registrationID},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, eoRepo *repomocks.MockEventOccurrenceRepository, sc *stripemocks.MockStripeClient) {
				regRepo.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(validRegistration("succeeded", time.Now().Add(48*time.Hour)), nil)

				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceID, mock.Anything).
					Return(validEventOccurrence(time.Now().Add(48*time.Hour)), nil)

				sc.On("RefundPayment", mock.Anything, mock.AnythingOfType("*models.RefundPaymentInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "stripe refund error"})
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
			tt.mockSetup(mockRegRepo, mockEORepo, mockStripeClient)

			handler := NewHandler(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient, nil)
			ctx := context.Background()

			result, err := handler.CancelRegistration(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "Registration cancelled successfully", result.Body.Message)
			}

			mockRegRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}
