package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) RefundPayment(ctx context.Context, input *models.RefundPaymentInput) (*models.RefundPaymentOutput, error) {
	params := &stripe.RefundCreateParams{
		PaymentIntent: stripe.String(input.PaymentIntentID),
	}

	result, err := sc.client.V1Refunds.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	output := &models.RefundPaymentOutput{}
	output.Body.RefundID = result.ID
	output.Body.Status = string(result.Status)
	output.Body.Amount = result.Amount
	output.Body.Currency = string(result.Currency)

	return output, nil
}
