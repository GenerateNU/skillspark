UPDATE event
SET title_en = $2, title_th = $3, description_en = $3, description_th = $4, organization_id = $5, age_range_min = $6, age_range_max = $7, category = $8, header_image_s3_key = $9, updated_at = NOW()
WHERE id = $1
RETURNING id, title_en, description_en, organization_id, age_range_min, age_range_max, category, header_image_s3_key, created_at, updated_at