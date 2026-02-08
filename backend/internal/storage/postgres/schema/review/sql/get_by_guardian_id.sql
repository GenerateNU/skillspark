SELECT id, registration_id, guardian_id, description, categories, created_at, updated_at
FROM review
WHERE review.guardian_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;