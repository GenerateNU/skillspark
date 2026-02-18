Insert into event(title_en, title_th, description_en, description_th, organization_id, age_range_min, age_range_max, category, header_image_s3_key)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, title_en, description_en, organization_id, age_range_min, age_range_max, category, header_image_s3_key, created_at, updated_at
