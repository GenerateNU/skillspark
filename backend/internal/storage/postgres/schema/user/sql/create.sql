INSERT INTO "user" (name, email, username, profile_picture_s3_key, language_preference, auth_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, email, username, profile_picture_s3_key, language_preference, auth_id, push_notifications, email_notifications, created_at, updated_at;
