package jobs

import (
	"skillspark/internal/models"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeScheduler(
	mockRegRepo *repomocks.MockRegistrationRepository,
	mockGuardianRepo *repomocks.MockGuardianRepository,
	mockEORepo *repomocks.MockEventOccurrenceRepository,
	mockOrgRepo *repomocks.MockOrganizationRepository,
	mockStripeClient *stripemocks.MockStripeClient,
) *JobScheduler {
	return &JobScheduler{
		repo: &storage.Repository{
			Registration:    mockRegRepo,
			Guardian:        mockGuardianRepo,
			EventOccurrence: mockEORepo,
			Organization:    mockOrgRepo,
		},
		stripeClient: mockStripeClient,
	}
}

func TestCreatePaymentIntentsJob_Success(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID := uuid.New()
	eoID := uuid.New()
	orgID := uuid.New()
	regID := uuid.New()
	customerID := "cus_test_123"
	accountID := "acct_test_123"
	pmID := "pm_test_123"

	pending := []models.RegistrationForPayment{{
		ID:                regID,
		GuardianID:        guardianID,
		EventOccurrenceID: eoID,
	}}

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return(pending, nil)

	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID, StripeCustomerID: &customerID}, nil)

	mockStripeClient.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
		Return(&models.GetPaymentMethodsByGuardianIDOutput{
			Body: struct {
				PaymentMethods []models.PaymentMethod `json:"payment_methods"`
			}{
				PaymentMethods: []models.PaymentMethod{{ID: pmID}},
			},
		}, nil)

	mockEORepo.On("GetEventOccurrenceByID", mock.Anything, eoID, "en-US").
		Return(&models.EventOccurrence{
			ID:        eoID,
			StartTime: time.Now().Add(2 * 24 * time.Hour),
			Price:     10000,
			Currency:  "usd",
			Event:     models.Event{OrganizationID: orgID},
		}, nil)

	mockOrgRepo.On("GetOrganizationByID", mock.Anything, orgID).
		Return(&models.Organization{ID: orgID, StripeAccountID: &accountID}, nil)

	mockStripeClient.On("CreatePaymentIntent", mock.Anything, mock.MatchedBy(func(input *models.CreatePaymentIntentInput) bool {
		return input.Body.GuardianStripeID == customerID &&
			input.Body.OrgStripeID == accountID &&
			input.Body.PaymentMethodID == pmID &&
			input.Body.Amount == 10000 &&
			input.Body.Currency == "usd" &&
			input.Body.PlatformFeePercentage == 10
	})).Return(&models.CreatePaymentIntentOutput{
		Body: struct {
			PaymentIntentID   string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
			ClientSecret      string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
			Status            string `json:"status" doc:"Payment intent status"`
			TotalAmount       int    `json:"total_amount" doc:"Total amount in cents"`
			ProviderAmount    int    `json:"provider_amount" doc:"Amount provider receives in cents"`
			PlatformFeeAmount int    `json:"platform_fee_amount" doc:"Platform fee in cents"`
			Currency          string `json:"currency" doc:"Currency code"`
		}{
			PaymentIntentID:   "pi_new_123",
			Status:            "requires_capture",
			TotalAmount:       10000,
			ProviderAmount:    9000,
			PlatformFeeAmount: 1000,
			Currency:          "usd",
		},
	}, nil)

	mockRegRepo.On("CreatePayment", mock.Anything, mock.MatchedBy(func(input *models.CreatePaymentData) bool {
		return input.RegistrationID == regID &&
			input.StripePaymentIntentID == "pi_new_123" &&
			input.StripeCustomerID == customerID &&
			input.OrgStripeAccountID == accountID &&
			input.StripePaymentMethodID == pmID
	})).Return(nil)

	scheduler.CreatePaymentIntentsJob()

	mockRegRepo.AssertExpectations(t)
	mockGuardianRepo.AssertExpectations(t)
	mockEORepo.AssertExpectations(t)
	mockOrgRepo.AssertExpectations(t)
	mockStripeClient.AssertExpectations(t)
}

