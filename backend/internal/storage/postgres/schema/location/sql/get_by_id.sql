SELECT id, latitude, longitude, street_number, street_name, secondary_address, city, state, postal_code, country, created_at, updated_at
FROM location
WHERE id = $1
