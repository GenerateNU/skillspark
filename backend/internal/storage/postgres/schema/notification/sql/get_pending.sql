SELECT 
    id,
    notification_type,
    recipient_email,
    recipient_push_token,
    subject,
    body,
    metadata,
    scheduled_for,
    sent_at,
    status,
    created_at,
    updated_at
FROM scheduled_notification
WHERE scheduled_for <= NOW()
  AND status = 'pending'
ORDER BY scheduled_for ASC
FOR UPDATE SKIP LOCKED;

