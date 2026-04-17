SELECT id, name, active, about_en, about_th, pfp_s3_key, location_id, links, stripe_account_id, stripe_account_activated, created_at, updated_at
FROM organization
WHERE id = $1;
