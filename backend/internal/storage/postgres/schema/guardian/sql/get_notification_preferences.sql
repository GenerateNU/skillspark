SELECT id, push_notifications, email_notifications
FROM guardian
WHERE id = ANY($1::uuid[]);
