INSERT INTO scheduled_notification (
    notification_type,
    recipient_email,
    recipient_push_token,
    subject,
    body,
    metadata,
    scheduled_for,
    status
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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

