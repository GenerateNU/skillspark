package webhook

import (
	"context"
	"encoding/json"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v84"
)

func makePaymentIntentEvent(piID string, status stripe.PaymentIntentStatus) stripe.Event {
	pi := stripe.PaymentIntent{
		ID:     piID,
		Status: status,
	}
	raw, _ := json.Marshal(pi)
	return stripe.Event{
		Type: "payment_intent.payment_failed",
		Data: &stripe.EventData{Raw: raw},
	}
}

func makeAccountEvent(accountID string, chargesEnabled, payoutsEnabled bool) stripe.Event {
	account := stripe.Account{
		ID:             accountID,
		ChargesEnabled: chargesEnabled,
		PayoutsEnabled: payoutsEnabled,
	}
	raw, _ := json.Marshal(account)
	return stripe.Event{
		Type: "account.updated",
		Data: &stripe.EventData{Raw: raw},
	}
}

func makePaymentMethodEvent(pmID, customerID string) stripe.Event {
	pm := stripe.PaymentMethod{
		ID: pmID,
		Customer: &stripe.Customer{
			ID: customerID,
		},
	}
	raw, _ := json.Marshal(pm)
	return stripe.Event{
		Type: "payment_method.attached",
		Data: &stripe.EventData{Raw: raw},
	}
}

func makePaymentMethodEventNoCustomer(pmID string) stripe.Event {
	pm := stripe.PaymentMethod{
		ID:       pmID,
		Customer: nil,
	}
	raw, _ := json.Marshal(pm)
	return stripe.Event{
		Type: "payment_method.attached",
		Data: &stripe.EventData{Raw: raw},
	}
}

func newHandler(
	regRepo *repomocks.MockRegistrationRepository,
	orgRepo *repomocks.MockOrganizationRepository,
	sc *stripemocks.MockStripeClient,
) *Handler {
	return &Handler{
		repo: &storage.Repository{
			Registration: regRepo,
			Organization: orgRepo,
		},
		stripeClient: sc,
	}
}

func TestHandler_HandlePaymentIntentFailed(t *testing.T) {
	piID := "pi_test_123"
	regID := uuid.MustParse("10000000-0000-0000-0000-000000000001")

	tests := []struct {
		name      string
		event     stripe.Event
		mockSetup func(*repomocks.MockRegistrationRepository)
		wantErr   bool
	}{
		{
			name:  "successful — cancels registration",
			event: makePaymentIntentEvent(piID, stripe.PaymentIntentStatusRequiresPaymentMethod),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID, "en-US").
					Return(&models.Registration{ID: regID}, nil)
				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(&models.CancelRegistrationOutput{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "registration not found — returns error",
			event: makePaymentIntentEvent(piID, stripe.PaymentIntentStatusRequiresPaymentMethod),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID, "en-US").
					Return(nil, &errs.HTTPError{Code: 404, Message: "registration not found"})
			},
			wantErr: true,
		},
		{
			name:  "cancel registration fails — returns error",
			event: makePaymentIntentEvent(piID, stripe.PaymentIntentStatusRequiresPaymentMethod),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID, "en-US").
					Return(&models.Registration{ID: regID}, nil)
				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(nil, &errs.HTTPError{Code: 500, Message: "db error"})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			handler := newHandler(mockRegRepo, mockOrgRepo, mockStripeClient)
			err := handler.handlePaymentIntentFailed(context.Background(), tt.event)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_HandleAccountUpdated(t *testing.T) {
	accountID := "acct_test_123"

	tests := []struct {
		name      string
		event     stripe.Event
		mockSetup func(*repomocks.MockOrganizationRepository)
		wantErr   bool
	}{
		{
			name:  "fully activated account — sets status to true",
			event: makeAccountEvent(accountID, true, true),
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository) {
				orgRepo.On("SetStripeAccountStatus", mock.Anything, accountID, true).
					Return(&models.Organization{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "charges disabled — sets status to false",
			event: makeAccountEvent(accountID, false, true),
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository) {
				orgRepo.On("SetStripeAccountStatus", mock.Anything, accountID, false).
					Return(&models.Organization{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "payouts disabled — sets status to false",
			event: makeAccountEvent(accountID, true, false),
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository) {
				orgRepo.On("SetStripeAccountStatus", mock.Anything, accountID, false).
					Return(&models.Organization{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "both disabled — sets status to false",
			event: makeAccountEvent(accountID, false, false),
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository) {
				orgRepo.On("SetStripeAccountStatus", mock.Anything, accountID, false).
					Return(&models.Organization{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "repo fails — returns error",
			event: makeAccountEvent(accountID, true, true),
			mockSetup: func(orgRepo *repomocks.MockOrganizationRepository) {
				orgRepo.On("SetStripeAccountStatus", mock.Anything, accountID, true).
					Return(nil, &errs.HTTPError{Code: 500, Message: "db error"})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockOrgRepo)

			handler := newHandler(mockRegRepo, mockOrgRepo, mockStripeClient)
			err := handler.handleAccountUpdated(context.Background(), tt.event)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockOrgRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_HandlePaymentMethodAttached(t *testing.T) {
	pmID := "pm_test_123"
	oldPMID := "pm_old_456"
	customerID := "cus_test_123"

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
		event     stripe.Event
		mockSetup func(*stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name:  "successfully detaches old payment methods, keeps new one",
			event: makePaymentMethodEvent(pmID, customerID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
					Return(makePMsOutput(pmID, oldPMID), nil)
				sc.On("DetachPaymentMethod", mock.Anything, oldPMID).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "no old payment methods — nothing to detach",
			event: makePaymentMethodEvent(pmID, customerID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
					Return(makePMsOutput(pmID), nil)
			},
			wantErr: false,
		},
		{
			name:      "no customer on payment method — returns nil early",
			event:     makePaymentMethodEventNoCustomer(pmID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {},
			wantErr:   false,
		},
		{
			name:  "fetching existing payment methods fails — returns error",
			event: makePaymentMethodEvent(pmID, customerID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
					Return(nil, &errs.HTTPError{Code: 500, Message: "stripe error"})
			},
			wantErr: true,
		},
		{
			name:  "detaching old payment method fails — returns error",
			event: makePaymentMethodEvent(pmID, customerID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
					Return(makePMsOutput(pmID, oldPMID), nil)
				sc.On("DetachPaymentMethod", mock.Anything, oldPMID).
					Return(&errs.HTTPError{Code: 500, Message: "stripe error"})
			},
			wantErr: true,
		},
		{
			name:  "multiple old payment methods — all detached except new one",
			event: makePaymentMethodEvent(pmID, customerID),
			mockSetup: func(sc *stripemocks.MockStripeClient) {
				sc.On("GetPaymentMethodsByCustomerID", mock.Anything, customerID).
					Return(makePMsOutput(pmID, oldPMID, "pm_old_789"), nil)
				sc.On("DetachPaymentMethod", mock.Anything, oldPMID).Return(nil)
				sc.On("DetachPaymentMethod", mock.Anything, "pm_old_789").Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockStripeClient)

			handler := newHandler(mockRegRepo, mockOrgRepo, mockStripeClient)
			err := handler.handlePaymentMethodAdditionSuccess(context.Background(), tt.event)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockStripeClient.AssertExpectations(t)
		})
	}
}
