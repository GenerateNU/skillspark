DELETE FROM child
WHERE id = $1
RETURNING id, school_id, birth_month, birth_year, interests, guardian_id, created_at, updated_at;