WITH new_user AS (
    INSERT INTO "user" (name, email, username, profile_picture_s3_key, language_preference)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference
),
new_manager AS (
    INSERT INTO manager (user_id, organization_id, "role")
    SELECT id, $6, $7 FROM new_user
    RETURNING id, user_id, organization_id, "role", created_at, updated_at
)
SELECT m.id, m.user_id, m.organization_id, m.role, u.name, u.email, u.username, u.profile_picture_s3_key, u.language_preference, m.created_at, m.updated_at
FROM new_manager m
JOIN new_user u ON m.user_id = u.id;