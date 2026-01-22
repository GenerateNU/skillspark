DELETE FROM guardian WHERE id = $1
RETURNING id, user_id, created_at, updated_at;