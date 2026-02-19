INSERT INTO notification (registration_id, type, payload, created_at)
VALUES ($1, $2, $3, NOW())
RETURNING id, registration_id, type, payload, created_at;
