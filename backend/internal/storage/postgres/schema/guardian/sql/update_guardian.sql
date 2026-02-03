UPDATE guardian 
SET updated_at = NOW()
WHERE id = $1
RETURNING user_id, created_at, updated_at;