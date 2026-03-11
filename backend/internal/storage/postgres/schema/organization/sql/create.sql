INSERT INTO organization (name, active, pfp_s3_key, location_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, active, pfp_s3_key, location_id, stripe_account_id, stripe_account_activated, created_at, updated_at;
