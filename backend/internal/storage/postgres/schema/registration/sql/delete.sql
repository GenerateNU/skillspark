WITH deleted AS (
    DELETE FROM registration
    WHERE id = $1
    RETURNING id, child_id, guardian_id, event_occurrence_id, status, created_at, updated_at
)
SELECT 
    d.id,
    d.child_id,
    d.guardian_id,
    d.event_occurrence_id,
    d.status,
    d.created_at,
    d.updated_at,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time
FROM deleted d
JOIN event_occurrence eo ON d.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;