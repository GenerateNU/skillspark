INSERT INTO review (registration_id, guardian_id, description, categories)
VALUES ($1, $2, $3, $4)
RETURNING id, registration_id, guardian_id, description, categories, created_at, updated_at