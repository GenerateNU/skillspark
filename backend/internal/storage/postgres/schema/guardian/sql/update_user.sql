UPDATE "user"
SET name = $2, email = $3, username = $4, profile_picture_s3_key = $5, language_preference = $6, updated_at = NOW()
WHERE id = $1
RETURNING name, email, username, profile_picture_s3_key, language_preference;