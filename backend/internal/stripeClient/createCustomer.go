package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreateCustomer(
    ctx context.Context,
    email string,
    name string,
) (*stripe.Customer, error) {
    params := &stripe.CustomerCreateParams{
        Email: stripe.String(email),
        Name:  stripe.String(name),
    }
    params.Context = ctx
    
    customer, err := sc.client.V1Customers.Create(ctx, params)
    if err != nil {
        return nil, err
    }
    
    return customer, nil
}