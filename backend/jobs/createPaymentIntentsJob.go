package jobs

import (
	"context"
	"log"
	"skillspark/internal/models"
)

func (j *JobScheduler) CreatePaymentIntentsJob() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("CreatePaymentIntentsJob panicked: %v", r)
		}
	}()

	ctx := context.Background()

	registrations, err := j.repo.Registration.GetRegistrationsForPaymentCreation(ctx)
	if err != nil {
		log.Printf("CreatePaymentIntentsJob: failed to get registrations: %v", err)
		return
	}

	for _, reg := range registrations {
		guardian, err := j.repo.Guardian.GetGuardianByID(ctx, reg.GuardianID)
		if err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to get guardian %s for registration %s: %v", reg.GuardianID, reg.ID, err)
			continue
		}
		if guardian.StripeCustomerID == nil {
			log.Printf("CreatePaymentIntentsJob: guardian %s has no Stripe customer ID, skipping registration %s", reg.GuardianID, reg.ID)
			continue
		}

		paymentMethods, err := j.stripeClient.GetPaymentMethodsByCustomerID(ctx, *guardian.StripeCustomerID)
		if err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to get payment methods for guardian %s: %v", reg.GuardianID, err)
			continue
		}
		if len(paymentMethods.Body.PaymentMethods) == 0 {
			log.Printf("CreatePaymentIntentsJob: guardian %s has no payment methods, skipping registration %s", reg.GuardianID, reg.ID)
			continue
		}
		paymentMethodID := paymentMethods.Body.PaymentMethods[0].ID

		eventOccurrence, err := j.repo.EventOccurrence.GetEventOccurrenceByID(ctx, reg.EventOccurrenceID, "en-US")
		if err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to get event occurrence %s for registration %s: %v", reg.EventOccurrenceID, reg.ID, err)
			continue
		}

		org, err := j.repo.Organization.GetOrganizationByID(ctx, eventOccurrence.Event.OrganizationID)
		if err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to get organization for registration %s: %v", reg.ID, err)
			continue
		}
		if org.StripeAccountID == nil {
			log.Printf("CreatePaymentIntentsJob: organization %s has no Stripe account ID, skipping registration %s", eventOccurrence.Event.OrganizationID, reg.ID)
			continue
		}

		piInput := models.CreatePaymentIntentInput{}
		piInput.Body.Amount = int64(eventOccurrence.Price)
		piInput.Body.Currency = eventOccurrence.Currency
		piInput.Body.GuardianStripeID = *guardian.StripeCustomerID
		piInput.Body.OrgStripeID = *org.StripeAccountID
		piInput.Body.PaymentMethodID = paymentMethodID
		piInput.Body.EventDate = eventOccurrence.StartTime
		piInput.Body.PlatformFeePercentage = 10

		paymentIntent, err := j.stripeClient.CreatePaymentIntent(ctx, &piInput)
		if err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to create payment intent for registration %s: %v", reg.ID, err)
			continue
		}

		paymentData := &models.CreatePaymentData{
			RegistrationID:        reg.ID,
			StripePaymentIntentID: paymentIntent.Body.PaymentIntentID,
			StripeCustomerID:      *guardian.StripeCustomerID,
			OrgStripeAccountID:    *org.StripeAccountID,
			StripePaymentMethodID: paymentMethodID,
			TotalAmount:           paymentIntent.Body.TotalAmount,
			ProviderAmount:        paymentIntent.Body.ProviderAmount,
			PlatformFeeAmount:     paymentIntent.Body.PlatformFeeAmount,
			Currency:              paymentIntent.Body.Currency,
			PaymentIntentStatus:   paymentIntent.Body.Status,
		}

		if err := j.repo.Registration.CreatePayment(ctx, paymentData); err != nil {
			log.Printf("CreatePaymentIntentsJob: failed to store payment for registration %s: %v", reg.ID, err)
			continue
		}

		log.Printf("CreatePaymentIntentsJob: created payment intent %s for registration %s", paymentIntent.Body.PaymentIntentID, reg.ID)
	}
}
