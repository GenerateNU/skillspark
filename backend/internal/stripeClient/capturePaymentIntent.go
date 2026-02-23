package stripeClient

import (
	"context"
	"errors"
	"log"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CapturePaymentIntent(ctx context.Context, input *models.CapturePaymentIntentInput) (*models.CapturePaymentIntentOutput, error) {
	params := &stripe.PaymentIntentCaptureParams{}

	log.Printf("Capturing payment intent %s", input.PaymentIntentID)

	pi, err := sc.client.V1PaymentIntents.Capture(ctx, input.PaymentIntentID, params)
	if err != nil {
		log.Printf("Stripe capture error for %s: %v", input.PaymentIntentID, err)
		return nil, err
	}
	if pi == nil {
		log.Printf("Stripe returned nil for payment intent %s", input.PaymentIntentID)
		return nil, errors.New("stripe returned nil payment intent")
	}

	log.Printf("Captured payment intent %s, status: %s", pi.ID, pi.Status)

	output := &models.CapturePaymentIntentOutput{}
	output.Body.PaymentIntentID = pi.ID
	output.Body.Status = string(pi.Status)
	output.Body.Amount = pi.Amount
	output.Body.Currency = string(pi.Currency)

	return output, nil
}