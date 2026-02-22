package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreateCustomerPaymentMethod(
	ctx context.Context,
	customerID string,
) (*stripe.PaymentMethod, error) {
	params := &stripe.PaymentMethodCreateParams{
		Customer: stripe.String(customerID),
	}
	params.Context = ctx

	pm, err := sc.client.V1PaymentMethods.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	return pm, nil
}