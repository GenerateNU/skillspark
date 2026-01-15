SELECT s.id, s.name, s.location_id, s.created_at, s.updated_at
FROM school s
ORDER BY s.created_at DESC
LIMIT $1 OFFSET $2