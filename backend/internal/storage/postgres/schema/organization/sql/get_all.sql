SELECT id, name, active, pfp_s3_key, location_id, stripe_account_id, stripe_account_activated, created_at, updated_at
FROM organization
ORDER BY created_at DESC
LIMIT $1 OFFSET $2