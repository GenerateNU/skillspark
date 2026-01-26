WITH updated_row AS (
    UPDATE event_occurrence eo
    SET 
        manager_id = COALESCE($2, eo.manager_id),
        event_id = COALESCE($3, eo.event_id),
        location_id = COALESCE($4, eo.location_id),
        start_time = COALESCE($5, eo.start_time),
        end_time = COALESCE($6, eo.end_time),
        max_attendees = COALESCE($7, eo.max_attendees),
        language = COALESCE($8, eo.language),
        curr_enrolled = COALESCE($9, eo.curr_enrolled),
        updated_at = NOW()
    WHERE eo.id = $1
    RETURNING eo.id, eo.manager_id, eo.event_id, eo.location_id, eo.start_time, eo.end_time, eo.max_attendees, eo.language, eo.curr_enrolled, eo.created_at, eo.updated_at
)
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
FROM updated_row eo
JOIN event e ON e.id = eo.event_id
JOIN location l ON l.id = eo.location_id;