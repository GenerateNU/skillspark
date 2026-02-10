INSERT INTO guardian_payment_methods (
    guardian_id,
    stripe_payment_method_id,
    card_brand,
    card_last4,
    card_exp_month,
    card_exp_year,
    is_default
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING 
    id,
    guardian_id,
    stripe_payment_method_id,
    card_brand,
    card_last4,
    card_exp_month,
    card_exp_year,
    is_default,
    created_at,
    updated_at;