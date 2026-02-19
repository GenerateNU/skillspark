SELECT 
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.created_at,
    r.updated_at,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time,
    r.reminder_sent,
    u.email AS guardian_email,
    u.name AS guardian_name,
    g.line_account_id
FROM registration r
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
LEFT JOIN guardian g ON r.guardian_id = g.id
LEFT JOIN "user" u ON g.user_id = u.id
WHERE r.reminder_sent = false
AND r.status = 'registered'
AND eo.start_time >= $1 AND eo.start_time <= $2;
