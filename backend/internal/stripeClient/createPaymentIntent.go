package stripeClient

import (
	"context"
	"errors"
	"skillspark/internal/models"
	"time"

	"github.com/stripe/stripe-go/v84"
)



func (sc *StripeClient) CreatePaymentIntent(ctx context.Context, input *models.CreatePaymentIntentInput) (*models.CreatePaymentIntentOutput, error) {
	
	const applicationFeePercentage = 10 // CHANGE THIS TO BE THE APPLICATION FEE PERCENTAGE

	applicationFeeTotal := (input.Body.Amount * int64(applicationFeePercentage)) / 100 
	organizationProfit := input.Body.Amount - applicationFeeTotal

	if input.Body.PaymentMethodID == nil || *input.Body.PaymentMethodID == "" {
		return nil, errors.New("payment method required for booking")
	}

	
	params := &stripe.PaymentIntentCreateParams{
  		Amount: stripe.Int64(input.Body.Amount),
  		Currency: stripe.String(input.Body.Currency),
		OnBehalfOf: stripe.String(input.Body.OrgStripeID),
		Customer: stripe.String(input.Body.GuardianStripeID),
		ApplicationFeeAmount: stripe.Int64(applicationFeeTotal),
		TransferData: &stripe.PaymentIntentCreateTransferDataParams{
			Destination: stripe.String(input.Body.OrgStripeID),
			Amount: stripe.Int64(organizationProfit),
		},
		OffSession: stripe.Bool(true),
		Metadata: map[string]string{
			"event_date":      input.Body.EventDate.Format(time.RFC3339),
		},
	}

	if input.Body.PaymentMethodID != nil {
		// reusing saved card
		params.PaymentMethod = stripe.String(*input.Body.PaymentMethodID)
		params.OffSession = stripe.Bool(true)
	} else {
		// New card - save for future use
		params.SetupFutureUsage = stripe.String("off_session")
	}

	intent, err := sc.client.V1PaymentIntents.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	
	output := &models.CreatePaymentIntentOutput{}
	output.Body.ClientSecret = intent.ClientSecret
	output.Body.PaymentIntentID = intent.ID
	output.Body.Status = string(intent.Status)
	output.Body.TotalAmount = int(intent.Amount)
	output.Body.ProviderAmount = int(intent.TransferData.Amount)
	output.Body.PlatformFeeAmount = int(intent.ApplicationFeeAmount)
	output.Body.Currency = input.Body.Currency
	
	return output, nil
}