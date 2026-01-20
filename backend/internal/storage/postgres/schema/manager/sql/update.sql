UPDATE manager
SET user_id = $2, organization_id = $3, "role" = $4
WHERE id = $1
RETURNING id, user_id, organization_id, "role", created_at, updated_at