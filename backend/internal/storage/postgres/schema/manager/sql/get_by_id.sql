SELECT m.id, m.user_id, m.organization_id, m.role, u.name, u.email, u.username, u.profile_picture_s3_key, u.language_preference, u.auth_id, m.created_at, m.updated_at
FROM manager m
JOIN "user" u ON m.user_id = u.id
WHERE m.id = $1