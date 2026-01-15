SELECT id, latitude, longitude, address_line1, address_line2, subdistrict, district, province, postal_code, country, created_at, updated_at
FROM location
WHERE id = $1
