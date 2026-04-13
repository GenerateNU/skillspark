UPDATE "user"
SET name = $1, email = $2, username = $3, profile_picture_s3_key = $4, language_preference = $5, auth_id = $6
WHERE id = $7
RETURNING id, name, email, username, profile_picture_s3_key, language_preference, auth_id, created_at, updated_at;
