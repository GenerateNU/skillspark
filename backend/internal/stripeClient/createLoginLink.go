package stripeClient

import (
	"context"
	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreateLoginLink(
	ctx context.Context,
	accountID string,
) (string, error) {
	params := &stripe.LoginLinkCreateParams{
		Account: stripe.String(accountID),
	}
	params.Context = ctx
	
	link, err := sc.client.V1LoginLinks.Create(ctx, params)
	if err != nil {
		return "", err
	}
	
	return link.URL, nil
}