SELECT r.id, r.registration_id, r.guardian_id, r.description_en, r.description_th, r.categories, r.created_at, r.updated_at
FROM review r
JOIN registration reg
ON r.registration_id= reg.id
JOIN event_occurrence eo 
ON reg.event_occurrence_id = eo.id
WHERE eo.event_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;