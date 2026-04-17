package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v84"
)

var (
	testOrgID      = uuid.MustParse("40000000-0000-0000-0000-000000000001")
	testGuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")

	stripeAccountID  = "acct_test_123"
	stripeCustomerID = "cus_test_123"
	testPMID         = "pm_test_123"

	validOrg = &models.Organization{
		ID:   testOrgID,
		Name: "Test Org",
	}
	orgWithStripe = &models.Organization{
		ID:              testOrgID,
		Name:            "Test Org",
		StripeAccountID: &stripeAccountID,
	}
	validManager = &models.Manager{
		ID:    uuid.New(),
		Email: "manager@testorg.com",
	}
	validLocation = &models.Location{
		ID:      uuid.New(),
		Country: "Thailand",
	}
	validGuardian = &models.Guardian{
		ID:    testGuardianID,
		Email: "guardian@example.com",
		Name:  "Test Guardian",
	}
	guardianWithStripe = &models.Guardian{
		ID:               testGuardianID,
		Email:            "guardian@example.com",
		Name:             "Test Guardian",
		StripeCustomerID: &stripeCustomerID,
	}
)

func newHandler(
	orgRepo *repomocks.MockOrganizationRepository,
	managerRepo *repomocks.MockManagerRepository,
	regRepo *repomocks.MockRegistrationRepository,
	locationRepo *repomocks.MockLocationRepository,
	guardianRepo *repomocks.MockGuardianRepository,
	sc *stripemocks.MockStripeClient,
) *Handler {
	return NewHandler(orgRepo, managerRepo, regRepo, locationRepo, guardianRepo, sc)
}

