SELECT address, raw_address, latitude, longitude, created_at
FROM geocode_cache
WHERE address = $1
