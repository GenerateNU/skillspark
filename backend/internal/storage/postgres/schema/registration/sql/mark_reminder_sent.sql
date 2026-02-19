UPDATE registration
SET reminder_sent = $2, updated_at = NOW()
WHERE id = $1
RETURNING id;
