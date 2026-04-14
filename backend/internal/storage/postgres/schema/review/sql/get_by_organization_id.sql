SELECT r.id, r.registration_id, r.guardian_id, eo.event_id, r.rating, r.description_en, r.description_th, r.categories, r.created_at, r.updated_at
FROM review r
JOIN registration reg
ON r.registration_id= reg.id
JOIN event_occurrence eo 
ON reg.event_occurrence_id = eo.id
JOIN event e ON e.id = eo.event_id
WHERE e.organization_id = $1
ORDER BY %s 
LIMIT $2 OFFSET $3;