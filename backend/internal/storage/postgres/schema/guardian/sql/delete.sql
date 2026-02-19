WITH guardian_registrations AS (
    SELECT r.id, r.guardian_id
    FROM registration r
    WHERE guardian_id = $1
),
cancelled_registrations AS (
    UPDATE registration r
    SET status = 'cancelled'
    WHERE r.id IN (
        SELECT gr.id
        FROM guardian_registrations gr
    )
),
deleted_guardian AS (
    DELETE FROM guardian WHERE id = $1
    RETURNING id, user_id, created_at, updated_at
),
deleted_user AS (
    DELETE FROM "user"
    WHERE id = (SELECT user_id FROM deleted_guardian)
    RETURNING id, name, email, username, profile_picture_s3_key, language_preference, auth_id
)
SELECT dg.id, dg.user_id, du.name, du.email, du.username, du.profile_picture_s3_key, du.language_preference, du.auth_id, dg.created_at, dg.updated_at
FROM deleted_guardian dg
JOIN deleted_user du ON dg.user_id = du.id;