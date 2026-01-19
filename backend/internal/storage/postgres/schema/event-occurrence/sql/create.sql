INSERT INTO event_occurrence (manager_id, event_id, location_id, start_time, end_time, max_attendees, language)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, manager_id, event_id, location_id, start_time, end_time, max_attendees, language, curr_enrolled, created_at, updated_at