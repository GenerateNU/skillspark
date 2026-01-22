DELETE from manager
WHERE id = $1
RETURNING id, user_id, organization_id, "role", created_at, updated_at