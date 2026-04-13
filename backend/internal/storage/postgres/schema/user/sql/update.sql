UPDATE "user"
SET name = $1, email = $2, username = $3, profile_picture_s3_key = $4, language_preference = $5, auth_id = $6, push_notifications = $7, email_notifications = $8
WHERE id = $9
RETURNING id, name, email, username, profile_picture_s3_key, language_preference, auth_id, push_notifications, email_notifications, created_at, updated_at;
