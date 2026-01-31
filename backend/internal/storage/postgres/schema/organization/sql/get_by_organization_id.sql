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
WHERE e.organization_id = $1;