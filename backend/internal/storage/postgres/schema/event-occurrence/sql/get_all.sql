SELECT 
    eo.id,
    eo.manager_id,
    eo.start_time,
    eo.end_time,
    eo.max_attendees,
    eo.language,
    eo.curr_enrolled,
    eo.created_at,
    eo.updated_at,
    eo.status,
    eo.price,

    e.id,
    e.title,
    e.description,
    e.organization_id,
    e.age_range_min,
    e.age_range_max,
    e.category,
    e.header_image_s3_key,
    e.created_at,
    e.updated_at,

    l.id,
    l.latitude,
    l.longitude,
    l.address_line1,
    l.address_line2,
    l.subdistrict,
    l.district,
    l.province,
    l.postal_code,
    l.country,
    l.created_at,
    l.updated_at
FROM event_occurrence eo
JOIN event e ON e.id = eo.event_id
JOIN location l ON l.id = eo.location_id

WHERE 1=1
AND ($3::text IS NULL OR e.title ILIKE '%' || $3 || '%' OR e.description ILIKE '%' || $3 || '%')
AND ($4::int IS NULL OR EXTRACT(EPOCH FROM (eo.end_time - eo.start_time))/60 >= $4)
AND ($5::int IS NULL OR EXTRACT(EPOCH FROM (eo.end_time - eo.start_time))/60 <= $5)
AND (
    $6::float IS NULL 
    OR $7::float IS NULL
    OR $8::float IS NULL
    OR earth_distance(
        ll_to_earth(l.latitude, l.longitude),
        ll_to_earth($6, $7)
    )/1000 <= $8
)
AND ($9::int IS NULL OR (e.age_range_min >= $9))
AND ($10::int IS NULL OR (e.age_range_max <= $10))
AND ($11::category IS NULL OR $11::category = ANY(e.category))
AND ($12::boolean IS NULL OR ((eo.curr_enrolled >= eo.max_attendees) = $12))
AND ($13::timestamp IS NULL OR eo.start_time >= $13)
AND ($14::timestamp IS NULL OR eo.end_time <= $14)

ORDER BY eo.id
LIMIT $1 OFFSET $2;