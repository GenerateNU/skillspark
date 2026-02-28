SELECT 
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.created_at,
    r.updated_at,
    e.title_en AS event_name,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
WHERE r.id = $1;