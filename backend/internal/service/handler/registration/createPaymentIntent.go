package registration

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreatePaymentIntent(ctx context.Context, input *models.CreatePaymentForRegistrationInput) (*models.CreatePaymentForRegistrationOutput, error) {
	regInput := &models.GetRegistrationByIDInput{
		AcceptLanguage: input.AcceptLanguage,
		ID:             input.RegistrationID,
	}
	reg, err := h.RegistrationRepository.GetRegistrationByID(ctx, regInput, nil)
	if err != nil {
		return nil, err
	}

	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, reg.Body.GuardianID)
	if err != nil {
		return nil, err
	}
	if guardian == nil {
		return nil, errors.New("guardian not found")
	}
	if guardian.StripeCustomerID == nil {
		return nil, errors.New("guardian must have a Stripe Customer ID before creating a payment")
	}

	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, reg.Body.EventOccurrenceID, "en-US")
	if err != nil {
		return nil, err
	}
	if eventOccurrence == nil {
		return nil, errors.New("event occurrence not found")
	}

	org, err := h.OrganizationRepository.GetOrganizationByID(ctx, eventOccurrence.Event.OrganizationID)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, errors.New("organization not found")
	}
	if org.StripeAccountID == nil {
		return nil, errors.New("organization must have a Stripe account ID before creating a payment")
	}

	piInput := models.CreatePaymentIntentInput{}
	piInput.Body.Amount = int64(eventOccurrence.Price)
	piInput.Body.Currency = eventOccurrence.Currency
	piInput.Body.GuardianStripeID = *guardian.StripeCustomerID
	piInput.Body.OrgStripeID = *org.StripeAccountID
	piInput.Body.PaymentMethodID = input.Body.PaymentMethodID
	piInput.Body.EventDate = eventOccurrence.StartTime
	piInput.Body.PlatformFeePercentage = 10

	paymentIntent, err := h.StripeClient.CreatePaymentIntent(ctx, &piInput)
	if err != nil {
		return nil, err
	}
	if paymentIntent == nil {
		return nil, errors.New("nil response from Stripe when creating payment intent")
	}

	paymentData := &models.CreatePaymentData{
		RegistrationID:        input.RegistrationID,
		StripePaymentIntentID: paymentIntent.Body.PaymentIntentID,
		StripeCustomerID:      *guardian.StripeCustomerID,
		OrgStripeAccountID:    *org.StripeAccountID,
		StripePaymentMethodID: input.Body.PaymentMethodID,
		TotalAmount:           paymentIntent.Body.TotalAmount,
		ProviderAmount:        paymentIntent.Body.ProviderAmount,
		PlatformFeeAmount:     paymentIntent.Body.PlatformFeeAmount,
		Currency:              paymentIntent.Body.Currency,
		PaymentIntentStatus:   paymentIntent.Body.Status,
	}

	if err := h.RegistrationRepository.CreatePayment(ctx, paymentData); err != nil {
		return nil, err
	}

	output := &models.CreatePaymentForRegistrationOutput{}
	output.Body.PaymentIntentID = paymentIntent.Body.PaymentIntentID
	output.Body.ClientSecret = paymentIntent.Body.ClientSecret
	output.Body.Status = paymentIntent.Body.Status
	output.Body.TotalAmount = paymentIntent.Body.TotalAmount
	output.Body.ProviderAmount = paymentIntent.Body.ProviderAmount
	output.Body.PlatformFeeAmount = paymentIntent.Body.PlatformFeeAmount
	output.Body.Currency = paymentIntent.Body.Currency

	return output, nil
}
