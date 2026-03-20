UPDATE scheduled_notification
SET 
    status = $2,
    sent_at = CASE WHEN $2 = 'sent' THEN NOW() ELSE sent_at END,
    updated_at = NOW()
WHERE id = $1
RETURNING 
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
    updated_at;

