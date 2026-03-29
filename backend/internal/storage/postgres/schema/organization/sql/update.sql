UPDATE organization
SET name = $1, active = $2, pfp_s3_key = $3, location_id = $4, links = $5, updated_at = Now()
WHERE id = $6
RETURNING id, name, active, pfp_s3_key, location_id, links, stripe_account_id, stripe_account_activated, created_at, updated_at;