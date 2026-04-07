CREATE OR REPLACE FUNCTION haversine_distance(
    lat1 FLOAT, lon1 FLOAT,
    lat2 FLOAT, lon2 FLOAT
) RETURNS FLOAT AS $$

DECLARE
    r        FLOAT := 6371;
    dlat     FLOAT;
    dlon     FLOAT;
    a        FLOAT;
    c        FLOAT;
    distance FLOAT;
BEGIN
    dlat:= RADIANS(lat2 - lat1);
    dlon := RADIANS(lon2 - lon1);

    a := SIN(dlat / 2) * SIN(dlat / 2) +
            COS(RADIANS(lat1)) * COS(RADIANS(lat2)) *
            SIN(dlon / 2) * SIN(dlon / 2);

    c := 2 * ATAN2(SQRT(a), SQRT(1 - a));

    distance := r * c;

    RETURN distance;
END;
$$ LANGUAGE plpgsql IMMUTABLE;