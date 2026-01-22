SELECT id, user_id, organization_id, "role", created_at, updated_at
FROM manager
WHERE organization_id = $1