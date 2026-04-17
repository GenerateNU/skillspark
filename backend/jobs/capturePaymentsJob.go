package jobs

import (
	"context"
	"log"
	"skillspark/internal/models"
	"time"
)

func (j *JobScheduler) CapturePaymentsJob() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("CapturePaymentsJob panicked: %v", r)
		}
	}()

	ctx := context.Background()

	now := time.Now()
	startWindow := now.Add(-24 * time.Hour)
	endWindow := now

	registrations, err := j.repo.Registration.GetRegistrationsForCapture(ctx, startWindow, endWindow)
	if err != nil {
		log.Printf("Failed to get registrations for capture: %v", err)
		return
	}

	for _, registration := range registrations {
		stripeInput := &models.CapturePaymentIntentInput{
			PaymentIntentID: registration.StripePaymentIntentID,
		}

		stripeOutput, err := j.stripeClient.CapturePaymentIntent(ctx, stripeInput)
		if err != nil {
			log.Printf("Failed to capture payment for registration %s, cancelling registration: %v", registration.ID, err)
			_, cancelErr := j.repo.Registration.CancelRegistration(ctx, &models.CancelRegistrationInput{ID: registration.ID})
			if cancelErr != nil {
				log.Printf("Failed to cancel registration %s after capture failure: %v", registration.ID, cancelErr)
			}
			continue
		}
		if stripeOutput == nil {
			log.Printf("Nil output for registration %s, skipping", registration.ID)
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
