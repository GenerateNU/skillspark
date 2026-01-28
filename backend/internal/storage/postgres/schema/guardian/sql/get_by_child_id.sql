SELECT g.id, g.user_id, u.name, u.email, u.username, u.profile_picture_s3_key, u.language_preference, g.created_at, g.updated_at
FROM guardian g
JOIN "user" u ON g.user_id = u.id
INNER JOIN child c ON c.guardian_id = g.id
WHERE c.id = $1;