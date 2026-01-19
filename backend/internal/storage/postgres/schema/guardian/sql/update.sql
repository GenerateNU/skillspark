UPDATE guardian 
SET user_id = $2 
WHERE id = $1
RETURNING id, user_id, created_at, updated_at;

