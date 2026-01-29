SELECT 
    r.id AS id,
    r.child_id AS child_id,
    r.guardian_id AS guardian_id,
    r.event_occurrence_id AS event_occurrence_id,
    r.status AS status,
    r.created_at AS created_at,
    r.updated_at AS updateed_at,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
WHERE $1 = r.guardian_id