func TestCreatePaymentIntentsJob_NoRegistrations(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return([]models.RegistrationForPayment{}, nil)

	scheduler.CreatePaymentIntentsJob()

	mockRegRepo.AssertExpectations(t)
	mockGuardianRepo.AssertNotCalled(t, "GetGuardianByID")
	mockStripeClient.AssertNotCalled(t, "CreatePaymentIntent")
	mockRegRepo.AssertNotCalled(t, "CreatePayment")
}

func TestCreatePaymentIntentsJob_FetchError(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return(nil, assert.AnError)

	scheduler.CreatePaymentIntentsJob()

	mockRegRepo.AssertExpectations(t)
	mockGuardianRepo.AssertNotCalled(t, "GetGuardianByID")
	mockStripeClient.AssertNotCalled(t, "CreatePaymentIntent")
}

func TestCreatePaymentIntentsJob_SkipsGuardianWithNoStripeID(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID := uuid.New()
	regID := uuid.New()
	eoID := uuid.New()

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return([]models.RegistrationForPayment{{ID: regID, GuardianID: guardianID, EventOccurrenceID: eoID}}, nil)

	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID, StripeCustomerID: nil}, nil)

	scheduler.CreatePaymentIntentsJob()

	mockStripeClient.AssertNotCalled(t, "GetPaymentMethodsByCustomerID")
	mockStripeClient.AssertNotCalled(t, "CreatePaymentIntent")
	mockRegRepo.AssertNotCalled(t, "CreatePayment")
}

func TestCreatePaymentIntentsJob_SkipsGuardianWithNoPaymentMethods(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID := uuid.New()
	regID := uuid.New()
	eoID := uuid.New()
	customerID := "cus_no_methods"

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return([]models.RegistrationForPayment{{ID: regID, GuardianID: guardianID, EventOccurrenceID: eoID}}, nil)

	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID, StripeCustomerID: &customerID}, nil)

	mockStripeClient.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
		Return(&models.GetPaymentMethodsByGuardianIDOutput{
			Body: struct {
				PaymentMethods []models.PaymentMethod `json:"payment_methods"`
			}{},
		}, nil)

	scheduler.CreatePaymentIntentsJob()

	mockStripeClient.AssertNotCalled(t, "CreatePaymentIntent")
	mockRegRepo.AssertNotCalled(t, "CreatePayment")
}

func TestCreatePaymentIntentsJob_SkipsOrgWithNoStripeAccount(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID := uuid.New()
	eoID := uuid.New()
	orgID := uuid.New()
	regID := uuid.New()
	customerID := "cus_test_123"
	pmID := "pm_test_123"

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return([]models.RegistrationForPayment{{ID: regID, GuardianID: guardianID, EventOccurrenceID: eoID}}, nil)

	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID, StripeCustomerID: &customerID}, nil)

	mockStripeClient.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
		Return(&models.GetPaymentMethodsByGuardianIDOutput{
			Body: struct {
				PaymentMethods []models.PaymentMethod `json:"payment_methods"`
			}{
				PaymentMethods: []models.PaymentMethod{{ID: pmID}},
			},
		}, nil)

	mockEORepo.On("GetEventOccurrenceByID", mock.Anything, eoID, "en-US").
		Return(&models.EventOccurrence{
			ID:    eoID,
			Event: models.Event{OrganizationID: orgID},
		}, nil)

	mockOrgRepo.On("GetOrganizationByID", mock.Anything, orgID).
		Return(&models.Organization{ID: orgID, StripeAccountID: nil}, nil)

	scheduler.CreatePaymentIntentsJob()

	mockStripeClient.AssertNotCalled(t, "CreatePaymentIntent")
	mockRegRepo.AssertNotCalled(t, "CreatePayment")
}

