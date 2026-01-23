DELETE FROM organization
WHERE id = $1
RETURNING id, name, active, pfp_s3_key, location_id, created_at, updated_at