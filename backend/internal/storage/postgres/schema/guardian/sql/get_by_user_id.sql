SELECT id, user_id, created_at, updated_at
FROM guardian
WHERE user_id = $1;