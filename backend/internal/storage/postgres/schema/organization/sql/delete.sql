DELETE FROM organization
WHERE id = $1
RETURNING id, name, active, pfp_s3_key, location_id, stripe_account_id, stripe_account_activated, created_at, updated_at;
