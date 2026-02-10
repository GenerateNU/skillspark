SELECT 
    id,
    guardian_id,
    stripe_payment_method_id,
    card_brand,
    card_last4,
    card_exp_month,
    card_exp_year,
    is_default,
    created_at,
    updated_at
FROM guardian_payment_methods
WHERE guardian_id = $1
ORDER BY is_default DESC, created_at DESC;