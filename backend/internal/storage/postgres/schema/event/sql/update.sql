UPDATE event
SET title = $2, description = $3, organization_id = $4, age_range_min = $5, age_range_max = $6, category = $7, header_image_s3_key = $8, updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, organization_id, age_range_min, age_range_max, category, header_image_s3_key, created_at, updated_at