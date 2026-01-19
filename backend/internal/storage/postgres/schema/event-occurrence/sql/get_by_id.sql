SELECT id, manager_id, event_id, location_id, start_time, end_time, max_attendees, language, curr_enrolled, created_at, updated_at
FROM event_occurrence
WHERE id = $1