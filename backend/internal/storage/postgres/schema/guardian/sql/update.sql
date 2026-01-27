WITH g AS (
    SELECT user_id FROM guardian WHERE id = $1
),
updated_user AS (
    UPDATE "user"
    SET name = $2, email = $3, username = $4, profile_picture_s3_key = $5, language_preference = $6, updated_at = NOW()
    WHERE id = (SELECT user_id FROM g)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference
),
updated_guardian AS (
    UPDATE guardian 
    SET updated_at = NOW()
    WHERE id = $1
    RETURNING id, user_id, created_at, updated_at
)
SELECT ug.id, ug.user_id, uu.name, uu.email, uu.username, uu.profile_picture_s3_key, uu.language_preference, ug.created_at, ug.updated_at
FROM updated_guardian ug
JOIN updated_user uu ON ug.user_id = uu.id;