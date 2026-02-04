DELETE FROM "user"
WHERE id = $1
RETURNING id, name, email, username, profile_picture_s3_key, language_preference, auth_id, created_at, updated_at;