func TestHandler_CreateOrgStripeAccount(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateOrgStripeAccountInput
		mockSetup func(*repomocks.MockOrganizationRepository, *repomocks.MockManagerRepository, *repomocks.MockLocationRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully creates stripe account",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(validOrg, nil)
				managerRepo.On("GetManagerByOrgID", mock.Anything, testOrgID).Return(validManager, nil)
				locationRepo.On("GetLocationByOrganizationID", mock.Anything, testOrgID).Return(validLocation, nil)
				sc.On("CreateOrganizationAccount", mock.Anything, validOrg.Name, validManager.Email, "TH").
					Return(&models.CreateOrgStripeAccountClientOutput{}, nil)
				orgRepo.On("SetStripeAccountID", mock.Anything, testOrgID, mock.AnythingOfType("string")).Return(orgWithStripe, nil)
			},
			wantErr: false,
		},
		{
			name:  "fails when org already has stripe account",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(orgWithStripe, nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when org not found",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when manager not found",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(validOrg, nil)
				managerRepo.On("GetManagerByOrgID", mock.Anything, testOrgID).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when location not found",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(validOrg, nil)
				managerRepo.On("GetManagerByOrgID", mock.Anything, testOrgID).Return(validManager, nil)
				locationRepo.On("GetLocationByOrganizationID", mock.Anything, testOrgID).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when stripe account creation fails",
			input: &models.CreateOrgStripeAccountInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, managerRepo *repomocks.MockManagerRepository, locationRepo *repomocks.MockLocationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(validOrg, nil)
				managerRepo.On("GetManagerByOrgID", mock.Anything, testOrgID).Return(validManager, nil)
				locationRepo.On("GetLocationByOrganizationID", mock.Anything, testOrgID).Return(validLocation, nil)
				sc.On("CreateOrganizationAccount", mock.Anything, validOrg.Name, validManager.Email, "TH").
					Return(nil, errors.New("stripe error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockOrgRepo, mockManagerRepo, mockLocationRepo, mockStripeClient)

			handler := newHandler(mockOrgRepo, mockManagerRepo, mockRegRepo, mockLocationRepo, mockGuardianRepo, mockStripeClient)
			result, err := handler.CreateOrgStripeAccount(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, testOrgID, result.Body.Account.ID)
			}

			mockOrgRepo.AssertExpectations(t)
			mockManagerRepo.AssertExpectations(t)
			mockLocationRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateOrgLoginLink(t *testing.T) {
	loginURL := "https://connect.stripe.com/login/test"

	tests := []struct {
		name      string
		input     *models.CreateOrgLoginLinkInput
		mockSetup func(*repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully creates login link",
			input: &models.CreateOrgLoginLinkInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(orgWithStripe, nil)
				sc.On("CreateLoginLink", mock.Anything, stripeAccountID).Return(loginURL, nil)
			},
			wantErr: false,
		},
		{
			name:  "fails when org not found",
			input: &models.CreateOrgLoginLinkInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when org has no stripe account",
			input: &models.CreateOrgLoginLinkInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(validOrg, nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when stripe returns error",
			input: &models.CreateOrgLoginLinkInput{OrganizationID: testOrgID},
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				orgRepo.On("GetOrganizationByID", mock.Anything, testOrgID, mock.Anything).Return(orgWithStripe, nil)
				sc.On("CreateLoginLink", mock.Anything, stripeAccountID).Return("", errors.New("stripe error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockOrgRepo, mockStripeClient)

			handler := newHandler(mockOrgRepo, mockManagerRepo, mockRegRepo, mockLocationRepo, mockGuardianRepo, mockStripeClient)
			result, err := handler.CreateOrgLoginLink(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, loginURL, result.Body.LoginURL)
			}

			mockOrgRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateStripeCustomer(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateStripeCustomerInput
		mockSetup func(*repomocks.MockGuardianRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully creates stripe customer",
			input: &models.CreateStripeCustomerInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(validGuardian, nil)
				sc.On("CreateCustomer", mock.Anything, validGuardian.Email, validGuardian.Name).Return(&stripe.Customer{ID: stripeCustomerID}, nil)
				guardianRepo.On("SetStripeCustomerID", mock.Anything, testGuardianID, stripeCustomerID).Return(guardianWithStripe, nil)
			},
			wantErr: false,
		},
		{
			name:  "fails when guardian not found",
			input: &models.CreateStripeCustomerInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when guardian already has stripe customer",
			input: &models.CreateStripeCustomerInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(guardianWithStripe, nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when stripe customer creation fails",
			input: &models.CreateStripeCustomerInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(validGuardian, nil)
				sc.On("CreateCustomer", mock.Anything, validGuardian.Email, validGuardian.Name).Return(nil, errors.New("stripe error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockStripeClient)

			handler := newHandler(mockOrgRepo, mockManagerRepo, mockRegRepo, mockLocationRepo, mockGuardianRepo, mockStripeClient)
			result, err := handler.CreateStripeCustomer(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, testGuardianID, result.Body.ID)
				assert.Equal(t, &stripeCustomerID, result.Body.StripeCustomerID)
			}

			mockGuardianRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateSetupIntent(t *testing.T) {
	clientSecret := "seti_test_secret_123"

	tests := []struct {
		name      string
		input     *models.CreateSetupIntentInput
		mockSetup func(*repomocks.MockGuardianRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully creates setup intent",
			input: &models.CreateSetupIntentInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(guardianWithStripe, nil)
				sc.On("CreateSetupIntent", mock.Anything, stripeCustomerID).Return(clientSecret, nil)
			},
			wantErr: false,
		},
		{
			name:  "fails when guardian not found",
			input: &models.CreateSetupIntentInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name:  "fails when guardian has no stripe customer ID",
			input: &models.CreateSetupIntentInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(validGuardian, nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when stripe returns error on setup intent creation",
			input: &models.CreateSetupIntentInput{GuardianID: testGuardianID},
			mockSetup: func(guardianRepo *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				guardianRepo.On("GetGuardianByID", mock.Anything, testGuardianID).Return(guardianWithStripe, nil)
				sc.On("CreateSetupIntent", mock.Anything, stripeCustomerID).Return("", errors.New("stripe error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockStripeClient)

			handler := newHandler(mockOrgRepo, mockManagerRepo, mockRegRepo, mockLocationRepo, mockGuardianRepo, mockStripeClient)
			result, err := handler.CreateSetupIntent(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, clientSecret, result.Body.ClientSecret)
			}

			mockGuardianRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHandler_DetachGuardianPaymentMethod(t *testing.T) {
	testGuardianID := uuid.MustParse("10000000-0000-0000-0000-000000000001")
	testStripeCustomerID := "cus_test123"

	makeInput := func(pmID string) *models.DetachPaymentMethodInput {
		i := &models.DetachPaymentMethodInput{}
		i.Body.PaymentMethodID = pmID
		i.Body.GuardianID = testGuardianID
		return i
	}

	makeGuardian := func() *models.Guardian {
		return &models.Guardian{
			ID:               testGuardianID,
			StripeCustomerID: &testStripeCustomerID,
		}
	}

	makePMsOutput := func(pmIDs ...string) *models.GetPaymentMethodsByGuardianIDOutput {
		pms := make([]models.PaymentMethod, len(pmIDs))
		for i, id := range pmIDs {
			pms[i] = models.PaymentMethod{ID: id}
		}
		out := &models.GetPaymentMethodsByGuardianIDOutput{}
		out.Body.PaymentMethods = pms
		return out
	}

	tests := []struct {
		name      string
		input     *models.DetachPaymentMethodInput
		mockSetup func(*repomocks.MockGuardianRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully detaches payment method",
			input: makeInput(testPMID),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(makeGuardian(), nil)
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, testStripeCustomerID).Return(makePMsOutput(testPMID), nil)
				sc.On("DetachPaymentMethod", mock.Anything, testPMID).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "fails when payment method does not belong to guardian",
			input: makeInput("pm_other"),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(makeGuardian(), nil)
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, testStripeCustomerID).Return(makePMsOutput(testPMID), nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when stripe returns error on detach",
			input: makeInput(testPMID),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(makeGuardian(), nil)
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, testStripeCustomerID).Return(makePMsOutput(testPMID), nil)
				sc.On("DetachPaymentMethod", mock.Anything, testPMID).Return(errors.New("stripe error"))
			},
			wantErr: true,
		},
		{
			name:  "fails when fetching payment methods from stripe errors",
			input: makeInput(testPMID),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(makeGuardian(), nil)
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, testStripeCustomerID).Return(nil, errors.New("stripe error"))
			},
			wantErr: true,
		},
		{
			name:  "fails when guardian has no stripe customer account",
			input: makeInput(testPMID),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(&models.Guardian{
					ID:               testGuardianID,
					StripeCustomerID: nil,
				}, nil)
			},
			wantErr: true,
		},
		{
			name:  "fails when guardian repo errors",
			input: makeInput(testPMID),
			mockSetup: func(gr *repomocks.MockGuardianRepository, sc *stripemocks.MockStripeClient) {
				gr.On("GetGuardianByID", mock.Anything, testGuardianID).Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockManagerRepo := new(repomocks.MockManagerRepository)
			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockLocationRepo := new(repomocks.MockLocationRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockStripeClient)

			handler := newHandler(mockOrgRepo, mockManagerRepo, mockRegRepo, mockLocationRepo, mockGuardianRepo, mockStripeClient)
			result, err := handler.DetachGuardianPaymentMethod(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			mockGuardianRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}
