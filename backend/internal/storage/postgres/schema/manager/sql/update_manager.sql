UPDATE manager
SET 
    organization_id = COALESCE($2, organization_id), 
    "role" = COALESCE($3, "role"), 
    updated_at = NOW()
WHERE id = $1
RETURNING user_id, organization_id, "role", created_at, updated_at;