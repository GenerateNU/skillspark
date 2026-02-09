package stripeClient

import (
    "github.com/stripe/stripe-go/v84"
)

type StripeClient struct {
    client *stripe.Client
}

func NewStripeClient(apiKey string) *StripeClient {
    return &StripeClient{
        client: stripe.NewClient(apiKey),
    }
}