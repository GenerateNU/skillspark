UPDATE registration
SET status = 'cancelled'
WHERE event_occurrence_id = $1;
