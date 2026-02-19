UPDATE guardian
SET 
    stripe_customer_id = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, stripe_customer_id, created_at, updated_at;