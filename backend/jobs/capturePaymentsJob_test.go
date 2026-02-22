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

func TestCapturePaymentsJob_Success(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	mockRepo := &storage.Repository{
		Registration: mockRegRepo,
	}

	scheduler := &JobScheduler{
		repo:         mockRepo,
		stripeClient: mockStripeClient,
	}

	reg1 := models.Registration{
		ID:                    uuid.New(),
		StripePaymentIntentID: "pi_test_123",
		OrgStripeAccountID:    "acct_test_123",
		PaymentIntentStatus:   "requires_capture",
		Status:                models.RegistrationStatusRegistered,
		OccurrenceStartTime:   time.Now().Add(24 * time.Hour),
	}

	reg2 := models.Registration{
		ID:                    uuid.New(),
		StripePaymentIntentID: "pi_test_456",
		OrgStripeAccountID:    "acct_test_456",
		PaymentIntentStatus:   "requires_capture",
		Status:                models.RegistrationStatusRegistered,
		OccurrenceStartTime:   time.Now().Add(24 * time.Hour),
	}

	mockRegRepo.On("GetRegistrationsForCapture", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return([]models.Registration{reg1, reg2}, nil)

	mockStripeClient.On("CapturePaymentIntent", mock.Anything, mock.MatchedBy(func(input *models.CapturePaymentIntentInput) bool {
		return input.PaymentIntentID == "pi_test_123"
	})).Return(&models.CapturePaymentIntentOutput{
		Body: struct {
			PaymentIntentID string `json:"payment_intent_id" doc:"Captured payment intent ID"`
			Status          string `json:"status" doc:"Payment intent status (should be 'succeeded')"`
			Amount          int64  `json:"amount" doc:"Amount captured in cents"`
			Currency        string `json:"currency" doc:"Currency code"`
		}{
			PaymentIntentID: "pi_test_123",
			Status:          "succeeded",
		},
	}, nil)

	mockStripeClient.On("CapturePaymentIntent", mock.Anything, mock.MatchedBy(func(input *models.CapturePaymentIntentInput) bool {
		return input.PaymentIntentID == "pi_test_456"
	})).Return(&models.CapturePaymentIntentOutput{
		Body: struct {
			PaymentIntentID string `json:"payment_intent_id" doc:"Captured payment intent ID"`
			Status          string `json:"status" doc:"Payment intent status (should be 'succeeded')"`
			Amount          int64  `json:"amount" doc:"Amount captured in cents"`
			Currency        string `json:"currency" doc:"Currency code"`
		}{
			PaymentIntentID: "pi_test_456",
			Status:          "succeeded",
		},
	}, nil)

	mockRegRepo.On("UpdateRegistrationPaymentStatus", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationPaymentStatusInput")).
		Return(&models.UpdateRegistrationPaymentStatusOutput{}, nil).Times(2)

	scheduler.CapturePaymentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertExpectations(t)
	mockRegRepo.AssertNumberOfCalls(t, "UpdateRegistrationPaymentStatus", 2)
}

func TestCapturePaymentsJob_NoRegistrations(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	mockRepo := &storage.Repository{
		Registration: mockRegRepo,
	}

	scheduler := &JobScheduler{
		repo:         mockRepo,
		stripeClient: mockStripeClient,
	}

	mockRegRepo.On("GetRegistrationsForCapture", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return([]models.Registration{}, nil)

	scheduler.CapturePaymentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertNotCalled(t, "CapturePaymentIntent")
	mockRegRepo.AssertNotCalled(t, "UpdateRegistrationPaymentStatus")
}

func TestCapturePaymentsJob_StripeCaptureFailure(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	mockRepo := &storage.Repository{
		Registration: mockRegRepo,
	}

	scheduler := &JobScheduler{
		repo:         mockRepo,
		stripeClient: mockStripeClient,
	}

	reg := models.Registration{
		ID:                    uuid.New(),
		StripePaymentIntentID: "pi_test_fail",
		OrgStripeAccountID:    "acct_test_123",
		PaymentIntentStatus:   "requires_capture",
		Status:                models.RegistrationStatusRegistered,
	}

	mockRegRepo.On("GetRegistrationsForCapture", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return([]models.Registration{reg}, nil)

	mockStripeClient.On("CapturePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CapturePaymentIntentInput")).
		Return(nil, assert.AnError)

	scheduler.CapturePaymentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertExpectations(t)
	mockRegRepo.AssertNotCalled(t, "UpdateRegistrationPaymentStatus")
}

func TestCapturePaymentsJob_DatabaseUpdateFailure(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	mockRepo := &storage.Repository{
		Registration: mockRegRepo,
	}

	scheduler := &JobScheduler{
		repo:         mockRepo,
		stripeClient: mockStripeClient,
	}

	reg := models.Registration{
		ID:                    uuid.New(),
		StripePaymentIntentID: "pi_test_123",
		OrgStripeAccountID:    "acct_test_123",
		PaymentIntentStatus:   "requires_capture",
		Status:                models.RegistrationStatusRegistered,
	}

	mockRegRepo.On("GetRegistrationsForCapture", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return([]models.Registration{reg}, nil)

	mockStripeClient.On("CapturePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CapturePaymentIntentInput")).
		Return(&models.CapturePaymentIntentOutput{
			Body: struct {
				PaymentIntentID string `json:"payment_intent_id" doc:"Captured payment intent ID"`
				Status          string `json:"status" doc:"Payment intent status (should be 'succeeded')"`
				Amount          int64  `json:"amount" doc:"Amount captured in cents"`
				Currency        string `json:"currency" doc:"Currency code"`
			}{
				PaymentIntentID: "pi_test_123",
				Status:          "succeeded",
			},
		}, nil)

	mockRegRepo.On("UpdateRegistrationPaymentStatus", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationPaymentStatusInput")).
		Return(nil, assert.AnError)

	scheduler.CapturePaymentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertExpectations(t)
}

func TestCapturePaymentsJob_FetchError(t *testing.T) {
	mockRegRepo := new(repomocks.MockRegistrationRepository)
	mockStripeClient := new(stripemocks.MockStripeClient)
	mockRepo := &storage.Repository{
		Registration: mockRegRepo,
	}

	scheduler := &JobScheduler{
		repo:         mockRepo,
		stripeClient: mockStripeClient,
	}

	mockRegRepo.On("GetRegistrationsForCapture", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
		Return(nil, assert.AnError)

	scheduler.CapturePaymentsJob()

	mockRegRepo.AssertExpectations(t)
	mockStripeClient.AssertNotCalled(t, "CapturePaymentIntent")
}