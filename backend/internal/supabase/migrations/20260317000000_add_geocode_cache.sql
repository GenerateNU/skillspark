CREATE TABLE geocode_cache (
    address     TEXT PRIMARY KEY,
    raw_address TEXT NOT NULL,
    latitude    DOUBLE PRECISION NOT NULL,
    longitude   DOUBLE PRECISION NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
