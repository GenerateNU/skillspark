UPDATE guardian
SET expo_push_token = $2, updated_at = NOW()
WHERE id = $1
RETURNING user_id, expo_push_token, created_at, updated_at;