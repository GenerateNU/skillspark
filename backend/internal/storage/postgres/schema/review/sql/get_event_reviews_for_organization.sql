SELECT
    e.id AS event_id,
    COUNT(r.id) AS total_reviews,
    COALESCE(AVG(r.rating), 0) AS average_rating,
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
LEFT JOIN event_occurrence eo ON eo.event_id = e.id
LEFT JOIN registration reg ON reg.event_occurrence_id = eo.id
LEFT JOIN review r ON r.registration_id = reg.id
WHERE e.organization_id = $1
GROUP BY e.id
ORDER BY %s
LIMIT $2 OFFSET $3;