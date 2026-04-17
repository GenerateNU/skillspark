SELECT o.id, o.name, o.active, o.about_en, o.about_th, o.pfp_s3_key, o.location_id, o.stripe_account_id, o.stripe_account_activated, o.created_at, o.updated_at,
       l.latitude::float8, l.longitude::float8, l.district
FROM organization o
LEFT JOIN location l ON l.id = o.location_id
ORDER BY o.created_at DESC
LIMIT $1 OFFSET $2
