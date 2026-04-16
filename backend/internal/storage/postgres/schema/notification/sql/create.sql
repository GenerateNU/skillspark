INSERT INTO scheduled_notification (
    notification_type,
    recipient_email,
    recipient_push_token,
    subject,
    body,
    metadata,
    scheduled_for,
    status,
    guardian_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
    guardian_id,
    created_at,
    updated_at;
