SELECT c.id, c.school_id, s.name as school_name, c.birth_month, c.birth_year, c.interests, c.guardian_id, c.created_at, c.updated_at
FROM child c
JOIN school s on c.school_id = s.id
where c.id = $1 