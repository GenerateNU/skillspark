SELECT 
    s.rating,
    COUNT(r.id) AS review_count
FROM generate_series(1, 5) AS s(rating)
LEFT JOIN review r 
    ON r.rating = s.rating
    AND r.registration_id IN (
        SELECT reg.id
        FROM registration reg
        JOIN event_occurrence eo ON reg.event_occurrence_id = eo.id
        WHERE eo.event_id = $1
    )
GROUP BY s.rating
ORDER BY s.rating;