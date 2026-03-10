package registration

import (
	"context"
	"errors"
	"skillspark/internal/models"
	"time"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
	_, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID, "en-US")
	if err != nil {
		return nil, err
	}

	if eventOccurrence.StartTime.Before(time.Now()) {
		return nil, errors.New("event occurrence has already started")
	}

	if eventOccurrence.CurrEnrolled >= eventOccurrence.MaxAttendees {
		return nil, errors.New("event occurrence has reached max registration")
	}

	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil {
		return nil, errors.New("guardian must have a Stripe Customer ID")
	}

	child, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID)
	if err != nil {
		return nil, err
	}

	if child.GuardianID != input.Body.GuardianID {
		return nil, errors.New("child does not belong to the specified guardian")
	}

	org, err := h.OrganizationRepository.GetOrganizationByID(ctx, eventOccurrence.Event.OrganizationID)
	if err != nil {
		return nil, err
	}

	piInput := models.CreatePaymentIntentInput{}
	piInput.Body.Amount = int64(eventOccurrence.Price)
	piInput.Body.Currency = eventOccurrence.Currency
	piInput.Body.GuardianStripeID = *guardian.StripeCustomerID
	piInput.Body.OrgStripeID = *org.StripeAccountID
	piInput.Body.PaymentMethodID = *input.Body.PaymentMethodID
	piInput.Body.EventDate = eventOccurrence.StartTime

	paymentIntent, err := h.StripeClient.CreatePaymentIntent(ctx, &piInput)
	if err != nil {
		return nil, err
	}

	completeRegistration := &models.CreateRegistrationWithPaymentData{
		ChildID:               input.Body.ChildID,
		GuardianID:            input.Body.GuardianID,
		EventOccurrenceID:     input.Body.EventOccurrenceID,
		Status:                input.Body.Status,
		StripePaymentIntentID: paymentIntent.Body.PaymentIntentID,
		StripeCustomerID:      *guardian.StripeCustomerID,
		OrgStripeAccountID:    *org.StripeAccountID,
		StripePaymentMethodID: *input.Body.PaymentMethodID,
		TotalAmount:           paymentIntent.Body.TotalAmount,
		ProviderAmount:        paymentIntent.Body.ProviderAmount,
		PlatformFeeAmount:     paymentIntent.Body.PlatformFeeAmount,
		Currency:              paymentIntent.Body.Currency,
		PaymentIntentStatus:   paymentIntent.Body.Status,
	}

	registration, err := h.RegistrationRepository.CreateRegistration(ctx, completeRegistration)
	if err != nil {
		return nil, err
	}

	return registration, nil
}