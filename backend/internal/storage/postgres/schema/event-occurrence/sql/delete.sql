BEGIN;

UPDATE registration
SET status = 'cancelled'
WHERE event_occurrence_id = $1;

DELETE FROM event_occurrence
WHERE id = $1;

COMMIT;