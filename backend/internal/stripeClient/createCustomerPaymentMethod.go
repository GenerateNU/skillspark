package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreateSetupIntent(
	ctx context.Context,
	stripeCustomerID string,
) (string, error) {
	params := &stripe.SetupIntentCreateParams{
		Customer: stripe.String(stripeCustomerID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	
	setupIntent, err := sc.client.V1SetupIntents.Create(ctx, params)
	if err != nil {
		return "", err
	}
	
	return setupIntent.ClientSecret, nil
}