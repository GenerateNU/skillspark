SELECT l.id, l.latitude, l.longitude, l.address_line1, l.address_line2, l.subdistrict, l.district, l.province, l.postal_code, l.country, l.created_at, l.updated_at
FROM location l
ORDER BY l.created_at DESC
LIMIT $1 OFFSET $2