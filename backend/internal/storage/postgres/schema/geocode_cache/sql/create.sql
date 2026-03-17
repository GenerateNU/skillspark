INSERT INTO geocode_cache (address, raw_address, latitude, longitude)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
    SET raw_address = EXCLUDED.raw_address,
        latitude    = EXCLUDED.latitude,
        longitude   = EXCLUDED.longitude,
        created_at  = NOW()
RETURNING address, raw_address, latitude, longitude, created_at
