SELECT s.id, s.name, s.location_id, s.created_at, s.updated_at
FROM school s
LEFT JOIN location l ON s.location_id = l.id;

