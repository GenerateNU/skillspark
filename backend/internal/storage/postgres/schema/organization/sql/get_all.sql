SELECT id, name, active, pfp_s3_key, location_id, created_at, updated_at
FROM organization
ORDER BY created_at DESC
LIMIT $1 OFFSET $2