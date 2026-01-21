INSERT INTO guardian (user_id) VALUES ($1) 
RETURNING id, user_id, created_at, updated_at;