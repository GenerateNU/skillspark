SELECT
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
FROM event e
WHERE 1=1
AND ($3::text IS NULL OR e.title_en ILIKE '%' || $3 || '%' OR e.description_en ILIKE '%' || $3 || '%' OR e.title_th ILIKE '%' || $3 || '%' OR e.description_th ILIKE '%' || $3 || '%')
AND ($4::text IS NULL OR $4 = '' OR e.category::text[] && string_to_array($4, ','))
AND ($5::int IS NULL OR e.age_range_max IS NULL OR e.age_range_max >= $5)
AND ($6::int IS NULL OR e.age_range_min IS NULL OR e.age_range_min <= $6)
ORDER BY e.created_at DESC
LIMIT $1 OFFSET $2;
