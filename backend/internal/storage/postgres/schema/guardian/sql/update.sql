UPDATE guardian 
SET user_id = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, created_at, updated_at;

