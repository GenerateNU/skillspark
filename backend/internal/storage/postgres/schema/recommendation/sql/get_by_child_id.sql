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
WHERE EXISTS (
    SELECT 1 FROM event_occurrence eo
    WHERE eo.event_id = e.id
      AND eo.status = 'scheduled'
      AND eo.start_time > NOW()
      AND ($5::timestamp IS NULL OR eo.start_time >= $5)
      AND ($6::timestamp IS NULL OR eo.start_time <= $6)
)
AND (
    $2::int IS NULL OR e.age_range_min IS NULL OR (EXTRACT(YEAR FROM NOW()) - $2) >= e.age_range_min
)
AND (
    $2::int IS NULL OR e.age_range_max IS NULL OR (EXTRACT(YEAR FROM NOW()) - $2) <= e.age_range_max
)
ORDER BY score DESC, e.created_at DESC, e.id DESC
LIMIT $3 OFFSET $4;
