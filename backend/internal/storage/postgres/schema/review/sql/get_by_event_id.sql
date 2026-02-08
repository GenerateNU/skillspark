SELECT r.id, r.registration_id, r.guardian_id, r.description, r.categories, r.created_at, r.updated_at
FROM review r
JOIN registration reg
ON r.registration_id= reg.id
JOIN event_occurrences eo 
ON reg.event_occurrence_id = eo.id
WHERE eo.event_id = $1;