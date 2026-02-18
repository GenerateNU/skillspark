SELECT id, registration_id, guardian_id, description_en, description_th, categories, created_at, updated_at
FROM review
WHERE review.guardian_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;