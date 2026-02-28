WITH updated AS (
    UPDATE registration
    SET
        child_id = $1,
        guardian_id = $2,
        event_occurrence_id = $3,
        status = $4,
        updated_at = NOW()
    WHERE id = $5
    RETURNING id, child_id, guardian_id, event_occurrence_id, status, created_at, updated_at
)
SELECT 
    u.id,
    u.child_id,
    u.guardian_id,
    u.event_occurrence_id,
    u.status,
    u.created_at,
    u.updated_at,
    e.title_en AS event_name,
    eo.start_time AS occurrence_start_time
FROM updated u
JOIN event_occurrence eo ON u.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;