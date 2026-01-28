WITH new_user AS (
    INSERT INTO "user" (name, email, username, profile_picture_s3_key, language_preference)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference
),
new_guardian AS (
    INSERT INTO guardian (user_id)
    SELECT id FROM new_user
    RETURNING id, user_id, created_at, updated_at
)
SELECT g.id, g.user_id, u.name, u.email, u.username, u.profile_picture_s3_key, u.language_preference, g.created_at, g.updated_at
FROM new_guardian g
JOIN new_user u ON g.user_id = u.id;