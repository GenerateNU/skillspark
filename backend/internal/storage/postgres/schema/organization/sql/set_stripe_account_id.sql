UPDATE organization
SET 
    stripe_account_id = $1,
    updated_at = NOW()
WHERE id = $2
RETURNING id, name, active, pfp_s3_key, location_id, stripe_account_id, stripe_account_activated, created_at, updated_at;