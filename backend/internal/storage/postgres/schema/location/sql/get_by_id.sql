SELECT id, latitude, longitude, address, city, state, zip_code, country, created_at, updated_at
FROM locations
WHERE id = $1
