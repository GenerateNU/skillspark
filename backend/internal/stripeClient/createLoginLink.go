package stripeClient

import (
	"context"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/loginlink"
)

func (sc *StripeClient) CreateLoginLink(
	ctx context.Context,
	accountID string,
) (string, error) {
	params := &stripe.LoginLinkParams{
		Account: stripe.String(accountID),
	}
	params.Context = ctx
	
	link, err := loginlink.New(params)
	if err != nil {
		return "", err
	}
	
	return link.URL, nil
}