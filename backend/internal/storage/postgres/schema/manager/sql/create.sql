Insert into manager (user_id, organization_id, "role")
VALUES ($1, $2, $3)
RETURNING id, user_id, organization_id, "role", created_at, updated_at