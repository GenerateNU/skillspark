DELETE FROM child c
USING school s
WHERE c.id = $1
  AND c.school_id = s.id
RETURNING 
    c.id,
    c.name,
    c.school_id,
    s.name AS school_name,
    c.birth_month,
    c.birth_year,
    c.interests,
    c.guardian_id,
    c.created_at,
    c.updated_at;
