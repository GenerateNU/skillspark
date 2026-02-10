UPDATE guardian 
SET updated_at = NOW()
WHERE id = $1
RETURNING user_id, stripe_customer_id, created_at, updated_at;