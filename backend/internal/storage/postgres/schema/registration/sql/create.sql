WITH inserted AS (
    INSERT INTO registration (child_id, guardian_id, event_occurrence_id, status)
    VALUES ($1, $2, $3, $4)
    RETURNING id, child_id, guardian_id, event_occurrence_id, status, created_at, updated_at
)
SELECT 
    i.id as id,
    i.child_id as child_id,
    i.guardian_id as guardian_id,
    i.event_occurrence_id as event_occurrence_id,
    i.status as status,
    i.created_at as created_at,
    i.updated_at as updated_at,
    e.title_en AS event_name,
    eo.start_time AS occurrence_start_time
FROM inserted i
JOIN event_occurrence eo ON i.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;