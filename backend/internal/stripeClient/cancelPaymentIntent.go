package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CancelPaymentIntent(ctx context.Context, input *models.CancelPaymentIntentInput) (*models.CancelPaymentIntentOutput, error) {
	params := &stripe.PaymentIntentCancelParams{}
	params.SetStripeAccount(input.StripeAccountID)

	pi, err := sc.client.V1PaymentIntents.Cancel(ctx, input.PaymentIntentID, params)
	if err != nil {
		return nil, err
	}

	output := &models.CancelPaymentIntentOutput{}
	output.Body.PaymentIntentID = pi.ID
	output.Body.Status = string(pi.Status)
	output.Body.Amount = pi.Amount
	output.Body.Currency = string(pi.Currency)

	return output, nil
}