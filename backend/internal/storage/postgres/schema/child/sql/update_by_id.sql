UPDATE child
SET
    name = COALESCE($1, name),
    school_id = COALESCE($2, school_id),
    birth_month = COALESCE($3, birth_month),
    birth_year = COALESCE($4, birth_year),
    interests = COALESCE($5, interests),
    guardian_id = COALESCE($6, guardian_id),
    updated_at = NOW()
WHERE id = $7
RETURNING id, school_id, birth_month, birth_year, interests, guardian_id, created_at, updated_at;
