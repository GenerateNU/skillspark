package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) DetachPaymentMethod(
	ctx context.Context,
	paymentMethodID string,
) error {
	params := &stripe.PaymentMethodDetachParams{}
	params.Context = ctx
	
	_, err := sc.client.V1PaymentMethods.Detach(ctx, paymentMethodID, params)
	return err
}