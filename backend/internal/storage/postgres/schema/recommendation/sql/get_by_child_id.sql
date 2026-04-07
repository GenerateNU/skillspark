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
    e.updated_at,
    (
        SELECT COUNT(*)
        FROM unnest(e.category::text[]) AS cat
        WHERE cat = ANY($1::text[])
    ) AS score
FROM event e
JOIN organization o ON o.id = e.organization_id
JOIN location l ON l.id = o.location_id
WHERE EXISTS (
    SELECT 1 FROM event_occurrence eo
    WHERE eo.event_id = e.id
      AND eo.status = 'scheduled'
      AND eo.start_time > NOW()
      AND ($5::timestamptz IS NULL OR eo.start_time >= $5)
      AND ($6::timestamptz IS NULL OR eo.start_time <= $6)
)
AND (
    $2::int IS NULL OR e.age_range_min IS NULL OR (EXTRACT(YEAR FROM NOW()) - $2) >= e.age_range_min
)
AND (
    $2::int IS NULL OR e.age_range_max IS NULL OR (EXTRACT(YEAR FROM NOW()) - $2) <= e.age_range_max
)
AND (
    $7::float IS NULL
    OR $8::float IS NULL
    OR $9::float IS NULL
    OR earth_distance(
        ll_to_earth(l.latitude, l.longitude),
        ll_to_earth($7, $8)
    )/1000 <= $9
)
ORDER BY score DESC, e.created_at DESC, e.id DESC
LIMIT $3 OFFSET $4;
