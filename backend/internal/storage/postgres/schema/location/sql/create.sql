Insert into location (latitude, longitude, street_number, street_name, secondary_address, city, state, postal_code, country)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, latitude, longitude, street_number, street_name, secondary_address, city, state, postal_code, country, created_at, updated_at