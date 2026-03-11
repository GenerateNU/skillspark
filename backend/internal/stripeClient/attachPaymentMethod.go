package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

// This function is for testing only
func (sc *StripeClient) AttachPaymentMethod(ctx context.Context, paymentMethodID string, customerID string) error {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	}

	_, err := sc.client.V1PaymentMethods.Attach(ctx, paymentMethodID, params)
	return err
}
