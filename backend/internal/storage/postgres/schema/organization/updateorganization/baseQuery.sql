UPDATE organization
SET name = $1, active = $2, pfp_s3_key = $3, location_id = $4, updated_at = $5
WHERE id = $6