Insert into location (latitude, longitude, address, city, state, zip_code, country)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, latitude, longitude, address, city, state, zip_code, country, created_at, updated_at