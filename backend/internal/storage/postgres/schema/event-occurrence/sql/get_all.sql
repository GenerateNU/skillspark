SELECT id, manager_id, event_id, location_id, start_time, end_time, max_attendees, language, curr_enrolled, created_at, updated_at
FROM event_occurrence
ORDER BY created_at DESC
LIMIT $1 OFFSET $2