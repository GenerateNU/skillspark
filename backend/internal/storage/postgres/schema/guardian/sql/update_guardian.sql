UPDATE guardian
SET 
    expo_push_token = $2,
    push_notifications = COALESCE($3, push_notifications),
    email_notifications = COALESCE($4, email_notifications),
    updated_at = NOW()
WHERE id = $1
RETURNING user_id, stripe_customer_id, expo_push_token, push_notifications, email_notifications, created_at, updated_at;