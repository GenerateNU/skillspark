package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (s *StripeClient) GetPaymentMethodsByCustomerID(ctx context.Context, customerID string) (*models.GetPaymentMethodsByGuardianIDOutput, error) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type:     stripe.String("card"),
	}
	params.Context = ctx

	pmList := s.client.V1PaymentMethods.List(ctx, params)

	var paymentMethods []models.PaymentMethod
	for pm, err := range pmList {
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, models.PaymentMethod{
			ID:   pm.ID,
			Type: string(pm.Type),
			Card: models.PaymentMethodCard{
				Brand:    string(pm.Card.Brand),
				Last4:    pm.Card.Last4,
				ExpMonth: pm.Card.ExpMonth,
				ExpYear:  pm.Card.ExpYear,
			},
		})
	}

	return &models.GetPaymentMethodsByGuardianIDOutput{
		Body: struct {
			PaymentMethods []models.PaymentMethod `json:"payment_methods"`
		}{
			PaymentMethods: paymentMethods,
		},
	}, nil
}
