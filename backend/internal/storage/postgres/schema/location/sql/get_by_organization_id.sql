SELECT 
    l.id,
    l.latitude,
    l.longitude,
    l.address_line1,
    l.address_line2,
    l.subdistrict,
    l.district,
    l.province,
    l.postal_code,
    l.country,
    l.created_at,
    l.updated_at
FROM location l
INNER JOIN organization o ON o.location_id = l.id
WHERE o.id = $1;