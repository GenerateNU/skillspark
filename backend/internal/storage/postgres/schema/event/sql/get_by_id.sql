SELECT 
    e.id,
    e.title,
    e.description,
    e.organization_id,
    e.age_range_min,
    e.age_range_max,
    e.category,
    e.header_image_s3_key,
    e.created_at,
    e.updated_at
FROM event e
WHERE e.id = $1;