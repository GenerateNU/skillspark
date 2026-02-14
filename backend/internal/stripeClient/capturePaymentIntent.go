package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CapturePaymentIntent(ctx context.Context, input *models.CapturePaymentIntentInput) (*models.CapturePaymentIntentOutput, error) {
	params := &stripe.PaymentIntentCaptureParams{}
	params.SetStripeAccount(input.StripeAccountID)

	pi, err := sc.client.V1PaymentIntents.Capture(ctx, input.PaymentIntentID, params)
	if err != nil {
		return nil, err
	}

	output := &models.CapturePaymentIntentOutput{}
	output.Body.PaymentIntentID = pi.ID
	output.Body.Status = string(pi.Status)
	output.Body.Amount = pi.Amount
	output.Body.Currency = string(pi.Currency)

	return output, nil
}