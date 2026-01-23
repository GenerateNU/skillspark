INSERT INTO organization (name, active, pfp_s3_key, location_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, active, pfp_s3_key, location_id, created_at, updated_at