Insert into event(title, description, organization_id, age_range_min, age_range_max, category, header_image_s3_key)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, description, organization_id, age_range_min, age_range_max, category, header_image_s3_key, created_at, updated_at
