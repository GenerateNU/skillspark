WITH inserted AS (
    INSERT INTO review (registration_id, guardian_id, rating, description_en, description_th, categories)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, registration_id, guardian_id, rating, description_en, description_th, categories, created_at, updated_at
)
SELECT i.id, i.registration_id, i.guardian_id, eo.event_id, i.rating, i.description_en, i.description_th, i.categories, i.created_at, i.updated_at
FROM inserted i
JOIN registration reg ON i.registration_id = reg.id
JOIN event_occurrence eo ON reg.event_occurrence_id = eo.id;