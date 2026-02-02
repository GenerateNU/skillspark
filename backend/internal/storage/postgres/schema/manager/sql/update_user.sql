UPDATE "user"
SET 
    name = COALESCE($2, name), 
    email = COALESCE($3, email), 
    username = COALESCE($4, username), 
    profile_picture_s3_key = COALESCE($5, profile_picture_s3_key), 
    language_preference = COALESCE($6, language_preference), 
    updated_at = NOW()
WHERE id = $1
RETURNING name, email, username, profile_picture_s3_key, language_preference;