WITH inserted AS (
    INSERT INTO child (name, school_id, birth_month, birth_year, interests, guardian_id, avatar_face, avatar_background)
    VALUES ($1, $2, $3, $4, $5::category[], $6, $7, $8)
    RETURNING *
)
SELECT i.id, i.name, i.school_id, s.name AS school_name, i.birth_month, i.birth_year,
       i.interests, i.guardian_id, i.avatar_face, i.avatar_background, i.created_at, i.updated_at
FROM inserted i
JOIN school s ON i.school_id = s.id;
