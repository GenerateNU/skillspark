package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {

	if _, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID); err != nil {
		return nil, errs.BadRequest("Invalid event_occurrence_id: event occurrence does not exist")
	}

	if _, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID); err != nil {
		return nil, errs.BadRequest("Invalid child_id: child does not exist")
	}

	if _, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID); err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	eventOccurrence, _ := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID)
	
	guardian, _ := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID)
	
	org, _ := h.OrganizationRepository.GetOrganizationByID(ctx, eventOccurrence.Event.OrganizationID)
	

	piInput := models.CreatePaymentIntentInput{}
	piInput.Body.Amount = int64(eventOccurrence.Price)
	piInput.Body.Currency = input.Body.Currency
	piInput.Body.GuardianStripeID = *guardian.StripeCustomerID
	piInput.Body.OrgStripeID = *org.StripeAccountID
	piInput.Body.PaymentMethodID = &input.Body.PaymentMethodID
	piInput.Body.EventDate = eventOccurrence.StartTime


	paymentIntent, err := h.StripeClient.CreatePaymentIntent(ctx, &piInput)
	if err != nil {
		return nil, errors.New("failed to create payment intent")
	}

	completeRegistration := &models.CreateRegistrationWithPaymentData{
		ChildID:               input.Body.ChildID,
		GuardianID:            input.Body.GuardianID,
		EventOccurrenceID:     input.Body.EventOccurrenceID,
		Status:                input.Body.Status,
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

	registration, err := h.RegistrationRepository.CreateRegistration(ctx, completeRegistration)
	if err != nil {
		return nil, err
	}

	return registration, nil
}
