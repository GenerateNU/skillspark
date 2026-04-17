SELECT
    r.id,
    r.guardian_id,
    r.event_occurrence_id
FROM registration r
LEFT JOIN payment p ON p.registration_id = r.id
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
WHERE r.status = 'registered'
  AND p.id IS NULL
  AND eo.start_time > NOW()
  AND eo.start_time <= NOW() + INTERVAL '4 days'
ORDER BY eo.start_time ASC;