func TestCreatePaymentIntentsJob_StripeCreateFailure(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID := uuid.New()
	eoID := uuid.New()
	orgID := uuid.New()
	regID := uuid.New()
	customerID := "cus_test_123"
	accountID := "acct_test_123"
	pmID := "pm_test_123"

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).
		Return([]models.RegistrationForPayment{{ID: regID, GuardianID: guardianID, EventOccurrenceID: eoID}}, nil)

	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID, StripeCustomerID: &customerID}, nil)

	mockStripeClient.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
		Return(&models.GetPaymentMethodsByGuardianIDOutput{
			Body: struct {
				PaymentMethods []models.PaymentMethod `json:"payment_methods"`
			}{
				PaymentMethods: []models.PaymentMethod{{ID: pmID}},
			},
		}, nil)

	mockEORepo.On("GetEventOccurrenceByID", mock.Anything, eoID, "en-US").
		Return(&models.EventOccurrence{
			ID:    eoID,
			Event: models.Event{OrganizationID: orgID},
		}, nil)

	mockOrgRepo.On("GetOrganizationByID", mock.Anything, orgID).
		Return(&models.Organization{ID: orgID, StripeAccountID: &accountID}, nil)

	mockStripeClient.On("CreatePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CreatePaymentIntentInput")).
		Return(nil, assert.AnError)

	scheduler.CreatePaymentIntentsJob()

	mockRegRepo.AssertNotCalled(t, "CreatePayment")
}

func TestCreatePaymentIntentsJob_ContinuesAfterSingleFailure(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockGuardianRepo := new(repomocks.MockGuardianRepository)
	mockEORepo := new(repomocks.MockEventOccurrenceRepository)
	mockOrgRepo := new(repomocks.MockOrganizationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	scheduler := makeScheduler(mockRegRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

	guardianID1 := uuid.New()
	guardianID2 := uuid.New()
	eoID1, eoID2 := uuid.New(), uuid.New()
	orgID := uuid.New()
	reg1ID, reg2ID := uuid.New(), uuid.New()
	customerID2 := "cus_test_456"
	accountID := "acct_test_123"
	pmID := "pm_test_123"

	pending := []models.RegistrationForPayment{
		{ID: reg1ID, GuardianID: guardianID1, EventOccurrenceID: eoID1},
		{ID: reg2ID, GuardianID: guardianID2, EventOccurrenceID: eoID2},
	}

	mockRegRepo.On("GetRegistrationsForPaymentCreation", mock.Anything).Return(pending, nil)

	// First registration: guardian lookup fails
	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID1).Return(nil, assert.AnError)

	// Second registration: succeeds
	mockGuardianRepo.On("GetGuardianByID", mock.Anything, guardianID2).
		Return(&models.Guardian{ID: guardianID2, StripeCustomerID: &customerID2}, nil)

	mockStripeClient.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID2).
		Return(&models.GetPaymentMethodsByGuardianIDOutput{
			Body: struct {
				PaymentMethods []models.PaymentMethod `json:"payment_methods"`
			}{
				PaymentMethods: []models.PaymentMethod{{ID: pmID}},
			},
		}, nil)

	mockEORepo.On("GetEventOccurrenceByID", mock.Anything, eoID2, "en-US").
		Return(&models.EventOccurrence{
			ID:    eoID2,
			Event: models.Event{OrganizationID: orgID},
		}, nil)

	mockOrgRepo.On("GetOrganizationByID", mock.Anything, orgID).
		Return(&models.Organization{ID: orgID, StripeAccountID: &accountID}, nil)

	mockStripeClient.On("CreatePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CreatePaymentIntentInput")).
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
				PaymentIntentID: "pi_new_456",
				Status:          "requires_capture",
			},
		}, nil)

	mockRegRepo.On("CreatePayment", mock.Anything, mock.MatchedBy(func(input *models.CreatePaymentData) bool {
		return input.RegistrationID == reg2ID
	})).Return(nil)

	scheduler.CreatePaymentIntentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertNumberOfCalls(t, "CreatePaymentIntent", 1)
	mockRegRepo.AssertNumberOfCalls(t, "CreatePayment", 1)
}
