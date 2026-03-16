SELECT id, guardian_id, event_occurrence_id, created_at, updated_at
FROM saved 
WHERE saved.guardian_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;