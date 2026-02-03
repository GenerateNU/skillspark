insert into school (name, location_id)
VALUES ($1, $2)
RETURNING id, name, location_id, created_at, updated_at