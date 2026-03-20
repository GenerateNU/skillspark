SELECT 
    s.id, 
    s.guardian_id, 
    s.created_at, 
    s.updated_at, 
        
    e.id,
    e.title_en,
    e.title_th,
    e.description_en,
    e.description_th,
    e.organization_id,
    e.age_range_min,
    e.age_range_max,
    e.category,
    e.header_image_s3_key,
    e.created_at,
    e.updated_at
FROM saved s 
JOIN event e ON e.id = s.event_id

WHERE s.guardian_id = $1
ORDER BY s.id
LIMIT $2 OFFSET $3;