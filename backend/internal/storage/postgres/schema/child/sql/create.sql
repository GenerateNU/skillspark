WITH inserted AS {
    insert into child (name, school_id, birth_month, birth_year, interests, guardian_id)
    VALUES ($1 $2 $3 $4 $5 $6)
    RETURNING *
}
SELECT i.id, i.name, i.school_id, s.name AS school_name, i.birth_month, i.birth_year,
       i.interests, i.guardian_id, i.created_at, i.updated_at
FROM inserted i
JOIN school s ON i.school_id = s.id;