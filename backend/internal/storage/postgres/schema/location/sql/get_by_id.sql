SELECT id, latitude, longitude, address, city, state, zip_code, country, created_at, updated_at
FROM location
WHERE id = $1
