INSERT INTO payment (
    registration_id,
    stripe_payment_intent_id,
    stripe_customer_id,
    org_stripe_account_id,
    stripe_payment_method_id,
    total_amount,
    provider_amount,
    platform_fee_amount,
    currency,
    payment_intent_status
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
