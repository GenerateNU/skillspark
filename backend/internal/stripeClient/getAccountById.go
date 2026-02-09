package stripeClient

import (
	"context"
	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) GetAccount(
	ctx context.Context,
	accountID string,
) (*stripe.V2CoreAccount, error) {
	params := &stripe.V2CoreAccountRetrieveParams{
		Include: []*string{
			stripe.String("defaults"),
			stripe.String("identity"),
			stripe.String("configuration.merchant"),
			stripe.String("configuration.recipient"),
		},
	}
	
	acct, err := sc.client.V2CoreAccounts.Retrieve(ctx, accountID, params)
	if err != nil {
		return nil, err
	}
	
	return acct, nil

}