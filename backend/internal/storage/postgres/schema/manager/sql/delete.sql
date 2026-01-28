WITH deleted_manager AS (
    DELETE FROM manager WHERE id = $1
    RETURNING id, user_id, organization_id, "role", created_at, updated_at
),
deleted_user AS (
    DELETE FROM "user" 
    WHERE id = (SELECT user_id FROM deleted_manager)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference
)
SELECT dm.id, dm.user_id, dm.organization_id, dm.role, du.name, du.email, du.username, du.profile_picture_s3_key, du.language_preference, dm.created_at, dm.updated_at
FROM deleted_manager dm
JOIN deleted_user du ON dm.user_id = du.id;