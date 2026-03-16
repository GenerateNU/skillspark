INSERT INTO saved (guardian_id, event_occurrence_id)
VALUES ($1, $2)
RETURNING id, guardian_id, event_occurrence_id, created_at, updated_at