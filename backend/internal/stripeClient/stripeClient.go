package stripeClient

import (
	"os"

	"github.com/stripe/stripe-go/v84"
)

type StripeClient struct {
    client *stripe.Client
}

func NewStripeClient(apiKey string) (*StripeClient, error) {
    if (apiKey == "") {
        apiKey = os.Getenv("STRIPE_SECRET_KEY")
    }
    return &StripeClient{
        client: stripe.NewClient(apiKey),
    }, nil
}