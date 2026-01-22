SELECT g.id, g.user_id, g.created_at, g.updated_at
FROM guardian g
INNER JOIN child c ON c.guardian_id = g.id
WHERE c.id = $1;