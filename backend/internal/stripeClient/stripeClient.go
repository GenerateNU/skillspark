package stripeClient

import (
	"errors"
	"log"
	"os"

	"github.com/stripe/stripe-go/v84"
)

type StripeClient struct {
    client *stripe.Client
}

func NewStripeClient(apiKey string) (*StripeClient, error) {
    if (apiKey == "") {
        apiKey = os.Getenv("STRIPE_SECRET_KEY")
	    if apiKey == "" {
		    log.Fatal("STRIPE_SECRET_KEY environment variable is required")
            return nil, errors.New("STRIPE_SECRET_KEY environment variable is required")
	    }
    }
    return &StripeClient{
        client: stripe.NewClient(apiKey),
    }, nil
}