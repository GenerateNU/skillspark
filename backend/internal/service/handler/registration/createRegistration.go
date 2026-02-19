package registration

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {

	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID)
	if err != nil {
		return nil, err
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

	paymentMethod := input.Body.PaymentMethodID
	if paymentMethod == nil {
		paymentMethod = guardian.StripeCustomerID
	}

	piInput := models.CreatePaymentIntentInput{}
	piInput.Body.Amount = int64(eventOccurrence.Price)
	piInput.Body.Currency = input.Body.Currency
	piInput.Body.GuardianStripeID = *guardian.StripeCustomerID
	piInput.Body.OrgStripeID = *org.StripeAccountID
	piInput.Body.PaymentMethodID = paymentMethod
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