WITH unset_default AS (
    UPDATE guardian_payment_methods
    SET is_default = false
    WHERE guardian_id = (
        SELECT guardian_id FROM guardian_payment_methods WHERE id = $1
    )
    AND is_default = true
    AND $2 = true
)
UPDATE guardian_payment_methods
SET 
    is_default = $2,
    updated_at = NOW()
WHERE id = $1
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