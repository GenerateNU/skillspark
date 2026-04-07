INSERT INTO organization (name, active, pfp_s3_key, location_id, links, about)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, active, about, pfp_s3_key, location_id, links, stripe_account_id, stripe_account_activated, created_at, updated_at;