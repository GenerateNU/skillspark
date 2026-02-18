INSERT INTO review (registration_id, guardian_id, description_en, description_th, categories)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, registration_id, guardian_id, description_en, categories, created_at, updated_at