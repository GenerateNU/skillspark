SELECT m.id, m.user_id, m.organization_id, m.role, m.created_at, m.updated_at
FROM manager m
INNER JOIN "user" u ON m.user_id = u.id
WHERE u.auth_id = $1;
