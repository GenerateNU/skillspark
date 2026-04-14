SELECT
    e.id AS event_id,
    COUNT(r.id) AS total_reviews,
    COALESCE(AVG(r.rating), 0) AS average_rating
FROM event e
LEFT JOIN event_occurrence eo ON eo.event_id = e.id
LEFT JOIN registration reg ON reg.event_occurrence_id = eo.id
LEFT JOIN review r ON r.registration_id = reg.id
WHERE e.organization_id = $1
GROUP BY e.id
ORDER BY %s
LIMIT $2 OFFSET $3;