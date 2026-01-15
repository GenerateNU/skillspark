Insert into location (latitude, longitude, address_line1, address_line2, subdistrict, district, province, postal_code, country)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, latitude, longitude, address_line1, address_line2, subdistrict, district, province, postal_code, country, created_at, updated_at