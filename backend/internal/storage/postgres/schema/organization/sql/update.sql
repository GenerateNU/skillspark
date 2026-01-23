UPDATE organization
SET name = $1, active = $2, pfp_s3_key = $3, location_id = $4, updated_at = Now()
WHERE id = $5
RETURNING id, name, active, pfp_s3_key, location_id, created_at, updated_at;
