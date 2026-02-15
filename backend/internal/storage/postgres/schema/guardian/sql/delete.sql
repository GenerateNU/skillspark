WITH guardian_registrations_and_times AS (
    SELECT r.id, r.guardian_id, eo.start_time, eo.end_time
    FROM registration r
    JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
    WHERE guardian_id = $1
),
cancelled_registrations AS (
    UPDATE registration r
    SET status = 'cancelled'
    WHERE r.id IN (
        SELECT grt.id
        FROM guardian_registrations_and_times grt
        WHERE grt.start_time > NOW()
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