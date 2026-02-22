package jobs

import (
	"context"
	"log"
	"skillspark/internal/models"
	"time"
)

func (j *JobScheduler) CapturePaymentsJob() {
	ctx := context.Background()
	
	now := time.Now()
	startWindow := now.Add(24 * time.Hour)
	endWindow := now.Add(25 * time.Hour)

	registrations, err := j.repo.Registration.GetRegistrationsForCapture(ctx, startWindow, endWindow)

	if err != nil {
		return
	}

	log.Printf("Found %d registrations to capture", len(registrations))

	for _, registration := range registrations {

		stripeInput := &models.CapturePaymentIntentInput{
			PaymentIntentID: registration.StripePaymentIntentID,
			StripeAccountID: registration.OrgStripeAccountID,
		}

		stripeOutput, err := j.stripeClient.CapturePaymentIntent(ctx, stripeInput)

		if err != nil {
			log.Printf("Failed to capture payment for registration %s: %v", registration.ID, err)
			continue
		}

		updateInput := &models.UpdateRegistrationPaymentStatusInput{
			ID: registration.ID,
		}

		updateInput.Body.PaymentIntentStatus = stripeOutput.Body.Status

		_, err = j.repo.Registration.UpdateRegistrationPaymentStatus(ctx, updateInput)
		if err != nil {
			log.Printf("Failed to update payment status for registration %s: %v", registration.ID, err)
		}
	}
}