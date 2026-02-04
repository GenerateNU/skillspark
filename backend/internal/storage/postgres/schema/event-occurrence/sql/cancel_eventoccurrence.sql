UPDATE event_occurrence
SET status = 'cancelled',
    updated_at = NOW()
WHERE id = $1;
