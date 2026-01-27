WITH m AS (
    SELECT user_id FROM manager WHERE id = $1
),
updated_user AS (
    UPDATE "user"
    SET name = $2, email = $3, username = $4, profile_picture_s3_key = $5, language_preference = $6, updated_at = NOW()
    WHERE id = (SELECT user_id FROM m)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference
),
updated_manager AS (
    UPDATE manager
    SET organization_id = $7, "role" = $8, updated_at = NOW()
    WHERE id = $1
    RETURNING id, user_id, organization_id, "role", created_at, updated_at
)
SELECT um.id, um.user_id, um.organization_id, um.role, uu.name, uu.email, uu.username, uu.profile_picture_s3_key, uu.language_preference, um.created_at, um.updated_at
FROM updated_manager um
JOIN updated_user uu ON um.user_id = uu.id;