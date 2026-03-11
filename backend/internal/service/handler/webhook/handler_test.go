package webhook

import (
	"context"
	"encoding/json"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
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

func newHandler(regRepo *repomocks.MockRegistrationRepository, orgRepo *repomocks.MockOrganizationRepository) *Handler {
	return &Handler{
		repo: &storage.Repository{
			Registration: regRepo,
			Organization: orgRepo,
		},
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
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID).
					Return(&models.Registration{
						ID: regID,
					}, nil)

				regRepo.On("CancelRegistration", mock.Anything, mock.AnythingOfType("*models.CancelRegistrationInput")).
					Return(&models.CancelRegistrationOutput{}, nil)
			},
			wantErr: false,
		},
		{
			name:  "registration not found — returns error",
			event: makePaymentIntentEvent(piID, stripe.PaymentIntentStatusRequiresPaymentMethod),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID).
					Return(nil, &errs.HTTPError{Code: 404, Message: "registration not found"})
			},
			wantErr: true,
		},
		{
			name:  "cancel registration fails — returns error",
			event: makePaymentIntentEvent(piID, stripe.PaymentIntentStatusRequiresPaymentMethod),
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository) {
				regRepo.On("GetRegistrationByPaymentIntentID", mock.Anything, piID).
					Return(&models.Registration{
						ID: regID,
					}, nil)

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
			tt.mockSetup(mockRegRepo)

			handler := newHandler(mockRegRepo, mockOrgRepo)

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
			tt.mockSetup(mockOrgRepo)

			handler := newHandler(mockRegRepo, mockOrgRepo)

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
