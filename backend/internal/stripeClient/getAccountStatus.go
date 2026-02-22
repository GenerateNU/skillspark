package stripeClient

import (
	"context"

	"github.com/stripe/stripe-go/v84"
)

type AccountStatus struct {
    AccountID          string
    ChargesEnabled     bool
    PayoutsEnabled     bool
    RequirementsErrors []error
    CurrentlyDue       []string
    EventuallyDue      []string
}

func (sc *StripeClient) GetAccountStatus(
    ctx context.Context,
    accountID string,
) (*AccountStatus, error) {
    params := &stripe.V2CoreAccountRetrieveParams{
		Include: []*string{
			stripe.String("defaults"),
			stripe.String("identity"),
			stripe.String("configuration.merchant"),
		},
	}

    // deprecated replace
    params.AddExpand("requirements")
    
    acct, err := sc.client.V2CoreAccounts.Retrieve(ctx, accountID, params)
    if err != nil {
        return nil, err
    }

    
    status := &AccountStatus{
        AccountID: acct.ID,
    }
    
    if acct.Configuration != nil && acct.Configuration.Merchant != nil {
        if cardPayments := acct.Configuration.Merchant.Capabilities.CardPayments; cardPayments != nil {
            status.ChargesEnabled = cardPayments.Status == "active"
        }
    }
    
    if acct.Configuration != nil && acct.Configuration.Recipient != nil {
        if stripeBalance := acct.Configuration.Recipient.Capabilities.StripeBalance; stripeBalance != nil {
            if stripeTransfers := stripeBalance.StripeTransfers; stripeTransfers != nil {
                status.PayoutsEnabled = stripeTransfers.Status == "active"
            }
        }
    }

    // TODO: add requirement info to status
    return status, nil
}