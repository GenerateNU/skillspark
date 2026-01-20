UPDATE child c
SET
    name = COALESCE($1, c.name),
    school_id = COALESCE($2, c.school_id),
    birth_month = COALESCE($3, c.birth_month),
    birth_year = COALESCE($4, c.birth_year),
    interests = COALESCE($5, c.interests),
    guardian_id = COALESCE($6, c.guardian_id),
    updated_at = NOW()
FROM school s
WHERE c.id = $7
  AND c.school_id = s.id
RETURNING c.id, c.name, c.school_id, s.name AS school_name, c.birth_month, c.birth_year, c.interests, c.guardian_id, c.created_at, c.updated_at;
