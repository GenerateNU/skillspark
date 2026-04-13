WITH nearby_orgs AS (
    SELECT o.id
    FROM organization o
    JOIN location l ON o.location_id = l.id
    WHERE 6371 * 2 * ATAN2(
            SQRT(
                SIN(RADIANS((l.latitude - $1) / 2)) ^ 2 +
                COS(RADIANS($1)) * COS(RADIANS(l.latitude)) *
                SIN(RADIANS((l.longitude - $2) / 2)) ^ 2
            ),
            SQRT(1 - (
                SIN(RADIANS((l.latitude - $1) / 2)) ^ 2 +
                COS(RADIANS($1)) * COS(RADIANS(l.latitude)) *
                SIN(RADIANS((l.longitude - $2) / 2)) ^ 2
            )) -- haversine distance function
        ) <= $3 -- in kilometers
),

event_popularity AS (
    SELECT
        eo.event_id,
        SUM(eo.curr_enrolled) AS total_enrolled
    FROM event_occurrence eo
    WHERE eo.start_time < NOW()
      AND eo.start_time > NOW() - INTERVAL '14 days'
    GROUP BY eo.event_id
),

ranked_occurrences AS (
    SELECT
        eo.id,
        eo.event_id,
        eo.manager_id,
        eo.start_time,
        eo.end_time,
        eo.max_attendees,
        eo.language,
        eo.curr_enrolled,
        eo.created_at,
        eo.updated_at,
        eo.status,
        eo.price,
        eo.currency,
        ROW_NUMBER() OVER (
            PARTITION BY eo.event_id
            ORDER BY eo.start_time
        ) AS rn
    FROM event_occurrence eo
    WHERE eo.start_time > NOW() + INTERVAL '1 day'
)

SELECT
    ro.id,
    ro.manager_id,
    ro.start_time,
    ro.end_time,
    ro.max_attendees,
    ro.language,
    ro.curr_enrolled,
    ro.created_at AS occurrence_created_at,
    ro.updated_at AS occurrence_updated_at,
    ro.status,
    ro.price,
    ro.currency,

    e.id AS event_id,
    e.title_en,
    e.title_th,
    e.description_en,
    e.description_th,
    e.organization_id,
    e.age_range_min,
    e.age_range_max,
    e.category,
    e.header_image_s3_key,
    e.created_at AS event_created_at,
    e.updated_at AS event_updated_at,

    l.id AS location_id,
    l.latitude,
    l.longitude,
    l.address_line1,
    l.address_line2,
    l.subdistrict,
    l.district,
    l.province,
    l.postal_code,
    l.country,
    l.created_at AS location_created_at,
    l.updated_at AS location_updated_at,

    o.links AS org_links

FROM ranked_occurrences ro
JOIN event e ON e.id = ro.event_id
JOIN organization o ON o.id = e.organization_id
JOIN location l ON l.id = o.location_id
JOIN nearby_orgs ON nearby_orgs.id = o.id
LEFT JOIN event_popularity ep ON ep.event_id = e.id

WHERE ro.rn = 1

ORDER BY COALESCE(ep.total_enrolled, 0) DESC, ro.start_time ASC
LIMIT $4;